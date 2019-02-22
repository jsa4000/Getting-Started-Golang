package http

import (
	"net/http"
	"time"
	log "webapp/core/logging"
)

var (
	filters = []string{"/debug/pprof/", "/swagger/"}
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
			Hdlr: LoggingHandler,
			Prio: PriorityLogging,
		},
	}
}

// NewCustomHeadersMiddleware creation
func NewCustomHeadersMiddleware() Middleware {
	return &CustomHeadersMiddleware{
		MiddlewareBase{
			Hdlr: CustomHeadersHandler,
			Prio: PriorityHeaders,
		},
	}
}

// LoggingHandler decorator (closure)
func LoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Received Request uri=%s args=%s ", RemoveParams(r.RequestURI), Vars(r))
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
		if !Matches(RemoveParams(r.RequestURI), filters) {
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
