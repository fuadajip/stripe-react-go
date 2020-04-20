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

	e.POST("/api/user/registration", handler.UserRegistration)
	e.POST("/api/user/login", handler.UserLogin)
}

func (h handlerUser) UserRegistration(c echo.Context) error {
	ac := c.(*util.CustomApplicationContext)

	payload := &models.UserRegistrationRequest{}

	if err := ac.Bind(payload); err != nil {
		msgError := fmt.Sprintf("[FAILED][EXP-SERVICE][USER][UserRegistration] Invalid request payload bind, err: %s", err.Error())
		logger.Error(msgError)
		return ac.CustomResponse("failed", nil, "Invalid request", msgError, http.StatusBadRequest, &models.ResponsePatternMeta{})
	}

	resp, err := h.usecase.UserRegistration(c, payload)
	if err != nil {
		if err.Error() == "EXIST_USERNAME" {
			msgError := "[FAILED][EXP-SERVICE][USER][UserRegistration] Failed user registration exist user, err: " + err.Error()
			return ac.CustomResponse("failed", nil, msgError, "Username not available", http.StatusBadRequest, &models.ResponsePatternMeta{})
		} else if err.Error() == "EXIST_EMAIL_PHONE" {
			msgError := "[FAILED][EXP-SERVICE][USER][UserRegistration] Failed user registration exist user, err: " + err.Error()
			return ac.CustomResponse("failed", nil, msgError, "Email or phone already registered", http.StatusBadRequest, &models.ResponsePatternMeta{})
		} else {
			msgError := "[FAILED][EXP-SERVICE][USER][UserRegistration] Failed user registration internal error, err: " + err.Error()
			return ac.CustomResponse("failed", nil, msgError, "Registration failed", http.StatusInternalServerError, &models.ResponsePatternMeta{})
		}

	}

	return ac.CustomResponse("success", resp, "Registration Success", "success", http.StatusOK, &models.ResponsePatternMeta{})
}

func (h handlerUser) UserLogin(c echo.Context) error {
	ac := c.(*util.CustomApplicationContext)

	payload := &models.UserLoginRequest{}

	if err := ac.Bind(payload); err != nil {
		msgError := fmt.Sprintf("[FAILED][EXP-SERVICE][USER][UserLogin] Invalid request payload bind, err: %s", err.Error())
		logger.Error(msgError)
		return ac.CustomResponse("failed", nil, "Invalid request", msgError, http.StatusBadRequest, &models.ResponsePatternMeta{})
	}

	resp, err := h.usecase.UserLogin(c, payload)
	if err != nil {
		if err.Error() == "INVALID_CREDENTIALS" {
			msgError := "[FAILED][EXP-SERVICE][USER][UserLogin] Failed user login, err: " + err.Error()
			return ac.CustomResponse("failed", nil, msgError, "Invalid credentials", http.StatusBadRequest, &models.ResponsePatternMeta{})
		} else if err.Error() == "USER_NOTFOUND" {
			msgError := "[FAILED][EXP-SERVICE][USER][UserLogin] Failed user login not found , err: " + err.Error()
			return ac.CustomResponse("failed", nil, msgError, "User not found", http.StatusBadRequest, &models.ResponsePatternMeta{})
		} else {
			msgError := "[FAILED][EXP-SERVICE][USER][UserLogin] Failed user login internal error, err: " + err.Error()
			return ac.CustomResponse("failed", nil, msgError, "Registration failed", http.StatusInternalServerError, &models.ResponsePatternMeta{})
		}

	}

	return ac.CustomResponse("success", resp, "Registration Success", "success", http.StatusOK, &models.ResponsePatternMeta{})
}
