package echo

import (
	"go_web_boilerplate/internal/middleware"
	"go_web_boilerplate/internal/pkg/auth"
)

func (r *rest) routing(auth auth.JWTAuth) {
	apiV1Group := r.echo.Group("api/v1/")
	publicGroup := apiV1Group.Group("public/")

	publicGroup.POST("login", r.handler.Login)

	privateGroup := apiV1Group.Group("private/", middleware.GetJwtMiddleWare(auth))
	privateGroup.GET("user", r.handler.GetAllUserList)

	//r.echo.POST("api/v1/register", r.handler.register)
}
