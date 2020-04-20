package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/fuadajip/stripe-react-go/server/shared/log"

	Error "github.com/fuadajip/stripe-react-go/server/shared/error"
	"github.com/spf13/viper"
)

var (
	logger = log.NewLog()
)

type (
	// ImmutableConfigInterface is an interface represent methods in config
	ImmutableConfigInterface interface {
		GetPort() int
		GetDatabaseHost() string
		GetDatabasePort() string
		GetDatabaseName() string
		GetDatabaseUser() string
		GetDatabasePassword() string
		GetRedisHost() string
		GetRedisPassword() string
		GetSecret() string
		GetStripeSecret() string
	}

	// im is a struct to mapping the structure of related value model
	im struct {
		Port             int    `mapstructure:"PORT"`
		DatabaseHost     string `mapstructure:"DATABASE_HOST"`
		DatabasePort     string `mapstructure:"DATABASE_PORT"`
		DatabaseName     string `mapstructure:"DATABASE_NAME"`
		DatabaseUser     string `mapstructure:"DATABASE_USER"`
		DatabasePassword string `mapstructure:"DATABASE_PASS"`
		RedisHost        string `mapstructure:"REDIS_HOST"`
		RedisPassword    string `mapstructure:"REDIS_PASS"`
		Secret           string `mapstructure:"SECRET"`
		StripeSecret     string `mapstructure:"STRIPE_SECRET"`
	}
)

func (i *im) GetPort() int {
	return i.Port
}

func (i *im) GetDatabaseHost() string {
	return i.DatabaseHost
}

func (i *im) GetDatabasePort() string {
	return i.DatabasePort
}

func (i *im) GetDatabaseName() string {
	return i.DatabaseName
}

func (i *im) GetDatabaseUser() string {
	return i.DatabaseUser
}

func (i *im) GetDatabasePassword() string {
	return i.DatabasePassword
}

func (i *im) GetRedisHost() string {
	return i.RedisHost
}

func (i *im) GetRedisPassword() string {
	return i.RedisPassword
}

func (i *im) GetSecret() string {
	return i.Secret
}

func (i *im) GetStripeSecret() string {
	return i.StripeSecret
}

var (
	imOnce    sync.Once
	myEnv     map[string]string
	immutable im
)

// NewImmutableConfig is a factory that return of its config implementation
func NewImmutableConfig() ImmutableConfigInterface {
	imOnce.Do(func() {

		v := viper.New()
		appEnv, exists := os.LookupEnv("APP_ENV")

		// TODO: revamp this
		if exists {
			if appEnv == "staging" {
				v.SetConfigName("app.config.staging")
				logger.Info(fmt.Sprintf("Reading app_env: %s", appEnv))
			} else if appEnv == "production" {
				v.SetConfigName("app.config.prod")
				logger.Info(fmt.Sprintf("Reading app_env: %s", appEnv))
			} else if appEnv == "development" {
				v.SetConfigName("app.config.dev")
				logger.Info(fmt.Sprintf("Reading app_env: %s", appEnv))
			}
		} else {
			v.SetConfigName("app.config.dev")
			logger.Info(fmt.Sprintf("Reading app_env: %s", "development"))
		}

		v.AddConfigPath(".")
		v.SetEnvPrefix("EXP")
		v.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			Error.Wrap(500, "[EXP-SYS-001]", err, "[CONFIG][missing] Failed to read app.config.* file", "failed")
		}

		v.Unmarshal(&immutable)
	})

	return &immutable
}
