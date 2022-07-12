package Grpc

import (
	"go_web_boilerplate/internal/transport/rpc/grpc/genproto/user/pb"
	"go_web_boilerplate/internal/transport/rpc/grpc/handler"
)

func (g GrpcApp) RegisterHandlers() {
	userHandler := handler.NewUserService(g.userService, g.Logger)
	pb.RegisterUserServer(g.grpcServer, userHandler)
}
