package auth

import "github.com/dgrijalva/jwt-go"

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
