package utils

import (
	"net/http"

	"github.com/mr-fame/pd-dome-api/models"
	"github.com/sirupsen/logrus"
)

// DataObject success oblect data
type DataObject struct {
	Title       *string     `json:"title,omitempty"`
	Description *string     `json:"description,omitempty"`
	Item        interface{} `json:"item,omitempty"`
	Items       interface{} `json:"items,omitempty"`
	NextOffset  int         `json:"nextOffset,omitempty"`
}

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// GetStatusCode return an error code
func GetStatusCode(err error) int {
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
