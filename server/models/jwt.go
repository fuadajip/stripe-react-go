package models

import (
	"github.com/dgrijalva/jwt-go"
)

type ClaimUserLoginJWT struct {
	Email *string `json:"email"`
	jwt.StandardClaims
}
