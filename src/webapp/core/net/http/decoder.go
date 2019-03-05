package http

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// DecodeParams decode the form params form the request
func DecodeParams(r *http.Request, data interface{}, tag string) error {
	v := reflect.ValueOf(data)
	t := v.Elem().Type()
	for i := 0; i < v.Elem().NumField(); i++ {
		fv := v.Elem().Field(i)
		ft := t.Field(i)
		if val, ok := ft.Tag.Lookup(tag); ok {
			fn := strings.SplitN(val, ",", 2)[0]
			if rv := r.FormValue(fn); len(rv) > 0 {
				switch fv.Kind() {
				case reflect.String:
					fv.SetString(rv)
				case reflect.Int, reflect.Int32, reflect.Int64:
					if intVal, err := strconv.Atoi(rv); err == nil {
						fv.SetInt(int64(intVal))
					}
				case reflect.Float32, reflect.Float64:
					if floatVal, err := strconv.ParseFloat(rv, 64); err == nil {
						fv.SetFloat(floatVal)
					}
				case reflect.Bool:
					if boolVal, err := strconv.ParseBool(rv); err == nil {
						fv.SetBool(boolVal)
					}
				}
			}
		}
	}
	return nil
}
