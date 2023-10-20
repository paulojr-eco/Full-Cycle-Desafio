package model_test

import (
	"testing"

	"github.com/paulojr-eco/full-cycle-desafio/domain/model"
	"github.com/stretchr/testify/require"
)

/* Should throw an error on validation fail */
func TestModel_NewWrongProduct_MissingParams(t *testing.T) {
	price := float32(15.99)

	_, err := model.NewProduct("", "", price)
	require.NotNil(t, err)
}

/* Should return a product on success */
func TestModel_NewProduct(t *testing.T) {
	name := "Product 01"
	description := "Description from product 01"
	price := float32(15.99)

	product, err := model.NewProduct(name, description, price)

	require.Nil(t, err)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.Name, name)
	require.Equal(t, product.Description, description)
	require.Equal(t, product.Price, price)
	require.NotNil(t, product.CreatedAt)
	require.True(t, product.UpdatedAt.IsZero())
}
