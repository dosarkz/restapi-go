package responses

import "net/http"

func ParamsMapFromRequest(r *http.Request) map[string]string {
	params := make(map[string]string)
	values := r.URL.Query()
	for k := range values {
		params[k] = values.Get(k)
	}
	return params
}
