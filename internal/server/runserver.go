package server

import (
	"fmt"
	"groupie-tracker/internal/handler"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "myapp_http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path", "method", "status_code"})
	http500Errors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_http_500_errors_total",
			Help: "Total number of HTTP 500 errors.",
		},
		[]string{"path", "method"},
	)
)

func init() {
	prometheus.MustRegister(httpDuration)
	prometheus.MustRegister(http500Errors)
}

func recordMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(r.URL.Path, r.Method, ""))
		defer timer.ObserveDuration()

		next.ServeHTTP(rw, r)

		status := rw.statusCode
		httpDuration.With(prometheus.Labels{"path": r.URL.Path, "method": r.Method, "status_code": fmt.Sprintf("%d", status)}).Observe(timer.ObserveDuration().Seconds())

		if status == http.StatusInternalServerError {
			http500Errors.With(prometheus.Labels{"path": r.URL.Path, "method": r.Method}).Inc()
		}
	})
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

func Runserver() {
	mux := http.NewServeMux()

	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/js"))))

	mux.Handle("/", recordMetrics(http.HandlerFunc(handler.IndexHandler)))
	mux.Handle("/groups/", recordMetrics(http.HandlerFunc(handler.GroupHandler)))
	mux.Handle("/filter", recordMetrics(http.HandlerFunc(handler.FilterHandler)))
	mux.Handle("/location/", recordMetrics(http.HandlerFunc(handler.LocationHandler)))
	mux.Handle("/search", recordMetrics(http.HandlerFunc(handler.SearchHandler)))

	//adding metrics
	mux.Handle("/metrics", promhttp.Handler())

	log.Println("Listening on: http://localhost:4000/")
	http.ListenAndServe(":4000", mux)
}
