package customerroom

import (
	"github.com/mr-fame/pd-dome-api/models"
)

// Repository represent the customer_room's repository contract
type Repository interface {
	Fetch() ([]*models.CustomerRoom, error)
}
