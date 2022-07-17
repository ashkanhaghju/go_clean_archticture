package app

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
	"go_web_boilerplate/internal/config"
	"go_web_boilerplate/internal/infra/db/mongodb"
	"go_web_boilerplate/internal/pkg/auth"
	"go_web_boilerplate/internal/pkg/logger/zap"
	"go_web_boilerplate/internal/repository/mongoImpl"
	"go_web_boilerplate/internal/service"
	"go_web_boilerplate/internal/transport/http/echo"
	"go_web_boilerplate/internal/usecase"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ServeCMD = &cli.Command{
	Name:    "serve-http",
	Aliases: []string{"serve-http"},
	Usage:   "serve-http",
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
	Action: serve,
}

func serve(c *cli.Context) error {
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

	jwtAuth := auth.NewJWTAuth(cfg.Auth.SecretKey, cfg.Auth.Audience, cfg.Auth.Issuer,
		time.Duration(cfg.Auth.AccessTokenDuration)*time.Minute,
		time.Duration(cfg.Auth.RefreshTokenDuration)*time.Hour,
	)

	dbClient, err := mongodb.NewClient(cfg.Mongo, logger)
	//dbClient, err := postgres.NewGorm(cfg.Postgres, logger)

	if err != nil {
		return err
	}
	// define repose
	userRepo := mongoImpl.NewUserRepository(*dbClient, jwtAuth, logger)
	//userRepo := gormImpl.NewUserRepository(*dbClient, jwtAuth, logger)

	// define userCases
	userUseCase := usecase.NewUserUseCase(userRepo, logger)
	tokenUseCase := usecase.NewTokenUseCase(userRepo, logger)

	// define services
	userService := service.NewUserService(userUseCase, tokenUseCase, logger)
	//httpOpts := app.HttpInterceptor()

	restServer := echo.New(userService, jwtAuth, logger)

	go func() {
		if err := restServer.Start(cfg.App.Rest.Address); err != nil {
			logger.Error(fmt.Sprintf("error happen while serving: %v", err))
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	if err := restServer.Shutdown(); err != nil {
		fmt.Println("\nRest server doesn't shutdown in 10 seconds")
	}

	defer func() {
		if err = dbClient.Disconnect(); err != nil {
			fmt.Println("\nDB server doesn't shutdown")
		}
	}()

	return nil
}
