package clients

import "net/http"

type AuthClient struct {
	client *http.Client
}

func NewAuthClient() *AuthClient {
	return &AuthClient{
		client: &http.Client{},
	}
}

func (a *AuthClient) ValidateToken(token string) (bool, error) {
	// Логика для запроса авторизации
	return true, nil
}
