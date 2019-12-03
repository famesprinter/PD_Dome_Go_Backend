package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/mr-fame/pd-dome-api/customer"
	"github.com/mr-fame/pd-dome-api/models"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

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
	// e.POST("/customers", handler.Store)
	// e.GET("/customers/:id", handler.GetByID)
	// e.DELETE("/customers/:id", handler.Delete)
}

// FetchCustomer will fetch the customer based on given params
func (ctm *CustomerHandler) FetchCustomer(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.QueryParam("offset")
	offset, _ := strconv.Atoi(offsetStr)
	print(offset)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	customers, nextOffset, err := ctm.CUsecase.Fetch(ctx, offset, limit)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	var content struct {
		Items      []*models.Customer `json:"items"`
		NextOffset int                `json:"nextOffset"`
	}
	content.Items = customers
	content.NextOffset = nextOffset
	return c.JSON(http.StatusOK, &content)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
