package core

import (
	"github.com/goui2/ui/base"
)

type Control interface {
	Element
	IsVisible() bool
}

type control struct {
	Element
	flexCSSClass []CSSClass
}

func TypeCheckControl(v interface{}) bool {
	_, ok := v.(Control)
	return ok
}

func (c *control) IsVisible() bool {
	b, _ := c.Property("visible").Get().(bool)
	return b
}

func constructorControl(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &control{
			flexCSSClass: make([]CSSClass, 0),
		}

		parentSettings := base.AdjustSelfSetting(mo, md, s)
		mo.Element = parent.New(id, parentSettings...).(Element)

		return mo
	}
}
