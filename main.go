package main

import (
	"git.gastrodon.io/imonke/monkebase"
	"git.gastrodon.io/imonke/monkelib"
	"github.com/gastrodon/groudon"

	"log"
	"net/http"
	"os"
)

func main() {
	monkebase.Connect(os.Getenv("DATABASE_CONNECTION"))
	groudon.RegisterHandler("GET", "^/nick/"+monkelib.NICK_PATTERN+"/?$", checkNick)
	groudon.RegisterHandler("GET", "^/email/"+monkelib.EMAIL_PATTERN+"/?$", checkEmail)
	http.Handle("/", http.HandlerFunc(groudon.Route))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
