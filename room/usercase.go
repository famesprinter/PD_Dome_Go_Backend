package room

import (
	"context"

	"github.com/mr-fame/pd-dome-api/models"
)

// Usecase represent the room's usecases
type Usecase interface {
	Fetch(ctx context.Context) ([]*models.Room, error)
	GetByID(ctx context.Context, id uint32) (*models.Room, error)
	Create(ctx context.Context, c *models.Room) error
	Update(ctx context.Context, c *models.Room) error
	Delete(ctx context.Context, id uint32) error
}
