package core

import (
	"strings"
	"sync"
	"syscall/js"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/model"
)

type Core interface {
	StartApplication(placeAt string, appl Control)
	RegisterCSS(css string)
	RegisterMetaData(n string, md base.ManagedObjectMetaData)
	GetMetaData(n string) (base.ManagedObjectMetaData, bool)
	SetModel(name string, m model.Model) model.Model
	GetModel(name string) (model.Model, bool)
	GetModelName(m model.Model) (string, bool)
}

type core struct {
	cssString []string
	document  js.Value
	style     js.Value
	head      js.Value
	models    map[string]model.Model
	metaData  map[string]base.ManagedObjectMetaData
	mutex     sync.Mutex
}

var mainCore = &core{
	cssString: make([]string, 0),
	models:    make(map[string]model.Model),
	metaData:  make(map[string]base.ManagedObjectMetaData),
}

func GetCore() Core {
	return mainCore
}
func init() {
	mainCore.document = js.Global().Get("document")
	mainCore.style = mainCore.document.Call("createElement", js.ValueOf("style"))
	mainCore.head = mainCore.document.Get("head")
	mainCore.head.Call("append", mainCore.style)
	styleText := strings.Join(mainCore.cssString, "")
	mainCore.style.Set("innerHTML", js.ValueOf(styleText))
	mainCore.SetModel("", model.NewStructModel(nil))

}

func (c *core) StartApplication(placeAt string, appl Control) {
	applRoot := c.document.Call("getElementById", placeAt)
	rm := NewRenderManager()
	appl.OnBeforeRendering()
	rm.WriteControl(appl)
	applHtml, _ := rm.Build()
	applRoot.Call("appendChild", applHtml)
	appl.OnAfterRendering(applHtml)
}

func (c *core) RegisterCSS(css string) {
	c.cssString = append(c.cssString, css)
	style := mainCore.document.Call("createElement", js.ValueOf("style"))
	mainCore.head.Call("append", style)
	mainCore.style.Set("innerHTML", js.ValueOf(css))
}

func (c *core) SetModel(name string, m model.Model) model.Model {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if old, ok := c.models[name]; ok {
		old.SetObject(m.GetObject())
		return old
	} else {
		c.models[name] = m
		return m
	}
}

func (c *core) GetModel(name string) (m model.Model, ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	m, ok = c.models[name]
	return
}

func (c *core) GetModelName(m model.Model) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, value := range c.models {
		if value == m {
			return key, true
		}
	}
	return "", false
}

func (c *core) RegisterMetaData(n string, md base.ManagedObjectMetaData) {
	c.metaData[n] = md
}
func (c *core) GetMetaData(n string) (md base.ManagedObjectMetaData, ok bool) {
	md, ok = c.metaData[n]
	return
}
