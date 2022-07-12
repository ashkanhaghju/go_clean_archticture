package handler

import (
	"go_web_boilerplate/internal/pkg/logger"
	"go_web_boilerplate/internal/service"
)

type Handler struct {
	Logger logger.Logger
	User   service.User
}
