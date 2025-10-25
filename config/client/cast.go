package client

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

func toStringMapFloat64E(i interface{}) (map[string]float64, error) {
	var m = map[string]float64{}
	if i == nil {
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]float64", i, i)
	}

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[cast.ToString(k)] = cast.ToFloat64(val)
		}
		return m, nil
	case map[string]interface{}:
		for k, val := range v {
			m[k] = cast.ToFloat64(val)
		}
		return m, nil
	case map[string]float64:
		return v, nil
	}

	if reflect.TypeOf(i).Kind() != reflect.Map {
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]float64", i, i)
	}
	mVal := reflect.ValueOf(m)
	v := reflect.ValueOf(i)
	for _, keyVal := range v.MapKeys() {
		val, err := cast.ToFloat64E(v.MapIndex(keyVal).Interface())
		if err != nil {
			return m, fmt.Errorf("unable to cast %#v of type %T to map[string]float64", i, i)
		}
		mVal.SetMapIndex(keyVal, reflect.ValueOf(val))
	}
	return m, nil
}
