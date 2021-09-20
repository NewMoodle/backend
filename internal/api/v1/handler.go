package v1

import (
	"github.com/ZhansultanS/myLMS/backend/internal/service"
	"github.com/ZhansultanS/myLMS/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger   logger.Logger
	services service.Service
}

func New(logger logger.Logger, services service.Service) *Handler {
	return &Handler{
		logger:   logger,
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initRoleRoutes(v1)
		h.initUserRoutes(v1)
	}
}
