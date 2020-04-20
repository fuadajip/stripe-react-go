package http

import (
	"fmt"
	"net/http"

	"github.com/fuadajip/stripe-react-go/server/domain/healthz"
	"github.com/fuadajip/stripe-react-go/server/models"

	"github.com/fuadajip/stripe-react-go/server/shared/log"
	"github.com/fuadajip/stripe-react-go/server/shared/util"
	"github.com/labstack/echo"
)

var (
	logger = log.NewLog()
)

type handlerHealtCheck struct {
	usecase healthz.Usecase
}

// AddHealthCheckHandler returns http handler for db session healthz
func AddHealthCheckHandler(e *echo.Echo, usecase healthz.Usecase) {
	handler := handlerHealtCheck{
		usecase: usecase,
	}

	e.GET("/api/healthz", handler.DoHeathCheck)
}

func (h handlerHealtCheck) DoHeathCheck(c echo.Context) error {
	ac := c.(*util.CustomApplicationContext)
	res, err := h.usecase.DoHealthCheck(c)
	if err != nil {
		msgError := fmt.Sprintf("[FAILED][EXP-SERVICE][HEALTHCHECK][DoHeathCheck] Failed to healthcheck, err: %s", err.Error())
		logger.Error(msgError)
		return ac.CustomResponse("failed", res, "System unhealthy", msgError, http.StatusInternalServerError, &models.ResponsePatternMeta{})
	}

	return ac.CustomResponse("success", res, "System healthy", "[SUCCESS][EXP-SERVICE][HEALTHCHECK][DoHeathCheck] Success system healthy", http.StatusOK, &models.ResponsePatternMeta{})
}
