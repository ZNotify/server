package controller

import (
	"net/http"

	"notify-api/db/model"
	"notify-api/db/util"
	"notify-api/server/types"
)

// Record godoc
//
//	@Summary		Get record
//	@Description	Get recent 30days message record of user
//	@Param			user_id	path	string	true	"user_id"
//	@Produce		json
//	@Success		200	{object}	types.Response[[]types.Message]
//	@Failure		401	{object}	types.UnauthorizedResponse
//	@Router			/{user_id}/record [get]
func Record(context *types.Ctx) {
	var messages []model.Message
	messages, err := util.MessageUtil.GetMessageInMonth(context.UserID)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	var ret []types.Message
	for _, message := range messages {
		ret = append(ret, types.FromModelMessage(message))
	}

	context.JSONResult(messages)
}
