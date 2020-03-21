package interfaces

import (
	"github.com/Flyewzz/group_preparation/models"
)

type UniversityStoreController interface {
	Add(u *models.University) (int, error)
	GetAll() ([]models.University, error)
	GetByPage(page int) ([]models.University, error)
	GetById(id int) (*models.University, error)
	RemoveById(id int) error
	RemoveAll() error
}
