package middleware

import (
	"rest/database/redis"
	"rest/router/lang"
	"net/http"
)

var languages = []string{"ru", "en"}

func Translate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accLang := r.Header.Get("content-language")
		if accLang != "" && findAvailableLang(accLang) {
			lang.GetTranslator().SetDefaultLocale(accLang)
			rConf := new(redis.Config)
			redisConn := redis.ConnectToRedis(rConf)
			_, err := redisConn.SetValue("locale", accLang, 0)
			if err != nil {
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func findAvailableLang(lang string) bool {
	found := false
	for _, value := range languages {
		if value == lang {
			found = true
		}
	}
	return found
}
