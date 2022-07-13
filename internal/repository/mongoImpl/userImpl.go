package mongoImpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go_web_boilerplate/internal/entity"
	"go_web_boilerplate/internal/infra/db/mongodb"
	"go_web_boilerplate/internal/pkg/auth"
	"go_web_boilerplate/internal/pkg/logger"
	"time"
)

type UserRepositoryImpl struct {
	dbClient mongodb.DBClient
	logger   logger.Logger
	auth     auth.JWTAuth
}

func NewUserRepository(dbClient mongodb.DBClient, jwtAuth auth.JWTAuth, logger logger.Logger) UserRepositoryImpl {
	return UserRepositoryImpl{
		dbClient: dbClient,
		logger:   logger,
		auth:     jwtAuth,
	}
}

func (u UserRepositoryImpl) GetAllUser() ([]entity.User, error) {

	var users []entity.User
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := u.dbClient.Db.Collection("users").Find(ctx, bson.D{})
	if err != nil {
		u.logger.Error("GetAllUser Error -> ", err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result []entity.User
		err := cur.All(context.Background(), &result)
		if err != nil {
			u.logger.Error("GetAllUser Decode -> ", err)
		}
		users = append(users, result[0], result[1])
		//users = append(users, result)
	}
	if err := cur.Err(); err != nil {
		u.logger.Error("GetAllUser cursor -> ", err)
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
