package controller

import (
	"github.com/tiagoncardoso/golang-api/internal/application/usecase"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository"
	"gorm.io/gorm"
	"net/http"
)

type ProductUseCases struct {
	CreateProduct usecase.GeneralInterface
}

func NewProductController(db *gorm.DB) *ProductUseCases {
	productDB := repository.NewProduct(db)
	createProductUsecase := usecase.NewProductHandler(productDB)

	return &ProductUseCases{
		CreateProduct: createProductUsecase,
	}
}

func (p *ProductUseCases) initializeRoutes() {
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		err := p.CreateProduct.Execute(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusFound)
		w.Write([]byte("Product created"))
	})
}
