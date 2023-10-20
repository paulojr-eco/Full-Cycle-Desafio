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
	price := 15.99
	product, _ := model.NewProduct(name, description, price)
	database, _ := db.ConnectDB()

	productRepository := repository.ProductRepositoryDb{Db: database}
	product, err := productRepository.AddProduct(product)

	require.NotNil(t, product)
	require.Nil(t, err)
}
