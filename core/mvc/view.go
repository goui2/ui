package mvc

import (
	"syscall/js"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
)

type ViewFactory interface {
	ViewId(string) string
}
type ContentFactory func(ViewFactory) []core.Control

type View interface {
	core.Control
}

type view struct {
	core.Control
	innerFrame       js.Value
	handleUpdateHTML com.EventHandler
}

func constructorView(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &view{}
		parentSettings := append(s, core.ISRenderer{mo.renderer})
		parentSettings = base.AdjustSelfSetting(mo, md, parentSettings)
		mo.Control = parent.New(id, parentSettings...).(core.Control)

		return mo
	}
}

func (p *view) renderer(rm core.RenderManager, item core.Element) {
	rm.OpenStartElement("div", p)
	rm.OpenEnd()
	rm.OpenStart("div")
	rm.WriteAttr("id", item.Id()+"-in")
	rm.OpenEnd()
	for pos, child := range p.Aggregation("content").Get() {
		ctrl := child.(core.Control)
		p.renderItem(rm, ctrl, pos)
	}
	rm.Close("div")
	rm.Close("div")
}

func (p *view) renderItem(rm core.RenderManager, item core.Control, pos int) {
	//	rm.OpenStart("div")
	//	rm.WriteAttr("id", p.childId(pos, "frame"))
	//	rm.OpenEnd()
	rm.WriteControl(item)
	//	rm.Close("div")
}

func (p *view) OnAfterRendering(v js.Value) {
	p.Control.OnAfterRendering(v)
	children := p.Aggregation("content").Get()
	p.innerFrame, _ = p.FindDomById(p.Id() + "-in")
	for _, abschild := range children {
		child := abschild.(core.Control)
		html, _ := p.FindDomChild(child)
		child.OnAfterRendering(html)
	}
	p.handleUpdateHTML = com.MakeHandler(p.updateHTML)
	p.Aggregation("content").AttachUpdateHTML(nil, p.handleUpdateHTML)
}

func (p *view) updateHTML(e com.Event) {
}

func (p *view) ElementById(id string) core.Element {
	if p.Id() == id {
		return p
	}
	viewId := p.Id() + "_" + id
	return p.Control.ElementById(viewId)
}
