package m

import (
	"github.com/goui2/ui/base"
	"github.com/goui2/ui/core"
)

const (
	Class_Button         = "goui.ui.m.Button"
	Class_InputBase      = "goui.ui.m.InputBase"
	Class_Input          = "goui.ui.m.Input"
	Class_Label          = "goui.ui.m.Label"
	Class_List           = "goui.ui.m.List"
	Class_ListBase       = "goui.ui.m.ListBase"
	Class_ListItemBase   = "goui.ui.m.ListItemBase"
	Class_Panel          = "goui.ui.m.Panel"
	Class_CustomListItem = "goui.ui.m.CustomListItem"

	Event_inputHTML = "inputHTML"
	Event_expand    = "expand"

	//Fired when the context menu is opened. When the context menu is opened, the binding context of the item is set to the given contextMenu.
	Event_beforeOpenContextMenu = "beforeOpenContextMenu"

	//Fires when delete icon is pressed by user.

	Event_delete = "delete"

	//Fires when an item is pressed unless the item's type property is Inactive.

	Event_itemPress = "itemPress"

	//Fires when selection is changed via user interaction inside the control.

	Event_selectionChange = "selectionChange"

	//Fires after user's swipe action and before the swipeContent is shown. On the swipe event handler, swipeContent can be changed according to the swiped item. Calling the preventDefault method of the event cancels the swipe action.

	Event_swipe = "swipe"

	// Fires after items binding is updated and processed by the control.
	Event_updateFinished = "updateFinished"

	//Fires before items binding is updated (e.g. sorting, filtering)
	//Note: Event handler should not invalidate the control.
	Event_updateStarted = "updateStarted"

	Event_press = "press"
)

var (
	MD_InputBase = base.Extend(core.MD_Control, Class_InputBase,
		base.MDSettingConstructor(constructorInputBase),
		base.MDSettingPropertyDef{base.PropertyDef{Name: "value", Type: base.PropTypeString, Required: false, DefaultValue: ""}},
		base.MDSettingPropertyDef{base.PropertyDef{Name: "name", Type: base.PropTypeString, Required: true, DefaultValue: ""}},
		base.MDSettingPropertyDef{base.PropertyDef{Name: "editable", Type: base.PropTypeBool, Required: true, DefaultValue: true}},
	).(core.ControlMetaData)
	MD_Input = base.Extend(MD_InputBase, Class_Input,
		base.MDSettingConstructor(constructorInput),
		base.MDSettingPropertyDef{base.PropertyDef{Name: "type", Type: PropTypeInputType, Required: true, DefaultValue: InputType_Text}},
		base.MDSettingEventDef{base.EventDef{Name: Event_inputHTML}},
	).(core.ControlMetaData)
	MD_Label = base.Extend(core.MD_Control, Class_Label,
		base.MDSettingConstructor(constructorLabel),
		base.MDSettingPropertyDef{base.PropertyDef{Name: "text", Type: base.PropTypeString, Required: false, DefaultValue: ""}},
	).(core.ControlMetaData)
	MD_ListBase = base.Extend(core.MD_Control, Class_ListBase,
		base.MDSettingConstructor(constructorListBase),
		base.MDSettingPropertyDef{base.PropertyDef{Name: "mode", Type: PropTypeListMode, DefaultValue: ListMode_None}},
		base.MDSettingAggregationDef{base.AggregationDef{Name: "items", Cardinality: "0..n", Type: TypeCheckListItemBase}},
		base.MDSettingEventDef{base.EventDef{Name: Event_beforeOpenContextMenu}},
		base.MDSettingEventDef{base.EventDef{Name: Event_delete}},
		base.MDSettingEventDef{base.EventDef{Name: Event_itemPress}},
		base.MDSettingEventDef{base.EventDef{Name: Event_selectionChange}},
		base.MDSettingEventDef{base.EventDef{Name: Event_swipe}},
		base.MDSettingEventDef{base.EventDef{Name: Event_updateFinished}},
		base.MDSettingEventDef{base.EventDef{Name: Event_updateStarted}},
	).(core.ControlMetaData)
	MD_ListItemBase = base.Extend(core.MD_Control, Class_ListItemBase,
		base.MDSettingConstructor(constructorListItemBase),
		base.MDSettingPropertyDef{base.PropertyDef{Name: "counter", Type: base.PropTypeInt, Required: true, DefaultValue: 0}},
		base.MDSettingPropertyDef{base.PropertyDef{Name: "selected", Type: base.PropTypeBool, Required: true, DefaultValue: false}},
		base.MDSettingPropertyDef{base.PropertyDef{Name: "unread", Type: base.PropTypeBool, Required: true, DefaultValue: false}},
		base.MDSettingPropertyDef{base.PropertyDef{Name: "visible", Type: base.PropTypeBool, Required: true, DefaultValue: false}},
	).(core.ControlMetaData)
	MD_List = base.Extend(MD_ListBase, Class_List,
		base.MDSettingConstructor(constructorList),
		base.MDSettingAggregationDef{base.AggregationDef{Name: "content", Cardinality: "0..n", Type: core.TypeCheckControl}},
		base.MDSettingEventDef{base.EventDef{Name: Event_expand}},
	).(core.ControlMetaData)
	MD_CustomListItem = base.Extend(MD_ListItemBase, Class_CustomListItem,
		base.MDSettingConstructor(constructorCustomListItem),
		base.MDSettingAggregationDef{base.AggregationDef{Name: "content", Cardinality: "0..n", Type: core.TypeCheckControl}},
	).(core.ControlMetaData)
	MD_Panel = base.Extend(core.MD_Control, Class_Panel,
		base.MDSettingConstructor(constructorPanel),
		base.MDSettingAggregationDef{base.AggregationDef{Name: "content", Cardinality: "0..n", Type: core.TypeCheckControl}},
	).(core.ControlMetaData)
	MD_Button = base.Extend(core.MD_Control, Class_Button,
		base.MDSettingConstructor(constructorButton),
		base.MDSettingEventDef{base.EventDef{Name: Event_press}},
		base.MDSettingPropertyDef{base.PropertyDef{Name: "text", Type: base.PropTypeString, Required: false, DefaultValue: ""}},
	).(core.ControlMetaData)
)
