package usecase

import (
	"context"
	"time"

	"github.com/mr-fame/pd-dome-api/customer"
	"github.com/mr-fame/pd-dome-api/models"
)

type customerUsecase struct {
	customerRepo   customer.Repository
	contextTimeout time.Duration
}

// NewCustomerUsecase will create new an customerUsecase object representation of customer.Usecase interface
func NewCustomerUsecase(ctm customer.Repository, timeout time.Duration) customer.Usecase {
	return &customerUsecase{
		customerRepo:   ctm,
		contextTimeout: timeout,
	}
}

func (ctm *customerUsecase) Fetch(ctx context.Context, offset int, limit int) ([]*models.Customer, int, error) {
	if limit == 0 {
		limit = 20
	}

	ctx, cancel := context.WithTimeout(ctx, ctm.contextTimeout)
	defer cancel()

	customers, nextOffset, err := ctm.customerRepo.Fetch(offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return customers, nextOffset, nil
}
