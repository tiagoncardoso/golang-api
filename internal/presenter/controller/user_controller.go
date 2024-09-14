package controller

import (
	"github.com/go-chi/chi"
	"github.com/tiagoncardoso/golang-api/internal/application/usecase"
	"github.com/tiagoncardoso/golang-api/internal/application/usecase/user"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository"
	"gorm.io/gorm"
	"net/http"
)

type UserUseCases struct {
	CreateUserHandler usecase.GeneralInterface
	Multiplexer       *chi.Mux
}

func NewUserController(db *gorm.DB, mux *chi.Mux) *UserUseCases {
	userDB := repository.NewUser(db)
	createUser := user.NewCreateUserHandler(userDB)

	return &UserUseCases{
		CreateUserHandler: createUser,
		Multiplexer:       mux,
	}
}

func (u *UserUseCases) createUser() {
	u.Multiplexer.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		err := u.CreateUserHandler.Execute(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created"))
	})
}

func (u *UserUseCases) Register() {
	u.createUser()
}
