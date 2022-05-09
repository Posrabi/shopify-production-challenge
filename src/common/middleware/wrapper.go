package middleware

import (
	"net/http"

	"github.com/Posrabi/shopify-backend-project/src/common/exception"
	"github.com/sirupsen/logrus"
)

type ErrorHandlerFunc func(http.ResponseWriter, *http.Request) error // http.Handler that will error

// ErrorWrapper logs out any error returned from the handler
func ErrorWrapper(name string, errorHandler ErrorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// this is where we could also use tracers to monitor performance.
		if err := errorHandler(w, r); err != nil {
			logrus.WithFields(logrus.Fields{
				"endpoint": name,
				"url":      r.URL.Path,
				"query":    r.URL.Query(),
			}).Error(exception.LogErr(err))
		}
	}
}
