package handlers

import (
	"coffeebeans-people-backend/constants"
	"coffeebeans-people-backend/models"
	"coffeebeans-people-backend/utility"
	"encoding/json"
	"net/http"
)

func UpdateProfile(apiSvc models.ApiSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userFromContext := r.Context().Value(constants.USER_KEY).(models.User)
		if len(userFromContext.Email) < 1 {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   "Unauthorized",
				Message: "Invalid token",
			}, http.StatusUnauthorized)
			return
		}

		var body models.User
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   err.Error(),
				Message: "Error decoding request body",
			}, http.StatusBadRequest)
			return
		}

		err = apiSvc.EditUser(r.Context(), body)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   err.Error(),
				Message: "Error querying mongo",
			}, http.StatusBadRequest)
			return
		}

		utility.NewJSONWriter(w).Write(models.Response{
			Error:   "",
			Message: "Profile updated successfully :)",
		}, http.StatusOK)
	}
}
