package server

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		rw.Header().Set("X-Request-ID", id)
		next.ServeHTTP(rw, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("[%s] %s [START]", r.Method, r.RequestURI)
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		logger.Infof("[%s] Completed in %s [STATUS %d]", r.Method, time.Since(start), rw.code)
	})
}

func (s *server) panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				s.logger.Error(err)
				s.error(rw, r, http.StatusInternalServerError, nil)
			}
		}()
		next.ServeHTTP(rw, r)
	})
}
