package usecase

import (
	"context"
	"time"

	"github.com/mr-fame/pd-dome-api/models"
	"github.com/mr-fame/pd-dome-api/room"
)

type roomUsecase struct {
	roomRepo       room.Repository
	contextTimeout time.Duration
}

// NewRoomUsecase will create new an roomUsecase object representation of room.Usecase interface
func NewRoomUsecase(rRep room.Repository, timeout time.Duration) room.Usecase {
	return &roomUsecase{
		roomRepo:       rRep,
		contextTimeout: timeout,
	}
}

func (room *roomUsecase) Fetch(ctx context.Context) ([]*models.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, room.contextTimeout)
	defer cancel()

	customers, err := room.roomRepo.Fetch()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (room *roomUsecase) GetByID(ctx context.Context, id uint32) (*models.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, room.contextTimeout)
	defer cancel()

	customer, err := room.roomRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (room *roomUsecase) Create(ctx context.Context, r *models.Room) error {
	ctx, cancel := context.WithTimeout(ctx, room.contextTimeout)
	defer cancel()

	err := room.roomRepo.Create(r)
	if err != nil {
		return err
	}
	return nil
}

func (room *roomUsecase) Update(ctx context.Context, r *models.Room) error {
	ctx, cancel := context.WithTimeout(ctx, room.contextTimeout)
	defer cancel()

	err := room.roomRepo.Update(r)
	if err != nil {
		return err
	}
	return nil
}

func (room *roomUsecase) Delete(ctx context.Context, id uint32) error {
	ctx, cancel := context.WithTimeout(ctx, room.contextTimeout)
	defer cancel()

	err := room.roomRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
