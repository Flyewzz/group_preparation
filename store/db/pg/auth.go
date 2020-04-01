package pg

import (
	"database/sql"

	"github.com/Flyewzz/group_preparation/models"
)

type AuthControllerPg struct {
	db *sql.DB
}

func NewAuthControllerPg(db *sql.DB) *AuthControllerPg {
	return &AuthControllerPg{
		db: db,
	}
}

func (ac *AuthControllerPg) SignUp(email, password string) (int, error) {
	var userId int
	err := ac.db.QueryRow("INSERT INTO users (email, password) "+
		"VALUES ($1, $2) RETURNING user_id", email, password).Scan(&userId)
	return userId, err
}

func (ac *AuthControllerPg) GetUser(email string) (*models.User, error) {
	var user *models.User = &models.User{}
	err := ac.db.QueryRow("SELECT user_id, email, password FROM users "+
		"WHERE email = $1", email).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)
	return user, err
}
