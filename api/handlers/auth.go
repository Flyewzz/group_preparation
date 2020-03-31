package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Flyewzz/group_preparation/auth"
)

func (hd *HandlerData) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var creds auth.Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
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
	w.Write([]byte("PASS"))
}
