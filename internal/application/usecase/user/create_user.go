package user

import (
	"encoding/json"
	"github.com/tiagoncardoso/golang-api/internal/application/dto"
	"github.com/tiagoncardoso/golang-api/internal/entity"
	"github.com/tiagoncardoso/golang-api/internal/infra/repository/interfaces"
	"net/http"
)

type CreateUserHandler struct {
	UserDB interfaces.UserInterface
}

func NewCreateUserHandler(db interfaces.UserInterface) *CreateUserHandler {
	return &CreateUserHandler{
		UserDB: db,
	}
}

func (h *CreateUserHandler) Execute(r *http.Request) error {
	var user dto.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return err
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	err = h.UserDB.Create(u)
	if err != nil {
		return err
	}

	return nil
}
