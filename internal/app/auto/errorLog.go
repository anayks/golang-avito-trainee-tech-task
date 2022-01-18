package auto

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *server) ErrorLog(r *http.Request, err error) {
	logger := s.logger.WithFields(logrus.Fields{
		"remote_addr": r.RemoteAddr,
		"request_id":  r.Context().Value(ctxKeyRequestID),
	})
	logger.Infof("[%s] Ended with error: %v", r.Method, err)
}
