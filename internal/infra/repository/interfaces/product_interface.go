package interfaces

import "github.com/tiagoncardoso/golang-api/internal/entity"

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) (*[]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
