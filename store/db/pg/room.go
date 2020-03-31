package pg

import "database/sql"

type RoomControllerPg struct {
	db *sql.DB
}

func NewRoomControllerPg(db *sql.DB) *RoomControllerPg {
	return &RoomControllerPg{
		db: db,
	}
}

// GetById(id int) (*models.Room, error)
// 	Add(name string, subjectId, authorId int) (int, error)
// 	Ban(userId, roomId int, status bool) error
// 	Join(userId, roomId int) error
// 	GetAll() ([]models.Room, error)

func (rc *RoomControllerPg) Add(name string, subjectId, authorId int) (int, string, error) {
	tx, err := rc.db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}
	var roomId int
	var uuid string
	err = tx.QueryRow("INSERT INTO rooms (name, subject_id, author_id) "+
		"VALUES ($1, $2, $3) RETURNING room_id, uuid",
		name, subjectId, authorId).Scan(&roomId, &uuid)
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}
	_, err = tx.Exec("INSERT INTO roomaccess (user_id, room_id) "+
		"($1, $2)", authorId, subjectId)
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}
	tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}
	return roomId, uuid, err
}
