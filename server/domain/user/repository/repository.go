package repository

import (
	"github.com/fuadajip/stripe-react-go/server/domain/user"
	"github.com/fuadajip/stripe-react-go/server/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type repoHandler struct {
	mysqlSess *gorm.DB
}

func NewUserRepository(mysqlSess *gorm.DB) user.Repository {
	return &repoHandler{
		mysqlSess: mysqlSess,
	}
}

func (r repoHandler) FindUserByEmailPhone(c echo.Context, payload *models.User) (*models.User, error) {
	res := &models.User{}
	db := r.mysqlSess.Where("email = ? OR phone = ?", payload.Email, payload.Phone).First(&res)
	return res, db.Error
}

func (r repoHandler) FindUserByUsername(c echo.Context, payload *models.User) (*models.User, error) {
	res := &models.User{}
	db := r.mysqlSess.Where("username = ?", payload.Username).First(&res)
	return res, db.Error
}

func (r repoHandler) TrxUserRegistration(trx *gorm.DB, payload *models.User) (*models.User, error) {
	db := trx.Create(&payload)
	return payload, db.Error

}
