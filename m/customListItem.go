package m

import (
	"syscall/js"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
)

type CustomerListItem interface {
	ListItemBase
}

type customerListItem struct {
	ListItemBase
	handleUpdateHTML com.EventHandler
	htmlChildren     []js.Value
}

func constructorCustomListItem(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &customerListItem{}
		parentSettings := append(s, core.ISRenderer{mo.renderer})
		mo.ListItemBase = parent.New(id, parentSettings...).(ListItemBase)

		return mo
	}
}

func (p *customerListItem) renderer(rm core.RenderManager, item core.Element) {
	rm.OpenStartElement("div", p)
	rm.OpenEnd()
	for _, child := range p.Aggregation("content").Get() {
		ctrl := child.(core.Control)
		rm.WriteControl(ctrl)
	}
	rm.Close("div")
}

func (p *customerListItem) OnAfterRendering(v js.Value) {
	p.ListItemBase.OnAfterRendering(v)
	children := p.Aggregation("content").Get()
	for _, abschild := range children {
		child := abschild.(ListItemBase)
		html, _ := p.FindDomChild(child)
		child.OnAfterRendering(html)
		p.htmlChildren = append(p.htmlChildren, html)
	}
	p.handleUpdateHTML = com.MakeHandler(func(e com.Event) { p.updateHTML(e) })
	p.Aggregation("content").AttachUpdateHTML(nil, p.handleUpdateHTML)
}

func (p *customerListItem) updateHTML(e com.Event) {
	logIt("updateHTML")
	parentRef, ok := p.DomRef()
	if !ok {
		return
	}
	for _, lbi := range p.htmlChildren {
		parentRef.Call("removeChild", lbi)
	}
	p.htmlChildren = make([]js.Value, 0)
	children := p.Aggregation("content").Get()
	for _, abschild := range children {
		child := abschild.(core.Control)
		rm := core.NewRenderManager()
		rm.WriteControl(child)
		html, _ := rm.Build()
		parentRef.Call("append", html)
		child.OnAfterRendering(html)
		p.htmlChildren = append(p.htmlChildren, html)
	}
}
