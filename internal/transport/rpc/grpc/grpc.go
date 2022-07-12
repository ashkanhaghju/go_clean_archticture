package Grpc

import (
	"go_web_boilerplate/internal/pkg/logger"
	"go_web_boilerplate/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcApp struct {
	grpcServer  *grpc.Server
	userService service.User
	Logger      logger.Logger
}

func NewGrpc(options grpc.ServerOption, userService service.User, logger logger.Logger) GrpcApp {
	return GrpcApp{
		grpcServer:  grpc.NewServer(options),
		userService: userService,
		Logger:      logger,
	}
}

func (g GrpcApp) Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	g.RegisterHandlers()
	err = g.grpcServer.Serve(listener)
	return err
}

func (g GrpcApp) Shutdown() error {
	g.grpcServer.Stop()
	return nil
}
