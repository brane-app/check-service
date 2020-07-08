package main

import (
	"github.com/gastrodon/groudon"
	"github.com/imonke/monkebase"

	"log"
	"net/http"
	"os"
)

func main() {
	monkebase.Connect(os.Getenv("MONKEBASE_CONNECTION"))
	groudon.RegisterHandler("GET", "^/nick/.+/?$", checkNick)
	groudon.RegisterHandler("GET", "^/email/.+/?$", checkEmail)
	http.Handle("/", http.HandlerFunc(groudon.Route))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
