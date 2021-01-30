package main

import (
	"coffeebeans-people-backend/api"
	"coffeebeans-people-backend/auth"
	"coffeebeans-people-backend/config"
	"coffeebeans-people-backend/constants"
	"coffeebeans-people-backend/dao"
	"coffeebeans-people-backend/router"
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/log"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	appConfig := config.GetAppConfiguration()
	mux := chi.NewMux()

	uri := constants.MONGO_BASE_URL + appConfig.MONGO_SERVER
	daoSvc, err := dao.NewService(ctx, uri, appConfig.MONGO_DATABASE)
	if err != nil {
		log.Fatal("unable to create mongo conn: ", err.Error())
	}

	apiSvc := api.ApiSvc{DbSvc: daoSvc}

	authSvc, err := auth.NewService(ctx, appConfig.SECRET_KEY)
	if err != nil {
		log.Error(err)
	}

	apiService := &router.API{
		DaoService:  daoSvc,
		ApiService:  apiSvc,
		AuthService: authSvc,
	}

	mux.Mount(appConfig.BASE_URL, router.APIMux(apiService))
	ListenAndServe(appConfig.PORT, mux)
}

// ListenAndServe runs the server.
func ListenAndServe(port string, handler http.Handler) {
	fmt.Println("Listening on:", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler); err != nil {
		panic(err)
	}
}
