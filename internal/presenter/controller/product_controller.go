package controller

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/tiagoncardoso/golang-api/internal/application/usecase"
	"github.com/tiagoncardoso/golang-api/internal/application/usecase/product"
	"github.com/tiagoncardoso/golang-api/internal/entity"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository"
	"gorm.io/gorm"
	"net/http"
)

type ProductUseCases struct {
	CreateProduct   usecase.GeneralInterface
	FindProductById usecase.ResponseWithData[*entity.Product]
	FindAllProducts usecase.ResponseWithData[*[]entity.Product]
	UpdateProduct   usecase.ResponseWithData[*entity.Product]
	DeleteProduct   usecase.GeneralInterface
	Multiplexer     *chi.Mux
}

func NewProductController(db *gorm.DB, mux *chi.Mux) *ProductUseCases {
	productDB := repository.NewProduct(db)
	createProductUsecase := product.NewCreateProductHandler(productDB)
	findProductById := product.NewFindProductHandler(productDB)
	findAllProducts := product.NewFindAllProductsHandler(productDB)
	udpateProduct := product.NewUpdateProductHandler(productDB)
	deleteProduct := product.NewDeleteProductHandler(productDB)

	return &ProductUseCases{
		CreateProduct:   createProductUsecase,
		FindProductById: findProductById,
		FindAllProducts: findAllProducts,
		UpdateProduct:   udpateProduct,
		DeleteProduct:   deleteProduct,
		Multiplexer:     mux,
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

func (p *ProductUseCases) findProduct() {
	p.Multiplexer.Get("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		product, err := p.FindProductById.Execute(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	})
}

func (p *ProductUseCases) findAllProducts() {
	p.Multiplexer.Get("/product", func(w http.ResponseWriter, r *http.Request) {
		products, err := p.FindAllProducts.Execute(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	})
}

func (p *ProductUseCases) updateProduct() {
	p.Multiplexer.Put("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		product, err := p.UpdateProduct.Execute(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	})
}

func (p *ProductUseCases) deleteProduct() {
	p.Multiplexer.Delete("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		err := p.DeleteProduct.Execute(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Product deleted"))
	})
}

func (p *ProductUseCases) InitializeRoutes() {
	p.createProduct()
	p.findProduct()
	p.findAllProducts()
	p.updateProduct()
	p.deleteProduct()
}
