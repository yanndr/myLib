package middlewares

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
)

var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "myLib_http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})

	count = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myLib_http_count",
		Help: "Count of HTTP requests",
	}, []string{"path"})
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		label := fmt.Sprintf("%s-%s", r.Method, r.RequestURI)
		count.WithLabelValues(label).Add(1)
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(label))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}
