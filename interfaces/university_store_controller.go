package interfaces

import (
	"github.com/Flyewzz/group_preparation/models"
)

type UniversityController interface {
	Add(name string) (int, error)
	GetAll() ([]models.University, error)
	GetByPage(page int) ([]models.University, error)
	GetById(id int) (*models.University, error)
	SearchByName(name int) ([]models.University, error)
	RemoveById(id int) error
	RemoveAll() error
	GetAllSubjects(universityId int) ([]models.Subject, error)
	GetSubjectsByPage(universityId, page int) ([]models.Subject, error)
}
