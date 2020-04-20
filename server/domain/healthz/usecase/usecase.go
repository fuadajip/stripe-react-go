package usecase

import (
	"github.com/fuadajip/stripe-react-go/server/domain/healthz"
	"github.com/labstack/echo"
)

type usecase struct {
	repository healthz.Repository
}

// NewHealthCheckUsecase is a factory that return implementation of methods in healthz.Usecase interface
func NewHealthCheckUsecase(repository healthz.Repository) healthz.Usecase {
	return &usecase{
		repository: repository,
	}
}

// DoHealthCheck is an method that implement healthz.Usecase
func (u usecase) DoHealthCheck(c echo.Context) (bool, error) {

	_, err := u.repository.RedisHealthCheck()
	if err != nil {
		return false, err
	}

	_, err = u.repository.MysqlHealthCheck()
	if err != nil {
		return false, err
	}

	return true, nil
}
