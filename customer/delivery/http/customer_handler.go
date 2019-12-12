package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v10"

	"github.com/mr-fame/pd-dome-api/customer"
	"github.com/mr-fame/pd-dome-api/models"
	"github.com/mr-fame/pd-dome-api/utils"
)

// CustomerHandler represent the httphandler for customer
type CustomerHandler struct {
	CUsecase customer.Usecase
}

// NewCustomerHandler will initialize the customer resources endpoint
func NewCustomerHandler(e *echo.Echo, us customer.Usecase) {
	handler := &CustomerHandler{
		CUsecase: us,
	}
	e.GET("/customers", handler.FetchCustomer)
	e.GET("/customers/:id", handler.GetByID)
	e.POST("/customers/create", handler.Create)
	e.POST("/customers/update/:id", handler.Update)

	// e.DELETE("/customers/:id", handler.Delete)
}

// FetchCustomer will fetch the customer based on given params
func (ctm *CustomerHandler) FetchCustomer(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.QueryParam("offset")
	offset, _ := strconv.Atoi(offsetStr)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	customers, nextOffset, err := ctm.CUsecase.Fetch(ctx, offset, limit)
	_ = nextOffset
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	title := "Customers"
	description := "Get customer"
	return c.JSON(http.StatusOK, utils.DataObject{
		Title:       &title,
		Description: &description,
		Items:       customers,
	})
}

// GetByID will fetch the customer based on customer id
func (ctm *CustomerHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	customer, err := ctm.CUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	title := "Customers"
	description := "Get customer by customer id"
	return c.JSON(http.StatusOK, utils.DataObject{
		Title:       &title,
		Description: &description,
		Item:        customer,
	})
}

// Create will create the customer by given request body
func (ctm *CustomerHandler) Create(c echo.Context) error {
	var customer models.Customer
	err := c.Bind(&customer)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	if ok, err := isRequestValid(&customer); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = ctm.CUsecase.Create(ctx, &customer)

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	title := "Customers"
	description := "Create customer success"
	return c.JSON(http.StatusOK, utils.DataObject{
		Title:       &title,
		Description: &description,
	})
}

// Update will create the customer by given request body
func (ctm *CustomerHandler) Update(c echo.Context) error {
	var customer models.Customer
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.Bind(&customer)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	if ok, err := isRequestValid(&customer); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	customer.ID = id
	err = ctm.CUsecase.Update(ctx, &customer)

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	title := "Customers"
	description := "Update customer success"
	return c.JSON(http.StatusOK, utils.DataObject{
		Title:       &title,
		Description: &description,
	})
}

// Private Function
func isRequestValid(m *models.Customer) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
