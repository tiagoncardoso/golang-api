package controller

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/tiagoncardoso/golang-api/internal/application/dto"
	"github.com/tiagoncardoso/golang-api/internal/application/usecase"
	"github.com/tiagoncardoso/golang-api/internal/application/usecase/user"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository"
	"gorm.io/gorm"
	"net/http"
)

type UserUseCases struct {
	CreateUserHandler  usecase.GeneralInterface
	GenJwtTokenHandler usecase.ResponseWithData[dto.JwtToken]
	Multiplexer        *chi.Mux
}

func NewUserController(db *gorm.DB, mux *chi.Mux, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserUseCases {
	userDB := repository.NewUser(db)
	createUser := user.NewCreateUserHandler(userDB)
	genJwtToken := user.NewCreateJwtTokenHandler(userDB, jwt, jwtExpiresIn)

	return &UserUseCases{
		CreateUserHandler:  createUser,
		GenJwtTokenHandler: genJwtToken,
		Multiplexer:        mux,
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

func (u *UserUseCases) genAccessToken() {
	u.Multiplexer.Post("/user/getToken", func(w http.ResponseWriter, r *http.Request) {
		jwtToken, err := u.GenJwtTokenHandler.Execute(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(jwtToken)
	})
}

func (u *UserUseCases) Register() {
	u.createUser()
	u.genAccessToken()
}
