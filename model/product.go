package model

import (
	"time"
)

// ProductRepository :nodoc:
type ProductRepository interface {
	Create(product *Product) error
	// FindAll(req ProductQuery) (products *[]Product, count int64, err error)
}

type ProductUsecase interface {
	Create(product *Product) (string, error)
	// FindAll(req ProductQuery) (products *[]Product, count int64, err error)
}

// Product :nodoc:
type Product struct {
	ID          string    `gorm:"type:varchar(128)" json:"id"`
	Name        string    `gorm:"type:varchar(128)" json:"name"`
	Price       int64     `json:"price"`
	Description string    `json:"description"`
	Quantity    int32     `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProductQuery :nodoc:
type ProductQuery struct {
	Sort string `json:"sort"`
}
