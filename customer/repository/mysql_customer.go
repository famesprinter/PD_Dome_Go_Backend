package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/mr-fame/pd-dome-api/customer"
	"github.com/mr-fame/pd-dome-api/models"
)

type mysqlCustomerRepository struct {
	Conn *gorm.DB
}

// NewMysqlCustomerRepository will create an object that represent the customer.Repository interface
func NewMysqlCustomerRepository(Conn *gorm.DB) customer.Repository {
	return &mysqlCustomerRepository{Conn}
}

func (m *mysqlCustomerRepository) Fetch(offset int, limit int) (res []*models.Customer, nextOffset int, err error) {
	customers := []*models.Customer{}
	m.Conn.Limit(limit).Offset(offset).Find(&customers)
	nextOffset = 0
	if offset == 0 {
		nextOffset = limit
	} else {
		nextOffset = offset + limit
	}
	return customers, nextOffset, nil
}
