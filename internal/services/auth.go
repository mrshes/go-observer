package services

import (
	"first-project/internal/database/models"
	"first-project/internal/database/repository"
	"first-project/pkg/hash"
	jwt2 "first-project/pkg/jwt"
	"fmt"
)

type AuthService struct {
	repo   repository.Users
	hasher hash.PasswordHasher
}

var jwtSecretKey = []byte("sdasdas3123asd123")

func newAuthService(repo repository.Users, hasher hash.PasswordHasher) *AuthService {
	return &AuthService{repo: repo, hasher: hasher}
}

// Регистрация
func (a *AuthService) Register(user *models.User) (*models.User, error) {
	ph, err := a.hasher.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = ph
	user, err = a.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Авторизация
func (a *AuthService) Login(user models.User) (string, error) {
	fmt.Println(user)
	jwt := jwt2.NewJwt(jwtSecretKey)
	token, err := jwt.GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
