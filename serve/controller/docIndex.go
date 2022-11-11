package controller

import (
	"net/http"

	"notify-api/serve/types"
)

// DocIndex godoc
//
//	@Summary  Redirect to docs
//	@Produce  plain
//	@Success  301  {string}  string  ""
//	@Router   /docs [get]
func DocIndex(ctx *types.Ctx) {
	ctx.Redirect(http.StatusMovedPermanently, "/docs/index.html")
}
