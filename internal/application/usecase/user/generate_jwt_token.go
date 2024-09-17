package user

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/jwtauth"
	"github.com/tiagoncardoso/golang-api/internal/application/dto"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository/interfaces"
	"net/http"
	"time"
)

type CreateJwtTokenUsecase struct {
	UserDB interfaces.UserInterface
}

func NewCreateJwtTokenUsecase(db interfaces.UserInterface) *CreateJwtTokenUsecase {
	return &CreateJwtTokenUsecase{
		UserDB: db,
	}
}

func (t *CreateJwtTokenUsecase) Execute(r *http.Request) (dto.JwtToken, error) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var user dto.GenerateTokenInput
	var token dto.JwtToken

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return token, err
	}

	u, err := t.UserDB.FindByEmail(user.Email)
	if err != nil {
		return token, err
	}

	if !u.ValidatePassword(user.Password) {
		return token, errors.New("invalid password")
	}

	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"id":  u.ID,
		"exp": time.Now().Add(time.Duration(jwtExpiresIn) * time.Second).Unix(),
	})

	token.AccessToken = tokenString

	return token, nil
}
