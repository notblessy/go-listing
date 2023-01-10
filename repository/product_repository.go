package repository

import (
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
