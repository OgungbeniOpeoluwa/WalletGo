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

func SendUnprocessableEntity(c *gin.Context, req any) {
	data := map[string]any{
		"message": "Unable to process request",
		"Data":    []any{req},
	}
	c.JSON(http.StatusUnprocessableEntity, data)
}
func SendSuccess(c *gin.Context, req any) {
	data := map[string]any{
		"message": "success",
		"Data":    []any{req},
	}
	c.JSON(http.StatusOK, data)
}
