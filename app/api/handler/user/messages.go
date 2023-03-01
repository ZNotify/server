package user

import (
	"errors"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"notify-api/app/api/common"
	"notify-api/app/api/models"
	"notify-api/app/db/dao"
)

// Messages godoc
//
//	@Summary      Get messages of user
//	@Id           getMessagesByUserSecret
//	@Tags         User
//	@Description  Get messages of user with pagination
//	@Param        user_secret  path   string  true   "Secret of user"
//	@Param        skip         query  int     false  "The number of records to skip"    default(0)   minimum(0)
//	@Param        limit        query  int     false  "The number of records to return"  default(20)  maximum(100)  minimum(0)
//	@Produce      json
//	@Success      200  {object}  common.Response[[]models.Message]
//	@Failure      401  {object}  common.UnauthorizedResponse
//	@Router       /{user_secret}/messages [get]
func Messages(context *common.Context) {
	skip, err := strconv.Atoi(context.DefaultQuery("skip", "0"))
	if err != nil {
		zap.S().Infof("can not parse skip %s to int", context.DefaultQuery("skip", "0"))
		context.JSONError(http.StatusBadRequest, err)
	}
	if skip < 0 {
		zap.S().Infof("skip can not be negative")
		context.JSONError(http.StatusBadRequest, errors.New("skip can not be negative"))
	}

	limit, err := strconv.Atoi(context.DefaultQuery("limit", "20"))
	if err != nil {
		zap.S().Infof("can not parse limit %s to int", context.DefaultQuery("limit", "20"))
		context.JSONError(http.StatusBadRequest, err)
	}
	if limit > 100 {
		zap.S().Infof("limit %d is too large", limit)
		context.JSONError(http.StatusBadRequest, errors.New("limit is too large, the maximum is 100"))
	}
	if limit < 0 {
		zap.S().Infof("limit %d can not be negative", limit)
		context.JSONError(http.StatusBadRequest, errors.New("limit is too small, the minimum is 0"))
	}

	messages, ok := dao.Message.GetUserMessagesPaginated(context, context.User, skip, limit)
	if !ok {
		context.JSONError(http.StatusInternalServerError, errors.New("can not get messages"))
	}

	ret := make([]models.Message, len(messages))
	for i, message := range messages {
		ret[i] = models.FromModelMessage(*message)
	}

	context.JSONResult(messages)
}
