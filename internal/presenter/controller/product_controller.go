package controller

import (
	"github.com/go-chi/chi"
	"github.com/tiagoncardoso/golang-api/internal/application/usecase"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository"
	"gorm.io/gorm"
	"net/http"
)

type ProductUseCases struct {
	CreateProduct usecase.GeneralInterface
	Multiplexer   *chi.Mux
}

func NewProductController(db *gorm.DB, mux *chi.Mux) *ProductUseCases {
	productDB := repository.NewProduct(db)
	createProductUsecase := usecase.NewProductHandler(productDB)

	return &ProductUseCases{
		CreateProduct: createProductUsecase,
		Multiplexer:   mux,
	}
}

func (p *ProductUseCases) createProduct() {
	p.Multiplexer.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		err := p.CreateProduct.Execute(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Product created"))
	})
}

func (p *ProductUseCases) InitializeRoutes() {
	p.createProduct()
}
