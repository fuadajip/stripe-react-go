package user

import "github.com/labstack/echo"

type Usecase interface {
	UserRegistration(c echo.Context, payload *models.UserRegistrationRequest) (*models.UserRegistrationResp, error)
}

type Repository interface {
	UserRegistration(c echo.Context, payload *models.UserRegistrationRequest) (*models.UserRegistrationResp, error)
}
