package user

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/google/go-github/v49/github"
	"go.uber.org/zap"

	"notify-api/ent/dao"
	"notify-api/server/types"
	"notify-api/setup/config"
	"notify-api/setup/oauth"
	"notify-api/utils"
)

// GitHub
//
//	@Summary	OAuth callback for GitHub, redirect to ui with user_secret
//	@Id			user.github
//	@Success	307
//	@Router		/login/github [get]
//	@Param		state	query		string	false	"should always be 'no_need_to_set_state'"
//	@Param		code	query		string	true	"access code"
//	@Failure	400		{object}	types.BadRequestResponse
//	@Failure	401		{object}	types.UnauthorizedResponse
func GitHub(context *types.Ctx) {
	state := context.Query("state")
	if state != "no_need_to_set_state" {
		zap.S().Warnf("state is not 'no_need_to_set_state': %s", state)
		context.JSONError(http.StatusBadRequest, errors.New("invalid state, should be 'no_need_to_set_state'"))
		return
	}

	code := context.Query("code")
	if code == "" {
		zap.S().Info("code is empty")
		context.JSONError(http.StatusBadRequest, errors.New("code is empty"))
		return
	}

	token, err := oauth.OAuthConf.Exchange(context, code)
	if err != nil {
		zap.S().Errorf("Failed to exchange code: %v", err)
		context.JSONError(http.StatusUnauthorized, err)
		return
	}

	oauthClient := oauth.OAuthConf.Client(context, token)
	client := github.NewClient(oauthClient)
	githubUser, _, err := client.Users.Get(context, "")
	if err != nil {
		zap.S().Errorf("Failed to get githubUser: %v", err)
		context.JSONError(http.StatusInternalServerError, err)
		return
	}

	user, ok := dao.User.EnsureUser(context, githubUser.GetID(), githubUser.GetName(), githubUser.GetLogin(), utils.OAuthTokenSerialize(token))
	if !ok {
		context.JSONError(http.StatusInternalServerError, errors.New("failed to ensure user"))
		return
	}

	redirectURL, err := url.Parse(config.Config.Server.URL)
	if err != nil {
		zap.S().Errorf("Failed to parse redirect url: %v", err)
		context.JSONError(http.StatusInternalServerError, err)
		return
	}
	redirectURL.Path = "/"
	q := redirectURL.Query()
	q.Set("user_secret", user.Secret)
	redirectURL.RawQuery = q.Encode()

	context.Redirect(http.StatusTemporaryRedirect, redirectURL.String())
}