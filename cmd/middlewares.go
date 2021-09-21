package main

import (
	"net/http"

	"github.com/RamiroCuenca/go-jwt-notes/auth"
	"github.com/RamiroCuenca/go-jwt-notes/common/handler"
)

// I'm sure that there are some provided by the community

// It's receives and returns a handler
func AuthenticationMiddleware(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := auth.ValidateToken(token) // auth is the package we created
		// If token is invalid
		if err != nil {
			forbidden(w, r)
			return
		}

		f(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	handler.SendResponse(w, http.StatusForbidden, []byte("It hasn't got authorization"))
}
