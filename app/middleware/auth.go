package middleware

import (
	"context"
	"github.com/casbin/casbin"
	"golang.org/x/exp/slices"
	"rest/app/helpers/auth"
	"rest/database/redis"
	"rest/domain/user/models"
	"rest/domain/user/repositories"
	"net/http"
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

			u, err := repo.Show(claims.UserID)
			if err != nil {
				http.Error(w, "user not found", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), UserCtxKey, u)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func RBAC(e *casbin.Enforcer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			user := CtxValue(r.Context())
			var role string
			if user == nil {
				role = "anonymous"
			} else {
				role = user.Role.Name
			}
			// casbin rule enforcing
			res, err := e.EnforceSafe(role, r.URL.Path, r.Method)
			if err != nil {
				//	fmt.Println(err)
				http.Error(w, "authorization error occurred", http.StatusInternalServerError)
				return
			}
			if res {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "current user has no access to requested resource", http.StatusForbidden)
				return
			}
		}

		return http.HandlerFunc(fn)
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
