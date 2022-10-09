package main

import (
	"github.com/brane-app/librane/tools"
	"github.com/gastrodon/groudon/v2"

	"os"
)

var (
	prefix = os.Getenv("PATH_PREFIX")

	checkNickRoute  = "^" + prefix + "/nick/" + tools.NICK_PATTERN + "/?$"
	checkEmailRoute = "^" + prefix + "/email/" + tools.EMAIL_PATTERN + "/?$"
)

func register_handlers() {
	groudon.AddHandler("GET", checkNickRoute, checkNick)
	groudon.AddHandler("GET", checkEmailRoute, checkEmail)
}
