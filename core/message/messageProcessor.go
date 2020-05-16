package message

import (
	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
)

const (
	messageChange = "MessageChange"
)

var (
	messageChangeEvents = []string{messageChange}
)

type MessageProcessor interface {
	base.EventProvider
	AttachMessageChange(data com.EventData, fn com.EventHandler)
	FireMessageChange(param com.EventParam)
	DetachMessageChange(fn com.EventHandler)
	HasMessageChange() bool
}

type messageProcessor struct {
	base.EventProvider
}

func constructorMessageProcessor(md base.MetaData, c base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &messageProcessor{}
		parentSettings := base.AdjustSelfSetting(mo, md, s)
		mo.EventProvider = c(id, parentSettings...).(base.EventProvider)
		return mo
	}
}

func (mp *messageProcessor) AttachMessageChange(data com.EventData, fn com.EventHandler) {
	mp.EventProvider.AttachEvent(messageChange, data, fn)
}
func (mp *messageProcessor) FireMessageChange(param com.EventParam) {
	mp.EventProvider.FireEvent(messageChange, param)
}
func (mp *messageProcessor) DetachMessageChange(fn com.EventHandler) {
	mp.EventProvider.DetachEvent(messageChange, fn)
}
func (mp *messageProcessor) HasMessageChange() bool {
	return mp.HasListener(messageChange)
}
