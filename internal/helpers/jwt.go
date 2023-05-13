package helpers

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
	"todo-app/config"
)

type TokenProvider struct {
	SecretKey       string
	ValiditySeconds int
	Header          string
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (p *TokenProvider) Validate(token string) (bool, error) {
	claims := &tokenClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(p.SecretKey), nil
	})

	if _, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return true, err
	} else {
		return false, err
	}
}

func (p *TokenProvider) Generate(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(p.ValiditySeconds) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(p.SecretKey))
}

func (p *TokenProvider) GetUserId(token string) (int, error) {
	claims := &tokenClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(p.SecretKey), nil
	})

	return claims.UserId, err
}

func (p *TokenProvider) parseToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(p.SecretKey), nil
}

func NewTokenProvider(cfg *config.Config) *TokenProvider {
	return &TokenProvider{
		SecretKey:       cfg.SecretKey,
		ValiditySeconds: cfg.ValiditySeconds,
		Header:          cfg.Header}
}
