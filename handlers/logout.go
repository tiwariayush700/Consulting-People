package handlers

import (
	"coffeebeans-people-backend/constants"
	"coffeebeans-people-backend/models"
	"coffeebeans-people-backend/utility"
	"net/http"
)

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userFromContext := r.Context().Value(constants.USER_KEY).(models.User)
		if len(userFromContext.Email) < 1 {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   "Unauthorized",
				Message: "Invalid token",
			}, http.StatusUnauthorized)
			return
		}

		utility.NewJSONWriter(w).Write([]byte("Logout successful"), http.StatusOK)
	}
}
