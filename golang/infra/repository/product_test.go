package repository_test

import (
	"testing"

	"github.com/paulojr-eco/full-cycle-desafio/domain/model"
	"github.com/paulojr-eco/full-cycle-desafio/infra/db"
	"github.com/paulojr-eco/full-cycle-desafio/infra/repository"
	"github.com/stretchr/testify/require"
)

/* Should add a new product on success */
func TestRepository_AddProduct(t *testing.T) {
	name := "Product 01"
	description := "Description from product 01"
	price := float32(15.99)
	product, _ := model.NewProduct(name, description, price)
	database, _ := db.ConnectDB()

	productRepository := repository.ProductRepositoryDb{Db: database}
	product, err := productRepository.AddProduct(product)

	require.NotNil(t, product)
	require.Nil(t, err)
}

/* Should return an empty array if there is no products */
func TestRepository_FindZeroProducts(t *testing.T) {
	database, _ := db.ConnectDB()

	productRepository := repository.ProductRepositoryDb{Db: database}
	products := productRepository.FindProducts()

	require.Equal(t, len(*products), 0)
}

/* Should find a product on success */
func TestRepository_FindProducts(t *testing.T) {
	name := "Product 01"
	description := "Description from product 01"
	price := float32(15.99)
	product, _ := model.NewProduct(name, description, price)
	database, _ := db.ConnectDB()

	productRepository := repository.ProductRepositoryDb{Db: database}
	productRepository.AddProduct(product)
	products := productRepository.FindProducts()

	require.NotNil(t, products)
	require.Equal(t, len(*products), 1)
}
