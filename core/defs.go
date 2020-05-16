package core

import "github.com/goui2/ui/base"

const (
	Class_Element = "goui.ui.core.Element"
	Class_Control = "goui.ui.core.Control"
)

var (
	PD_Controle_visible = base.MDSettingPropertyDef{base.PropertyDef{Name: "visible", Type: base.PropTypeBool, Required: true, DefaultValue: true}}
)

var (
	MD_Element = base.Extend(base.MD_ManagedObject, Class_Element,
		base.MDSettingConstructor(constructorElement),
		base.MetaModelBuilder(elementMetaDataBuilder),
	)
	MD_Control = base.Extend(MD_Element, Class_Control,
		base.MDSettingConstructor(constructorControl),
		base.MetaModelBuilder(controlMetaDataBuilder),
		PD_Controle_visible,
	)
)
