package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

var totalRequest = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "Number of get request",
	},
	[]string{"path"},
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		totalRequest.With(prometheus.Labels{"path": r.RequestURI}).Inc()
	})
}

func init() {
	prometheus.Register(totalRequest)
}
