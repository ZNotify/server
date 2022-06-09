package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Alive(context *gin.Context) {
	context.String(http.StatusNoContent, "")
}
