package com

type EventData interface {
	Data() interface{}
}

type EventParam interface {
	Param() interface{}
}

type NewValue interface {
	NewValue() interface{}
}
type OldValue interface {
	OldValue() interface{}
}

type Event interface {
	Type() string
	Data() EventData
	Param() EventParam
}
