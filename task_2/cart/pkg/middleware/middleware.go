package middleware

import (
	"log"
	"net/http"
	"time"
)

// loggingResponseWriter - структура, которая помогает получить статус код
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware - логирующий миддлваре
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		log.Printf("> > > HTTP-REQUEST: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(lrw, r)

		log.Printf("< < < HTTP-RESPONSE: %s %s %d %s", r.Method, r.URL.Path, lrw.statusCode, time.Since(start))
	})
}
