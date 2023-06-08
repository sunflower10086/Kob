package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Message string `json:"message"`
	Date    any    `json:"date"`
}

func Success(data any) *Result {
	return &Result{Date: data}
}

func Fail(Err error) *Result {
	return &Result{Message: Err.Error()}
}

func SendResult(c *gin.Context, r *Result) {
	c.JSON(http.StatusOK, r)
}
