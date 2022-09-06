package responses

import (
	"encoding/json"
	"net/http"
)

type ResponseWithPagination interface {
	any
}

func GenerateErrorResponse(w http.ResponseWriter, message string, errors map[string]interface{}, code ...int) {
	var statusCode int
	response := ResponseSingle{
		Data:   message,
		Errors: errors,
	}

	if code == nil {
		statusCode = 422
	} else {
		statusCode = code[0]
	}

	err := CreateResponse(w, response, statusCode)
	if err != nil {
		return
	}
}

func CreateResponse(w http.ResponseWriter, response any, status int) error {
	responseToSend, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, resErr := w.Write(responseToSend)
	if resErr != nil {
		return resErr
	}
	return err
}
