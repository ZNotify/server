package controller

import (
	"net/http"

	"notify-api/db/entity"
	"notify-api/db/model"
	"notify-api/serve/types"
)

// Record godoc
//
//	@Summary		Get record
//	@Description	Get recent 30days message record of user
//	@Param			user_id	path	string	true	"user_id"
//	@Produce		json
//	@Success		200	{object}	types.Response[[]entity.Message]
//	@Failure		401	{object}	types.UnauthorizedResponse
//	@Router			/{user_id}/record [get]
func Record(context *types.Ctx) {
	var messages []entity.Message
	messages, err := model.MessageUtils.GetMessageInMonth(context.UserID)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSONResult(messages)
}
