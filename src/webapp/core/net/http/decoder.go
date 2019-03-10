package http

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

const (
	paramTag  = "param"
	varTag    = "var"
	omitEmpty = "omitempty"
)

// DecodeOptions struct to specify the request fields to decode
type DecodeOptions struct {
	Body,
	Params,
	Vars,
	Validate bool
}

// NewDecodeOptions struct to specify the request fields to decode
func NewDecodeOptions(body, params, vars, validate bool) *DecodeOptions {
	return &DecodeOptions{
		Body:     body,
		Params:   params,
		Vars:     vars,
		Validate: validate,
	}
}

// DecodeJSON decode the form params form the request
func DecodeJSON(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

// DecodeParams decode the form params form the request
func DecodeParams(r *http.Request, data interface{}) error {
	v := reflect.ValueOf(data)
	t := v.Elem().Type()
	for i := 0; i < v.Elem().NumField(); i++ {
		fv := v.Elem().Field(i)
		ft := t.Field(i)
		if val, ok := ft.Tag.Lookup(paramTag); ok {
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

// DecodeVars decode the form params form the request
func DecodeVars(r *http.Request, data interface{}) error {
	vars := Vars(r)
	v := reflect.ValueOf(data)
	t := v.Elem().Type()
	for i := 0; i < v.Elem().NumField(); i++ {
		fv := v.Elem().Field(i)
		ft := t.Field(i)
		if val, ok := ft.Tag.Lookup(varTag); ok {
			fn := strings.SplitN(val, ",", 2)[0]
			if rv, exist := vars[fn]; exist {
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
