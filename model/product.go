package model

import (
	"time"
)

// ProductSortValidValues valid sort values
var ProductSortValidValues = map[string]bool{
	"price": true,
	"name":  true,
}

// DefaultProductSort a default product sort
var DefaultProductSort = "created_at desc"

// ProductRepository :nodoc:
type ProductRepository interface {
	Create(product *Product) error
	FindAll(req *ProductQuery) (products *[]Product, err error)
	FindByID(id string) (product *Product, err error)
}

// ProductUsecase :nodoc:
type ProductUsecase interface {
	Create(product *Product) (string, error)
	FindAll(req *ProductQuery) (products *[]Product, err error)
}

// Product :nodoc:
type Product struct {
	ID          string    `gorm:"type:varchar(128)" json:"id"`
	Name        string    `gorm:"type:varchar(128)" json:"name" validate:"required"`
	Price       int64     `json:"price" validate:"required"`
	Description string    `json:"description"`
	Quantity    int32     `json:"quantity" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProductQuery :nodoc:
type ProductQuery struct {
	Sort string `json:"sort"`
}
