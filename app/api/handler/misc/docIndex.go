package misc

import (
	"net/http"

	"github.com/ZNotify/server/app/api/common"
)

// DocRedirect godoc
//
//	@Summary  Redirect to docs
//	@Id       docRedirect
//	@Tags     UI
//	@Produce  plain
//	@Success  301  {string}  string  ""
//	@Router   /docs [get]
func DocRedirect(ctx *common.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "/docs/index.html")
}
