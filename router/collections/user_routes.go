package collections

import (
	"net/http"
	"rest/config/initial"

	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", initial.GetControllers().UserController.Create).Methods(http.MethodPost)
	router.HandleFunc("/users", initial.GetControllers().UserController.Index).Methods(http.MethodGet)
	router.HandleFunc("/profile", initial.GetControllers().UserController.Update).Methods(http.MethodPut)
	router.HandleFunc("/users/{id:[0-9]+}", initial.GetControllers().UserController.Show).Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}", initial.GetControllers().UserController.Delete).Methods(http.MethodDelete)

	return router
}
