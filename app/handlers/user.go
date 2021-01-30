package handlers

import (
	"coffeebeans-people-backend/constants"
	"coffeebeans-people-backend/models"
	"coffeebeans-people-backend/utility"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateUser(apiSvc models.ApiSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   err.Error(),
				Message: "Error decoding request body",
			}, http.StatusBadRequest)
			return
		}

		user.Role = "ADMIN" //TODO ACCEPT FROM UI
		err = apiSvc.RegisterUser(context.TODO(), user)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   "Mongo error",
				Message: "Employee Id or email already exists",
			}, http.StatusBadRequest)
			return
		}

		utility.NewJSONWriter(w).Write(models.Response{
			Error:   "",
			Message: "User Created Successsfully",
		}, http.StatusOK)
	}
}

func GetUsers(apiSvc models.ApiSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userFromContext := r.Context().Value(constants.USER_KEY).(models.User)
		if len(userFromContext.Email) < 1 {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   "Unauthorized",
				Message: "Incorrect access token",
			}, http.StatusBadRequest)
			return
		}
		params := make(map[string]interface{})
		skill := r.FormValue("skill")
		id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)

		if len(skill) > 0 {
			params["skill"] = skill
		}

		if len(r.FormValue("id")) > 0 {
			params["id"] = id
		}

		users, err := apiSvc.GetUsers(r.Context(), params)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   err.Error(),
				Message: "Mongo error",
			}, http.StatusBadRequest)
		}

		utility.NewJSONWriter(w).Write(users, http.StatusOK)
	}
}
