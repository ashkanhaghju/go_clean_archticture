package app

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
	"go_web_boilerplate/internal/config"
	"go_web_boilerplate/internal/infra/db/postgres"
	app "go_web_boilerplate/internal/middleware"
	"go_web_boilerplate/internal/pkg/auth"
	"go_web_boilerplate/internal/pkg/logger/zap"
	"go_web_boilerplate/internal/repository/gormImpl"
	"go_web_boilerplate/internal/service"
	Grpc "go_web_boilerplate/internal/transport/rpc/grpc"
	"go_web_boilerplate/internal/usecase"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ServeGrpcCMD = &cli.Command{
	Name:    "serve-grpc",
	Aliases: []string{"serve-grpc"},
	Usage:   "serve grpc",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "config , c",
			Usage: "config file path",
		},
		&cli.StringFlag{
			Name:  "log , l",
			Usage: "log file path",
			Value: "/logs/app.log",
		},
	},
	Action: serveGrpc,
}

func serveGrpc(c *cli.Context) error {
	configFilePath := c.String("config")
	logFilePath := c.String("log")
	cfg := new(config.Config)
	err := config.ReadYAML(configFilePath, cfg)
	err = config.ReadEnv(cfg)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}

	logger := zap.New(f, zapcore.InfoLevel)
	jwtAuth := auth.NewJWTAuth("G-KaPdSgVkYp3s6v9y$B&E)H+MbQeThW", "testProject", "ashkan", 2000*time.Minute, 70000*time.Hour)

	pgRepoImpl, err := postgres.NewGorm(cfg.Postgres, logger)
	if err != nil {
		return err
	}
	// define repose
	userRepo := gormImpl.NewUserRepository(*pgRepoImpl, jwtAuth, logger)

	// define userCases
	userUserCase := usecase.NewUserUseCase(userRepo, logger)
	tokenUserCase := usecase.NewTokenUseCase(userRepo, logger)
	userService := service.NewUserService(userUserCase, tokenUserCase, logger)

	grpcOpts := app.GrpcInterceptor()
	gServer := Grpc.NewGrpc(grpcOpts, userService, logger)

	go func() {
		err = gServer.Start(cfg.App.Grpc.Address)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	if err := gServer.Shutdown(); err != nil {
		fmt.Println("\nRest server doesn't shutdown in 10 seconds")
	}

	return nil
}
