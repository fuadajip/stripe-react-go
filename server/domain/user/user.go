package user

import (
	"github.com/fuadajip/stripe-react-go/server/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Usecase interface {
	UserRegistration(c echo.Context, payload *models.UserRegistrationRequest) (*models.UserRegistrationResponse, error)
	UserLogin(c echo.Context, payload *models.UserLoginRequest) (*models.UserLoginResponse, error)
}

type Repository interface {
	FindUserByUsername(c echo.Context, payload *models.User) (*models.User, error)
	FindUserByEmailPhone(c echo.Context, payload *models.User) (*models.User, error)
	TrxUserRegistration(trx *gorm.DB, payload *models.User) (*models.User, error)
}
