package model

import (
	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
)

type contextBinding struct {
	com.Binding
}

func (c *contextBinding) Self() base.Object {
	return c
}

func NewContextBinding(ctx com.Context) com.ContextBinding {
	return MD_ContextBinding.GetClass().New("ctxbind#", ISContext{ctx}).(com.ContextBinding)
}

func constructorContextBinding(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &contextBinding{}
		mo.Binding = parent.New(id, s...).(com.Binding)
		return mo
	}
}

func (cb *contextBinding) GetPath() string {
	return cb.GetContext().GetPath()
}

func (cb *contextBinding) handlePropertyChange(e com.Event) {
	val := cb.GetContext().GetObject("")
	cb.FireChange(eventData{val})
}

func (cb *contextBinding) handleChanged(e com.Event) {

}

func (cb *contextBinding) Handle(e com.Event) {
	//log.Printf("contextBinding.Handle(%#v)\n", e)
	eventType := e.Type()
	switch eventType {
	case "changed":
		cb.handleChanged(e)
	case "propertyChange":
		cb.handlePropertyChange(e)
	}
}
func (b *contextBinding) Metadata() base.MetaData {
	return b.Binding.(*binding).EventProvider.Metadata()
}
