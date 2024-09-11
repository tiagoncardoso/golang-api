package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tiagoncardoso/golang-api/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func dbConnectProduct() *gorm.DB {
	dsn := "file::memory:"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		panic(err)
	}

	return db
}

func TestProduct_Create(t *testing.T) {
	db := dbConnectProduct()
	product, err := entity.NewProduct("Product 1", 10.5)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)

	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db := dbConnectProduct()
	productDB := NewProduct(db)

	for i := range 25 {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)

		db.Create(product)
	}

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 0", products[0].Name)
	assert.Equal(t, "Product 9", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 10", products[0].Name)
	assert.Equal(t, "Product 19", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "Product 20", products[0].Name)
	assert.Equal(t, "Product 24", products[4].Name)
	assert.Panics(t, func() {
		_ = products[5].Name
	})
}
