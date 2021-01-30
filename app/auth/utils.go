package auth

import (
	"coffeebeans-people-backend/constants"
	"github.com/dgrijalva/jwt-go"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/log"
)

func InvalidTokenErrorMsgs(jwtToken *jwt.Token, err error) (bool, string, error) {
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			log.Errorf(constants.INVALID_TOKEN_MSG)
			return false, constants.INVALID_TOKEN_MSG, err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			log.Errorf(constants.INVALID_TOKEN_MSG_TOKEN_EXPIRED)
			return false, constants.INVALID_TOKEN_MSG_TOKEN_EXPIRED, err
		} else {
			log.Errorf(constants.INVALID_TOKEN_MSG)
			return false, constants.INVALID_TOKEN_MSG, err
		}
	}

	return false, constants.INVALID_TOKEN_MSG, err
}
