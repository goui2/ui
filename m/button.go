package m

import (
	"syscall/js"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
)

type Button interface {
	core.Control
	AttachPress(data com.EventData, fn com.EventHandler)
	AttachPressOnce(data com.EventData, fn com.EventHandler)
	FirePress(param com.EventParam)
	DetachPress(fn com.EventHandler)
	HasPress() bool
}

type button struct {
	core.Control
	onPressCB js.Func
}

func constructorButton(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &button{}
		parentSettings := append(s, core.ISRenderer{mo.renderer})
		parentSettings = base.AdjustSelfSetting(mo, md, parentSettings)
		mo.Control = parent.New(id, parentSettings...).(core.Control)
		return mo
	}
}

func (l *button) renderer(rm core.RenderManager, item core.Element) {
	rm.OpenStartElement("button", l)
	rm.OpenEnd()
	rm.WriteText(asString(l.Property("text").Get()))
	rm.Close("button")

}

func (b *button) Param() interface{} {
	return nil
}

func (b *button) OnAfterRendering(v js.Value) {
	b.Control.OnAfterRendering(v)
	if dom, ok := b.DomRef(); ok {
		b.onPressCB = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			b.FirePress(b)
			return nil
		})
		dom.Call("addEventListener", "click", b.onPressCB)
	}
}

func (b *button) AttachPress(data com.EventData, fn com.EventHandler) {
	b.Control.AttachEvent(Event_press, data, fn)
}
func (b *button) AttachPressOnce(data com.EventData, fn com.EventHandler) {
	b.Control.AttachEventOnce(Event_press, data, fn)
}
func (b *button) FirePress(param com.EventParam) {
	b.Control.FireEvent(Event_press, param)
}
func (b *button) DetachPress(fn com.EventHandler) {
	b.Control.DetachEvent(Event_press, fn)
}
func (b *button) HasPress() bool {
	return b.Control.HasListener(Event_press)
}
