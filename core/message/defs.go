package message

import (
	"github.com/goui2/ui/base"
)

const (
	Class_MessageProcessor = "goui.ui.core.message.MessageProcessor"
)

var (
	MD_MessageProcessor_default = []base.MetaDataSetting{
		base.MDSettingConstructor(constructorMessageProcessor),
		Event_MessageChange.AsSet(Class_MessageProcessor),
	}

	MD_MessageProcessor = base.Extend(base.MD_EventProvider, Class_MessageProcessor,
		MD_MessageProcessor_default...,
	).(base.EventProviderMetadata)
)

func init() {

}
