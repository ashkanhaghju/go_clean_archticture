package usecase

import "go_web_boilerplate/internal/entity"

type User interface {
	GetAllUsers() ([]entity.User, error)
}
