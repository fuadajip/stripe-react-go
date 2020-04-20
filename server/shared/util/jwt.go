package util

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/dgrijalva/jwt-go"
	"github.com/fuadajip/stripe-react-go/server/models"
	"github.com/labstack/echo"
)

// CreateUserJWT will generate custom jwt key to create session
func CreateUserJWT(c echo.Context, claims *models.ClaimUserLoginJWT) (*string, error) {
	ac := c.(*CustomApplicationContext)
	conf := ac.SharedConf
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString([]byte(conf.GetSecret()))
	if err != nil {
		return nil, err
	}
	return aws.String(tokenString), nil

}
