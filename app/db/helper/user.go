package helper

import (
	"notify-api/app/db/ent/generate"
)

func GetReadableName(u *generate.User) string {
	if u.GithubName != "" {
		return u.GithubName
	}
	return u.GithubLogin
}
