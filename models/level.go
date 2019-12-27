package models

import "time"

// Level represent the Level model
type Level struct {
	ID        uint32     `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	Name      *string    `gorm:"TYPE:varchar(255);NOT NULL" json:"name,omitempty"`
	CreateAt  *time.Time `gorm:"DEFAULT:now()" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"DEFAULT:now()" json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}