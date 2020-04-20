package userprofile

import (
	"github.com/fuadajip/stripe-react-go/server/models"
	"github.com/jinzhu/gorm"
)

type Usecase interface {
	TrxInsertUserProfile(trx *gorm.DB, payload *models.UserProfileRequest) (*models.UserProfile, error)
}

type Repository interface {
	TrxInsertUserProfile(trx *gorm.DB, payload *models.UserProfile) (*models.UserProfile, error)
}
