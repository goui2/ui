package core

import (
	"syscall/js"
)

type RenderManager interface {
	OpenStart(tag string)
	OpenStartElement(tag string, item Element)
	OpenEnd()
	Close(tag string)
	WriteControl(item Control)
	WriteText(txt string)
	WriteAttr(name string, value interface{})
	WriteStyle(name, value interface{})
	WriteClass(cls ...string)

	Build() (js.Value, error)
}

type Renderer func(rm RenderManager, item Element)

type renderManager struct {
	domParser js.Value
	content   string
}

func NewRenderManager() RenderManager {
	rm := renderManager{

		domParser: js.Global().Get("DOMParser").New(),
	}
	return &rm
}

func (rm *renderManager) OpenStartElement(tag string, id Element) {
	rm.content += "<" + tag + " id=\"" + id.Id() + "\""
}

func (rm *renderManager) OpenStart(tag string) {
	rm.content += "<" + tag
}
func (rm *renderManager) OpenEnd() {
	rm.content += ">"
}
func (rm *renderManager) Close(tag string) {
	rm.content += "</" + tag + ">"
}
func (rm *renderManager) WriteControl(item Control) {
	renderer := item.Renderer()
	//	logIt("writeControl %#v\n", renderer)
	renderer(rm, item)
}
func (rm *renderManager) WriteText(txt string) {
	rm.content += txt
}
func (rm *renderManager) WriteAttr(name string, value interface{}) {
	rm.content += " " + name + "=\"" + toString(value) + "\""
}
func (rm *renderManager) WriteStyle(name, value interface{}) {

}
func (rm *renderManager) WriteClass(cls ...string) {

}

func (rm *renderManager) Build() (js.Value, error) {
	//logIt("%s\n", rm.content)
	child := rm.domParser.Call("parseFromString", rm.content, "text/html")
	//	logIt("parsed %v\n", child)
	child = child.Get("body")
	//	logIt("selected %v\n", child)
	child = child.Get("children").Index(0)
	//	logIt("selected %v\n", child)
	return child, nil
}

func toString(v interface{}) string {
	switch i := v.(type) {
	case string:
		return i
	default:
		return "[default]"
	}
}
