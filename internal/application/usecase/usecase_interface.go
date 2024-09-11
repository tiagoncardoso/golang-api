package usecase

import "net/http"

type GeneralInterface interface {
	Execute(r *http.Request) error
}
