package utils

import (
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, username string) {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:    "session_token",
		Value:   username,
		Expires: expiration,
		Path:    "/*",
	}
	http.SetCookie(w, &cookie)
}
