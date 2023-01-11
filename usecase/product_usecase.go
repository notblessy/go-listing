package usecase

import (
	"github.com/notblessy/go-listing/model"
	"github.com/notblessy/go-listing/utils"
	"github.com/sirupsen/logrus"
)

type productUsecase struct {
	productRepo model.ProductRepository
}

// NewProductUsecase :nodoc:
func NewProductUsecase(p model.ProductRepository) model.ProductUsecase {
	return &productUsecase{
		productRepo: p,
	}
}

// Create :nodoc:
func (u *productUsecase) Create(product *model.Product) (string, error) {
	if product.ID == "" {
		product.ID = utils.GenerateID()
	}

	err := u.productRepo.Create(product)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"product": utils.Dump(product),
		}).Error(err)

		return "", err
	}

	return product.ID, nil
}

// FindAll :nodoc:
func (u *productUsecase) FindAll(req *model.ProductQuery) (*[]model.Product, error) {
	products, err := u.productRepo.FindAll(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"productQuery": utils.Dump(req),
		}).Error(err)

		return nil, err
	}

	return products, nil
}
