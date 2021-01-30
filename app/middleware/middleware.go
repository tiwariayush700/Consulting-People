package middleware

import (
	"coffeebeans-people-backend/auth"
	"coffeebeans-people-backend/constants"
	"coffeebeans-people-backend/models"
	"coffeebeans-people-backend/utility"
	"context"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/log"
	"net/http"
)

func AuthenticateTokenMiddlewareHandler(authSvc auth.AuthSvc) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			user, msg, err := authSvc.AuthenticateToken(token)
			if err != nil {
				log.Errorf(err.Error())
				utility.NewJSONWriter(w).Write(models.Response{
					Error:   err.Error(),
					Message: msg,
				}, http.StatusUnauthorized)
				return
			}

			ctx := ContextWithUser(r.Context(), *user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// ContextWithUser returns a context with a Employee.
func ContextWithUser(ctx context.Context, user models.User) context.Context {
	return context.WithValue(ctx, constants.USER_KEY, user)
}
