package api

import (
	v1 "github.com/ZhansultanS/myLMS/backend/internal/api/v1"
	"github.com/ZhansultanS/myLMS/backend/internal/service"
	"github.com/ZhansultanS/myLMS/backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h *Handler) Init(mode string) *gin.Engine {
	gin.SetMode(mode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.New(h.logger, h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
