package usecase

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/tiagoncardoso/golang-api/internal/entity"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository/interfaces"
	entityPkg "github.com/tiagoncardoso/golang-api/pkg/entity"
	"net/http"
)

type DeleteProductHandler struct {
	ProductDB interfaces.ProductInterface
}

func NewDeleteProductHandler(db interfaces.ProductInterface) *DeleteProductHandler {
	return &DeleteProductHandler{
		ProductDB: db,
	}
}

func (h *DeleteProductHandler) Execute(r *http.Request) error {
	var product *entity.Product

	id := chi.URLParam(r, "id")
	if id == "" {
		return errors.New("id is required")
	}

	_, err := entityPkg.ParseID(id)
	if err != nil {
		return err
	}

	product, err = h.ProductDB.FindByID(id)
	if err != nil {
		return err
	}

	err = h.ProductDB.Delete(product.ID.String())

	return nil
}
