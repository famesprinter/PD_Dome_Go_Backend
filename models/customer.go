package models

import (
	"time"
)

// Customer represent the customer model
type Customer struct {
	ID             int       `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	FirstName      string    `json:"firstName,omitempty"`
	LastName       string    `json:"lastName,omitempty"`
	PhoneNumber    int       `json:"phoneNumber,omitempty"`
	Address        string    `json:"address"`
	IDCardImageURL string    `json:"idCardImageURL"`
	CreatedAt      time.Time `gorm:"DEFAULT:now()" json:"createdAt,omitempty"`
	UpdatedAt      time.Time `gorm:"DEFAULT:now()" json:"updatedAt"`
}
