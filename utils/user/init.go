package user

import "notify-api/utils/config"

func Init() {
	users = config.Config.Users
}
