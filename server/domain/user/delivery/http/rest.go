package http

import (
	"fmt"
	"net/http"

	"github.com/fuadajip/stripe-react-go/server/domain/user"
	"github.com/fuadajip/stripe-react-go/server/models"
	"github.com/fuadajip/stripe-react-go/server/shared/log"
	"github.com/fuadajip/stripe-react-go/server/shared/util"
	"github.com/labstack/echo"
)

var (
	logger = log.NewLog()
)

type handlerUser struct {
	usecase user.Usecase
}

func AddUserHandler(e *echo.Echo, usecase user.Usecase) {
	handler := handlerUser{
		usecase: usecase,
	}

	e.POST("/api/user/registrations", handler.UserRegistration)
}

func (h handlerUser) UserRegistration(c echo.Context) error {
	ac := c.(*util.CustomApplicationContext)

	payload := &models.UserRegistrationRequest{}

	if err := ac.Bind(payload); err != nil {
		msgError := fmt.Sprintf("[FAILED][EXP-SERVICE][USER][UserRegistration] Invalid request payload bind, err: %s", err.Error())
		logger.Error(msgError)
		return ac.CustomResponse("failed", nil, "Invalid request", msgError, http.StatusBadRequest, &models.QoalaResponsePatternMeta{})
	}

	resp, err := h.usecase.UserRegistration(c, payload)
	if err != nil {
		msgError := "[FAILED][EXP-SERVICE][USER][UserRegistration] Failed user registration internal server error, err: " + err.Error()
		return ac.CustomResponse("failed", nil, msgError, "Registration failed", http.StatusInternalServerError, &models.ResponsePatternMeta{})
	}

	return ac.CustomResponse("success", resp, nil, "Registration Success", http.StatusOK, &models.ResponsePatternMeta{})
}
