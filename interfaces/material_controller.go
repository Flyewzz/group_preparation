package interfaces

import (
	"github.com/Flyewzz/group_preparation/models"
)

type MaterialController interface {
	GetById(id int) (*models.MaterialData, error)
	Add(subjectId int, name string, typeId, authorId int) (int, error)
	RemoveById(id int) error
	RemoveAll(subjectId int) error
	Search(subjectId int, name string, typeId, page int) ([]models.MaterialData, error)
	GetAllMaterials(subjectId, page int) ([]models.MaterialData, error)
	GetElementsCount(subjectId int) (int, error)
	GetItemsPerPageCount() int
}
