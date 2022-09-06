package collections

import (
	"github.com/gorilla/mux"
	"rest/config/initial"
	"net/http"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/login", initial.GetControllers().AuthController.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/refresh-token", initial.GetControllers().AuthController.RefreshToken).Methods(http.MethodGet)
	router.HandleFunc("/auth/logout", initial.GetControllers().AuthController.Logout).Methods(http.MethodGet)
	router.HandleFunc("/auth/profile", initial.GetControllers().UserController.Profile).Methods(http.MethodGet)
	return router
}
