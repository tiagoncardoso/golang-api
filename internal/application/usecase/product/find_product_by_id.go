package product

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/tiagoncardoso/golang-api/internal/entity"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository/interfaces"
	entityPkg "github.com/tiagoncardoso/golang-api/pkg/entity"
	"net/http"
)

type FindProductByIdUsecase struct {
	ProductDB interfaces.ProductInterface
}

func NewFindProductByIdUsecase(db interfaces.ProductInterface) *FindProductByIdUsecase {
	return &FindProductByIdUsecase{
		ProductDB: db,
	}
}

func (h *FindProductByIdUsecase) Execute(r *http.Request) (*entity.Product, error) {
	var product *entity.Product

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

	return product, nil
}
