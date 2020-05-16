package core

import "github.com/goui2/ui/base"

type ISRenderer struct {
	Renderer Renderer
}

func (i ISRenderer) IType() base.SettingType {
	return base.Unspecific
}
func (i ISRenderer) Mtype() base.SettingType {
	return base.Unspecific
}
