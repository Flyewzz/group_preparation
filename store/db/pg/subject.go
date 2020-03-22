package pg

import (
	"database/sql"
	// . "github.com/Flyewzz/group_preparation/models"
)

type SubjectControllerPg struct {
	itemsPerPage int
	db           *sql.DB
}

func NewSubjectControllerPg(itemsPerPage int, db *sql.DB) *SubjectControllerPg {
	return &SubjectControllerPg{
		itemsPerPage: itemsPerPage,
		db:           db,
	}
}

// func (ssc *SubjectControllerPg) GetAll() ([]Subject, error) {
// 	rows, err := ssc.db.Query("SELECT subject_id, name FROM subjects")
// 	var subjects []Subject
// 	if err != nil {
// 		// 'subjects' is an empty struct
// 		return subjects, err
// 	}
// 	for rows.Next() {
// 		var s Subject
// 		rows.Scan(&s.Id, &u.Name)
// 		subjects = append(subjects, u)
// 	}
// 	defer rows.Close()
// 	return subjects, nil
// }

// func (ssc *SubjectControllerPg) GetByPage(page int) ([]Subject, error) {
// 	itemsPerPage := ssc.itemsPerPage
// 	rows, err := ssc.db.Query("SELECT subject_id, name FROM subjects LIMIT $1 OFFSET $2",
// 		itemsPerPage, itemsPerPage*(page-1))
// 	if err != nil {
// 		return nil, err
// 	}
// 	var subjects []Subject
// 	for rows.Next() {
// 		var s Subject
// 		rows.Scan(&u.Id, &u.Name)
// 		subjects = append(subjects, u)
// 	}

// 	return subjects, nil
// }

// func (ssc *SubjectControllerPg) GetById(id int) (*University, error) {
// 	row := ssc.db.QueryRow("SELECT subject_id, name FROM subjects WHERE subject_id = $1", id)
// 	var u *University
// 	err := row.Scan(&u.Id, &u.Name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return u, nil
// }

// func (ssc *SubjectControllerPg) Add(name string) (int, error) {
// 	var idUniversity int
// 	err := ssc.db.QueryRow("INSERT INTO subjects (name) VALUES ($1) RETURNING subject_id", name).Scan(&idUniversity)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return idUniversity, nil
// }

// func (ssc *SubjectControllerPg) RemoveById(id int) error {
// 	_, err := ssc.db.Exec("DELETE FROM univerisities WHERE subject_id = $1", id)
// 	return err
// }

// func (ssc *SubjectControllerPg) RemoveAll() error {

// 	_, err := ssc.db.Exec("TRUNCATE TABLE subjects;")
// 	return err
// }
