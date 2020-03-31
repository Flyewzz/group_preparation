package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Flyewzz/group_preparation/auth"
	"github.com/spf13/viper"
)

func (hd *HandlerData) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var creds auth.Credentials
	// Get the JSON body and decode into credentials
	params := r.FormValue
	creds.Email = params("email")
	creds.Password = params("password")
	if creds.Email == "" || creds.Password == "" {
		// If the structure of the body is wrong, return an HTTP error
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	userId, err := hd.AuthController.SignUp(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("User with id %d was added\n", userId)))
}

func (hd *HandlerData) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var creds auth.Credentials
	// Get the JSON body and decode into credentials
	params := r.FormValue
	creds.Email = params("email")
	creds.Password = params("password")
	if creds.Email == "" || creds.Password == "" {
		// If the structure of the body is wrong, return an HTTP error
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	user, err := hd.AuthController.GetUser(creds.Email)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	// Check Authentication
	if creds.Password != user.Password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	expSeconds := viper.GetInt("auth.exp_time")
	expirationTime := time.Now().Add(time.Duration(expSeconds) * time.Second)
	token, err := auth.NewToken(
		&creds,
		expirationTime,
		viper.GetString("auth.secret_key"),
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
}
