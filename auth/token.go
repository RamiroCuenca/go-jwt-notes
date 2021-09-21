package auth

import (
	"time"

	"github.com/RamiroCuenca/go-rest-notesApi/users/models"
	"github.com/golang-jwt/jwt/v4"
)

// Generates a JWT. It receives the data from the user who has logged in!
// The JWT is a string
func GenerateToken(data models.User) (string, error) {
	claim := models.Claim{
		Username: data.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Ramiro Cuenca",
		},
	}

	// Prepared the token to be signed with RS256 method and including the claim
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)

	// Sign the token with our private key
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	// Return the generated token
	return signedToken, nil
}
