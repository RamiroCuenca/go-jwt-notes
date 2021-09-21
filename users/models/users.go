package models

import "github.com/golang-jwt/jwt/v4"

// All users must have a username and a password.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claim is the information that will be sent through the JWT
// In this case we are going to add only the username. The
// other parameters are going to be automatically added by
// the library that we are going to use.
type Claim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
