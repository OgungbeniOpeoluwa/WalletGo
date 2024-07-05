package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendInternalSeverError(err error) map[string]any {
	return map[string]any{
		"status":  http.StatusInternalServerError,
		"message": err,
	}
}

func SendBadRequestError(c *gin.Context, req any) {
	data := map[string]any{
		"message": "Bad request",
		"Data":    []any{req},
	}
	c.JSON(http.StatusBadRequest, data)
}
