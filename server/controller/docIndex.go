package controller

import (
	"net/http"

	"notify-api/server/types"
)

// DocRedirect godoc
//
//	@Summary	Redirect to docs
//	@Produce	plain
//	@Success	301	{string}	string	""
//	@Router		/docs [get]
func DocRedirect(ctx *types.Ctx) {
	ctx.Redirect(http.StatusMovedPermanently, "/docs/index.html")
}
