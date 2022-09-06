package router

import (
	"github.com/casbin/casbin"
	"rest/app/middleware"
	"rest/config/initial"
	"rest/router/collections"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router Router

type Router struct {
	driver *mux.Router
}

func Init() {
	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Printf("read policy file: %v", err)
	}

	router.driver = mux.NewRouter()
	c := initial.GetControllers()
	// General routing
	//	GetRouter().Use(middleware.PrometheusMiddleware)
	//	GetRouter().Handle("/metrics", promhttp.Handler())

	GetRouter().HandleFunc("/", c.MainController.Index)
	GetRouter().NotFoundHandler = http.HandlerFunc(c.MainController.NotFound)
	GetRouter().Use(middleware.Translate)
	router.driver.PathPrefix("/storage/").Handler(http.StripPrefix("/storage/", http.FileServer(http.Dir("./storage/"))))
	api := GetRouter().PathPrefix("/api/v1").Subrouter()

	api.Use(middleware.LimitRequestsByIP)
	api.Use(middleware.Auth(initial.GetRepositories().UserRepo))
	if authEnforcer != nil {
		api.Use(middleware.RBAC(authEnforcer))
	}

	api = collections.SetAuthenticationRoutes(api)
	api = collections.SetUserRoutes(api)
}
func GetRouter() *mux.Router {
	return router.driver
}
