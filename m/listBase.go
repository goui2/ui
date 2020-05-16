package m

import (
	"log"
	"strconv"
	"strings"
	"sync"
	"syscall/js"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
	"github.com/goui2/ui/model"
)

type ListMode string

func (_ ListMode) IType() base.SettingType {
	return base.Unspecific
}

const (
	ListMode_None         ListMode = "none"
	ListMode_SingleSelect ListMode = "SingleSelect"
	ListMode_MultiSelect  ListMode = "MultiSelect"

//	ListMode_Delete       = "Delete"
)

func TypeCheckListMode(v interface{}) bool {
	_, ok := v.(ListMode)
	return ok
}
func PropTypeListMode(v interface{}) (interface{}, bool) {
	if TypeCheckListMode(v) {
		return v, true
	} else if base.TypeCheckString(v) {
		return ListMode(v.(string)), true
	} else {
		return v, false
	}
}

type ListBase interface {
	core.Control
	AttachBeforeOpenContextMenu(data com.EventData, fn com.EventHandler)
	AttachBeforeOpenContextMenuOnce(data com.EventData, fn com.EventHandler)
	FireBeforeOpenContextMenu(param com.EventParam)
	DetachBeforeOpenContextMenu(fn com.EventHandler)
	HasBeforeOpenContextMenu() bool
	AttachDelete(data com.EventData, fn com.EventHandler)
	AttachDeleteOnce(data com.EventData, fn com.EventHandler)
	FireDelete(param com.EventParam)
	DetachDelete(fn com.EventHandler)
	HasDelete() bool
	AttachItemPress(data com.EventData, fn com.EventHandler)
	AttachItemPressOnce(data com.EventData, fn com.EventHandler)
	FireItemPress(param com.EventParam)
	DetachItemPress(fn com.EventHandler)
	HasItemPress() bool
	AttachSelectionChange(data com.EventData, fn com.EventHandler)
	AttachSelectionChangeOnce(data com.EventData, fn com.EventHandler)
	FireSelectionChange(param com.EventParam)
	DetachSelectionChange(fn com.EventHandler)
	HasSelectionChange() bool
	AttachSwipe(data com.EventData, fn com.EventHandler)
	AttachSwipeOnce(data com.EventData, fn com.EventHandler)
	FireSwipe(param com.EventParam)
	DetachSwipe(fn com.EventHandler)
	HasSwipe() bool
	AttachUpdateFinished(data com.EventData, fn com.EventHandler)
	AttachUpdateFinishedOnce(data com.EventData, fn com.EventHandler)
	FireUpdateFinished(param com.EventParam)
	DetachUpdateFinished(fn com.EventHandler)
	HasUpdateFinished() bool
	AttachUpdateStarted(data com.EventData, fn com.EventHandler)
	AttachUpdateStartedOnce(data com.EventData, fn com.EventHandler)
	FireUpdateStarted(param com.EventParam)
	DetachUpdateStarted(fn com.EventHandler)
	HasUpdateStarted() bool
	GetSelectedItems() []ListItemBase
}

type EventListSelectionChange struct {
	selection []ListItemBase
}

type EventListItemPresse struct {
	item ListItemBase
}

type listBase struct {
	core.Control
	itemFactory      base.ObjectFactory
	innerFrame       js.Value
	handleUpdateHTML com.EventHandler
	itemFrames       []js.Value
	itemControls     []ListItemBase
	binding          com.Binding
	changeHandler    com.EventHandler
	mutex            sync.Mutex
	selectionCB      js.Func
	selection        map[int]bool
}

func constructorListBase(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &listBase{
			itemControls: make([]ListItemBase, 0),
			itemFrames:   make([]js.Value, 0),
			//			itemFactory:  s.ItemFactory,
			selection: make(map[int]bool),
		}
		listMode := ListMode_SingleSelect
		parentSettings := base.SelectInstanceSettings(s,
			func(item base.InstanceSetting) bool {
				switch d := item.(type) {
				case base.ObjectFactorySetting:
					mo.itemFactory = d.Factory
					return false
				case ListMode:
					listMode = d
					return false
				}
				return true
			},
		)
		parentSettings = append(parentSettings, core.ISRenderer{mo.renderer})
		parentSettings = base.AdjustSelfSetting(mo, md, parentSettings)
		mo.Control = parent.New(id, parentSettings...).(core.Control)
		mo.Property("mode").Set(listMode)
		updateHTML := com.MakeHandler(mo.updateHTML)
		mo.Aggregation("items").AttachUpdateHTML(nil, updateHTML)

		return mo
	}
}

