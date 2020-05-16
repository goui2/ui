package base

import (
	"fmt"
	"sync"

	"github.com/goui2/ui/com"
)

type Property interface {
	EventProvider
	Get() interface{}
	Set(interface{})
	Bind(com.Binding)
	AttachChanged(data com.EventData, fn com.EventHandler)
	AttachChangedOnce(data com.EventData, fn com.EventHandler)
	FireChanged(param com.EventParam)
	DetachChanged(fn com.EventHandler)
	HasChanged() bool
	AttachUpdateHTML(data com.EventData, fn com.EventHandler)
	AttachUpdateHTMLOnce(data com.EventData, fn com.EventHandler)
	FireUpdateHTML(param com.EventParam)
	DetachUpdateHTML(fn com.EventHandler)
	HasUpdateHTML() bool
}

type PropertyDef struct {
	Name         string
	Type         func(interface{}) (interface{}, bool)
	Required     bool
	DefaultValue interface{}
}

type PropertyChanged struct {
	Old interface{}
	New interface{}
}

func (pc PropertyChanged) Param() interface{} {
	//return pc
	panic("should not be used")
}
func (pc PropertyChanged) NewValue() interface{} { return pc.New }

func (pc PropertyChanged) Previous() com.Event { return nil }

type property struct {
	EventProvider
	mO        *managedObject
	name      string
	value     interface{}
	typeCheck func(interface{}) (interface{}, bool)
	binding   com.Binding
	mutex     sync.Mutex
}

func (p *property) Get() interface{} {
	return p.value
}
func (p *property) Set(v interface{}) {
	var pv interface{}
	switch uw := v.(type) {
	case PropertyChanged:
		pv = uw.New
	default:
		pv = v
	}
	if cv, ok := p.typeCheck(pv); ok {
		if p.value != cv {
			pc := PropertyChanged{p.value, cv}
			p.value = cv
			p.FireChanged(pc)
		}
	} else {
		panic(fmt.Sprintf("property %s doesn't accept value %v", p.name, pv))
	}
	//logIt("Property(%s).Set(%#v - %#v)", p.name, p.value, pv)
}

func (p *property) AttachChanged(data com.EventData, fn com.EventHandler) {
	p.EventProvider.AttachEvent("changed", data, fn)
}
func (p *property) AttachChangedOnce(data com.EventData, fn com.EventHandler) {
	p.EventProvider.AttachEventOnce("changed", data, fn)
}
func (p *property) FireChanged(param com.EventParam) {
	p.EventProvider.FireEvent("changed", param)
}
func (p *property) DetachChanged(fn com.EventHandler) {
	p.EventProvider.DetachEvent("changed", fn)
}
func (p *property) HasChanged() bool {
	return p.EventProvider.HasListener("changed")
}
func (p *property) GetSource() EventProvider {
	return p
}
func (p *property) AttachUpdateHTML(data com.EventData, fn com.EventHandler) {
	p.EventProvider.AttachEvent("updateHTML", data, fn)
}
func (p *property) AttachUpdateHTMLOnce(data com.EventData, fn com.EventHandler) {
	p.EventProvider.AttachEventOnce("updateHTML", data, fn)
}
func (p *property) FireUpdateHTML(param com.EventParam) {
	p.EventProvider.FireEvent("updateHTML", param)
}
func (p *property) DetachUpdateHTML(fn com.EventHandler) {
	p.EventProvider.DetachEvent("updateHTML", fn)
}
func (p *property) HasUpdateHTML() bool {
	return p.EventProvider.HasListener("updateHTML")
}

func (p *property) Handle(e com.Event) {
	switch e.Type() {
	case "change":
		p.handleChange(e)
	}
}

func (p *property) handleChange(e com.Event) {
	//logIt("property.Handle1(%s, %#v", p.name, e)
	//logIt("property.Handle3(%s, %#v", p.name, ep.Parameter())
	var newValue interface{}
	newValue = e.Param()
	switch ed := newValue.(type) {
	case com.NewValue:
		newValue = ed.NewValue()
	default:
		newValue = nil
	}
	vc := PropertyChanged{
		Old: p.value,
		New: newValue,
	}
	p.value = newValue
	p.FireUpdateHTML(vc)
}

type bindDataChange string

func (b bindDataChange) Data() interface{} { return string(b) }

func (p *property) Bind(b com.Binding) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.binding != nil {
		p.binding.DetachChange(p)
		p.DetachChanged(p.binding)
	}
	p.binding = b
	b.AttachChange(bindDataChange("bDC"), p)
	p.AttachChanged(nil, b)
	vc := PropertyChanged{
		Old: p.value,
		New: b.Get(),
	}
	p.value = vc.New
	p.FireUpdateHTML(vc)
}
