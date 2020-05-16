package core

import "github.com/goui2/ui/base"

type ControlMetaData interface {
	ElementMetaData
	CSSClasses() []string
}

type controlMetaData struct {
	ElementMetaData
	cssClasses []string
}

func (c *controlMetaData) CSSClasses() []string {
	return c.cssClasses
}

func controlMetaDataBuilder(s string, ms ...base.MetaDataSetting) []base.MetaDataSetting {
	epm := &controlMetaData{
		cssClasses: make([]string, 0),
	}
	parentSettings := base.SelectMetadataSettings(ms, func(i base.MetaDataSetting) bool {
		switch s := i.(type) {
		case CSSClass:
			epm.cssClasses = append(epm.cssClasses, s.Class())
			return false
		}
		return true
	})
	parentSettings = append(parentSettings,
		base.MDSettingCallback{
			5,
			func(m base.MetaData, prev base.MetaData) base.MetaData {
				epm.ElementMetaData = m.(ElementMetaData)
				if prevMD, ok := prev.(*controlMetaData); ok {
					epm.cssClasses = append(epm.cssClasses, prevMD.cssClasses...)
				}
				return epm
			},
		},
	)
	return parentSettings

}