func (l *listBase) AttachBeforeOpenContextMenu(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEvent(Event_beforeOpenContextMenu, data, fn)
}
func (l *listBase) AttachBeforeOpenContextMenuOnce(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEventOnce(Event_beforeOpenContextMenu, data, fn)
}
func (l *listBase) FireBeforeOpenContextMenu(param com.EventParam) {
	l.Control.FireEvent(Event_beforeOpenContextMenu, param)
}
func (l *listBase) DetachBeforeOpenContextMenu(fn com.EventHandler) {
	l.Control.DetachEvent(Event_beforeOpenContextMenu, fn)
}
func (l *listBase) HasBeforeOpenContextMenu() bool {
	return l.Control.HasListener(Event_beforeOpenContextMenu)
}

func (l *listBase) AttachDelete(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEvent(Event_delete, data, fn)
}
func (l *listBase) AttachDeleteOnce(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEventOnce(Event_delete, data, fn)
}
func (l *listBase) FireDelete(param com.EventParam) {
	l.Control.FireEvent(Event_delete, param)
}
func (l *listBase) DetachDelete(fn com.EventHandler) {
	l.Control.DetachEvent(Event_delete, fn)
}
func (l *listBase) HasDelete() bool {
	return l.Control.HasListener(Event_delete)
}

func (l *listBase) AttachItemPress(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEvent(Event_itemPress, data, fn)
}
func (l *listBase) AttachItemPressOnce(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEventOnce(Event_itemPress, data, fn)
}
func (l *listBase) FireItemPress(param com.EventParam) {
	l.Control.FireEvent(Event_itemPress, param)
}
func (l *listBase) DetachItemPress(fn com.EventHandler) {
	l.Control.DetachEvent(Event_itemPress, fn)
}
func (l *listBase) HasItemPress() bool {
	return l.Control.HasListener(Event_itemPress)
}

func (l *listBase) AttachSelectionChange(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEvent(Event_selectionChange, data, fn)
}
func (l *listBase) AttachSelectionChangeOnce(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEventOnce(Event_selectionChange, data, fn)
}
func (l *listBase) FireSelectionChange(param com.EventParam) {
	l.Control.FireEvent(Event_selectionChange, param)
}
func (l *listBase) DetachSelectionChange(fn com.EventHandler) {
	l.Control.DetachEvent(Event_selectionChange, fn)
}
func (l *listBase) HasSelectionChange() bool {
	return l.Control.HasListener(Event_selectionChange)
}

func (l *listBase) AttachSwipe(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEvent(Event_swipe, data, fn)
}
func (l *listBase) AttachSwipeOnce(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEventOnce(Event_swipe, data, fn)
}
func (l *listBase) FireSwipe(param com.EventParam) {
	l.Control.FireEvent(Event_swipe, param)
}
func (l *listBase) DetachSwipe(fn com.EventHandler) {
	l.Control.DetachEvent(Event_swipe, fn)
}
func (l *listBase) HasSwipe() bool {
	return l.Control.HasListener(Event_swipe)
}
func (l *listBase) AttachUpdateFinished(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEvent(Event_updateFinished, data, fn)
}
func (l *listBase) AttachUpdateFinishedOnce(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEventOnce(Event_updateFinished, data, fn)
}
func (l *listBase) FireUpdateFinished(param com.EventParam) {
	l.Control.FireEvent(Event_updateFinished, param)
}
func (l *listBase) DetachUpdateFinished(fn com.EventHandler) {
	l.Control.DetachEvent(Event_updateFinished, fn)
}
func (l *listBase) HasUpdateFinished() bool {
	return l.Control.HasListener(Event_updateFinished)
}

func (l *listBase) AttachUpdateStarted(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEvent(Event_updateStarted, data, fn)
}
func (l *listBase) AttachUpdateStartedOnce(data com.EventData, fn com.EventHandler) {
	l.Control.AttachEventOnce(Event_updateStarted, data, fn)
}
func (l *listBase) FireUpdateStarted(param com.EventParam) {
	l.Control.FireEvent(Event_updateStarted, param)
}
func (l *listBase) DetachUpdateStarted(fn com.EventHandler) {
	l.Control.DetachEvent(Event_updateStarted, fn)
}
func (l *listBase) HasUpdateStarted() bool {
	return l.Control.HasListener(Event_updateStarted)
}

