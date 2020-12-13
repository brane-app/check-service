package main

import (
	"github.com/gastrodon/groudon"
	"git.gastrodon.io/imonke/monkebase"
	"git.gastrodon.io/imonke/monkelib"

	"log"
	"net/http"
	"os"
)

func main() {
	monkebase.Connect(os.Getenv("MONKEBASE_CONNECTION"))
	groudon.RegisterHandler("GET", "^/nick/"+monkelib.NICK_PATTERN+"/?$", checkNick)
	groudon.RegisterHandler("GET", "^/email/"+monkelib.EMAIL_PATTERN+"/?$", checkEmail)
	http.Handle("/", http.HandlerFunc(groudon.Route))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
