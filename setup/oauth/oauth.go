package oauth

import (
	"golang.org/x/oauth2"
	githubOAuth "golang.org/x/oauth2/github"

	"notify-api/setup/config"
)

var Conf *oauth2.Config

func Init() {
	Conf = &oauth2.Config{
		ClientID:     config.Config.User.SSO.GitHub.ClientID,
		ClientSecret: config.Config.User.SSO.GitHub.ClientSecret,
		Scopes:       make([]string, 0),
		Endpoint:     githubOAuth.Endpoint,
		RedirectURL:  config.Config.Server.URL + "/login/github",
	}
}
