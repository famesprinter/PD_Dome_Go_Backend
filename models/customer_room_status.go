package models

// CustomerRoomStatus represent the CustomerRoomStatus model
type CustomerRoomStatus struct {
	ID   uint32  `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	Name *string `gorm:"TYPE:varchar(45);NOT NULL" json:"name,omitempty"`
}
