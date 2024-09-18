package controller

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
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
	findProductById := product.NewFindProductByIdUsecase(productDB)
	findAllProducts := product.NewFindAllProductsUsecase(productDB)
	udpateProduct := product.NewUpdateProductUsecase(productDB)
	deleteProduct := product.NewDeleteProductUsecase(productDB)

	return &ProductUseCases{
		CreateProduct:   createProductUsecase,
		FindProductById: findProductById,
		FindAllProducts: findAllProducts,
		UpdateProduct:   udpateProduct,
		DeleteProduct:   deleteProduct,
		Multiplexer:     mux,
	}
}

// Create Product godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param request body dto.CreateProductInput true "product request"
// @Success 201 {string} string "Product created"
// @Failure 500 {string} string "Internal server error"
// @Router /product [post]
// @Security apiKeyAuth
func (p *ProductUseCases) createProductHandler(w http.ResponseWriter, r *http.Request) {
	err := p.CreateProduct.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product created"))
}

// Find Product by ID godoc
// @Summary Find product by ID
// @Description Find product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 400 {string} {object} Error
// @Failure 500 {string} {object} Error
// @Router /product/{id} [get]
// @Security apiKeyAuth
func (p *ProductUseCases) findProductHandler(w http.ResponseWriter, r *http.Request) {
	prd, err := p.FindProductById.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prd)
}

// Find All Products godoc
// @Summary Find all products
// @Description Find all products
// @Tags Product
// @Accept json
// @Produce json
// @Param page query string false "page number"
// @Param limit query string false "limit of products"
// @Success 200 {array} entity.Product
// @Failure 400 {string} {object} Error
// @Failure 500 {string} {object} Error
// @Router /product [get]
// @Security apiKeyAuth
func (p *ProductUseCases) findAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := p.FindAllProducts.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Update Product godoc
// @Summary Update a product
// @Description Update a product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body dto.CreateProductInput true "product request"
// @Success 200 {object} entity.Product
// @Failure 400 {string} {object} Error
// @Failure 500 {string} {object} Error
// @Router /product/{id} [put]
// @Security apiKeyAuth
func (p *ProductUseCases) updateProductHandler(w http.ResponseWriter, r *http.Request) {
	prd, err := p.UpdateProduct.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prd)
}

// Delete Product godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {string} string "Product deleted"
// @Failure 400 {string} {object} Error
// @Failure 500 {string} {object} Error
// @Router /product/{id} [delete]
// @Security apiKeyAuth
func (p *ProductUseCases) deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	err := p.DeleteProduct.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted"))
}

func (p *ProductUseCases) Register(jwt *jwtauth.JWTAuth) {
	p.Multiplexer.Route("/product", func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwt))
		r.Use(jwtauth.Authenticator)

		r.Post("/", p.createProductHandler)
		r.Get("/{id}", p.findProductHandler)
		r.Get("/product", p.findAllProductsHandler)
		r.Put("/product/{id}", p.updateProductHandler)
		r.Delete("/product/{id}", p.deleteProductHandler)
	})
}
