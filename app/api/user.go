package api

import (
	"coffeebeans-people-backend/models"
	"context"
	"encoding/json"
	"errors"
	"log"
	"reflect"
)

type ApiSvc struct {
	DbSvc models.Dao
}

func (apiSvc *ApiSvc) RegisterUser(ctx context.Context, user models.User) error {

	users, err := apiSvc.DbSvc.GetAllUsers(ctx, nil)
	if err != nil {
		log.Println("Error querying mongo", err)
		return err
	}

	isUnique := checkUniqueEmail(users, user.Email)
	if !isUnique {
		return errors.New("Email already exists try with new email")
	}

	err = apiSvc.DbSvc.CreateUser(ctx, user)
	if err != nil {
		log.Println("Error querying mongo", err)
		return err
	}

	return nil
}

func checkUniqueEmail(users []models.User, email string) bool {
	for _, user := range users {
		if user.Email == email {
			return false
		}
	}

	return true
}

func (apiSvc *ApiSvc) LoginUser(ctx context.Context, email string, password string) (models.User, bool, error) {
	var isProfileComplete bool
	user, err := apiSvc.DbSvc.GetUserByCredentials(ctx, email, password)
	if err != nil {
		return user, isProfileComplete, err
	}

	isProfileComplete = isProfileCompleted(user)

	return user, isProfileComplete, nil
}

func isProfileCompleted(user models.User) bool {
	var userMandatoryFields models.UserMandatoryFields

	marshalledData, _ := json.Marshal(&user)

	json.Unmarshal(marshalledData, &userMandatoryFields)

	return !reflect.DeepEqual(userMandatoryFields, models.UserMandatoryFields{})

}

func (apiSvc *ApiSvc) GetUsers(ctx context.Context, params map[string]interface{}) ([]models.User, error) {
	users, err := apiSvc.DbSvc.GetAllUsers(ctx, params)
	if err != nil {
		return users, err
	}

	return users, nil
}
