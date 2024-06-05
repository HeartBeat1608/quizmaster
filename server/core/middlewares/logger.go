package middlewares

import (
	"log"
	"net/http"
)

type extendedResponseWriter struct {
	http.ResponseWriter
	statusCode    int
	contentLength int
}

func (w *extendedResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *extendedResponseWriter) Write(b []byte) (int, error) {
	w.contentLength = len(b)
	return w.ResponseWriter.Write(b)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ew := &extendedResponseWriter{w, 0, 0}
		next.ServeHTTP(ew, r)
		defer log.Printf("[%s] %s - %d - %d", r.Method, r.URL.Path, ew.statusCode, ew.contentLength)
	})
}
