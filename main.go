package main

import (
	"github.com/gastrodon/groudon"
	"github.com/imonke/monkebase"
	"github.com/imonke/monkelib"

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
