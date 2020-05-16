package m

import (
	"strconv"
	"syscall/js"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
)

type Panel interface {
	core.Control
}

type panel struct {
	core.Control
	innerFrame       js.Value
	children         []panelItem
	handleUpdateHTML com.EventHandler
}

type panelItem struct {
	ctrl    core.Control
	frame   js.Value
	visible bool
	pos     int
}

func constructorPanel(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &panel{}
		parentSettings := append(s, core.ISRenderer{mo.renderer})
		parentSettings = base.AdjustSelfSetting(mo, md, parentSettings)
		mo.Control = parent.New(id, parentSettings...).(core.Control)
		return mo
	}
}

func (p *panel) renderer(rm core.RenderManager, item core.Element) {
	rm.OpenStartElement("div", p)
	rm.OpenEnd()
	rm.OpenStart("div")
	rm.WriteAttr("id", item.Id()+"_in")
	rm.OpenEnd()
	for pos, child := range p.Aggregation("content").Get() {
		ctrl := child.(core.Control)
		p.renderItem(rm, pos, ctrl)
	}
	rm.Close("div")
	rm.Close("div")
}

func (p *panel) frameId(pos int) string {
	return p.Id() + "_p" + strconv.Itoa(pos) + "_frame"
}

func (p *panel) renderItem(rm core.RenderManager, pos int, item core.Control) {
	rm.OpenStart("div")
	rm.WriteAttr("id", p.frameId(pos))
	rm.OpenEnd()
	rm.WriteControl(item)
	rm.Close("div")
}

func (p *panel) renderContentItem(pos int, item core.Control) panelItem {
	pi := panelItem{
		ctrl: item,
		pos:  pos,
	}
	rm := core.NewRenderManager()
	p.renderItem(rm, pos, pi.ctrl)
	pi.frame, _ = rm.Build()
	html := pi.frame.Call("querySelector", "#"+pi.ctrl.Id())
	pi.ctrl.OnAfterRendering(html)
	pi.visible = true
	return pi
}

func (p *panel) OnAfterRendering(v js.Value) {
	logIt("panel.OnAfterRendering %s - %#v", p.Id(), v)
	p.Control.OnAfterRendering(v)
	content := p.Aggregation("content").Get()
	p.children = make([]panelItem, len(content))
	p.innerFrame, _ = p.FindDomById(p.Id() + "_in")
	for i, item := range content {
		ctrl := item.(core.Control)
		p.children[i].ctrl = ctrl
		html, _ := p.FindDomChild(ctrl)
		ctrl.OnAfterRendering(html)
		html, _ = p.FindDomById(p.frameId(i))
		p.children[i].frame = html
		p.children[i].visible = true
	}
	p.handleUpdateHTML = com.MakeHandler(func(e com.Event) { p.updateHTML(e) })
	p.Aggregation("content").AttachUpdateHTML(nil, p.handleUpdateHTML)
}

func (p *panel) updateHTML(e com.Event) {
	logIt("panel.updateHTML")
	ctrlMap := make(map[core.Control]panelItem)
	for _, c := range p.children {
		ctrlMap[c.ctrl] = c
	}
	devChildren := p.Aggregation("content").Get()
	newChildren := make([]panelItem, len(devChildren))
	for i, dc := range devChildren {
		devCtrl := dc.(core.Control)
		if oldItem, ok := ctrlMap[devCtrl]; ok {
			newChildren[i] = oldItem
			newChildren[i].visible = true
		} else {
			newChildren[i] = p.renderContentItem(i, devCtrl)
		}
	}
	for _, c := range p.children {
		p.innerFrame.Call("removeChild", c.frame)
	}
	p.children = newChildren
	for _, c := range p.children {
		p.innerFrame.Call("append", c.frame)
	}
}
