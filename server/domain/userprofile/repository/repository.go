package repository

import (
	"github.com/fuadajip/stripe-react-go/server/domain/userprofile"
	"github.com/fuadajip/stripe-react-go/server/models"
	"github.com/jinzhu/gorm"
)

type repoHandler struct {
	mysqlSess *gorm.DB
}

func NewUserProfileRepository(mysqlSess *gorm.DB) userprofile.Repository {
	return &repoHandler{
		mysqlSess: mysqlSess,
	}
}

func (r repoHandler) TrxInsertUserProfile(trx *gorm.DB, payload *models.UserProfile) (*models.UserProfile, error) {
	db := trx.Create(&payload)
	return payload, db.Error
}