func (l *listBase) BindAggregation2(n string, b com.BindingDescription, of base.ObjectFactory) {
	l.Control.BindAggregation(n, b, of)
}

func (l *listBase) BindObjectOld(ub com.Binding) {
	lb, ok := ub.(com.ListBinding)
	if !ok {
		panic("binding must be from type ListBinding")
	}
	if l.binding != nil && l.changeHandler != nil {
		l.binding.DetachChange(l.changeHandler)
	}
	if l.changeHandler == nil {
		l.changeHandler = com.MakeHandler(func(e com.Event) {
			l.handleChange()
		})
	}
	l.binding = lb
	l.binding.AttachChange(nil, l.changeHandler)
	l.handleChange()
}

func (l *listBase) handleChange() {
	logIt("listBase.handleChange")
	//	l.mutex.Lock()
	//	defer l.mutex.Unlock()
	ctx := l.binding.GetContext()
	length := ctx.GetValue("").ArraySize()
	mdl := ctx.GetModel().(model.Model)
	path := model.NewPath(ctx.GetPath())
	itemControls := make([]ListItemBase, 0)
	itemControls = append(itemControls, l.itemControls...)
	for i := len(l.itemControls); i < length; i++ {
		itemPath := path.Append(model.NewPath(strconv.Itoa(i)))
		itemBind := mdl.BindContext(itemPath.String())
		itemCtrl := l.itemFactory(itemBind).(ListItemBase)
		itemControls = append(itemControls, itemCtrl)
	}
	intA := make([]interface{}, length)
	for i := 0; i < length; i++ {
		intA[i] = itemControls[i]
	}
	logIt("listBase.handleChange setAggregation items %#v", itemControls)
	l.Aggregation("items").Set(intA)
	l.updateHTML(nil)
}

func (p *listBase) renderer(rm core.RenderManager, item core.Element) {
	rm.OpenStartElement("div", p)
	rm.OpenEnd()
	rm.OpenStart("ul")
	rm.WriteAttr("id", item.Id()+"-in")
	rm.OpenEnd()
	for pos, child := range p.Aggregation("items").Get() {
		ctrl := child.(core.Control)
		p.renderItem(rm, ctrl, pos)
	}
	rm.Close("ul")
	rm.Close("div")
}

func (p *listBase) childId(pos int, s string) string {
	return p.Id() + "-" + s + "-" + strconv.Itoa(pos)
}

func (p *listBase) renderItem(rm core.RenderManager, item core.Control, pos int) {
	rm.OpenStart("li")
	rm.WriteAttr("id", p.childId(pos, "frame"))
	rm.OpenEnd()
	rm.WriteControl(item)
	rm.Close("li")
}

func (p *listBase) OnAfterRendering(v js.Value) {
	logIt("listBase.OnAfterRendering %s - %#v", p.Id(), v)
	p.Control.OnAfterRendering(v)
	children := p.Aggregation("items").Get()
	p.innerFrame, _ = p.FindDomById(p.Id() + "-in")
	for i, abschild := range children {
		child := abschild.(ListItemBase)
		html, _ := p.FindDomChild(child)
		p.itemControls = append(p.itemControls, child)
		child.OnAfterRendering(html)
		html, _ = p.FindDomById(p.childId(i, "frame"))
		p.itemFrames = append(p.itemFrames, html)
	}
	p.handleUpdateHTML = com.MakeHandler(p.updateHTML)
	p.Aggregation("items").AttachUpdateHTML(nil, p.handleUpdateHTML)
	p.selectionHandler()
	logIt("listBase.OnAfterRendering %s - %#v exit", p.Id(), v)
}

