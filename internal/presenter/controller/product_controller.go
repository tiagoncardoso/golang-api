package controller

import (
	"github.com/tiagoncardoso/golang-api/internal/application/usecase"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository"
	"gorm.io/gorm"
	"net/http"
)

type ProductUseCases struct {
	CreateProduct usecase.GeneralInterface
	Multiplexer   *http.ServeMux
}

func NewProductController(db *gorm.DB, mux *http.ServeMux) *ProductUseCases {
	productDB := repository.NewProduct(db)
	createProductUsecase := usecase.NewProductHandler(productDB)

	return &ProductUseCases{
		CreateProduct: createProductUsecase,
		Multiplexer:   mux,
	}
}

func (p *ProductUseCases) InitializeRoutes() {
	p.Multiplexer.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		err := p.CreateProduct.Execute(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Product created"))
	})
}
