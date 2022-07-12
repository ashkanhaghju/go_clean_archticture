package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go_web_boilerplate/internal/pkg/auth"
	"net/http"
	"strings"
)

func GetJwtMiddleWare(jwtAuth auth.JWTAuth) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed jwt")
			}
			token := strings.Split(authorization, " ")[1]
			tokenData, err := jwtAuth.ParseToken(token)
			if err != nil {
				fmt.Println(err)
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed jwt")
			}
			c.Set("user", tokenData)

			return next(c)
		}
	}

}
