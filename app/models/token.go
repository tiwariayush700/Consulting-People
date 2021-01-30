package models

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTTokenClaims struct {
	User *User
	jwt.StandardClaims
}
