package core

import (
	"github.com/goui2/ui/base"
)

type ElementMetaData interface {
	base.ManagedObjectMetaData
	GetElementName() string
	Renderer() Renderer
}

type elementMetaData struct {
	base.ManagedObjectMetaData
	renderer Renderer
}

func (e elementMetaData) GetElementName() string {
	return e.GetName()
}

func (e elementMetaData) Renderer() Renderer {
	return e.renderer
}

func elementMetaDataBuilder(s string, ms ...base.MetaDataSetting) []base.MetaDataSetting {
	epm := &elementMetaData{}
	parentSettings := base.SelectMetadataSettings(ms, func(i base.MetaDataSetting) bool {
		switch x := i.(type) {
		case ISRenderer:
			epm.renderer = x.Renderer
			return false
		}
		return true
	})
	parentSettings = append(ms,
		base.MDSettingCallback{
			4,
			func(m base.MetaData, prev base.MetaData) base.MetaData {
				epm.ManagedObjectMetaData = m.(base.ManagedObjectMetaData)
				if prevMD, ok := prev.(*elementMetaData); ok {
					if epm.renderer == nil {
						epm.renderer = prevMD.renderer
					}
				}
				return epm
			},
		},
	)
	return parentSettings

}
