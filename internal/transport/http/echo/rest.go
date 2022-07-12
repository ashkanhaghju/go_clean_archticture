package echo

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_web_boilerplate/internal/pkg/auth"
	"go_web_boilerplate/internal/pkg/logger"
	"go_web_boilerplate/internal/service"
	"go_web_boilerplate/internal/transport/http"
	"go_web_boilerplate/internal/transport/http/echo/handler"
	"time"
)

type rest struct {
	echo    *echo.Echo
	handler *handler.Handler
	auth    auth.JWTAuth
}

type CustomValidator struct {
	validator *validator.Validate
}

func New(userService service.User, auth auth.JWTAuth, logger logger.Logger) http.Rest {
	echoService := echo.New()
	echoService.Validator = &CustomValidator{validator: validator.New()}
	return &rest{
		echo: echoService,
		auth: auth,
		handler: &handler.Handler{
			Logger: logger,
			User:   userService,
		}}
}

func (r *rest) Start(address string) error {
	r.echo.Use(middleware.Recover())

	r.routing(r.auth)
	return r.echo.Start(address)
}

func (r *rest) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // use config for time
	defer cancel()

	return r.echo.Shutdown(ctx)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}
