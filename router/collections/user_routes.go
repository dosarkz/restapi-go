package collections

import (
	"rest/config/initial"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/user", initial.GetControllers().UserController.Create).Methods(http.MethodPost)
	router.HandleFunc("/user", initial.GetControllers().UserController.List).Methods(http.MethodGet)
	router.HandleFunc("/auth/profile", initial.GetControllers().UserController.Update).Methods(http.MethodPut)
	router.HandleFunc("/user/{id:[0-9]+}", initial.GetControllers().UserController.Show).Methods(http.MethodGet)
	router.HandleFunc("/user/{id:[0-9]+}", initial.GetControllers().UserController.Delete).Methods(http.MethodDelete)

	return router
}
