package service

import (
	"go_web_boilerplate/internal/entity"
	"go_web_boilerplate/internal/pkg/auth"
	"go_web_boilerplate/internal/pkg/logger"
	"go_web_boilerplate/internal/usecase"
)

type UserService struct {
	userUseCase  usecase.User
	tokenUseCase usecase.Token
	logger       logger.Logger
}

func NewUserService(useCase usecase.User, tokenUserCase usecase.Token, logger logger.Logger) UserService {
	return UserService{
		userUseCase:  useCase,
		tokenUseCase: tokenUserCase,
		logger:       logger,
	}
}

func (us UserService) GetAllUsers() ([]entity.User, error) {
	return us.userUseCase.GetAllUsers()
}

func (us UserService) GenerateToken(username string, password string) (*auth.JwtResponse, error) {
	return us.tokenUseCase.GenerateToken(username, password)
}
