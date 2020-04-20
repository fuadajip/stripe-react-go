package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.elastic.co/apm/module/apmecho"
	"gopkg.in/go-playground/validator.v9"

	Config "github.com/fuadajip/stripe-react-go/server/shared/config"
	Container "github.com/fuadajip/stripe-react-go/server/shared/container"
	Database "github.com/fuadajip/stripe-react-go/server/shared/database"
	Logger "github.com/fuadajip/stripe-react-go/server/shared/log"
	Util "github.com/fuadajip/stripe-react-go/server/shared/util"

	healthzHandler "github.com/fuadajip/stripe-react-go/server/domain/healthz/delivery/http"
	healthzRepository "github.com/fuadajip/stripe-react-go/server/domain/healthz/repository"
	healthzUsecase "github.com/fuadajip/stripe-react-go/server/domain/healthz/usecase"

	userHandler "github.com/fuadajip/stripe-react-go/server/domain/user/delivery/http"
	userRepository "github.com/fuadajip/stripe-react-go/server/domain/user/repository"
	userUsecase "github.com/fuadajip/stripe-react-go/server/domain/user/usecase"

	userprofileRepository "github.com/fuadajip/stripe-react-go/server/domain/userprofile/repository"
	userprofileUsecase "github.com/fuadajip/stripe-react-go/server/domain/userprofile/usecase"
)

var (
	logger = Logger.NewLog()
)

func main() {
	e := echo.New()
	container := Container.DefaultContainer()
	conf := container.MustGet("shared.config").(Config.ImmutableConfigInterface)
	redis := container.MustGet("shared.redis").(Database.RedisInterface)
	mysql := container.MustGet("shared.mysql").(Database.MysqlInterface)

	redisSess, err := redis.OpenRedisConn()
	if err != nil {
		msgError := fmt.Sprintf("Failed to open redis connection, err: %s", err.Error())
		logger.Errorf(msgError)
	}

	mysqlSess, err := mysql.OpenMysqlConn()
	if err != nil {
		msgError := fmt.Sprintf("Failed to open mysql connection, err: %s", err.Error())
		logger.Errorf(msgError)
		panic(msgError)
	}

	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())

	// provides protection against cross-site scripting (XSS) attack, content type sniffing,
	// clickjacking, insecure connection and other code injection attacks.
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	// The middleware will recover panics and send them to Elastic APM
	// so you do not need to install the echo/middleware.Recover middleware.
	e.Use(apmecho.Middleware())

	e.Validator = &Util.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = func(err error, e echo.Context) {
		Util.CustomHTTPErrorHandler(err, e)
	}

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := &Util.CustomApplicationContext{
				Context:      c,
				Container:    container,
				SharedConf:   conf,
				RedisSession: redisSess,
				MysqlSession: mysqlSess,
			}
			return h(ac)
		}
	})

	healthzRepo := healthzRepository.NewHealthCheckRepository(redisSess, mysqlSess)
	userRepo := userRepository.NewUserRepository(mysqlSess)
	userprofileRepo := userprofileRepository.NewUserProfileRepository(mysqlSess)

	healthzUcase := healthzUsecase.NewHealthCheckUsecase(healthzRepo)
	userUcase := userUsecase.NewUserUsecase(userRepo, userprofileRepo)
	_ = userprofileUsecase.NewUserProfileUsecase(userprofileRepo)

	// delivery handler implementation
	healthzHandler.AddHealthCheckHandler(e, healthzUcase)
	userHandler.AddUserHandler(e, userUcase)

	e.Logger.Info(e.Start(fmt.Sprintf(":%d", conf.GetPort())))
}
