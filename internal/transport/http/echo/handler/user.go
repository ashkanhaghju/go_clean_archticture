package handler

import (
	"github.com/labstack/echo/v4"
	"go_web_boilerplate/internal/transport/http/response"
	"net/http"
)

func (h *Handler) GetAllUserList(c echo.Context) error {
	// start span with context
	users, err := h.User.GetAllUsers()
	if err != nil {
		h.Logger.Error("GetUserList Error ->", err)
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest})
	}

	var result []response.User
	for i := 0; i < len(users); i++ {
		result = append(result, response.User{
			ID:   users[i].ID,
			Name: users[i].Name,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{Message: "ok", Status: http.StatusOK, Result: result})
}

func (h *Handler) Login(c echo.Context) error {
	tokenResult, err := h.User.GenerateToken("", "")
	if err != nil {
		h.Logger.Error("GetUserList Error ->", err)
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest})
	}

	result := response.JwtResponse{
		AccessToken:  tokenResult.AccessToken,
		RefreshToken: tokenResult.RefreshToken,
	}

	return c.JSON(http.StatusOK, response.BaseResponse{Message: "ok", Status: http.StatusOK, Result: result})
}
