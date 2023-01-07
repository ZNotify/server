package user

import (
	"net/http"

	"golang.org/x/oauth2"

	"notify-api/server/setup/oauth"
	"notify-api/server/types"
)

// Login
//
//	@Summary	Login with GitHub
//	@Success	307
//	@Router		/login [get]
func Login(context *types.Ctx) {
	url := oauth.OAuthConf.AuthCodeURL("no_need_to_set_state", oauth2.AccessTypeOnline)
	context.Redirect(http.StatusTemporaryRedirect, url)
}
