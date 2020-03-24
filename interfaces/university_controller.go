package interfaces

import (
	"github.com/Flyewzz/group_preparation/models"
)

type UniversityController interface {
	Add(name, fullName string) (int, error)
	GetAll(page int) ([]models.University, error)
	GetById(id int) (*models.University, error)
	Search(name string, page int) ([]models.University, error)
	RemoveById(id int) error
	RemoveAll() error

	GetItemsPerPageCount() int
}
