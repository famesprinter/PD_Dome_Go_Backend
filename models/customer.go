package models

import "time"

// Customer represent the customer model
type Customer struct {
	ID             uint32     `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	FirstName      *string    `gorm:"TYPE:varchar(255);NOT NULL" json:"firstName,omitempty"`
	LastName       *string    `gorm:"TYPE:varchar(255);NOT NULL" json:"lastName,omitempty"`
	PhoneNumber    *int       `gorm:"NOT NULL" json:"phoneNumber,omitempty"`
	Address        *string    `gorm:"TYPE:varchar(255)" json:"address,omitempty"`
	IDCardImageURL *string    `gorm:"TYPE:varchar(255)" json:"idCardImageUrl,omitempty"`
	CreatedAt      *time.Time `gorm:"DEFAULT:now()" json:"createdAt,omitempty"`
	UpdatedAt      *time.Time `gorm:"DEFAULT:now()" json:"updatedAt,omitempty"`
	DeletedAt      *time.Time `json:"deletedAt,omitempty"`
}
