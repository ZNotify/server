package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/web"
)

var UI = http.FS(web.UI)

func Index(ctx *gin.Context) {
	ctx.FileFromFS("/", http.FS(web.UI))
	// hardcode index.html, use this as a trick to get html file
	// https://github.com/golang/go/blob/a7e16abb22f1b249d2691b32a5d20206282898f2/src/net/http/fs.go#L587
}
