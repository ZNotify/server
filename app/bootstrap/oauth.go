package bootstrap

import (
	"net/url"
	"strings"

	"go.uber.org/zap"
	"golang.org/x/oauth2"
	githubOAuth "golang.org/x/oauth2/github"

	"github.com/ZNotify/server/app/global"
)

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func initializeOauth() {
	var redirectURL string
	if global.App.Config.Server.URL == "" {
		zap.S().Fatalf("Server URL is not set")
	}

	if strings.HasSuffix(global.App.Config.Server.URL, "/") {
		redirectURL = global.App.Config.Server.URL + "login/github"
	} else {
		redirectURL = global.App.Config.Server.URL + "/login/github"
	}

	if !isUrl(redirectURL) {
		zap.S().Fatalf("Server URL is not valid")
	}

	global.App.OAuth = &oauth2.Config{
		ClientID:     global.App.Config.User.SSO.GitHub.ClientID,
		ClientSecret: global.App.Config.User.SSO.GitHub.ClientSecret,
		Scopes:       make([]string, 0),
		Endpoint:     githubOAuth.Endpoint,
		RedirectURL:  redirectURL,
	}
}
