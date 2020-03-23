package pg

import (
	"database/sql"

	"github.com/Flyewzz/group_preparation/models"
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

func (uc *UniversityControllerPg) GetAllSubjects(universityId int) ([]models.Subject, error) {
	rows, err := uc.db.Query("SELECT s.subject_id, s.name, s.semester from subjects s "+
		"INNER JOIN universities u ON s.university_id = u.university_id AND u.university_id = $1", universityId)
	if err != nil {
		return nil, err
	}
	var subjects []models.Subject
	defer rows.Close()
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

func (uc *UniversityControllerPg) GetSubjectsByPage(universityId, page int) ([]models.Subject, error) {
	itemsPerPage := uc.itemsPerPage
	rows, err := uc.db.Query("SELECT s.subject_id, s.name, s.semester FROM subjects s "+
		"INNER JOIN universities u ON s.university_id = u.university_id LIMIT $1 OFFSET $2",
		itemsPerPage, itemsPerPage*(page-1))
	if err != nil {
		return nil, err
	}
	var subjects []models.Subject
	defer rows.Close()
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

func (uc *UniversityControllerPg) GetAll() ([]University, error) {
	rows, err := uc.db.Query("SELECT university_id, name FROM universities")
	var universities []University
	if err != nil {
		// 'universities' is an empty struct
		return universities, err
	}
	defer rows.Close()
	for rows.Next() {
		var u University
		rows.Scan(&u.Id, &u.Name)
		universities = append(universities, u)
	}
	return universities, nil
}

func (uc *UniversityControllerPg) GetByPage(page int) ([]University, error) {
	itemsPerPage := uc.itemsPerPage
	rows, err := uc.db.Query("SELECT university_id, name FROM universities LIMIT $1 OFFSET $2",
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

func (uc *UniversityControllerPg) GetById(id int) (*University, error) {
	row := uc.db.QueryRow("SELECT university_id, name FROM universities WHERE university_id = $1", id)
	var u *University
	err := row.Scan(&u.Id, &u.Name)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UniversityControllerPg) Add(name string) (int, error) {
	var idUniversity int
	err := uc.db.QueryRow("INSERT INTO universities (name) VALUES ($1) RETURNING university_id", name).Scan(&idUniversity)
	if err != nil {
		return 0, err
	}
	return idUniversity, nil
}

func (uc *UniversityControllerPg) RemoveById(id int) error {
	_, err := uc.db.Exec("DELETE FROM universities WHERE university_id = $1", id)
	return err
}

func (uc *UniversityControllerPg) RemoveAll() error {
	// #! Removed all the subjects too. Warning!
	_, err := uc.db.Exec("TRUNCATE TABLE universities CASCADE")
	return err
}

func (uc *UniversityControllerPg) SearchByName(name string) ([]models.University, error) {
	rows, err := uc.db.Query("SELECT university_id, name FROM universities "+
		"WHERE LOWER(name) LIKE '%' || $1 || '%' "+
		"ORDER BY name ASC", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var universities []models.University
	for rows.Next() {
		var u models.University
		err = rows.Scan(&u.Id, &u.Name)
		if err != nil {
			continue
		}
		universities = append(universities, u)
	}
	return universities, nil
}
