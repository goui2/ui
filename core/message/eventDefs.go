package message

import (
	"github.com/goui2/ui/base"
)

type EventSettingFactor base.EventDef

func (f EventSettingFactor) AsSet(s string) base.MDSettingEventDef {
	return base.MDSettingEventDef{base.EventDef(f)}
}

var (
	Event_MessageChange = EventSettingFactor(base.EventDef{Name: "MessageChange"})
)
