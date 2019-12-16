package models

import "time"

// Room represent the Room model
type Room struct {
	ID             int       `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	RoomNumber     string    `json:"roomNumber,omitempty"`
	RendID         int       `gorm:"ForeignKey" json:"rendID,omitempty"`
	InsuranceFeeID int       `gorm:"ForeignKey" json:"insuranceFeeID,omitempty"`
	LevelID        int       `gorm:"ForeignKey" json:"levelID,omitempty"`
	CreatedAt      time.Time `gorm:"DEFAULT:now()" json:"createdAt,omitempty"`
	UpdatedAt      time.Time `gorm:"DEFAULT:now()" json:"updatedAt"`
	DeletedAt      time.Time `gorm:"DEFAULT:now()" json:"deletedAt"`
}
