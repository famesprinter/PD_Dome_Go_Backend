package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	h := CustomerHandler{}
	h.Initialize()

	e.GET("/customers", h.GetAllCustomer)
	e.GET("/customers/:id", h.GetCustomer)
	e.POST("/customers", h.CreateCustomer)
	// e.PUT("/customers/:id", h.UpdateCustomer)
	// e.DELETE("/customers/:id", h.DeleteCustomer)

	e.Logger.Fatal(e.Start(":8080"))
}

// CustomerHandler structure
type CustomerHandler struct {
	DB *gorm.DB
}

// Initialize Custormer
func (h *CustomerHandler) Initialize() {
	db, err := gorm.Open("mysql", "root:Root1234@/pd_dome_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Customer{})

	h.DB = db
}

// Customer structure
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

// GetAllCustomer get all customer
func (h *CustomerHandler) GetAllCustomer(c echo.Context) error {
	customers := []Customer{}

	h.DB.Find(&customers)

	return c.JSON(http.StatusOK, customers)
}

// GetCustomer get customer by id
func (h *CustomerHandler) GetCustomer(c echo.Context) error {
	id := c.Param("id")
	customer := Customer{}

	if err := h.DB.Find(&customer, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, customer)
}

// CreateCustomer create customer
func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	customer := Customer{}

	if err := c.Bind(&customer); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Create(&customer).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, customer)
}
