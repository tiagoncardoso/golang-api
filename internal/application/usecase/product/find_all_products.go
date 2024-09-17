package product

import (
	"fmt"
	"github.com/tiagoncardoso/golang-api/internal/entity"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository/interfaces"
	"net/http"
	"strconv"
)

type FindAllProductUsecase struct {
	ProductDB interfaces.ProductInterface
}

func NewFindAllProductsUsecase(db interfaces.ProductInterface) *FindAllProductUsecase {
	return &FindAllProductUsecase{
		ProductDB: db,
	}
}

func (h *FindAllProductUsecase) Execute(r *http.Request) (*[]entity.Product, error) {
	var products *[]entity.Product

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	fmt.Println("Finding all products", "page", page, "limit", limit, "sort", sort)

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	products, err = h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		return products, err
	}

	return products, nil
}
