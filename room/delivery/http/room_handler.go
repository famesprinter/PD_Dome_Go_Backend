package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v10"

	"github.com/mr-fame/pd-dome-api/models"
	"github.com/mr-fame/pd-dome-api/room"
	"github.com/mr-fame/pd-dome-api/utils"
)

// RoomHandler represent the httphandler for customer
type RoomHandler struct {
	RUsecase room.Usecase
}

// NewRoomHandler will initialize the room resources endpoint
func NewRoomHandler(e *echo.Echo, rUse room.Usecase) {
	handler := &RoomHandler{
		RUsecase: rUse,
	}
	e.GET("/rooms", handler.FetchRoom)
	e.GET("/rooms/:id", handler.GetByID)
	e.POST("/rooms/create", handler.Create)
	e.POST("/rooms/update/:id", handler.Update)
	e.DELETE("/rooms/delete/:id", handler.Delete)
}

// FetchRoom will fetch the room based on given params
func (handler *RoomHandler) FetchRoom(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	rooms, err := handler.RUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	title := "Rooms"
	description := "Get room"
	return c.JSON(http.StatusOK, utils.DataObject{
		Title:       &title,
		Description: &description,
		Items:       rooms,
	})
}

// GetByID will fetch the customer based on customer id
func (handler *RoomHandler) GetByID(c echo.Context) error {
	u64, _ := strconv.ParseUint(c.Param("id"), 2, 32)
	id := uint32(u64)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	room, err := handler.RUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	title := "Rooms"
	description := "Get room by customer id"
	return c.JSON(http.StatusOK, utils.DataObject{
		Title:       &title,
		Description: &description,
		Item:        room,
	})
}

// Create will create the customer by given request body
func (handler *RoomHandler) Create(c echo.Context) error {
	var room models.Room
	err := c.Bind(&room)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	if ok, err := isRequestValid(&room); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = handler.RUsecase.Create(ctx, &room)

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	title := "Room"
	description := "Create room success"
	return c.JSON(http.StatusOK, utils.DataObject{
		Title:       &title,
		Description: &description,
	})
}

// Update will create the customer by given request body
func (handler *RoomHandler) Update(c echo.Context) error {
	var room models.Room
	u64, _ := strconv.ParseUint(c.Param("id"), 2, 32)
	id := uint32(u64)
	err := c.Bind(&room)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	if ok, err := isRequestValid(&room); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	room.ID = id
	err = handler.RUsecase.Update(ctx, &room)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	title := "Rooms"
	description := "Update room success"
	return c.JSON(http.StatusOK, utils.DataObject{
		Title:       &title,
		Description: &description,
	})
}

// Delete will delelete the customer by given request body
func (handler *RoomHandler) Delete(c echo.Context) error {
	var room models.Room
	u64, _ := strconv.ParseUint(c.Param("id"), 2, 32)
	id := uint32(u64)
	err := c.Bind(&room)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	if ok, err := isRequestValid(&room); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = handler.RUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}

	title := "Rooms"
	description := "Delete room success"
	return c.JSON(http.StatusOK, utils.DataObject{
		Title:       &title,
		Description: &description,
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
