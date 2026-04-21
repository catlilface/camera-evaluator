package utils

import (
	"net/http"
	api "photo-upload-service/internal/pkg/api/photo"
	"reflect"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, res interface{}) {
	v := reflect.ValueOf(res)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct {
		statusField := v.FieldByName("Status")
		if statusField.IsValid() && statusField.CanSet() && statusField.Kind() == reflect.String {
			statusField.SetString(http.StatusText(http.StatusOK))
		}
	}

	c.JSON(http.StatusOK, res)
}

func AbortWithStatus(c *gin.Context, statusCode int, err error) {
	_ = c.Error(err)
	c.AbortWithStatusJSON(statusCode, api.ErrorResponse{
		Error:  err.Error(),
		Status: http.StatusText(statusCode),
	})
}
