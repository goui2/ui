package model

import (
	"reflect"

	"github.com/goui2/ui/com"
)

type Value struct {
	parent  reflect.Value
	name    string
	value   reflect.Value
	context com.Context
}

func (v Value) String() string {
	return v.value.String()
}

func (v Value) IsValid() bool {
	return v.value.IsValid()
}

func (v Value) Interface() interface{} {
	return v.value.Interface()
}

func (v Value) ArraySize() int {
	switch k := v.value.Kind(); k {
	case reflect.Array, reflect.Slice:
		return v.value.Len()
	default:
		return 0
	}
}

func (v Value) Set(s interface{}) {
	switch k := v.parent.Kind(); k {
	case reflect.Array, reflect.Slice:
	case reflect.Struct:
		v.value.Set(reflect.ValueOf(s))
	case reflect.Map:
	}
}
