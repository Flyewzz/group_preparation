package pg

import (
	"database/sql"

	. "github.com/Flyewzz/group_preparation/models"
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

func (sc *SubjectControllerPg) GetByPage(universityId, page int) ([]Subject, error) {
	itemsPerPage := sc.itemsPerPage
	rows, err := sc.db.Query("SELECT subject_id, university_id, name, semester FROM subjects "+
		"LIMIT $1 OFFSET $2",
		itemsPerPage, itemsPerPage*(page-1))
	if err != nil {
		return nil, err
	}
	var subjects []Subject
	for rows.Next() {
		var s Subject
		err := rows.Scan(&s.Id, &s.UniversityId, &s.Name, &s.Semester)
		if err != nil {
			continue
		}
		subjects = append(subjects, s)
	}
	return subjects, nil
}

func (sc *SubjectControllerPg) GetById(id int) (*Subject, error) {
	row := sc.db.QueryRow("SELECT subject_id, university_id, name, semester FROM subjects "+
		"WHERE subject_id = $1", id)
	var s *Subject
	err := row.Scan(&s.Id, &s.UniversityId, &s.Name, &s.Semester)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (sc *SubjectControllerPg) Add(universityId int, name, semester string) (int, error) {
	var idSubject int
	err := sc.db.QueryRow("INSERT INTO subjects (university_id, name, semester) VALUES ($1, $2, $3) RETURNING subject_id",
		universityId, name, semester).Scan(&idSubject)
	if err != nil {
		return 0, err
	}
	return idSubject, nil
}

func (sc *SubjectControllerPg) RemoveById(id int) error {
	_, err := sc.db.Exec("DELETE FROM subjects WHERE subject_id = $1", id)
	return err
}

func (sc *SubjectControllerPg) RemoveAll(universityId int) error {
	// #! Removed all the materials too. Warning!
	_, err := sc.db.Exec("DELETE FROM subjects WHERE university_id = $1", universityId)
	return err
}
