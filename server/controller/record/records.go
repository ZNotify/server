package record

import (
	"errors"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"notify-api/ent/dao"
	"notify-api/server/types"
)

// Records godoc
//
//	@Summary		Get record
//	@Description	Get records
//	@Param			user_secret	path	string	true	"Secret of user"
//	@Param			skip		query	int		false	"The number of records to skip"
//	@Param			limit		query	int		false	"The number of records to return"
//	@Produce		json
//	@Success		200	{object}	types.Response[[]types.Message]
//	@Failure		401	{object}	types.UnauthorizedResponse
//	@Router			/{user_id}/record [get]
func Records(context *types.Ctx) {
	skip, err := strconv.Atoi(context.DefaultQuery("skip", "0"))
	if err != nil {
		zap.S().Infof("can not parse skip %s to int", context.DefaultQuery("skip", "0"))
		context.JSONError(http.StatusBadRequest, err)
	}
	limit, err := strconv.Atoi(context.DefaultQuery("limit", "20"))
	if err != nil {
		zap.S().Infof("can not parse limit %s to int", context.DefaultQuery("limit", "20"))
		context.JSONError(http.StatusBadRequest, err)
	}

	messages, ok := dao.Message.GetUserMessagesPaginated(context, context.User, skip, limit)
	if !ok {
		context.JSONError(http.StatusInternalServerError, errors.New("can not get messages"))
	}

	ret := make([]types.Message, len(messages))
	for i, message := range messages {
		ret[i] = types.FromModelMessage(*message)
	}

	context.JSONResult(messages)
}
