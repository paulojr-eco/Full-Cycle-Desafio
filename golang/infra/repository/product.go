package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/paulojr-eco/full-cycle-desafio/domain/model"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (p ProductRepositoryDb) AddProduct(product *model.Product) (*model.Product, error) {
	err := p.Db.Create(product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p ProductRepositoryDb) FindProducts() *[]model.Product {
	var products []model.Product

	p.Db.Find(&products)

	return &products
}
