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
	createProductUsecase := product.NewCreateProductUsecase(productDB)
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

func (p *ProductUseCases) createProductHandler(w http.ResponseWriter, r *http.Request) {
	err := p.CreateProduct.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product created"))
}

func (p *ProductUseCases) findProductHandler(w http.ResponseWriter, r *http.Request) {
	prd, err := p.FindProductById.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prd)
}

func (p *ProductUseCases) findAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := p.FindAllProducts.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (p *ProductUseCases) updateProductHandler(w http.ResponseWriter, r *http.Request) {
	prd, err := p.UpdateProduct.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prd)
}

func (p *ProductUseCases) deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	err := p.DeleteProduct.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted"))
}

func (p *ProductUseCases) Register() {
	p.Multiplexer.Route("/product", func(r chi.Router) {
		r.Post("/", p.createProductHandler)
		r.Get("/{id}", p.findProductHandler)
		r.Get("/product", p.findAllProductsHandler)
		r.Put("/product/{id}", p.updateProductHandler)
		r.Delete("/product/{id}", p.deleteProductHandler)
	})
}
