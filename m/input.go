package m

import (
	"syscall/js"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
)

type Input interface {
	InputBase
}

type InputType string

const (
	InputType_Text   InputType = "text"
	InputType_Passwd           = "password"
	InputType_Tel              = "tel"
	InputType_Url              = "url"
	InputType_EMail            = "email"
)

func TypeCheckInputType(v interface{}) bool {
	_, ok := v.(InputType)
	return ok
}

func PropTypeInputType(v interface{}) (interface{}, bool) {
	if TypeCheckInputType(v) {
		return v, true
	} else if base.TypeCheckString(v) {
		return InputType(v.(string)), true
	} else {
		return v, false
	}
}

var inputEvents = []string{Event_inputHTML}

type input struct {
	InputBase
	valueCB js.Func
}

func constructorInput(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &input{}
		parentSettings := append(s, core.ISRenderer{mo.renderer})
		parentSettings = base.AdjustSelfSetting(mo, md, parentSettings)
		mo.InputBase = parent.New(id, parentSettings...).(InputBase)

		return mo
	}
}

func asString(i interface{}) string {
	if i == nil {
		return ""
	} else {
		return i.(string)
	}
}

func (l *input) renderer(rm core.RenderManager, item core.Element) {
	rm.OpenStartElement("input", l)
	inputtype := l.Property("type").Get().(InputType)
	rm.WriteAttr("type", string(inputtype))
	rm.WriteAttr("name", asString(l.Property("name").Get()))
	rm.WriteAttr("value", asString(l.Property("value").Get()))
	rm.OpenEnd()
	rm.Close("input")

}

type inputChanged string

func (ic inputChanged) Data() interface{} {
	return ic
}

type inputHtml string

func (ih inputHtml) Data() interface{} {
	return ih
}

/*
func (ih inputHtml) Param() interface{} {
	return ih
}
*/

func (l *input) OnAfterRendering(v js.Value) {
	logIt("Input.OnAfterRendering %v", v)
	l.InputBase.OnAfterRendering(v)
	propValue := l.Property("value")
	propValue.AttachUpdateHTML(inputChanged("inputChanged"), l)
	if dom, ok := l.DomRef(); ok {
		l.valueCB = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			data := dom.Get("value").String()
			logIt("Input.OnAfterrendering.HtmlCallback(%v)", data)
			propValue.Set(data)
			return nil
		})
		dom.Call("addEventListener", "input", l.valueCB)
	}
	logIt("Input.OnAfterRendering exit")
}

func (l *input) Handle(e com.Event) {
	logIt("input.Handle 1 %v -- %s", e, l)
	switch e.Data().(type) {
	case inputChanged:
		logIt("input.Handle(inputChanged 1, %v) -- %p", e, l)
		v := l.Property("value").Get()
		logIt("input.Handle(inputChanged 2, %#v) -- %p", v, l)
		r, _ := l.DomRef()
		r.Set("value", v)
	case inputHtml:
		logIt("input.Handle(inputHtml1, %v) -- %p", e, l)
		r, _ := l.DomRef()
		s := r.Get("value").String()
		logIt("input.Handle(inputHtml2, %v) -- %s", e, s)
		l.Property("value").Set(s)
	}
}

func (l *input) Destroy() {
	l.valueCB.Release()

}

func (l *input) GetSource() base.EventProvider {
	return l
}
