package _interface

import "github.com/tiagoncardoso/golang-api/internal/entity"

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id int) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
