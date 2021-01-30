package auth

import (
	"coffeebeans-people-backend/constants"
	"coffeebeans-people-backend/models"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func (svc *Service) GenerateToken(user *models.User) (string, error) {

	claims := models.JWTTokenClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "coffeebeans-people",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(svc.SecretKey))

	return tokenString, err
}

func (svc *Service) AuthenticateToken(jwtTokenString string) (*models.User, string, error) {
	if len(jwtTokenString) == 0 {
		return nil, constants.EMPTY_TOKEN_MSG, errors.New("forbidden")
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtTokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(svc.SecretKey), nil
	})

	if !token.Valid {
		_, msg, err := InvalidTokenErrorMsgs(token, err)
		return nil, msg, err
	}

	user := GetUserFromTokenClaims(claims)

	return &user, constants.SUCCESS_TOKEN, err
}

func GetUserFromTokenClaims(claims map[string]interface{}) models.User {
	var user models.User
	employeeMapFromToken := claims["User"].(map[string]interface{})

	employeeEmailFromToken := fmt.Sprintf("%v", employeeMapFromToken["email"])
	employeeNameFromToken := fmt.Sprintf("%v", employeeMapFromToken["name"])
	employeeIDFromToken := int64(employeeMapFromToken["employee_id"].(float64))
	employeeRoleFromToken := fmt.Sprintf("%v", employeeMapFromToken["role"])

	user.EmployeeId = employeeIDFromToken
	user.Email = employeeEmailFromToken
	user.Name = employeeNameFromToken
	user.Role = employeeRoleFromToken

	return user
}
