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

func (m *mysqlCustomerRepository) Fetch(offset int, limit int) ([]*models.Customer, error) {
	customers := []*models.Customer{}
	if limit == 0 {
		limit = -1
	}
	// db := m.Conn.Limit(limit).Offset(offset).Find(&customers)
	db := m.Conn.Find(&customers)
	if db.Error != nil {
		return nil, db.Error
	}
	return customers, nil
}

func (m *mysqlCustomerRepository) GetByID(id uint32) (*models.Customer, error) {
	customer := models.Customer{}
	db := m.Conn.First(&customer, id)
	if db.Error != nil {
		return nil, db.Error
	}
	return &customer, nil
}

func (m *mysqlCustomerRepository) Create(ctm *models.Customer) error {
	db := m.Conn.Create(&ctm)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (m *mysqlCustomerRepository) Update(ctm *models.Customer) error {
	db := m.Conn.Model(&ctm).Updates(&ctm)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (m *mysqlCustomerRepository) Delete(id uint32) error {
	db := m.Conn.Where("id = ?", id).Delete(&models.Customer{})
	if db.Error != nil {
		return db.Error
	}
	return nil
}
