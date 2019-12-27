package models

import "time"

// CustomerRoom represent the CustomerRoom model
type CustomerRoom struct {
	ID                   uint32     `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	CustomerID           *uint32    `gorm:"NOT NULL" json:"customerID,omitempty"`
	RoomID               *uint32    `gorm:"NOT NULL" json:"roomID,omitempty"`
	CustomerRoomStatusID *uint32    `gorm:"NOT NULL" json:"customerRoomStatusID,omitempty"`
	CreatedAt            *time.Time `gorm:"DEFAULT:now()" json:"createdAt,omitempty"`
	UpdatedAt            *time.Time `gorm:"DEFAULT:now()" json:"updatedAt,omitempty"`
	DeletedAt            *time.Time `json:"deletedAt,omitempty"`
	CheckInAt            *time.Time `json:"checkInAt,omitempty"`
	CheckOutAt           *time.Time `json:"checkOutAt,omitempty"`

	Customer           *Customer
	Room               *Room
	CustomerRoomStatus *CustomerRoomStatus
}
