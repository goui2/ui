package model

import (
	"reflect"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core/message"
)

type Model interface {
	message.MessageProcessor
	SetObject(v interface{})
	GetObject() interface{}
	AttachParseError(data com.EventData, fn com.EventHandler)
	AttachParseErrorOnce(data com.EventData, fn com.EventHandler)
	FireParseError(param com.EventParam)
	DetachParseError(fn com.EventHandler)
	HasParseError() bool
	AttachPropertyChange(data com.EventData, fn com.EventHandler)
	AttachPropertyChangeOnce(data com.EventData, fn com.EventHandler)
	FirePropertyChange(param com.EventParam)
	DetachPropertyChange(fn com.EventHandler)
	HasPropertyChange() bool
	AttachRequestCompleted(data com.EventData, fn com.EventHandler)
	AttachRequestCompletedOnce(data com.EventData, fn com.EventHandler)
	FireRequestCompleted(param com.EventParam)
	DetachRequestCompleted(fn com.EventHandler)
	HasRequestCompleted() bool
	AttachRequestSend(data com.EventData, fn com.EventHandler)
	AttachRequestSendOnce(data com.EventData, fn com.EventHandler)
	FireRequestSend(param com.EventParam)
	DetachRequestSend(fn com.EventHandler)
	HasRequestSend() bool
	AttachRequestFailed(data com.EventData, fn com.EventHandler)
	AttachRequestFailedOnce(data com.EventData, fn com.EventHandler)
	FireRequestFailed(param com.EventParam)
	DetachRequestFailed(fn com.EventHandler)
	HasRequestFailed() bool
	BindProperty(p string) com.PropertyBinding
	BindContext(p string) com.ContextBinding
	BindList(p string) com.ListBinding
}

type pathBindings struct {
	path     string
	property com.PropertyBinding
	context  com.ContextBinding
	list     com.ListBinding
}

type model struct {
	message.MessageProcessor
	data     interface{}
	bindings map[string]pathBindings
}

func constructorModel(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &model{}
		mo.bindings = make(map[string]pathBindings)
		mo.MessageProcessor = parent.New(id, s...).(message.MessageProcessor)
		return mo
	}
}

func (m *model) SetObject(v interface{}) {
	mcd := base.PropertyChanged{m.data, v}
	switch va := v.(type) {
	case reflect.Value:
		mcd.New = va.Interface()
	case com.Value:
		mcd.New = va.Interface()
	}
	m.data = v
	m.FirePropertyChange(mcd)
	//	}
}
func (m *model) GetObject() interface{} {
	return m.data
}

func (m *model) GetSource() base.EventProvider { return m }

func (m *model) AttachParseError(data com.EventData, fn com.EventHandler) {
	m.MessageProcessor.AttachEvent(Event_parseError, data, fn)
}
func (m *model) AttachParseErrorOnce(data com.EventData, fn com.EventHandler) {
	m.MessageProcessor.AttachEventOnce(Event_parseError, data, fn)
}
func (m *model) FireParseError(param com.EventParam) {
	m.MessageProcessor.FireEvent(Event_parseError, param)
}
func (m *model) DetachParseError(fn com.EventHandler) {
	m.MessageProcessor.DetachEvent(Event_parseError, fn)
}
func (m *model) HasParseError() bool {
	return m.MessageProcessor.HasListener(Event_parseError)
}

func (m *model) AttachPropertyChange(data com.EventData, fn com.EventHandler) {
	m.MessageProcessor.AttachEvent(Event_propertyChange, data, fn)
}
func (m *model) AttachPropertyChangeOnce(data com.EventData, fn com.EventHandler) {
	m.MessageProcessor.AttachEventOnce(Event_propertyChange, data, fn)
}
func (m *model) FirePropertyChange(param com.EventParam) {
	m.MessageProcessor.FireEvent(Event_propertyChange, param)
}
func (m *model) DetachPropertyChange(fn com.EventHandler) {
	m.MessageProcessor.DetachEvent(Event_propertyChange, fn)
}
func (m *model) HasPropertyChange() bool {
	return m.MessageProcessor.HasListener(Event_propertyChange)
}

