package auth

// Create a struct to read the username and password from the request body
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
