package middleware

import (
	"net/http"

	"github.com/Flyewzz/group_preparation/auth"
	"github.com/Flyewzz/group_preparation/errs"
	"github.com/spf13/viper"
)

// Define our struct
type AuthenticationMiddleware struct {
	excludedRoutes map[string][]string
}

func NewAuthenticationMiddleware() *AuthenticationMiddleware {
	return &AuthenticationMiddleware{}
}

func (amw *AuthenticationMiddleware) SetUpExcludedRoutes() {
	amw.excludedRoutes = make(map[string][]string)
	exRoutes := amw.excludedRoutes
	exRoutes["/signup"] = []string{
		"POST",
	}
	exRoutes["/signin"] = []string{
		"POST",
	}
}

// Middleware function, which will be called for each request
func (amw *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Check for public (a token is unnecessary) routes
		path := r.URL.Path
		if methods, ok := amw.excludedRoutes[path]; ok {
			for _, method := range methods {
				if r.Method == method {
					next.ServeHTTP(w, r)
					return
				}
			}
		}
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tokenString := c.Value
		_, err = auth.DecodeToken(tokenString, viper.GetString("auth.secret_key"))
		if err != nil {
			if err == errs.TokenIsNotValid {
				http.Error(w, "Forbidden", http.StatusForbidden)
			} else {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
			return
		}
		// Finally
		// user's email given in the token
		// We found the token
		next.ServeHTTP(w, r)
	})
}
