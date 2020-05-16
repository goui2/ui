package m

import (
	"github.com/goui2/ui/base"
	"github.com/goui2/ui/core"
)

type ListItemBase interface {
	core.Control
}

type listItemBase struct {
	core.Control
}

func constructorListItemBase(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &listItemBase{}
		parentSettings := s //append(s, core.ISRenderer{mo.renderer})
		parentSettings = base.AdjustSelfSetting(mo, md, parentSettings)
		mo.Control = parent.New(id, parentSettings...).(core.Control)

		return mo
	}
}

func TypeCheckListItemBase(v interface{}) bool {
	//	logIt("TypeCheckListItemBase %#v", v)
	_, ok := v.(ListItemBase)
	return ok
}
