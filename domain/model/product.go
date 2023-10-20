package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Product struct {
	Base        `valid:"required"`
	Name        string  `json:"name" gorm:"type:varchar(50);not null" valid:"notnull"`
	Description string  `json:"description" gorm:"type:varchar(255);not null" valid:"notnull"`
	Price       float64 `json:"price" gorm:"type:float" valid:"notnull"`
}

func NewProduct(name string, description string, price float64) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	product.ID = uuid.NewV4().String()
	product.CreatedAt = time.Now()

	return &product, nil
}
