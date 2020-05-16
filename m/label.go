package m

import (
	"syscall/js"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
)

type Label interface {
	core.Control
}

type label struct {
	core.Control
}

func constructorLabel(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &label{}
		parentSettings := append(s, core.ISRenderer{mo.renderer})
		parentSettings = base.AdjustSelfSetting(mo, md, parentSettings)
		mo.Control = parent.New(id, parentSettings...).(core.Control)

		return mo
	}
}

func (l *label) renderer(rm core.RenderManager, item core.Element) {
	rm.OpenStartElement("span", l)
	rm.OpenEnd()
	rm.WriteText(asString(l.Property("text").Get()))
	rm.Close("span")

}

func (l *label) OnAfterRendering(v js.Value) {
	l.Control.OnAfterRendering(v)
	propValue := l.Property("text")
	propValue.AttachUpdateHTML(inputChanged("inputChanged"), l)
}

func (l *label) Handle(e com.Event) {
	switch e.Data().(type) {
	case inputChanged:
		v := l.Property("text").Get()
		r, _ := l.DomRef()
		r.Set("innerText", v)
	}
}
