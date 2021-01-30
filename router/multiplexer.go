package router

import (
	"coffeebeans-people-backend/api"
	"coffeebeans-people-backend/auth"
	"coffeebeans-people-backend/dao"
	"coffeebeans-people-backend/handlers"
	"coffeebeans-people-backend/middleware"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"net/http"
)

// API is an API configuration.
type API struct {
	DaoService  *dao.Service
	ApiService  api.ApiSvc
	AuthService auth.AuthSvc
}

// APIMux returns an API multiplexer.
func APIMux(api *API) *chi.Mux {
	mux := chi.NewMux()
	corsSettings := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Accept",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"X-XSRF-Token",
			"X-HTTP-Method-Override",
			"X-Requested-With",
		},
		AllowCredentials: true,
		MaxAge:           86400,
	})

	mux.Use(corsSettings.Handler)

	mux.Route("/ping", func(r chi.Router) {
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("pong"))
		})
	})

	mux.Route("/register", func(r chi.Router) {
		r.Post("/", handlers.CreateUser(&api.ApiService))
	})

	mux.Route("/login", func(r chi.Router) {
		r.Post("/", handlers.Login(&api.ApiService, api.AuthService))
	})

	mux.Route("/edit", func(r chi.Router) {
		r.Use(middleware.AuthenticateTokenMiddlewareHandler(api.AuthService))
		r.Post("/profile", handlers.UpdateProfile(&api.ApiService))
	})

	mux.Route("/tokenCheck", func(r chi.Router) {
		r.Use(middleware.AuthenticateTokenMiddlewareHandler(api.AuthService))
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("token correct"))
		})
	})

	mux.Route("/logout", func(r chi.Router) {
		r.Use(middleware.AuthenticateTokenMiddlewareHandler(api.AuthService))
		r.Post("/", handlers.Logout())
	})

	mux.Route("/create", func(r chi.Router) {
		r.Use(middleware.AuthenticateTokenMiddlewareHandler(api.AuthService))
		r.Post("/project", handlers.CreateProject(&api.ApiService))
	})

	mux.Route("/users", func(r chi.Router) {
		r.Use(middleware.AuthenticateTokenMiddlewareHandler(api.AuthService))
		r.Get("/", handlers.GetUsers(&api.ApiService))
	})

	return mux
}
