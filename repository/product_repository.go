package repository

import (
	"fmt"
	"strings"

	"github.com/notblessy/go-listing/model"
	"github.com/notblessy/go-listing/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository :nodoc:
func NewProductRepository(d *gorm.DB) model.ProductRepository {
	return &productRepository{
		db: d,
	}
}

// Create :nodoc:
func (r *productRepository) Create(product *model.Product) error {
	err := r.db.Create(&product).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"product": utils.Dump(product),
		}).Error(err)

		return err
	}

	return err
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

	err = qb.Find(&products).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return products, err
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
