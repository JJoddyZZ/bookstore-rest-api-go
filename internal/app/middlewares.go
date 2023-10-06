package app

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewStatusResponseWriter
func NewStatusResponseWriter(w http.ResponseWriter) *statusResponseWriter {
	return &statusResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

// WriteHeader "overrides" the "inherited" method from http.ResponseWriter
func (sw *statusResponseWriter) WriteHeader(statusCode int) {
	sw.statusCode = statusCode // catches the response status
	sw.ResponseWriter.WriteHeader(statusCode)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := NewStatusResponseWriter(w)

		defer func() {
			log.Info().
				Str("method", r.Method).
				Str("path", r.RequestURI).
				Str("host", r.Host).
				Int("status", sw.statusCode).
				Msg("")
		}()

		next.ServeHTTP(sw, r)
	})
}

func setContentHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(contentTypeHeaderKey, contentTypeHeaderValueApplicationJSON)
		next.ServeHTTP(w, r)
	})
}
