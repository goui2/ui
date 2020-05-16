package model

import (
	"github.com/goui2/ui/base"
)

type EventSettingFactor base.EventDef

func (f EventSettingFactor) AsSet(s string) base.MDSettingEventDef {
	return base.MDSettingEventDef{base.EventDef(f)}
}

const (
	Event_parseError         = "parseError"
	Event_propertyChange     = "propertyChange"
	Event_requestCompleted   = "requestCompleted"
	Event_requestFailed      = "requestFailed"
	Event_requestSend        = "requestSend"
	Event_changeEvent        = "change"
	Event_dataReceivedEvent  = "dataReceived"
	Event_dataRequestedEvent = "dataRequested"
	Event_filterEvent        = "filter"
	Event_sortEvent          = "sort"
)

var (
	Event_parseError_       = EventSettingFactor(base.EventDef{Name: "parseError"})
	Event_propertyChange_   = EventSettingFactor(base.EventDef{Name: "propertyChange"})
	Event_requestCompleted_ = EventSettingFactor(base.EventDef{Name: "requestCompleted"})
	Event_requestFailed_    = EventSettingFactor(base.EventDef{Name: "requestFailed"})
	Event_requestSend_      = EventSettingFactor(base.EventDef{Name: "requestSend"})
	// Event for Binding
	Event_changeEvent_ = EventSettingFactor(base.EventDef{Name: "change"})
	// Event for Binding
	Event_dataReceivedEvent_ = EventSettingFactor(base.EventDef{Name: "dataReceived"})
	// Event for Binding
	Event_dataRequestedEvent_ = EventSettingFactor(base.EventDef{Name: "dataRequested"})
	// Event for ListBinding
	Event_filterEvent_ = EventSettingFactor(base.EventDef{Name: "filter"})
	Event_sortEvent_   = EventSettingFactor(base.EventDef{Name: "sort"})
)
