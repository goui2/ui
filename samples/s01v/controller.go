package main

import (
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core/mvc"
)

type controller struct {
	view mvc.View
}

func (c *controller) Init() {

}
func (c *controller) Handler(name string) (com.EventHandler, bool) {
	switch name {
	default:
		return nil, false
	}
}

func (c *controller) Destroy() {}

func simpleController(v mvc.View) mvc.Controller {

	return &controller{view: v}
}
