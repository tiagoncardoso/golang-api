package usecase

import (
	"encoding/json"
	"github.com/tiagoncardoso/golang-api/internal/application/dto"
	"github.com/tiagoncardoso/golang-api/internal/entity"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository/interfaces"
	"net/http"
)

type ProductHandler struct {
	ProductDB interfaces.ProductInterface
}

func NewProductHandler(db interfaces.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) Execute(r *http.Request) error {
	var product dto.CreateProductInput

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		return err
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		return err
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		return err
	}

	return nil
}
