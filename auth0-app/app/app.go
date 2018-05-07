package app

import (
	"github.com/gorilla/sessions"
	"encoding/gob"
)

var (
	Store *sessions.CookieStore
)

func Init() error {
	Store = sessions.NewCookieStore([]byte("something-very-secret"))
	gob.Register(map[string]interface{}{})
	return nil
}