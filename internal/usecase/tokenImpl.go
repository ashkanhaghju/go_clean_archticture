package usecase

import (
	"go_web_boilerplate/internal/pkg/auth"
	"go_web_boilerplate/internal/pkg/logger"
	"go_web_boilerplate/internal/repository"
)

type TokenImpl struct {
	repo repository.UserRepository
}

func NewTokenUseCase(repo repository.UserRepository, logger logger.Logger) TokenImpl {
	return TokenImpl{
		repo: repo,
	}
}

func (t TokenImpl) GenerateToken(user string, password string) (*auth.JwtResponse, error) {
	return t.repo.GenerateToken(user, password)
}