func (m *model) AttachRequestCompleted(data com.EventData, fn com.EventHandler) {
	m.MessageProcessor.AttachEvent(Event_requestCompleted, data, fn)
}
func (m *model) AttachRequestCompletedOnce(data com.EventData, fn com.EventHandler) {
	m.MessageProcessor.AttachEventOnce(Event_requestCompleted, data, fn)
}
func (m *model) FireRequestCompleted(param com.EventParam) {
	m.MessageProcessor.FireEvent(Event_requestCompleted, param)
}
func (m *model) DetachRequestCompleted(fn com.EventHandler) {
	m.MessageProcessor.DetachEvent(Event_requestCompleted, fn)
}
func (m *model) HasRequestCompleted() bool {
	return m.MessageProcessor.HasListener(Event_requestCompleted)
}

func (m *model) AttachRequestFailed(data com.EventData, fn com.EventHandler) {
	m.MessageProcessor.AttachEvent(Event_requestFailed, data, fn)
}

func (m *model) AttachRequestFailedOnce(data com.EventData, fn com.EventHandler) {
	m.MessageProcessor.AttachEventOnce(Event_requestFailed, data, fn)
}

func (m *model) FireRequestFailed(param com.EventParam) {
	m.MessageProcessor.FireEvent(Event_requestFailed, param)
}

func (m *model) DetachRequestFailed(fn com.EventHandler) {
	m.MessageProcessor.DetachEvent(Event_requestFailed, fn)
}

func (m *model) HasRequestFailed() bool {
	return m.MessageProcessor.HasListener(Event_requestFailed)
}

func (m *model) AttachRequestSend(data com.EventData, fn com.EventHandler) {
	m.MessageProcessor.AttachEvent(Event_requestSend, data, fn)
}
func (m *model) AttachRequestSendOnce(data com.EventData, fn com.EventHandler) {
	m.MessageProcessor.AttachEventOnce(Event_requestSend, data, fn)
}
func (m *model) FireRequestSend(param com.EventParam) {
	m.MessageProcessor.FireEvent(Event_requestSend, param)
}
func (m *model) DetachRequestSend(fn com.EventHandler) {
	m.MessageProcessor.DetachEvent(Event_requestSend, fn)
}
func (m *model) HasRequestSend() bool {
	return m.MessageProcessor.HasListener(Event_requestSend)
}

func (sm *model) BindContext(p string) com.ContextBinding {
	if pathB, ok := sm.bindings[p]; ok {
		if pathB.context != nil {
			return pathB.context
		} else {
			pathB.context = NewContextBinding(NewContext(sm, p))
			sm.AttachPropertyChange(nil, pathB.context)
			sm.bindings[p] = pathB
			return pathB.context
		}
	} else {
		pathB := pathBindings{
			path:    p,
			context: NewContextBinding(NewContext(sm, p)),
		}
		sm.AttachPropertyChange(nil, pathB.context)
		sm.bindings[p] = pathB
		return pathB.context
	}
}

func (sm *model) BindProperty(p string) com.PropertyBinding {
	if pathB, ok := sm.bindings[p]; ok {
		if pathB.property != nil {
			return pathB.property
		} else {
			pathB.property = NewPropertyBinding(NewContext(sm, p))
			sm.AttachPropertyChange(nil, pathB.property)
			sm.bindings[p] = pathB
			return pathB.property
		}
	} else {
		pathB := pathBindings{
			path:     p,
			property: NewPropertyBinding(NewContext(sm, p)),
		}
		sm.AttachPropertyChange(nil, pathB.property)
		sm.bindings[p] = pathB
		return pathB.property
	}
}

func (sm *model) BindList(p string) com.ListBinding {
	if pathB, ok := sm.bindings[p]; ok {
		if pathB.list != nil {
			return pathB.list
		} else {
			pathB.list = NewListBinding(NewContext(sm, p))
			sm.AttachPropertyChange(nil, pathB.list)
			sm.bindings[p] = pathB
			return pathB.list
		}
	} else {
		pathB := pathBindings{
			path: p,
			list: NewListBinding(NewContext(sm, p)),
		}
		sm.AttachPropertyChange(nil, pathB.list)
		sm.bindings[p] = pathB
		return pathB.list
	}
}
