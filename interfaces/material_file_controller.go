package interfaces

import (
	"github.com/Flyewzz/group_preparation/models"
)

type MaterialFileController interface {
	GetById(id int) (*models.MaterialData, error)
	// Add(name, path string, materialId int) (int, error)
	GetAll(materialId int) ([]models.MaterialFile, error)
}
