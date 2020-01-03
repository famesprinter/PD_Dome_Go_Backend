package customerroom

import (
	"context"

	"github.com/mr-fame/pd-dome-api/models"
)

// Usecase represent the customer_room's usecases
type Usecase interface {
	Fetch(ctx context.Context) ([]*models.CustomerRoom, error)
}
