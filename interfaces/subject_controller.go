package interfaces

import (
	"github.com/Flyewzz/group_preparation/models"
)

type SubjectController interface {
	GetById(id int) (*models.Subject, error)
	Add(universityId int, name, semester string) (int, error)
	RemoveById(id int) error
	RemoveAll(universityId int) error
	Search(universityId int, name, semester string, page int) ([]models.Subject, error)
	GetAllSubjects(universityId, page int) ([]models.Subject, error)
	GetItemsPerPageCount() int
}
