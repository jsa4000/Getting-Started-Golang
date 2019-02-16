package http

import (
	"net/http"
	"strings"
	"time"
	log "webapp/core/logging"
)

const (
	pprofPreffix   = "/debug/pprof/"
	swaggerPreffix = "/swagger"
)

// LoggingMiddleware returns LogginMiddleware struct
type LoggingMiddleware struct {
	MiddlewareBase
}

// CustomHeadersMiddleware returns LogginMiddleware struct
type CustomHeadersMiddleware struct {
	MiddlewareBase
}

// NewLoggingMiddleware creation
func NewLoggingMiddleware() Middleware {
	return &LoggingMiddleware{
		MiddlewareBase{
			handler:  LoggingHandler,
			priority: PriorityLogging,
		},
	}
}

// NewCustomHeadersMiddleware creation
func NewCustomHeadersMiddleware() Middleware {
	return &CustomHeadersMiddleware{
		MiddlewareBase{
			handler:  CustomHeadersHandler,
			priority: PriorityHeaders,
		},
	}
}

// LoggingHandler decorator (closure)
func LoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Received Request uri=%s args=%s ", r.RequestURI, Vars(r))
		start := time.Now()
		defer func() {
			log.Debugf("Processed Response in %d ns", time.Since(start).Nanoseconds())
		}()
		next.ServeHTTP(w, r)
	})
}

// CustomHeadersHandler decorator (closure)
func CustomHeadersHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Setting custom headers for request uri=%s", r.RequestURI)
		if !strings.Contains(r.RequestURI, pprofPreffix) && !strings.Contains(r.RequestURI, swaggerPreffix) {
			w.Header().Set("Content-Type", "application/json")
		}
		//defaultHeaders(w)
		enableCors(w)

		next.ServeHTTP(w, r)
	})
}

func defaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
