package base

type EventSettingFactor EventDef

func (f EventSettingFactor) AsSet(s string) MDSettingEventDef {
	return MDSettingEventDef{EventDef(f)}
}

var (
	Event_formatError        = EventSettingFactor(EventDef{Name: "formatError"})
	Event_modelContextChange = EventSettingFactor(EventDef{Name: "modelContextChange"})
	Event_parseError         = EventSettingFactor(EventDef{Name: "parseError"})
	Event_validationError    = EventSettingFactor(EventDef{Name: "validationError"})
	Event_validationSuccess  = EventSettingFactor(EventDef{Name: "validationSuccess"})
	Event_DataReceived       = EventSettingFactor(EventDef{Name: "DataReceived"})
	Event_DataRequested      = EventSettingFactor(EventDef{Name: "DataRequested"})
	Event_changed            = EventSettingFactor(EventDef{Name: "changed"})
	Event_updateHTML         = EventSettingFactor(EventDef{Name: "updateHTML"})
)
