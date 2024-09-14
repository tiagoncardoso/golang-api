package usecase

import "net/http"

type GeneralInterface interface {
	Execute(r *http.Request) error
}

type ResponseWithData[T interface{}] interface {
	Execute(r *http.Request) (T, error)
}
