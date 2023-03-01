package user

import (
	"net/http"

	"notify-api/app/api/common"
	"notify-api/app/global"

	"golang.org/x/oauth2"
)

// Login
//
//	@Summary  Login with GitHub
//	@Id       userLogin
//	@Tags     User
//	@Success  307
//	@Router   /login [get]
func Login(context *common.Context) {
	url := global.App.OAuth.AuthCodeURL("no_need_to_set_state", oauth2.AccessTypeOnline)
	context.Redirect(http.StatusTemporaryRedirect, url)
}
