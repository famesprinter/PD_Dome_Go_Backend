package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mr-fame/pd-dome-api/models"
	"github.com/mr-fame/pd-dome-api/room"
)

type mysqlRoomRepository struct {
	Conn *gorm.DB
}

// NewMysqlRoomRepository will create an object that represent the room.Repository interface
func NewMysqlRoomRepository(Conn *gorm.DB) room.Repository {
	return &mysqlRoomRepository{Conn}
}

func (m *mysqlRoomRepository) Fetch() ([]*models.Room, error) {
	rooms := []*models.Room{}
	// db := m.Conn.Find(&rooms)
	db := m.Conn.Preload("Rent").Find(&rooms)
	if db.Error != nil {
		return nil, db.Error
	}
	return rooms, nil
}

func (m *mysqlRoomRepository) GetByID(id uint32) (*models.Room, error) {
	room := models.Room{}
	db := m.Conn.First(&room, id)
	if db.Error != nil {
		return nil, db.Error
	}
	return &room, nil
}

func (m *mysqlRoomRepository) Create(room *models.Room) error {
	fmt.Print(room)
	db := m.Conn.Create(room)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (m *mysqlRoomRepository) Update(room *models.Room) error {
	db := m.Conn.Model(&room).Updates(&room)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (m *mysqlRoomRepository) Delete(id uint32) error {
	db := m.Conn.Where("id = ?", id).Delete(&models.Customer{})
	if db.Error != nil {
		return db.Error
	}
	return nil
}
