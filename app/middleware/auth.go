package middleware

import (
	"context"
	"golang.org/x/exp/slices"
	"net/http"
	"rest/app/helpers/auth"
	"rest/database/redis"
	"rest/domain/user/models"
	"rest/domain/user/repositories"
)

var UserCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Auth(repo repositories.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearer := r.Header.Get("Authorization")
			whitelist := []string{"/api/v1/auth/login"}
			if bearer == "" {
				if slices.Contains(whitelist, r.URL.Path) {
					next.ServeHTTP(w, r)
					return
				}
				http.Error(w, "token not provided", http.StatusUnauthorized)
				return
			}

			claims, err := auth.ValidateRequest(r)
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			if IsInBlacklist(bearer) {
				http.Error(w, "Token expired", http.StatusUnauthorized)
				return
			}

			u, showErr := repo.FindById(claims.UserID)
			if showErr != nil {
				http.Error(w, "user not found", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), UserCtxKey, u)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func IsInBlacklist(token string) bool {
	rConf := new(redis.Config)
	redisConn := redis.ConnectToRedis(rConf)
	redisToken, _ := redisConn.GetValue(token)
	return redisToken != ""
}

func CtxValue(ctx context.Context) *models.User {
	raw, _ := ctx.Value(UserCtxKey).(*models.User)
	return raw
}
