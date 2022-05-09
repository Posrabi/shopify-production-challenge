package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(begin time.Time) {
			logrus.Infof("took: %s", time.Since(begin).String())
		}(time.Now())

		next.ServeHTTP(w, r)
	})
}
