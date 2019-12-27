package models

import "time"

// InsuranceFee represent the InsuranceFee model
type InsuranceFee struct {
	ID        uint32     `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	Price     *float64   `gorm:"NOT NULL" json:"price,omitempty"`
	IsActive  *bool      `gorm:"NOT NULL" json:"isActive,omitempty"`
	Note      *string    `gorm:"TYPE:varchar(255)" json:"note,omitempty"`
	CreateAt  *time.Time `gorm:"DEFAULT:now()" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"DEFAULT:now()" json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}
