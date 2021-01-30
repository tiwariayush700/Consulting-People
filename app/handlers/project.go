package handlers

import (
	"coffeebeans-people-backend/models"
	"coffeebeans-people-backend/utility"
	"context"
	"encoding/json"
	"net/http"
)

func CreateProject(apiSvc models.ApiSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var project models.Project
		err := json.NewDecoder(r.Body).Decode(&project)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   err.Error(),
				Message: "Error decoding request body",
			}, http.StatusBadRequest)
			return
		}

		err = apiSvc.CreateProjectByAdmin(context.TODO(), project)
		if err != nil {
			utility.NewJSONWriter(w).Write(models.Response{
				Error:   "Mongo error",
				Message: "Project Id already exists. Check you project id",
			}, http.StatusBadRequest)
			return
		}

		utility.NewJSONWriter(w).Write(models.Response{
			Error:   "",
			Message: "Project Created Successsfully",
		}, http.StatusOK)
	}
}
