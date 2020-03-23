package pg

import (
	"database/sql"

	"github.com/Flyewzz/group_preparation/models"
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

func (sc *SubjectControllerPg) SearchByNameAndSemester(universityId int, name, semester string) ([]models.Subject, error) {
	rows, err := sc.db.Query("SELECT subject_id, name, semester FROM subjects "+
		"WHERE LOWER(name) LIKE '%' || $1 || '%' AND semester = $2 AND university_id = $3"+
		"ORDER BY name ASC", name, semester, universityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var subjects []models.Subject
	for rows.Next() {
		var subject models.Subject
		err = rows.Scan(&subject.Id, &subject.Name, &subject.Semester)
		if err != nil {
			continue
		}
		subjects = append(subjects, subject)
	}
	return subjects, nil
}

func (sc *SubjectControllerPg) SearchByName(universityId int, name string) ([]models.Subject, error) {
	rows, err := sc.db.Query("SELECT subject_id, name, semester FROM subjects "+
		"WHERE LOWER(name) LIKE '%' || $1 || '%' AND university_id = $2"+
		"ORDER BY name ASC", name, universityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var subjects []models.Subject
	for rows.Next() {
		var subject models.Subject
		err = rows.Scan(&subject.Id, &subject.Name, &subject.Semester)
		if err != nil {
			continue
		}
		subjects = append(subjects, subject)
	}
	return subjects, nil
}
