package models

import "time"

// Receipt represent the Receipt model
type Receipt struct {
	ID                uint32     `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	ReceiptStatusID   *uint32    `gorm:"NOT NULL" json:"receiptStatusID,omitempty"`
	RoomID            *uint32    `gorm:"NOT NULL" json:"roomID,omitempty"`
	ElectricUnitCalID *uint32    `gorm:"NOT NULL" json:"electricUnitCalID,omitempty"`
	WaterUnitCalID    *uint32    `gorm:"NOT NULL" json:"waterUnitCalID,omitempty"`
	ElectricUnits     *float64   `gorm:"NOT NULL" json:"electricUnits,omitempty"`
	WaterUnits        *float64   `gorm:"NOT NULL" json:"waterUnits,omitempty"`
	Note              *string    `gorm:"TYPE:varchar(255)" json:"note,omitempty"`
	CreatedAt         *time.Time `gorm:"DEFAULT:now()" json:"createdAt,omitempty"`
	UpdatedAt         *time.Time `gorm:"DEFAULT:now()" json:"updatedAt,omitempty"`
	DeletedAt         *time.Time `json:"deletedAt,omitempty"`

	ReceiptStatus   *ReceiptStatus
	Room            *Room
	ElectricUnitCal *ElectricUnitCal
	WaterUnitCal    *WaterUnitCal
}
