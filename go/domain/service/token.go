package service

import (
	"fmt"
	"time"
	"wing/infrastructure/auth"

	"github.com/dgrijalva/jwt-go"
)

type TokenService interface {
	CreateToken(id uint, name string) (*auth.TokenDetails, error)
}

type tokenService struct {
	// userRepo repository.UserRepository
}

func NewTokenService() TokenService {
	return &tokenService{}
}

func (ts *tokenService) CreateToken(id uint, name string) (*auth.TokenDetails, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	fmt.Println(t)
	return nil, nil
}
