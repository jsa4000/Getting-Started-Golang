package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"webapp/core/errors"
	log "webapp/core/logging"
)

// Error Sets the error from inner layers
func Error(w http.ResponseWriter, err error) {
	herr, ok := err.(*errors.Error)
	if !ok {
		herr = ErrInternalServer.From(err)
	}
	JSON(w, herr, herr.Code)
	log.Error(herr)
}

// JSON Sets the error from inner layers
func JSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set(HeaderContentType, JSONMime)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// URL decode the form params form the request
func URL(r *http.Request, data interface{}) error {
	q := r.URL.Query()
	v := reflect.ValueOf(data)
	t := v.Elem().Type()
	for i := 0; i < v.Elem().NumField(); i++ {
		ft := t.Field(i)
		fv := v.Elem().Field(i)
		if tv, exist := ft.Tag.Lookup(paramTag); exist {
			val := fmt.Sprintf("%v", fv)
			if !(len(val) == 0 && strings.Contains(tv, omitEmpty)) {
				q.Add(strings.SplitN(tv, ",", 2)[0], val)
			}
		}
	}
	r.URL.RawQuery = q.Encode()
	return nil
}
