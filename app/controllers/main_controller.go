package controllers

import (
	"fmt"
	"net/http"
	"rest/app/helpers/responses"
	"rest/router/lang"
)

type Controller struct {
}

func (*Controller) Index(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Welcome to rest API")
	if err != nil {
		return
	}
}

func (*Controller) NotFound(w http.ResponseWriter, _ *http.Request) {
	text := lang.GetTranslator().Trans("messages.resource_not_found")
	responses.GenerateErrorResponse(w, text, nil, 404)
}
