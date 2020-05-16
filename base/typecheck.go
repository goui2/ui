package base

import (
	"strconv"
	"strings"
)

func TypeCheckBool(v interface{}) bool {
	_, ok := v.(bool)
	return ok
}

func TypeCheckString(v interface{}) bool {
	_, ok := v.(string)
	return ok
}
func TypeCheckInt(v interface{}) bool {
	_, ok := v.(int)
	return ok
}
func TypeCheckManagedObject(v interface{}) bool {
	_, ok := v.(ManagedObject)
	return ok
}

func PropTypeBool(v interface{}) (interface{}, bool) {
	if TypeCheckBool(v) {
		return v, true
	} else if TypeCheckString(v) {
		vs := v.(string)
		vs = strings.ToLower(vs)
		if vs == "true" {
			return true, true
		} else {
			return false, true
		}
	} else {
		return v, false
	}
}

func PropTypeString(v interface{}) (interface{}, bool) {
	if TypeCheckString(v) {
		return v, true
	} else {
		return v, false
	}
}

func PropTypeInt(v interface{}) (interface{}, bool) {
	if TypeCheckInt(v) {
		return v, true
	} else if TypeCheckString(v) {
		i, err := strconv.Atoi(v.(string))
		if err == nil {
			return i, true
		} else {
			return v, false
		}
	} else {
		return v, false
	}
}
