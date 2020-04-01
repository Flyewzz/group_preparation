package interfaces

import (
	"github.com/Flyewzz/group_preparation/room"
)

type RoomController interface {
	GetById(id int) (*room.RoomData, error)
	Add(name string, subjectId, typeId, authorId int) (int, string, error)
	Ban(userId, roomId int, status bool) error
	Join(userId int, uuid string) error
	GetAll(userId int) ([]room.RoomData, error)
	GetAuthorId(roomId int) (int, error)
	IsBanned(userId, roomId int) (bool, error)
}
