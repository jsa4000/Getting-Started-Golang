package http

import (
	"net/http"
	"time"
	log "webapp/core/logging"
)

// LoggingMiddleware returns LogginMiddleware struct
type LoggingMiddleware struct {
	MiddlewareBase
}

// NewLoggingMiddleware creation
func NewLoggingMiddleware() Middleware {
	return &LoggingMiddleware{
		MiddlewareBase{
			Hdlr: LoggingHandler,
			Prio: PriorityLogging,
		},
	}
}

// LoggingHandler decorator (closure)
func LoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Received Request uri=%s args=%s ", RemoveURLParams(r.RequestURI), Vars(r))
		start := time.Now()
		defer func() {
			log.Debugf("Processed Response in %d ns", time.Since(start).Nanoseconds())
		}()
		next.ServeHTTP(w, r)
	})
}
