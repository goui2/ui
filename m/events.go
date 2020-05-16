package m

import (
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
)

type eventNotation string

func (en eventNotation) Data() interface{} {
	return string(en)
}

type commonEvent struct {
	source core.Control
	data   interface{}
}

func (c commonEvent) Source() core.Control {
	return c.source
}

func (c commonEvent) Param() interface{} {
	return c.data
}

func NewEvent(c core.Control, data interface{}) com.EventParam {
	return commonEvent{c, data}
}
