package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

func AddCookie(w http.ResponseWriter, name string, domain string) string {

	var expiration = time.Now().Add(365 * 24 * time.Hour)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: name, Value: state, Expires: expiration, Domain: domain}
	http.SetCookie(w, &cookie)

	return state
}
