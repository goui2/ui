package com

const (

	// Event from Binding to Listener shows changing of data
	Event_change = "change"
)

type Context interface {
	// gibt immer ein Model zurück
	GetModel() interface{}
	GetObject(p string) interface{}
	GetValue(p string) Value
	GetPath() string
}

type Binding interface {
	AttachChange(data EventData, fn EventHandler)
	AttachChangeOnce(data EventData, fn EventHandler)
	FireChange(param EventParam)
	DetachChange(fn EventHandler)
	HasChange() bool

	AttachDataReceived(data EventData, fn EventHandler)
	AttachDataReceivedOnce(data EventData, fn EventHandler)
	FireDataReceived(param EventParam)
	DetachDataReceived(fn EventHandler)
	HasDataReceived() bool

	AttachDataRequested(data EventData, fn EventHandler)
	AttachDataRequestedOnce(data EventData, fn EventHandler)
	FireDataRequested(param EventParam)
	DetachDataRequested(fn EventHandler)
	HasDataRequested() bool

	GetPath() string
	GetContext() Context
	Get() interface{}
	Handle(Event)
}

type ContextBinding interface {
	Binding
}

type PropertyBinding interface {
	Binding
}

type ListBinding interface {
	Binding
	AttachSort(data EventData, fn EventHandler)
	AttachSortOnce(data EventData, fn EventHandler)
	FireSort(param EventParam)
	DetachSort(fn EventHandler)
	HasSort() bool
	AttachFilter(data EventData, fn EventHandler)
	AttachFilterOnce(data EventData, fn EventHandler)
	FireFilter(param EventParam)
	DetachFilter(fn EventHandler)
	HasFilter() bool
}

type Value interface {
	String() string
	IsValid() bool
	Interface() interface{}
	ArraySize() int
	Set(s interface{})
}
