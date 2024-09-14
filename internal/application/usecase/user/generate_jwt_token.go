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

type CreateJwtTokenHandler struct {
	UserDB       interfaces.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewCreateJwtTokenHandler(db interfaces.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *CreateJwtTokenHandler {
	return &CreateJwtTokenHandler{
		UserDB:       db,
		Jwt:          jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

func (t *CreateJwtTokenHandler) Execute(r *http.Request) (dto.JwtToken, error) {
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

	_, tokenString, err := t.Jwt.Encode(map[string]interface{}{
		"id":  u.ID,
		"exp": time.Now().Add(time.Duration(t.JwtExpiresIn) * time.Second).Unix(),
	})

	token.AccessToken = tokenString

	return token, nil
}
