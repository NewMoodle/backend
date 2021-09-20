package v1

import (
	"github.com/ZhansultanS/myLMS/backend/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initUserRoutes(v1 *gin.RouterGroup) {
	roles := v1.Group("/users")
	{
		roles.GET("/", h.usersGetAll)
		roles.POST("/", h.usersCreateOne)
	}
}

func (h *Handler) usersCreateOne(c *gin.Context) {
	var userDto dto.UserCreateDto
	if err := c.BindJSON(&userDto); err != nil {
		h.errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Users.Create(c, userDto); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) usersGetAll(c *gin.Context) {
	users, err := h.services.Users.FindAll(c)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	h.successResponse(c, http.StatusOK, map[string]interface{}{"users": users})
}
