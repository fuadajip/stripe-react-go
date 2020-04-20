package util

import (
	"fmt"
	"net/http"

	"github.com/fgrosse/goldi"
	"github.com/fuadajip/stripe-react-go/server/models"
	"github.com/fuadajip/stripe-react-go/server/shared/config"
	"github.com/fuadajip/stripe-react-go/server/shared/log"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var (
	logger = log.NewLog()
)

type (
	// CustomApplicationContext return qoala custom application context
	CustomApplicationContext struct {
		echo.Context
		Container    *goldi.Container
		SharedConf   config.ImmutableConfigInterface
		RedisSession *redis.Client
		MysqlSession *gorm.DB
	}

	// CustomValidator return qoala custom validator
	CustomValidator struct {
		Validator *validator.Validate
	}
)

// CustomResponse is a method that returns custom object response
func (c *CustomApplicationContext) CustomResponse(status string, data interface{}, message string, systemMessage string, code int, meta *models.ResponsePatternMeta) error {
	return c.JSON(code, &models.ResponsePattern{
		Status:        status,
		Data:          data,
		Message:       message,
		SystemMessage: systemMessage,
		Code:          code,
		Meta:          *meta,
	})
}

// Validate will validate given input with related struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// CustomHTTPErrorHandler will return custom echo http error handler
func CustomHTTPErrorHandler(err error, e echo.Context) {

	report, ok := err.(*echo.HTTPError)
	var msgError string

	if !ok {
		msgError = "[Generic] Internal server error, error [" + err.Error() + "]"
		report = echo.NewHTTPError(http.StatusInternalServerError, msgError)
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		msgError = "[Validation] Invalid validation, error [ field: %s is %s ]"
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				msgError = fmt.Sprintf(msgError, err.Field(), "is required")
				report = echo.NewHTTPError(http.StatusBadRequest, msgError)
			case "email":
				msgError = fmt.Sprintf(msgError, err.Field(), "is not valid email")
				report = echo.NewHTTPError(http.StatusBadRequest, msgError)
				break
			}
		}

	}

	logger.Error(msgError)
	qr := &models.ResponsePattern{
		Code:    report.Code,
		Data:    nil,
		Message: fmt.Sprintf("%+v", report.Message),
		Meta:    models.ResponsePatternMeta{},
		Status:  "failed",
	}

	e.JSON(report.Code, qr)
}
