package usecase

import (
	"github.com/fuadajip/stripe-react-go/server/domain/userprofile"
	"github.com/fuadajip/stripe-react-go/server/models"
	"github.com/fuadajip/stripe-react-go/server/shared/log"
	"github.com/jinzhu/gorm"
)

var (
	logger = log.NewLog()
)

type usecase struct {
	repository userprofile.Repository
}

func NewUserProfileUsecase(
	repository userprofile.Repository,
) userprofile.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u usecase) TrxInsertUserProfile(trx *gorm.DB, payload *models.UserProfileRequest) (*models.UserProfile, error) {

	profilePayload := &models.UserProfile{
		Avatar:   payload.Avatar,
		Bio:      payload.Bio,
		Fullname: payload.Fullname,
		Gender:   payload.Gender,
		UserID:   payload.UserID,
	}
	resp, err := u.repository.TrxInsertUserProfile(trx, profilePayload)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
