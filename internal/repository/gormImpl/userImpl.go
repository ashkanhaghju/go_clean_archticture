package gormImpl

import (
	"go_web_boilerplate/internal/entity"
	"go_web_boilerplate/internal/infra/db/postgres"
	"go_web_boilerplate/internal/pkg/auth"
	"go_web_boilerplate/internal/pkg/logger"
)

type UserRepositoryImpl struct {
	db     postgres.DB
	logger logger.Logger
	auth   auth.JWTAuth
}

func NewUserRepository(db postgres.DB, jwtAuth auth.JWTAuth, logger logger.Logger) UserRepositoryImpl {
	return UserRepositoryImpl{
		db:     db,
		logger: logger,
		auth:   jwtAuth,
	}
}

func (u UserRepositoryImpl) GetAllUser() ([]entity.User, error) {

	var users []entity.User

	result := u.db.Db.Find(&users)
	if result.Error != nil {
		u.logger.Error("GetAllUser Error -> ", result.Error.Error())
		return nil, result.Error
	}

	return users, nil

}

func (u UserRepositoryImpl) GenerateToken(user string, _ string) (*auth.JwtResponse, error) {
	return u.auth.GenerateJWTToken(auth.JWTAccessTokenPayload{
		Id:   1,
		Name: user,
		Role: []string{"admin"},
	})
}
