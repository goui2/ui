package model

import (
	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
)

const (
	filterEvent = "filter"
	sortEvent   = "sort"
)

type listBinding struct {
	com.Binding
}

func (c *listBinding) Self() base.Object {
	return c
}

func NewListBinding(ctx com.Context) com.ListBinding {
	return MD_ListBinding.GetClass().New("listbind#", ISContext{ctx}).(com.ListBinding)
}
func constructorListBinding(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &listBinding{}
		mo.Binding = parent.New(id, s...).(com.Binding)
		return mo
	}
}

func (b *listBinding) AttachFilter(data com.EventData, fn com.EventHandler) {
	b.Binding.(*binding).EventProvider.AttachEvent(filterEvent, data, fn)
}
func (b *listBinding) AttachFilterOnce(data com.EventData, fn com.EventHandler) {
	b.Binding.(*binding).EventProvider.AttachEventOnce(filterEvent, data, fn)
}
func (b *listBinding) FireFilter(param com.EventParam) {
	b.Binding.(*binding).EventProvider.FireEvent(filterEvent, param)
}
func (b *listBinding) DetachFilter(fn com.EventHandler) {
	b.Binding.(*binding).EventProvider.DetachEvent(filterEvent, fn)
}
func (b *listBinding) HasFilter() bool {
	return b.Binding.(*binding).EventProvider.HasListener(filterEvent)
}

func (b *listBinding) AttachSort(data com.EventData, fn com.EventHandler) {
	b.Binding.(*binding).EventProvider.AttachEvent(sortEvent, data, fn)
}
func (b *listBinding) AttachSortOnce(data com.EventData, fn com.EventHandler) {
	b.Binding.(*binding).EventProvider.AttachEventOnce(sortEvent, data, fn)
}
func (b *listBinding) FireSort(param com.EventParam) {
	b.Binding.(*binding).EventProvider.FireEvent(sortEvent, param)
}
func (b *listBinding) DetachSort(fn com.EventHandler) {
	b.Binding.(*binding).EventProvider.DetachEvent(sortEvent, fn)
}
func (b *listBinding) HasSort() bool {
	return b.Binding.(*binding).EventProvider.HasListener(sortEvent)
}

func (b *listBinding) handlePropertyChange(e com.Event) {
	val := b.GetContext().GetObject("")
	b.FireChange(eventData{val})
}

func (b *listBinding) handleChanged(e com.Event) {
}

func (b *listBinding) Handle(e com.Event) {
	eventType := e.Type()
	switch eventType {

	case "changed":
		b.handleChanged(e)
	case "propertyChange":
		b.handlePropertyChange(e)
	}
}
func (b *listBinding) Metadata() base.MetaData {
	return b.Binding.(*binding).EventProvider.Metadata()
}
