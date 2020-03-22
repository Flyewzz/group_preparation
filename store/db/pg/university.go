package pg

import (
	"database/sql"

	. "github.com/Flyewzz/group_preparation/models"
)

type UniversityControllerPg struct {
	itemsPerPage int
	db           *sql.DB
}

func NewUniversityControllerPg(itemsPerPage int, db *sql.DB) *UniversityControllerPg {
	return &UniversityControllerPg{
		itemsPerPage: itemsPerPage,
		db:           db,
	}
}

func (usc *UniversityControllerPg) GetAll() ([]University, error) {
	rows, err := usc.db.Query("SELECT university_id, name FROM universities")
	var universities []University
	if err != nil {
		// 'universities' is an empty struct
		return universities, err
	}
	for rows.Next() {
		var u University
		rows.Scan(&u.Id, &u.Name)
		universities = append(universities, u)
	}
	defer rows.Close()
	return universities, nil
}

func (usc *UniversityControllerPg) GetByPage(page int) ([]University, error) {
	itemsPerPage := usc.itemsPerPage
	rows, err := usc.db.Query("SELECT university_id, name FROM universities LIMIT $1 OFFSET $2",
		itemsPerPage, itemsPerPage*(page-1))
	if err != nil {
		return nil, err
	}
	var universities []University
	for rows.Next() {
		var u University
		rows.Scan(&u.Id, &u.Name)
		universities = append(universities, u)
	}

	return universities, nil
}

func (usc *UniversityControllerPg) GetById(id int) (*University, error) {
	row := usc.db.QueryRow("SELECT university_id, name FROM universities WHERE university_id = $1", id)
	var u *University
	err := row.Scan(&u.Id, &u.Name)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (usc *UniversityControllerPg) Add(name string) (int, error) {
	var idUniversity int
	err := usc.db.QueryRow("INSERT INTO universities (name) VALUES ($1) RETURNING university_id", name).Scan(&idUniversity)
	if err != nil {
		return 0, err
	}
	return idUniversity, nil
}

func (usc *UniversityControllerPg) RemoveById(id int) error {
	_, err := usc.db.Exec("DELETE FROM univerisities WHERE university_id = $1", id)
	return err
}

func (usc *UniversityControllerPg) RemoveAll() error {

	_, err := usc.db.Exec("TRUNCATE TABLE universities;")
	return err
}
