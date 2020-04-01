package auth

import (
	"time"

	"github.com/Flyewzz/group_preparation/errs"
	"github.com/Flyewzz/group_preparation/models"
	"github.com/dgrijalva/jwt-go"
)

func NewToken(user *models.User, expirationTime time.Time,
	secretKey string) (string, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserId: user.Id,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}

func DecodeToken(strToken, secretKey string) (*models.User, error) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(strToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errs.TokenIsNotValid
	}
	u := &models.User{
		Id:    claims.UserId,
		Email: claims.Email,
	}
	return u, nil
}
