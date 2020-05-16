package model

import (
	"reflect"
	"strconv"

	"github.com/goui2/ui/com"
)

type context struct {
	model Model
	path  Path
}

func NewContext(m Model, s string) com.Context {
	p := NewPath(s)
	if !p.IsAbsolut() {
		panic("path is not absolute: " + p.String())
	}
	return &context{
		model: m,
		path:  p,
	}
}

func (c *context) GetModel() interface{} {
	return c.model
}

func (c *context) GetObject(s string) interface{} {
	if v := c.GetValue(s); v.IsValid() {
		return v.Interface()
	} else {
		return nil
	}
}

var (
	reflectInvalid = reflect.ValueOf(nil)
)

func (c *context) GetValue(sp string) (val com.Value) {
	p := NewPath(sp)
	var rv, pv reflect.Value
	rv = reflect.ValueOf(nil)
	pv = rv
	v := c.GetModel().(Model).GetObject()
	if v == nil {
		return Value{reflectInvalid, "", reflectInvalid, c}
	}
	var ok bool
	var s string
	defer func() {
		if t := recover(); t != nil {
			val = Value{pv, s, reflectInvalid, c}
		}
	}()
	dp := c.path.Append(p)
	if rv, ok = v.(reflect.Value); !ok {
		rv = reflect.ValueOf(v)
	}
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	for idx, _ := range dp[1:] {
		pv = rv
		s = dp.Name(uint8(idx + 1))
		switch k := rv.Kind(); k {
		case reflect.Struct:
			rv = rv.FieldByName(s)
		case reflect.Map:
			rv = rv.MapIndex(reflect.ValueOf(s))
			if rv.IsValid() {
				rv = rv.Elem()
			}
		case reflect.Array, reflect.Slice:
			if idx, err := strconv.Atoi(s); err == nil {
				rv = rv.Index(idx)
			} else {
				return Value{pv, s, reflectInvalid, c}
			}
		default:
			return Value{pv, s, reflectInvalid, c}
		}
	}
	return Value{pv, s, rv, c}
}

func (c *context) GetPath() string {
	return c.path.String()
}
