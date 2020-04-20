package repository

import (
	"github.com/fuadajip/stripe-react-go/server/domain/healthz"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type repoHandler struct {
	redisSess *redis.Client
	mysqlSess *gorm.DB
}

// NewHealthCheckRepository returns implementation of methods in auth.Repository
func NewHealthCheckRepository(redisSess *redis.Client, mysqlSess *gorm.DB) healthz.Repository {
	return &repoHandler{
		redisSess: redisSess,
		mysqlSess: mysqlSess,
	}
}

// RedisHealthCheck is method that implement healthz.Repository
func (r repoHandler) RedisHealthCheck() (bool, error) {
	_, err := r.redisSess.Ping().Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

// MysqlHealthCheck is method that implement healthz.Repository
func (r repoHandler) MysqlHealthCheck() (bool, error) {

	err := r.mysqlSess.DB().Ping()
	if err != nil {
		return false, err
	}

	return true, nil

}
