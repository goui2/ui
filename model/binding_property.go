package model

import (
	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
)

type propertyBinding struct {
	com.Binding
	actualValue interface{}
}

func (p *propertyBinding) Self() base.Object {
	return p
}

func NewPropertyBinding(ctx com.Context) com.PropertyBinding {
	return MD_PropertyBinding.GetClass().New("propbind#", ISContext{ctx}).(com.PropertyBinding)
}
func constructorPropertyBinding(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &propertyBinding{}
		mo.Binding = parent.New(id, s...).(com.Binding)
		return mo
	}
}

func (pb *propertyBinding) Handle(e com.Event) {
	//log.Printf("event %v\n", e)
	switch e.Type() {
	case "changed":
		pb.handleChanged(e)
	case Event_propertyChange:
		pb.handlePropertyChange(e)
	}
}

func (pb *propertyBinding) handleChanged(e com.Event) {
	ep := e.Param()
	ctx := pb.GetContext()
	val := ctx.GetValue("")
	newValue := unwrapData(ep)
	if val.Interface() != newValue {
		val.Set(newValue)
		pb.FireChange(base.PropertyChanged{val.Interface(), newValue})
	}
}

func (pb *propertyBinding) handlePropertyChange(e com.Event) {
	ctx := pb.GetContext()
	val := ctx.GetValue("")
	if val.Interface() != pb.actualValue {
		changed := base.PropertyChanged{pb.actualValue, val.Interface()}
		pb.actualValue = val.Interface()
		pb.FireChange(changed)
	}
}

//func (pb *propertyBinding) GetSource() base.EventProvider { return pb }
func (b *propertyBinding) Metadata() base.MetaData {
	return b.Binding.(*binding).EventProvider.Metadata()
}
