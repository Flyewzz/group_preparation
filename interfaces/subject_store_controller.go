package interfaces

import (
	"github.com/Flyewzz/group_preparation/models"
)

type SubjectController interface {
	GetById(id int) (*models.University, error)
	Add(universityId int, name, semester string) (int, error)
	RemoveById(id int) error
	RemoveAll(universityId int) error
	GetByPage(page int) ([]models.University, error)
}
