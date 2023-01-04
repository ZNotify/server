package user

import (
	"notify-api/utils/user"

	"notify-api/server/types"
)

// Check godoc
//
//	@Summary	Check if the user_id is valid
//	@Produce	json
//	@Param		user_id	query		string	true	"user_id"
//	@Success	200		{object}	types.Response[bool]
//	@Router		/check [get]
func Check(context *types.Ctx) {
	userID := context.Query("user_id")
	result := user.Is(userID)
	context.JSONResult(result)
}
