package service

import (
	"go_web_boilerplate/internal/entity"
	"go_web_boilerplate/internal/pkg/auth"
)

type User interface {
	GetAllUsers() ([]entity.User, error)
	GenerateToken(user string, password string) (*auth.JwtResponse, error)
}
