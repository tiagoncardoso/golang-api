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
	CreateUser  usecase.GeneralInterface
	GenJwtToken usecase.ResponseWithData[dto.JwtToken]
	Multiplexer *chi.Mux
}

func NewUserController(db *gorm.DB, mux *chi.Mux, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserUseCases {
	userDB := repository.NewUser(db)
	createUser := user.NewCreateUserUsecase(userDB)
	genJwtToken := user.NewCreateJwtTokenUsecase(userDB, jwt, jwtExpiresIn)

	return &UserUseCases{
		CreateUser:  createUser,
		GenJwtToken: genJwtToken,
		Multiplexer: mux,
	}
}

func (u *UserUseCases) createUserHandler(w http.ResponseWriter, r *http.Request) {
	err := u.CreateUser.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}

func (u *UserUseCases) genAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	jwtToken, err := u.GenJwtToken.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jwtToken)
}

func (u *UserUseCases) Register() {
	u.Multiplexer.Route("/user", func(r chi.Router) {
		r.Post("/", u.createUserHandler)
		r.Post("/getToken", u.genAccessTokenHandler)
	})
}
