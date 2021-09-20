package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initRoleRoutes(v1 *gin.RouterGroup) {
	roles := v1.Group("/roles")
	{
		roles.GET("/", h.rolesGetAll)
	}
}

func (h *Handler) rolesGetAll(c *gin.Context) {
	roles, err := h.services.Roles.FindAll(c)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, http.StatusOK, map[string]interface{}{"roles": roles})
}
