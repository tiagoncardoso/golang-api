package controller

import (
	"encoding/json"
	"github.com/go-chi/chi"
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

type Error struct {
	Message string `json:"message"`
}

func NewUserController(db *gorm.DB, mux *chi.Mux) *UserUseCases {
	userDB := repository.NewUser(db)
	createUser := user.NewCreateUserUsecase(userDB)
	genJwtToken := user.NewCreateJwtTokenUsecase(userDB)

	return &UserUseCases{
		CreateUser:  createUser,
		GenJwtToken: genJwtToken,
		Multiplexer: mux,
	}
}

// Create user godoc
// @Summary Create user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param request body dto.CreateUserInput true "user request"
// @Success 201 {string} string "User created"
// @Failure 500 {string} {object} Error
// @Router /user [post]
func (u *UserUseCases) createUserHandler(w http.ResponseWriter, r *http.Request) {
	err := u.CreateUser.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}

// Generate access token godoc
// @Summary Generate access token
// @Description Generate a new access token
// @Tags user
// @Accept json
// @Produce json
// @Param request body dto.GenerateTokenInput true "user credentials"
// @Success 200 {object} dto.JwtToken
// @Failure 401 {string} {object} Error
// @Router /user/getToken [post]
func (u *UserUseCases) genAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	jwtToken, err := u.GenJwtToken.Execute(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)

		return
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
