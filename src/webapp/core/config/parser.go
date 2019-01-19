package config

import (
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Tag used for inference path and default value
// i.e  Name   string  `config:"app.name:ServerApp"`
const Tag = "config"

// PathSeparator to split the tag into `config:"path:default"`
const PathSeparator = ":"

// Parser interface to read from config files or environment
type Parser interface {
	ReadFields(data interface{})
	Get(path string) (interface{}, error)
	GetInt(path string) int
	GetBool(path string) bool
	GetFloat64(path string) float64
	GetString(path string) string
	SetDefault(key string, value interface{})
}

// GetTagValue gets a tag from a field and from a Type
func GetTagValue(t reflect.Type, field string) string {
	f, _ := t.FieldByName(field)
	value, _ := f.Tag.Lookup(Tag)
	return value
}

//ProcessPath split the path in two parts, returns only the path
func ProcessPath(parser Parser, path string) string {
	items := strings.SplitN(path, PathSeparator, 2)
	if len(items) == 1 {
		return items[0]
	}
	parser.SetDefault(items[0], items[1])
	return items[0]
}

// ReadFields set automatically the config using a parser and a struct with tags 'config'
func ReadFields(parser Parser, data interface{}) {
	v := reflect.ValueOf(data)
	t := v.Elem().Type()
	for i := 0; i < v.Elem().NumField(); i++ {
		fv := v.Elem().Field(i)
		ft := t.Field(i)
		tag, ok := ft.Tag.Lookup(Tag)
		if ok {
			log.WithFields(log.Fields{"Name": ft.Name, "Kind": fv.Kind(), "Type": fv.Type(), "Tag": tag}).
				Debugf("Read Default Configuration from '%s/%s'", t.PkgPath(), t.Name())
			switch fv.Kind() {
			case reflect.String:
				fv.SetString(parser.GetString(tag))
			case reflect.Int, reflect.Int32, reflect.Int64:
				fv.SetInt(int64(parser.GetInt(tag)))
			case reflect.Float32, reflect.Float64:
				fv.SetFloat(parser.GetFloat64(tag))
			case reflect.Bool:
				fv.SetBool(parser.GetBool(tag))
			default:
				log.WithFields(log.Fields{"Name": ft.Name, "Kind": fv.Kind(), "Type": fv.Type(), "Tag": tag}).
					Warning("Kind is not supported")
			}
		}
	}
}
