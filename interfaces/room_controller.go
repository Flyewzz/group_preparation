package interfaces

import "github.com/Flyewzz/group_preparation/models"

type RoomController interface {
	GetById(id int) (*models.Room, error)
	Add(name string, subjectId, authorId int) (int, error)
	Ban(userId, roomId int, status bool) error
	Join(userId, roomId int) error
	GetAll() ([]models.Room, error)
}
