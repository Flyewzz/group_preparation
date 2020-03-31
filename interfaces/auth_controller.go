package interfaces

import (
	"github.com/Flyewzz/group_preparation/models"
)

type AuthController interface {
	// Returns the created user's id and error
	SignUp(email, password string) (int, error)
	GetUser(email string) (*models.User, error)
}
