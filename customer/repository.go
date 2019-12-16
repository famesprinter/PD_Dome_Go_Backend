package customer

import (
	"github.com/mr-fame/pd-dome-api/models"
)

// Repository represent the article's repository contract
type Repository interface {
	Fetch(offset int, limit int) ([]*models.Customer, error)
	GetByID(id int) (*models.Customer, error)
	Create(ctm *models.Customer) error
	Update(ctm *models.Customer) error
	Delete(id int) error
}
