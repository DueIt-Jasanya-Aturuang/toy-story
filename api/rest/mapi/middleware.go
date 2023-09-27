package mapi

import (
	"net/http"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/api/rest/helper"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("Username")
		password := r.Header.Get("Password")

		if username != "admin" || password != "password" {
			helper.ErrorResponseEncode(w, _error.HttpErrString("authorization", response.CM04))
			return
		}

		next.ServeHTTP(w, r)
	})
}
