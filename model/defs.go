package model

import (
	"github.com/goui2/ui/base"
	"github.com/goui2/ui/core/message"
)

const (
	Class_Model           = "goui.model.Model"
	Class_StructModel     = "goui.model.StructModel"
	Class_Binding         = "goui.model.Binding"
	Class_PropertyBinding = "goui.model.PropertyBinding"
	Class_ContextBinding  = "goui.model.ContextBinding"
	Class_ListBinding     = "goui.model.ListBinding"
)

var (
	MD_Model           base.EventProviderMetadata
	MD_StructModel     base.EventProviderMetadata
	MD_Binding         base.EventProviderMetadata
	MD_PropertyBinding base.EventProviderMetadata
	MD_ContextBinding  base.EventProviderMetadata
	MD_ListBinding     base.EventProviderMetadata
)

func init() {
	MD_Model = base.Extend(message.MD_MessageProcessor, Class_Model,
		base.MDSettingConstructor(constructorModel),
		Event_parseError_.AsSet(Class_Model),
		Event_propertyChange_.AsSet(Class_Model),
		Event_requestCompleted_.AsSet(Class_Model),
		Event_requestFailed_.AsSet(Class_Model),
		Event_requestSend_.AsSet(Class_Model),
		base.MDSettingAbstract(true),
	).(base.EventProviderMetadata)
	MD_StructModel = base.Extend(MD_Model, Class_StructModel,
		base.MDSettingConstructor(constructorStructModel),
	).(base.EventProviderMetadata)
	MD_Binding = base.Extend(base.MD_EventProvider, Class_Binding,
		Event_changeEvent_.AsSet(Class_Binding),
		Event_dataReceivedEvent_.AsSet(Class_Binding),
		Event_dataRequestedEvent_.AsSet(Class_Binding),
		base.MDSettingConstructor(constructorBinding),
	).(base.EventProviderMetadata)
	MD_PropertyBinding = base.Extend(MD_Binding, Class_PropertyBinding,
		base.MDSettingConstructor(constructorPropertyBinding),
	).(base.EventProviderMetadata)
	MD_ContextBinding = base.Extend(MD_Binding, Class_ContextBinding,
		base.MDSettingConstructor(constructorContextBinding),
	).(base.EventProviderMetadata)
	MD_ListBinding = base.Extend(MD_Binding, Class_ListBinding,
		base.MDSettingConstructor(constructorListBinding),
		Event_filterEvent_.AsSet(Class_ListBinding),
		Event_sortEvent_.AsSet(Class_ListBinding),
	).(base.EventProviderMetadata)
}
