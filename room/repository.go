package room

import (
	"github.com/mr-fame/pd-dome-api/models"
)

// Repository represent the room's repository contract
type Repository interface {
	Fetch() ([]*models.Room, error)
	GetByID(id uint32) (*models.Room, error)
	Create(room *models.Room) error
	Update(room *models.Room) error
	Delete(id uint32) error
}
