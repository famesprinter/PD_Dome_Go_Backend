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

func (ctm *customerUsecase) Fetch(ctx context.Context, offset int, limit int) ([]*models.Customer, *int, error) {
	ctx, cancel := context.WithTimeout(ctx, ctm.contextTimeout)
	defer cancel()

	customers, err := ctm.customerRepo.Fetch(offset, limit)
	if err != nil {
		return nil, nil, err
	}

	nextOffset := offset + limit
	if nextOffset == 0 || len(customers) <= nextOffset {
		return customers, nil, nil
	}
	return customers, &nextOffset, nil
}

func (ctm *customerUsecase) GetByID(ctx context.Context, id int) (*models.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, ctm.contextTimeout)
	defer cancel()

	customer, err := ctm.customerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (ctm *customerUsecase) Create(ctx context.Context, c *models.Customer) error {
	ctx, cancel := context.WithTimeout(ctx, ctm.contextTimeout)
	defer cancel()

	err := ctm.customerRepo.Create(c)
	if err != nil {
		return err
	}
	return nil
}

func (ctm *customerUsecase) Update(ctx context.Context, c *models.Customer) error {
	ctx, cancel := context.WithTimeout(ctx, ctm.contextTimeout)
	defer cancel()

	err := ctm.customerRepo.Update(c)
	if err != nil {
		return err
	}
	return nil
}
