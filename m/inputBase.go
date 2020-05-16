package m

import (
	"github.com/goui2/ui/base"
	"github.com/goui2/ui/core"
)

type InputBase interface {
	core.Control
}

type inputBase struct {
	core.Control
}

func (l *inputBase) renderer(rm core.RenderManager, item core.Element) {

}

func constructorInputBase(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &inputBase{}
		parentSettings := append(s, core.ISRenderer{mo.renderer})
		parentSettings = base.AdjustSelfSetting(mo, md, parentSettings)
		mo.Control = parent.New(id, parentSettings...).(core.Control)

		return mo
	}
}
