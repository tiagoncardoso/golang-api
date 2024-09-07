package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10.5)

	assert.Nil(t, err)
	assert.NotEmpty(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10.5, p.Price)
}

func TestProductWhenNameIsEmpty(t *testing.T) {
	p, err := NewProduct("", 10.5)

	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsEmpty(t *testing.T) {
	p, err := NewProduct("Product 1", 0)

	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsNegative(t *testing.T) {
	p, err := NewProduct("Product 1", -10)

	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProduct_Validate(t *testing.T) {
	p, err := NewProduct("Product 1", 10.5)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
