package base

import (
	"sync"

	"github.com/goui2/ui/com"
)

type ObjectFactory func(b com.Binding) ManagedObject

type ObjectFactorySetting struct {
	Factory            ObjectFactory
	BindingDescription com.BindingDescription
	Aggregation        string
}

func (_ ObjectFactorySetting) IType() SettingType {
	return Unspecific
}

type Aggregation interface {
	EventProvider
	Get() []interface{}
	Set([]interface{})
	Add(interface{})
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

type AggregationDef struct {
	Name        string
	Type        func(interface{}) bool
	Cardinality string
}

type aggregation struct {
	EventProvider
	mO        *managedObject
	name      string
	values    []interface{}
	typeCheck func(interface{}) bool
	min       int
	max       int
	binding   com.Binding
	mutex     sync.Mutex
}

type setObject interface {
	SetParent(Object)
}

func (a *aggregation) Get() []interface{} { return a.values }
func (a *aggregation) Set(v []interface{}) {
	//	logIt("aggregation.Set")
	for _, i := range v {
		if !a.typeCheck(i) {
			panic("invalid type: " + a.name + "  ")
		}
		if os, ok := i.(setObject); ok {
			os.SetParent(a.mO.Self())
		}
	}
	pc := PropertyChanged{a.values, v}
	a.values = v
	a.FireUpdateHTML(pc)
	//	logIt("aggregation.Set exit")
}
func (a *aggregation) Add(v interface{}) {
	if !a.typeCheck(v) {
		panic("invalid type")
	}
	pc := PropertyChanged{a.values, nil}
	a.values = append(a.values, v)
	pc.New = a.values
	if os, ok := v.(setObject); ok {
		os.SetParent(a.mO.Self())
	}
	a.FireEvent("changed", pc)
}

func (a *aggregation) AttachChanged(data com.EventData, fn com.EventHandler) {
	a.EventProvider.AttachEvent("changed", data, fn)
}
func (a *aggregation) AttachChangedOnce(data com.EventData, fn com.EventHandler) {
	a.EventProvider.AttachEventOnce("changed", data, fn)
}
func (a *aggregation) FireChanged(param com.EventParam) {
	a.EventProvider.FireEvent("changed", param)
}
func (a *aggregation) DetachChanged(fn com.EventHandler) {
	a.EventProvider.DetachEvent("changed", fn)
}
func (a *aggregation) HasChanged() bool {
	return a.EventProvider.HasListener("changed")
}

func (a *aggregation) AttachUpdateHTML(data com.EventData, fn com.EventHandler) {
	a.EventProvider.AttachEvent("updateHTML", data, fn)
}
func (a *aggregation) AttachUpdateHTMLOnce(data com.EventData, fn com.EventHandler) {
	a.EventProvider.AttachEventOnce("updateHTML", data, fn)
}
func (a *aggregation) FireUpdateHTML(param com.EventParam) {
	a.EventProvider.FireEvent("updateHTML", param)
}
func (a *aggregation) DetachUpdateHTML(fn com.EventHandler) {
	a.EventProvider.DetachEvent("updateHTML", fn)
}
func (a *aggregation) HasUpdateHTML() bool {
	return a.EventProvider.HasListener("updateHTML")
}

func (a *aggregation) Handle(e com.Event) {
	switch e.Type() {
	case "change":
		a.handleChange(e)
	}
}

func (a *aggregation) handleChange(e com.Event) {
	var newValue interface{}
	newValue = e.Param()
	switch ed := newValue.(type) {
	case PropertyChanged:
		newValue = ed.New
	default:
		newValue = nil
	}
	vc := PropertyChanged{
		Old: a.values,
		New: newValue,
	}
	var newValues []interface{}
	var ok bool
	if newValues, ok = newValue.([]interface{}); ok {
		a.values = newValues
		a.FireUpdateHTML(vc)
	}

}

func (a *aggregation) Bind(b com.Binding) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.binding != nil {
		a.binding.DetachChange(a)
		a.DetachChanged(a.binding)
	}
	a.binding = b
	b.AttachChange(bindDataChange("bDC"), a)
	a.AttachChanged(nil, b)
}
