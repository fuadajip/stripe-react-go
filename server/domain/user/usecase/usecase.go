package usecase

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/fuadajip/stripe-react-go/server/domain/user"
	"github.com/fuadajip/stripe-react-go/server/domain/userprofile"
	"github.com/fuadajip/stripe-react-go/server/models"
	Error "github.com/fuadajip/stripe-react-go/server/shared/error"
	"github.com/fuadajip/stripe-react-go/server/shared/log"
	"github.com/fuadajip/stripe-react-go/server/shared/util"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var (
	logger = log.NewLog()
)

type usecase struct {
	repository      user.Repository
	userprofileRepo userprofile.Repository
}

func NewUserUsecase(
	repository user.Repository,
	userprofileRepo userprofile.Repository,
) user.Usecase {
	return &usecase{
		repository:      repository,
		userprofileRepo: userprofileRepo,
	}
}

func (u usecase) UserRegistration(c echo.Context, payload *models.UserRegistrationRequest) (*models.UserRegistrationResponse, error) {

	ac := c.(*util.CustomApplicationContext)

	userPayload := &models.User{
		Email:    payload.Email,
		IsActive: aws.Bool(true),
		Phone:    payload.Phone,
		Type:     aws.String("GENERAL"),
		Username: payload.Username,
	}

	username, err := u.repository.FindUserByUsername(c, userPayload)
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
	}

	if username.Username != nil {
		return nil, Error.New("EXIST_USERNAME")
	}

	emailPhoneResp, err := u.repository.FindUserByEmailPhone(c, userPayload)
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
	}

	if emailPhoneResp.Email != nil || emailPhoneResp.Phone != nil {
		return nil, Error.New("EXIST_EMAIL_PHONE")
	}

	hashedPassowrd, err := util.GenerateHashPassword(payload.Password)
	if err != nil {
		return nil, err
	}
	userPayload.Password = hashedPassowrd

	trx := ac.MysqlSession.Begin()
	userResp, err := u.repository.TrxUserRegistration(trx, userPayload)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	userProfilePayload := &models.UserProfile{
		UserID:   userResp.ID,
		Fullname: payload.Fullname,
	}
	profileResp, err := u.userprofileRepo.TrxInsertUserProfile(trx, userProfilePayload)
	if err != nil {
		trx.Rollback()
		return nil, err
	}
	trx.Commit()

	mappedResp := &models.UserRegistrationResponse{
		Email:    userResp.Email,
		Fullname: profileResp.Fullname,
		Phone:    userResp.Phone,
		Username: userResp.Username,
	}

	return mappedResp, nil
}

func (u usecase) UserLogin(c echo.Context, payload *models.UserLoginRequest) (*models.UserLoginResponse, error) {
	userPayload := &models.User{
		Email: payload.Email,
	}
	emailPhoneResp, err := u.repository.FindUserByEmailPhone(c, userPayload)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, Error.New("USER_NOTFOUND")
		}
		return nil, err
	}

	isValidPass := util.CompareHashPassword(payload.Password, emailPhoneResp.Password)
	if !isValidPass {
		return nil, Error.New("INVALID_CREDENTIALS")
	}

	claimJWT := &models.ClaimUserLoginJWT{
		Email: emailPhoneResp.Email,
	}

	token, err := util.CreateUserJWT(c, claimJWT)
	if err != nil {
		return nil, err
	}

	mappedResp := &models.UserLoginResponse{
		Email: emailPhoneResp.Email,
		Token: token,
	}

	return mappedResp, nil
}
