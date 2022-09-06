package responses

import "net/http"

type ResponseSingle struct {
	Data   any                    `json:"data"`
	Errors map[string]interface{} `json:"errors"`
}

func ResponseResource[T ResponseWithPagination](w http.ResponseWriter, s T) {
	response := ResponseSingle{
		Data: s,
	}
	err := CreateResponse(w, response, http.StatusOK)
	if err != nil {
		return
	}
}
