package user

import (
	"github.com/Flyewzz/group_preparation/models",
	"github.com/gorilla/websocket",
)

type UserClient struct {
	User *models.User
	Conn *websocket.Conn
}

func NewUserClient(u *models.User, ws *websocket.Conn) *UserClient {
	return &UserClient{
		User: u,
		Conn: ws,
	}
}