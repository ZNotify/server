package controller

import (
	"net/http"

	"notify-api/server/types"
)

// Alive godoc
//
//	@Summary		Server Heartbeat
//	@Description	If the server is alive
//	@Produce		plain
//	@Success		204	{string}	string	""
//	@Router			/alive [get]
func Alive(context *types.Ctx) {
	context.String(http.StatusNoContent, "")
}
