package usecase

import "github.com/paulojr-eco/full-cycle-desafio/domain/model"

type ProductUseCase struct {
	ProductRepository model.ProductRepositoryInterface
}

func (p *ProductUseCase) AddProduct(name string, description string, price float64) (*model.Product, error) {
	product, err := model.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	product, err = p.ProductRepository.AddProduct(product)
	if product.ID == "" {
		return nil, err
	}

	return product, nil
}

func (p *ProductUseCase) FindProducts() *[]model.Product {
	products := p.ProductRepository.FindProducts()

	return products
}
