package usecase

import (
	"go_web_boilerplate/internal/entity"
	"go_web_boilerplate/internal/pkg/logger"
	"go_web_boilerplate/internal/repository"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository, logger logger.Logger) UserUseCase {
	return UserUseCase{repo: repo}
}

func (u UserUseCase) GetAllUsers() ([]entity.User, error) {
	return u.repo.GetAllUser()
}
