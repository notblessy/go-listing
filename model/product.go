package model

import (
	"time"
)

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
