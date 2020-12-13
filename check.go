package main

import (
	"git.gastrodon.io/imonke/monkebase"

	"fmt"
	"net/http"
	"strings"
)

func pathSplit(it rune) (ok bool) {
	ok = it == '/'
	return
}

func query(path string) (want string) {
	var parts []string = strings.FieldsFunc(path, pathSplit)
	want = parts[len(parts)-1]
	return
}

func exists(key string, request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var query string = query(request.URL.Path)

	var exists bool
	switch key {
	case "nick":
		_, exists, err = monkebase.ReadSingleUserNick(query)
	case "email":
		_, exists, err = monkebase.ReadSingleUserEmail(query)
	default:
		err = fmt.Errorf("I have no idea how I got key %s", key)
	}

	code = 200
	r_map = map[string]interface{}{"exists": exists}
	return
}

func checkNick(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	code, r_map, err = exists("nick", request)
	return
}

func checkEmail(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	code, r_map, err = exists("email", request)
	return
}
