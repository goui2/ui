package base

const (
	Class_ManagedObject     = "goui.base.ManagedObject"
	Class_EventProvider     = "goui.base.EventProvider"
	Class_EventProviderProp = "goui.base.EventProvider$Prop"
	Class_EventProviderAggr = "goui.base.EventProvider$Aggr"
)

var (
	MD_Object            = createObject()
	MD_EventProvider     EventProviderMetadata
	MD_EventProviderProp EventProviderMetadata
	MD_EventProviderAggr EventProviderMetadata
	MD_ManagedObject     ManagedObjectMetaData

	MD_EventProvider_Defaults = []MetaDataSetting{
		MDSettingConstructor(constructEventProvider),
		MetaModelBuilder(eventProviderMetadataBuilder),
	}
	MD_EventProviderProp_Defaults = []MetaDataSetting{
		Event_changed.AsSet(Class_EventProviderProp),
		Event_updateHTML.AsSet(Class_EventProviderProp),
	}
	MD_EventProviderAggr_Defaults = []MetaDataSetting{
		Event_changed.AsSet(Class_EventProviderAggr),
		Event_updateHTML.AsSet(Class_EventProviderAggr),
	}

	MD_ManagedObject_Defaults = []MetaDataSetting{
		Event_formatError.AsSet(Class_ManagedObject),
		Event_modelContextChange.AsSet(Class_ManagedObject),
		Event_parseError.AsSet(Class_ManagedObject),
		Event_validationError.AsSet(Class_ManagedObject),
		Event_validationSuccess.AsSet(Class_ManagedObject),
		MDSettingConstructor(constructorManagedObject),
		MetaModelBuilder(managedObjectMetaDataBuilder),
	}
)

func init() {
	MD_EventProvider = Extend(MD_Object, Class_EventProvider, MD_EventProvider_Defaults...).(EventProviderMetadata)
	MD_EventProviderProp = Extend(MD_EventProvider, Class_EventProviderProp, MD_EventProviderProp_Defaults...).(EventProviderMetadata)
	MD_EventProviderAggr = Extend(MD_EventProvider, Class_EventProviderAggr, MD_EventProviderAggr_Defaults...).(EventProviderMetadata)
	MD_ManagedObject = Extend(MD_EventProvider, Class_ManagedObject, MD_ManagedObject_Defaults...).(ManagedObjectMetaData)

}

func createObject() MetaData {
	m := &metaData{
		name:     "goui.base.Object",
		parent:   nil,
		abstract: false,
		final:    false,
	}
	m.constructor = func(id string, s ...InstanceSetting) Object {
		obj := &object{
			metaData: m,
		}
		obj.self = obj
		for _, o := range s {
			if self, ok := o.(SelfSetting); ok {
				obj.self = self.Object
				obj.metaData = self.MetaData
				return obj
			}
		}
		return obj
	}
	m.builder = baseObjectBuilder
	return m
}
