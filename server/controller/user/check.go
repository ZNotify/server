package user

import (
	"errors"
	"net/http"

	"notify-api/ent/dao"
	"notify-api/server/types"
)

// Check godoc
//
//	@Summary	Check if the user secret is valid
//	@Id			checkUserSecret
//	@Tags		User
//	@Produce	json
//	@Param		user_secret	query		string	true	"Secret of user"
//	@Success	200			{object}	types.Response[bool]
//	@Router		/check [get]
func Check(context *types.Ctx) {
	userSecret := context.Query("user_secret")
	if userSecret == "" {
		context.JSONError(http.StatusBadRequest, errors.New("user_secret is empty"))
		return
	}
	_, ok := dao.User.GetUserBySecret(context, userSecret)
	context.JSONResult(ok)
}
