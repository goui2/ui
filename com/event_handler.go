package com

type EventHandler interface {
	Handle(Event)
}

type eventHdl struct {
	fn func(Event)
}

func (eh *eventHdl) Handle(e Event) {
	eh.fn(e)
}

func MakeHandler(fn func(Event)) EventHandler {
	return &eventHdl{fn}
}
