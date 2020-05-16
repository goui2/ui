package core

import (
	"github.com/goui2/ui/base"
)

type CSSClass string

func (c CSSClass) Class() string {
	return string(c)
}

func (c CSSClass) Mtype() base.SettingType {
	return base.Unspecific
}
