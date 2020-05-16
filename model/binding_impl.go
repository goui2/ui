package model

import (
	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
)

type binding struct {
	base.EventProvider
	context com.Context
}

func NewBinding(ctx com.Context) com.Binding {
	return MD_Binding.GetClass().New("bind#", ISContext{ctx}).(com.Binding)
}

func constructorBinding(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &binding{}
		mo.EventProvider = parent.New(id, s...).(base.EventProvider)
		for _, v := range s {
			switch c := v.(type) {
			case ISContext:
				mo.context = c.Context
			}
		}
		return mo
	}
}

func (b *binding) AttachChange(data com.EventData, fn com.EventHandler) {
	b.EventProvider.AttachEvent(Event_changeEvent, data, fn)
}
func (b *binding) AttachChangeOnce(data com.EventData, fn com.EventHandler) {
	b.EventProvider.AttachEventOnce(Event_changeEvent, data, fn)
}
func (b *binding) FireChange(param com.EventParam) {
	b.EventProvider.FireEvent(Event_changeEvent, param)
}
func (b *binding) DetachChange(fn com.EventHandler) {
	b.EventProvider.DetachEvent(Event_changeEvent, fn)
}
func (b *binding) HasChange() bool {
	return b.EventProvider.HasListener(Event_changeEvent)
}

func (b *binding) AttachDataReceived(data com.EventData, fn com.EventHandler) {
	b.EventProvider.AttachEvent(Event_dataReceivedEvent, data, fn)
}
func (b *binding) AttachDataReceivedOnce(data com.EventData, fn com.EventHandler) {
	b.EventProvider.AttachEventOnce(Event_dataReceivedEvent, data, fn)
}
func (b *binding) FireDataReceived(param com.EventParam) {
	b.EventProvider.FireEvent(Event_dataReceivedEvent, param)
}
func (b *binding) DetachDataReceived(fn com.EventHandler) {
	b.EventProvider.DetachEvent(Event_dataReceivedEvent, fn)
}
func (b *binding) HasDataReceived() bool {
	return b.EventProvider.HasListener(Event_dataReceivedEvent)
}

func (b *binding) AttachDataRequested(data com.EventData, fn com.EventHandler) {
	b.EventProvider.AttachEvent(Event_dataRequestedEvent, data, fn)
}
func (b *binding) AttachDataRequestedOnce(data com.EventData, fn com.EventHandler) {
	b.EventProvider.AttachEventOnce(Event_dataRequestedEvent, data, fn)
}
func (b *binding) FireDataRequested(param com.EventParam) {
	b.EventProvider.FireEvent(Event_dataRequestedEvent, param)
}
func (b *binding) DetachDataRequested(fn com.EventHandler) {
	b.EventProvider.DetachEvent(Event_dataRequestedEvent, fn)
}
func (b *binding) HasDataRequested() bool {
	return b.EventProvider.HasListener(Event_dataRequestedEvent)
}

func (b *binding) GetPath() string {
	return b.context.GetPath()
}

func (b *binding) GetContext() com.Context {
	return b.context
}

func (b *binding) Metadata() base.MetaData {
	return b.EventProvider.Metadata()
}

func (b *binding) Handle(e com.Event) {
	panic("method Handle is abstract")
}

func (b *binding) Get() interface{} {
	ctx := b.GetContext()
	return ctx.GetObject("")
}
