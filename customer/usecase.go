package customer

import (
	"context"

	"github.com/mr-fame/pd-dome-api/models"
)

// Usecase represent the article's usecases
type Usecase interface {
	Fetch(ctx context.Context, offset int, limit int) ([]*models.Customer, *int, error)
	GetByID(ctx context.Context, id int) (*models.Customer, error)
	Create(ctx context.Context, c *models.Customer) error
	Update(ctx context.Context, c *models.Customer) error
	Delete(ctx context.Context, id int) error
}
