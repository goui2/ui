package core

import (
	"github.com/goui2/ui/base"
)

type CSSSize string

func CSSSizeOf(s string) CSSSize {
	return CSSSize(s)
}

func (s CSSSize) String() string {
	return string(s)
}

func TypeCheckCSSSize(v interface{}) bool {
	_, ok := v.(CSSSize)
	return ok
}

func PropTypeCSSSize(v interface{}) (interface{}, bool) {
	if TypeCheckCSSSize(v) {
		return v, true
	} else if base.TypeCheckString(v) {
		return CSSSize(v.(string)), true
	} else {
		return v, false
	}
}