func (l *listBase) selectionHandler() {
	listId := l.innerFrame.Get("id").String()
	modeProp := l.Property("mode")
	if dom, ok := l.DomRef(); ok {
		l.selectionCB = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if modeProp.Get() == ListMode_None {
				return nil
			}
			element := args[0].Get("srcElement")
			parent := element.Get("parentElement")
			for !parent.IsNull() && !parent.IsUndefined() && parent.Get("id").String() != listId {
				element = parent
				parent = element.Get("parentElement")
			}
			if !element.IsNull() && !element.IsUndefined() {
				parts := strings.Split(element.Get("id").String(), "-")
				posParts := len(parts)
				if posParts > 1 {
					pos, _ := strconv.Atoi(parts[posParts-1])
					if pos < len(l.itemControls) {
						l.FireItemPress(NewEvent(l, EventListItemPresse{l.itemControls[pos]}))
					}
					switch modeProp.Get().(ListMode) {
					case ListMode_MultiSelect:
						l.selectionMultiSelect(pos)
					case ListMode_SingleSelect:
						l.selectionSingleSelect(pos)
					}
				}
			}
			return nil
		})
		dom.Call("addEventListener", "click", l.selectionCB, js.ValueOf(true))
	}
}

func (l *listBase) selectionSingleSelect(pos int) {
	log.Printf("listBase.selectionSingleSelect(%d)", pos)
	for k := range l.selection {
		delete(l.selection, k)
	}
	if pos < len(l.itemFrames) {
		l.selection[pos] = true
	}
	l.selectionUpdate()
	if pos < len(l.itemFrames) {
		l.FireSelectionChange(NewEvent(l, EventListSelectionChange{l.GetSelectedItems()}))
	}
}

func (l *listBase) selectionUpdate() {
	jsSelected := js.ValueOf("selected")
	for _, html := range l.itemFrames {
		classList := html.Get("classList")
		classList.Call("remove", jsSelected)
	}
	for pos, html := range l.itemFrames {
		if selected, ok := l.selection[pos]; ok && selected {
			classList := html.Get("classList")
			classList.Call("add", jsSelected)
		}
	}
}

func (l *listBase) selectionMultiSelect(pos int) {
	if selected, ok := l.selection[pos]; ok {
		l.selection[pos] = !selected
	} else {
		l.selection[pos] = true
	}
	l.selectionUpdate()
	l.FireSelectionChange(NewEvent(l, EventListSelectionChange{l.GetSelectedItems()}))
}

func (l *listBase) GetSelectedItems() []ListItemBase {
	result := make([]ListItemBase, 0)
	for i, ctrl := range l.itemControls {
		if selected, ok := l.selection[i]; ok && selected {
			result = append(result, ctrl)
		}
	}
	return result
}

func (p *listBase) updateHTML(e com.Event) {
	logIt("listBase.updateHTML")
	children := p.Aggregation("items").Get()
	oldItemControls := p.itemControls
	oldItemFrames := p.itemFrames
	p.itemControls = make([]ListItemBase, len(children))
	p.itemFrames = make([]js.Value, len(children))
	for pos, child := range children {
		ctrl := child.(ListItemBase)
		if pos < len(oldItemControls) {
			p.itemFrames[pos] = oldItemFrames[pos]
			if oldItemControls[pos] != ctrl {
				html, _ := oldItemControls[pos].DomRef()
				p.itemFrames[pos].Call("removeChild", html)
				rm := core.NewRenderManager()
				rm.WriteControl(ctrl)
				html, _ = rm.Build()
				p.itemFrames[pos].Call("append", html)
				p.itemControls[pos] = ctrl
				ctrl.OnAfterRendering(html)
			} else {
				p.itemControls[pos] = oldItemControls[pos]
			}
		} else {
			logIt("create new %d %v", pos, ctrl)
			rm := core.NewRenderManager()
			rm.OpenStart("ul")
			rm.OpenEnd()
			p.renderItem(rm, ctrl, pos)
			rm.Close("ul")
			html, _ := rm.Build()
			htmlElement := htmlElement{html}
			htmlElmnt2, _ := htmlElement.FindElementById(p.childId(pos, "frame"))
			p.itemFrames[pos] = htmlElmnt2.Html
			p.innerFrame.Call("append", p.itemFrames[pos])
			html, _ = p.FindDomChild(ctrl)
			ctrl.OnAfterRendering(html)
			p.itemControls[pos] = ctrl
		}
	}
	for i := len(p.itemFrames); i < len(oldItemFrames); i++ {
		p.innerFrame.Call("removeChild", oldItemFrames[i])
	}
	logIt("listBase.updateHTML exit")
}

func (l *listBase) Destroy() {
	for _, ctrl := range l.itemControls {
		ctrl.Destroy()
	}
	l.selectionCB.Release()
}

var listBaseCss = `
   li.selected { color : red }
`

func init() {
	core.GetCore().RegisterCSS(listBaseCss)
}
