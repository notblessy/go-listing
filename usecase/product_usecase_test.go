package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/notblessy/go-listing/model"
	"github.com/notblessy/go-listing/model/mock"
	"github.com/notblessy/go-listing/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPostgreMock() (db *gorm.DB, mock sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	db, err = gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return
}

func TestProductUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRepo := mock.NewMockProductRepository(ctrl)

	pr := productUsecase{
		productRepo: productRepo,
	}

	product := &model.Product{
		ID:          utils.GenerateID(),
		Name:        "Oud Vibrant Leather",
		Price:       399999,
		Description: "A Perfume from Zara",
		Quantity:    10,
	}

	t.Run("Success", func(t *testing.T) {
		productRepo.EXPECT().Create(product).
			Times(1).
			Return(nil)

		err := pr.productRepo.Create(product)
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		productRepo.EXPECT().Create(product).
			Times(1).
			Return(errors.New("Internal server error"))

		err := pr.productRepo.Create(product)
		assert.Error(t, err)
	})
}

func TestProductUsecase_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRepo := mock.NewMockProductRepository(ctrl)

	pr := productUsecase{
		productRepo: productRepo,
	}

	req := &model.ProductQuery{
		Sort: "price-desc",
	}

	product := model.Product{
		ID:          "12345ABC",
		Name:        "Oud Vibrant Leather",
		Price:       399999,
		Description: "A Perfume from Zara",
		Quantity:    10,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	results := []model.Product{}
	results = append(results, product)

	t.Run("Success", func(t *testing.T) {
		productRepo.EXPECT().FindAll(req).
			Times(1).
			Return(&results, nil)

		res, err := pr.productRepo.FindAll(req)
		assert.NoError(t, err)
		assert.Equal(t, &results, res)

	})

	t.Run("Error", func(t *testing.T) {
		productRepo.EXPECT().FindAll(req).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		_, err := pr.productRepo.FindAll(req)
		assert.Error(t, err)
	})
}
