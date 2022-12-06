package controller

import (
	"notify-api/serve/types"
	"notify-api/web"
)

// WebIndex godoc
//
//	@Summary		Web Index
//	@Description	Provide UI
//	@Produce		html
//	@Success		200	{string}	string	"html"
//	@Router			/ [get]
func WebIndex(ctx *types.Ctx) {
	ctx.FileFromFS("/", web.StaticHttpFS)
	// hardcode index.html, use this as a trick to get html file
	// https://github.com/golang/go/blob/a7e16abb22f1b249d2691b32a5d20206282898f2/src/net/http/fs.go#L587
}
