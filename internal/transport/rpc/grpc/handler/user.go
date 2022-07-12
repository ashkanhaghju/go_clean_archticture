package handler

import (
	"context"
	"go_web_boilerplate/internal/pkg/logger"
	"go_web_boilerplate/internal/service"
	"go_web_boilerplate/internal/transport/rpc/grpc/genproto/user/pb"
)

type UserService struct {
	Logger      logger.Logger
	UserService service.User
}

func NewUserService(us service.User, logger logger.Logger) UserService {
	return UserService{
		UserService: us,
		Logger:      logger,
	}
}

func (u UserService) GetAllUsers(_ context.Context, _ *pb.Empty) (*pb.UsersResponse, error) {
	u.Logger.Info("GetAllUsers called")
	users, err := u.UserService.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := pb.UsersResponse{}
	var usersList []*pb.UserResponse
	for _, user := range users {
		usersList = append(usersList, &pb.UserResponse{
			Id:   user.ID,
			Name: user.Name,
		})
	}
	response.Users = usersList

	return &response, nil

}
