package services

import (
	"context"
	"first-project/internal/config"
	"first-project/internal/database/models"
	"first-project/internal/database/repository"
	"first-project/pkg/hash"
)

type Auth interface {
	Register(user *models.User) (*models.User, error)
	Login(user models.User) (string, error)
}

type Services struct {
	Auth Auth
}

type Deps struct {
	Ctx    context.Context
	Repos  *repository.Repositories
	Hasher hash.PasswordHasher
	Cfg    config.Config
}

func NewServices(deps Deps) Services {
	return Services{
		Auth: newAuthService(deps.Repos.Users, deps.Hasher),
	}
}
