package models

import "time"

// Rent represent the Rent model
type Rent struct {
	ID        uint32     `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	Price     *float64   `gorm:"NOT NULL" json:"price,omitempty"`
	IsActive  *bool      `gorm:"NOT NULL" json:"isActive,omitempty"`
	CreateAt  *time.Time `gorm:"DEFAULT:now()" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"DEFAULT:now()" json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}