package router

import (
	"net/http"
	"rest/app/middleware"
	"rest/config/initial"
	"rest/router/collections"

	"github.com/gorilla/mux"
)

var router Router

type Router struct {
	driver *mux.Router
}

func Init() {
	router.driver = mux.NewRouter()
	c := initial.GetControllers()
	//registerLogger()

	// general routes
	GetRouter().HandleFunc("/", c.MainController.Index)
	GetRouter().NotFoundHandler = http.HandlerFunc(c.MainController.NotFound)
	GetRouter().Use(middleware.Translate)
	GetRouter().PathPrefix("/storage/").Handler(http.StripPrefix("/storage/", http.FileServer(http.Dir("./storage/"))))

	api := GetRouter().PathPrefix("/api/v1").Subrouter()

	// middlewares
	api.Use(middleware.LimitRequestsByIP)
	api.Use(middleware.Auth(initial.GetRepositories().UserRepo))

	api = collections.SetAuthenticationRoutes(api)
	api = collections.SetUserRoutes(api)
}

func registerLogger() {
	//	GetRouter().Use(middleware.PrometheusMiddleware)
	//	GetRouter().Handle("/metrics", promhttp.Handler())
}

func GetRouter() *mux.Router {
	return router.driver
}
