package model

import (
	"log"
	"math"
	"sync"
	"time"

	"github.com/asaskevich/govalidator"
)

type ProductRepositoryInterface interface {
	AddProduct(product *Product) (*Product, error)
	FindProducts() *[]Product
}

type Counter struct {
	counter int32
}

var (
	instance *Counter
	once     sync.Once
)

func GetInstance() *Counter {
	once.Do(func() {
		instance = &Counter{counter: math.MinInt32}
	})
	return instance
}

func (c *Counter) Increment() int32 {
	if c.counter == math.MaxInt32 {
		c.counter = math.MinInt32
		log.Fatal("exceeded limit of product creation")
	}
	c.counter++
	return c.counter
}

func (c *Counter) GetValue() int32 {
	return c.counter
}

type Product struct {
	Base        `valid:"required"`
	Name        string  `json:"name" gorm:"type:varchar(50);not null" valid:"notnull"`
	Description string  `json:"description" gorm:"type:varchar(255);not null" valid:"notnull"`
	Price       float32 `json:"price" gorm:"type:float" valid:"notnull"`
}

func (product *Product) isValid() error {
	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return err
	}

	return nil
}

func NewProduct(name string, description string, price float32) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	counterSingleton := GetInstance()
	product.ID = counterSingleton.GetValue()
	counterSingleton.Increment()

	product.CreatedAt = time.Now()

	err := product.isValid()
	if err != nil {
		return nil, err
	}

	return &product, nil
}
