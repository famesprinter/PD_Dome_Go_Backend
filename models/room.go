package models

import "time"

// Room represent the Room model
type Room struct {
	ID             uint32     `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	RoomNumber     *string    `gorm:"TYPE:varchar(45);NOT NULL" json:"roomNumber,omitempty"`
	RentID         *uint32    `gorm:"NOT NULL" json:"rentID,omitempty"`
	InsuranceFeeID *uint32    `gorm:"NOT NULL" json:"insuranceFeeID,omitempty"`
	LevelID        *uint32    `gorm:"NOT NULL" json:"levelID,omitempty"`
	CreatedAt      *time.Time `gorm:"DEFAULT:now()" json:"createdAt,omitempty"`
	UpdatedAt      *time.Time `gorm:"DEFAULT:now()" json:"updatedAt,omitempty"`
	DeletedAt      *time.Time `json:"deletedAt,omitempty"`

	Rent *Rent
	InsuranceFee *InsuranceFee
	Level *Level
}
