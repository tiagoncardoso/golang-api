package usecase

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/tiagoncardoso/golang-api/internal/entity"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository/interfaces"
	entityPkg "github.com/tiagoncardoso/golang-api/pkg/entity"
	"net/http"
)

type UpdateProductHandler struct {
	ProductDB interfaces.ProductInterface
}

func NewUpdateProductHandler(db interfaces.ProductInterface) *UpdateProductHandler {
	return &UpdateProductHandler{
		ProductDB: db,
	}
}

func (h *UpdateProductHandler) Execute(r *http.Request) (*entity.Product, error) {
	var product *entity.Product
	var tmpProduct *entity.Product

	id := chi.URLParam(r, "id")
	if id == "" {
		return product, errors.New("id is required")
	}

	_, err := entityPkg.ParseID(id)
	if err != nil {
		return product, err
	}

	product, err = h.ProductDB.FindByID(id)
	if err != nil {
		return product, err
	}

	err = json.NewDecoder(r.Body).Decode(&tmpProduct)
	if err != nil {
		return product, err
	}

	tmpProduct.ID = product.ID
	err = h.ProductDB.Update(tmpProduct)
	if err != nil {
		return product, err
	}

	return tmpProduct, nil
}
