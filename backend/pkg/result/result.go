package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Message string `json:"message"`
	Date    any    `json:"data"`
}

func Success(message string, data any) *Result {
	return &Result{Message: message, Date: data}
}

func Fail(Err error) *Result {
	return &Result{Message: Err.Error()}
}

func SendResult(c *gin.Context, r *Result) {
	c.JSON(http.StatusOK, r)
}
