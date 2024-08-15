package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: writer,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, request)
		log.Println(wrapped.statusCode, request.Method, request.URL.Path, time.Since(start))
	})
}

func EnsureAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: writer,
			statusCode:     http.StatusCreated,
		}

		next.ServeHTTP(wrapped, request)
		log.Println(wrapped.statusCode, request.Method, request.URL.Path, time.Since(start), "from admin panel")
	})
}
