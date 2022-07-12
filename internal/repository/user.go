package repository

import (
	"go_web_boilerplate/internal/entity"
	"go_web_boilerplate/internal/pkg/auth"
)

type UserRepository interface {
	GetAllUser() ([]entity.User, error)
	GenerateToken(user string, password string) (*auth.JwtResponse, error)
}
