package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v10"

	customerroom "github.com/mr-fame/pd-dome-api/customer_room"
	"github.com/mr-fame/pd-dome-api/models"
	"github.com/mr-fame/pd-dome-api/utils"
)

// CustomerRoomHandler represent the httphandler for customer rooms
type CustomerRoomHandler struct {
	CRUsecase customerroom.Usecase
}

// NewCustomerRoomHandler will initialize the room resources endpoint
func NewCustomerRoomHandler(e *echo.Echo, crUse customerroom.Usecase) {
	handler := &CustomerRoomHandler{
		CRUsecase: crUse,
	}
	e.GET("/customer/rooms", handler.FetchCustomerRoom)
}

// FetchCustomerRoom will fetch the customer rooms based on given params
func (handler *CustomerRoomHandler) FetchCustomerRoom(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	customerRooms, err := handler.CRUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	title := "Customer Rooms"
	description := "Get customer room"
	return c.JSON(http.StatusOK, utils.DataObject{
		Title:       &title,
		Description: &description,
		Items:       customerRooms,
	})
}

// Private Function
func isRequestValid(m *models.Room) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
