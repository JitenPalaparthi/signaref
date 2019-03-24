package database

import (
	"errors"
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2"
)

var (
	DBName string
)

const (
	ERRConnectionstr = "connection string must be provided"
)

//GetConnection is to get database connection
func GetConnection(connections ...string) (interface{}, error) {
	if len(connections) < 1 {
		return nil, errors.New(ERRConnectionstr)
	}
	session, err := mgo.Dial(connections[0])
	if err != nil {
		return nil, err
	}
	return session, nil
}

// ToMap converts a struct to a map using the struct's tags.
//
// ToMap uses tags on struct fields to decide which fields to add to the
// returned map.
func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil
}
