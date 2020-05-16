package m

import (
	"github.com/goui2/ui/base"
)

type List interface {
	ListBase
}

type list struct {
	ListBase
}

func constructorList(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &list{}
		parentSettings := s // append(s)
		parentSettings = base.AdjustSelfSetting(mo, md, parentSettings)
		mo.ListBase = parent.New(id, parentSettings...).(ListBase)

		return mo
	}
}
