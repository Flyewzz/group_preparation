package interfaces

import (
	"github.com/Flyewzz/group_preparation/models"
)

type SubjectController interface {
	GetById(id int) (*models.Subject, error)
	Add(universityId int, name, semester string) (int, error)
	RemoveById(id int) error
	RemoveAll(universityId int) error
	SearchByNameAndSemester(universityId int, name, semester string, page int) ([]models.Subject, error)
	SearchByName(universityId int, name string, page int) ([]models.Subject, error)
}
