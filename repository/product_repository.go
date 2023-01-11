package repository

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/notblessy/go-listing/config"
	"github.com/notblessy/go-listing/model"
	"github.com/notblessy/go-listing/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type productRepository struct {
	db    *gorm.DB
	cache redis.Conn
}

// NewProductRepository :nodoc:
func NewProductRepository(d *gorm.DB, r redis.Conn) model.ProductRepository {
	return &productRepository{
		db:    d,
		cache: r,
	}
}

// Create :nodoc:
func (p *productRepository) Create(product *model.Product) error {
	logger := logrus.WithFields(logrus.Fields{
		"product": utils.Dump(product),
	})

	err := p.db.Create(&product).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	cacheKey := p.newProductCacheKey(product.ID)

	err = p.storeCache(cacheKey, product)
	if err != nil {
		logger.Error(err)
		return err
	}

	return err
}

// FindByID :nodoc:
func (p *productRepository) FindByID(id string) (product *model.Product, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"id": utils.Dump(id),
	})

	err = p.db.First(&product, id).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	cacheKey := p.newProductCacheKey(product.ID)

	err = p.storeCache(cacheKey, product)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return product, err
}

// FindAll :nodoc:
func (p *productRepository) FindAll(req *model.ProductQuery) (products *[]model.Product, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"memoryRequest": req,
	})

	qb := p.db.Model(products)

	if req.Sort != "" {
		sorts := p.sortHandler(req.Sort)
		for _, s := range sorts {
			if s.Field != "" {
				qb.Order(fmt.Sprintf("%s %s", s.Field, s.Type))
			}
		}
	} else {
		qb.Order(model.DefaultProductSort)
	}

	var ids []string
	err = qb.Pluck("id", &ids).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	var results []model.Product

	for _, id := range ids {
		product := &model.Product{}
		cacheKey := p.newProductCacheKey(id)

		// Find from cache
		product, err = p.findByIDFromCache(cacheKey)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		if product != nil {
			results = append(results, *product)
			continue
		}

		// Find from database if cache is not found
		product, err = p.FindByID(id)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		results = append(results, *product)
	}

	return &results, err
}

func (p *productRepository) sortHandler(sortReq string) (sortResults []utils.Sort) {
	sorts := strings.Split(sortReq, ",")

	if len(sorts) > 0 {
		for _, s := range sorts {
			sortParams := strings.Split(s, "-")

			if len(sortParams) == 2 {
				if _, isSortValidValueExists := model.ProductSortValidValues[sortParams[0]]; isSortValidValueExists {
					sortResults = append(sortResults, utils.Sort{
						Field: sortParams[0],
						Type:  strings.ToLower(sortParams[1]),
					})
				}
			}
		}
	}

	return sortResults
}

func (p *productRepository) newProductCacheKey(id string) string {
	return fmt.Sprintf("product:%s", id)
}

func (p *productRepository) findByIDFromCache(cacheKey string) (product *model.Product, err error) {
	res, err := p.cache.Do("GET", cacheKey)
	if err != nil {
		return nil, err
	}

	switch res {
	case nil:
		return nil, nil
	default:
		err := json.Unmarshal(res.([]byte), &product)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (p *productRepository) storeCache(cacheKey string, product *model.Product) error {
	b, err := json.Marshal(product)
	_, err = p.cache.Do("SETEX", cacheKey, config.RedisTTL(), b)
	if err != nil {
		return err
	}

	return nil
}
