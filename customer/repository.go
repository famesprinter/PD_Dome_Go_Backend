package customer

import (
	"github.com/mr-fame/pd-dome-api/models"
)

// Repository represent the article's repository contract
type Repository interface {
	Fetch(offset int, limit int) ([]*models.Customer, error)
	GetByID(id int) (*models.Customer, error)
	// Update(ctx context.Context, c *models.Customer) error
	// Store(ctx context.Context, c *models.Customer) error
	// Delete(ctx context.Context, id int) error
}
