package common

import (
	"reflect"
)

func StructToMap(str interface{}, data map[int]map[string]string) {

	var (
		t   interface{}
		tag string
	)

	t = reflect.TypeOf(&str).Elem()

	for k := 0; k < t.NumField(); k++ {
		tag = t.Field(k).Tag.Get("json")

	}

}
