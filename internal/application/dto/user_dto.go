package dto

type CreateUserInput struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type GenerateTokenInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtToken struct {
	AccessToken string `json:"access_token"`
}
