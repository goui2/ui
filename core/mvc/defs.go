package mvc

import (
	"github.com/goui2/ui/base"
	"github.com/goui2/ui/core"
)

const (
	Class_View    = "goui.ui.core.mvc.View"
	Class_XmlView = "goui.ui.core.mvc.XMLView"
)

var (
	MD_View = base.Extend(core.MD_Control, Class_View,
		base.MDSettingConstructor(constructorView),
		base.MDSettingAggregationDef{base.AggregationDef{Name: "content", Cardinality: "0..n", Type: core.TypeCheckControl}},
		base.MDSettingPropertyDef{base.PropertyDef{Name: "width", Type: core.PropTypeCSSSize, DefaultValue: core.CSSSizeOf("100%")}},
		base.MDSettingPropertyDef{base.PropertyDef{Name: "height", Type: core.PropTypeCSSSize, DefaultValue: core.CSSSizeOf("100%")}},
	).(core.ControlMetaData)
	MD_XmlView = base.Extend(MD_View, Class_XmlView,
		base.MDSettingConstructor(constructorXmlView),
	).(core.ControlMetaData)
)
