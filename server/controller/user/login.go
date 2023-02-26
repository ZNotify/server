package user

import (
	"net/http"

	"golang.org/x/oauth2"

	"notify-api/server/types"
	"notify-api/setup/oauth"
)

// Login
//
//	@Summary	Login with GitHub
//	@Id			userLogin
//	@Tags		User
//	@Success	307
//	@Router		/login [get]
func Login(context *types.Ctx) {
	url := oauth.Conf.AuthCodeURL("no_need_to_set_state", oauth2.AccessTypeOnline)
	context.Redirect(http.StatusTemporaryRedirect, url)
}
