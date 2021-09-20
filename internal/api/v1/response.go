package v1

import "github.com/gin-gonic/gin"

type errorEnvelop struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

type successEnvelop struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

func (h *Handler) errorResponse(c *gin.Context, statusCode int, error string) {
	h.logger.Error(error)
	c.AbortWithStatusJSON(statusCode, errorEnvelop{
		StatusCode: statusCode,
		Error:      error,
	})
}

func (h *Handler) successResponse(c *gin.Context, statusCode int, message interface{}) {
	c.JSON(statusCode, successEnvelop{
		StatusCode: statusCode,
		Data:       message,
	})
}
