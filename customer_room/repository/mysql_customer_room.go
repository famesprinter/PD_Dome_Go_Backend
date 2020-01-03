package repository

import (
	"github.com/jinzhu/gorm"
	customerroom "github.com/mr-fame/pd-dome-api/customer_room"
	"github.com/mr-fame/pd-dome-api/models"
)

type mysqlCustomerRoomRepository struct {
	Conn *gorm.DB
}

// NewMysqlCustomerRoomRepository will create an object that represent the room.Repository interface
func NewMysqlCustomerRoomRepository(Conn *gorm.DB) customerroom.Repository {
	return &mysqlCustomerRoomRepository{Conn}
}

func (m *mysqlCustomerRoomRepository) Fetch() ([]*models.CustomerRoom, error) {
	customerRooms := []*models.CustomerRoom{}
	db := m.Conn.Preload("Customer").Preload("Room").Preload("CustomerRoomStatus").Find(&customerRooms)
	if db.Error != nil {
		return nil, db.Error
	}
	return customerRooms, nil
}
