package healthz

import (
	"github.com/labstack/echo"
)

// Usecase is an interface of healthcheck domain that implement Healthcheck's Usecase methods
type Usecase interface {
	DoHealthCheck(c echo.Context) (bool, error)
}

// Repository is an interface of healthcheck domain that implement Healthcheck's repository methods
type Repository interface {
	RedisHealthCheck() (bool, error)
	MysqlHealthCheck() (bool, error)
}
