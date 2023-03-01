package misc

import (
	"net/http"

	"notify-api/app/api/common"
)

// Alive godoc
//
//	@Summary      Server Heartbeat
//	@Id           alive
//	@Tags         Health
//	@Description  If the server is alive
//	@Produce      plain
//	@Success      204  {string}  string  ""
//	@Router       /alive [get]
func Alive(context *common.Context) {
	context.String(http.StatusNoContent, "")
}
