package responses

import (
	"net/http"
)

var perPage = 10

type MetaData struct {
	PerPage  int `json:"per_page"`
	LastPage int `json:"last_page"`
	Total    int `json:"total"`
}

type ResponseList struct {
	Meta MetaData `json:"meta"`
	Data any      `json:"data"`
}

func ResponseCollection[T ResponseWithPagination](w http.ResponseWriter, s []T) {
	meta := MetaData{
		PerPage:  perPage,
		LastPage: len(s)/perPage + 1,
		Total:    len(s),
	}
	response := ResponseList{
		Data: s,
		Meta: meta,
	}
	if len(s) == 0 {
		empty := make([]string, 0)
		response.Data = empty
	}
	err := CreateResponse(w, response, http.StatusOK)
	if err != nil {
		return
	}
}
