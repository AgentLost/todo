package service_impl

import "todo-app/internal/helpers"

type Auth struct {
	provider *helpers.TokenProvider
}

func (a *Auth) Validate(token string) (bool, error) {
	return a.provider.Validate(token)
}

func (a *Auth) GetHeader() string {
	return a.provider.Header
}

func (a *Auth) GetUserId(token string) (int, error) {
	return a.provider.GetUserId(token)
}

func (a *Auth) Generate(userId int) (string, error) {
	return a.provider.Generate(userId)
}

func NewAuthService(p *helpers.TokenProvider) *Auth {
	return &Auth{provider: p}
}
