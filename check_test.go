package main

import (
	"github.com/brane-app/database-library"
	"github.com/brane-app/types-library"

	"fmt"
	"net/http"
	"testing"
)

const (
	nick  = "foobar"
	email = "foo@bar.com"
)

var (
	user types.User
)

func TestMain(main *testing.M) {
	user = types.NewUser(nick, "", email)
	database.WriteUser(user.Map())
}

func testCheck(test *testing.T, key, query string, expected bool) {
	var request *http.Request
	var err error
	if request, err = http.NewRequest("GET", fmt.Sprintf("https://imonke.co/%s/%s", key, query), nil); err != nil {
		test.Fatal(err)
	}

	var code int
	var r_map map[string]interface{}
	switch key {
	case "nick":
		code, r_map, err = checkNick(request)
	case "email":
		code, r_map, err = checkEmail(request)
	}

	if err != nil {
		test.Fatal(err)
	}

	if code != 200 {
		test.Errorf("Got code %d", code)
	}

	var exists, ok bool
	if exists, ok = r_map["exists"].(bool); !ok || exists != expected {
		test.Errorf("exists not %t in %#v!", expected, r_map)
	}
}

func Test_checkThings(test *testing.T) {
	var sets []map[string]interface{} = []map[string]interface{}{
		map[string]interface{}{
			"key":      "email",
			"value":    email,
			"expected": true,
		},
		map[string]interface{}{
			"key":      "nick",
			"value":    nick,
			"expected": true,
		},
		map[string]interface{}{
			"key":      "email",
			"value":    "email",
			"expected": false,
		},
		map[string]interface{}{
			"key":      "nick",
			"value":    "nick",
			"expected": false,
		},
	}

	var value map[string]interface{}
	for _, value = range sets {
		testCheck(
			test,
			value["key"].(string),
			value["value"].(string),
			value["expected"].(bool),
		)
	}
}

func Test_exists_badkey(test *testing.T) {
	var request *http.Request
	var err error
	if request, err = http.NewRequest("GET", "https://imonke.co/foo/bar", nil); err != nil {
		test.Fatal(err)
	}

	if _, _, err = exists("foo", request); err == nil {
		test.Errorf("Expected err but got none!")
	}
}
