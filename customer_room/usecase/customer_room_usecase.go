package usecase

import (
	"context"
	"time"

	customerroom "github.com/mr-fame/pd-dome-api/customer_room"
	"github.com/mr-fame/pd-dome-api/models"
)

type customerRoomUsecase struct {
	customerRoomRepo customerroom.Repository
	contextTimeout   time.Duration
}

// NewCustomerRoomUsecase will create new an customerRoomUsecase object representation of customerroom.Usecase interface
func NewCustomerRoomUsecase(crRep customerroom.Repository, timeout time.Duration) customerroom.Usecase {
	return &customerRoomUsecase{
		customerRoomRepo: crRep,
		contextTimeout:   timeout,
	}
}

func (customerRoom *customerRoomUsecase) Fetch(ctx context.Context) ([]*models.CustomerRoom, error) {
	ctx, cancel := context.WithTimeout(ctx, customerRoom.contextTimeout)
	defer cancel()

	customerRooms, err := customerRoom.customerRoomRepo.Fetch()
	if err != nil {
		return nil, err
	}
	return customerRooms, nil
}
