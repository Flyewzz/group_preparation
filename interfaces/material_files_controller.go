package interfaces

import (
	"github.com/Flyewzz/group_preparation/models"
)

type MaterialController interface {
	GetById(id int) (*models.MaterialData, error)
	Add(name, path string, materialId int) (int, error)
	RemoveById(id int) error
	RemoveAll(subjectId int) error
	GetAll(materialId int) ([]models.MaterialData, error)

	GetItemsPerPageCount() int
}
