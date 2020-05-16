package model

import (
	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
)

type ISContext struct {
	com.Context
}

func (x ISContext) IType() base.SettingType {
	return base.Required
}
