package pg

import (
	"database/sql"

	"github.com/Flyewzz/group_preparation/room"
)

type RoomControllerPg struct {
	db *sql.DB
}

func NewRoomControllerPg(db *sql.DB) *RoomControllerPg {
	return &RoomControllerPg{
		db: db,
	}
}

func (rc *RoomControllerPg) Add(name string, subjectId, typeId, authorId int) (int, string, error) {
	tx, err := rc.db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}
	var roomId int
	var uuid string
	err = tx.QueryRow("INSERT INTO rooms (name, subject_id, type_id, author_id) "+
		"VALUES ($1, $2, $3, $4) RETURNING room_id, uuid",
		name, subjectId, typeId, authorId).Scan(&roomId, &uuid)
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}
	_, err = tx.Exec("INSERT INTO roomaccess (user_id, room_id) "+
		"VALUES ($1, $2)", authorId, roomId)
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

func (rc *RoomControllerPg) Ban(userId, roomId int, status bool) error {
	result, err := rc.db.Exec("UPDATE roomaccess SET banned = $1 "+
		"WHERE user_id = $2, room_id = $3", status, userId, roomId)
	if err != nil {
		return err
	}
	count, _ := result.RowsAffected()
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (rc *RoomControllerPg) Join(userId int, uuid string) error {
	var roomId int
	// Search the room by uuid
	err := rc.db.QueryRow("SELECT room_id FROM rooms "+
		"WHERE uuid = $1", uuid).Scan(&roomId)
	if err != nil {
		return err
	}
	// And after that join the user to this
	_, err = rc.db.Exec("INSERT INTO roomaccess (user_id, room_id) "+
		"VALUES ($1, $2)", roomId, userId)
	return nil
}

func (rc *RoomControllerPg) GetAll(userId int) ([]room.RoomData, error) {
	rows, err := rc.db.Query(
		"SELECT r.room_id, r.name, "+
			"r.uuid, wt.name, u.email, ra.banned FROM rooms r "+
			"INNER JOIN roomaccess ra ON ra.user_id = $1 "+
			"AND r.room_id = ra.room_id "+
			"INNER JOIN users u ON r.author_id = u.user_id "+
			"INNER JOIN worktypes wt ON r.type_id = wt.type_id ", userId)
	if err != nil {
		return nil, err
	}
	var rooms []room.RoomData
	for rows.Next() {
		var room room.RoomData
		err = rows.Scan(
			&room.RoomId,
			&room.Name,
			&room.UUID,
			&room.Type,
			&room.AuthorEmail,
			&room.Banned,
		)
		if err != nil {
			continue
		}
		rooms = append(rooms, room)
	}
	return rooms, err
}

func (rc *RoomControllerPg) GetById(id int) (*room.RoomData, error) {
	r := &room.RoomData{}
	err := rc.db.QueryRow("SELECT r.room_id, r.name, r.uuid "+
		"s.name, u.email FROM rooms r "+
		"INNER JOIN subjects s ON r.subject_id = s.subject_id "+
		"INNER JOIN users u ON r.author_id = u.user_id "+
		"WHERE r.room_id = $1", id).Scan(
		&r.RoomId,
		&r.Name,
		&r.UUID,
		&r.SubjectName,
		&r.AuthorEmail,
	)
	return r, err
}
