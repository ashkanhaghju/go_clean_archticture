package usecase

import "go_web_boilerplate/internal/pkg/auth"

type Token interface {
	GenerateToken(user string, password string) (*auth.JwtResponse, error)
}
