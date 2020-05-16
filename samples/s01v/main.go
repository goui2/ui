package main

import (
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
	"github.com/goui2/ui/core/mvc"
	"github.com/goui2/ui/m"
	"github.com/goui2/ui/model"
)

type Data1 struct {
	Field1 string
}

const view01 = `
<mvc:View controllerName="simpleController" xmlns:core="goui.core"
xmlns:mvc="goui.ui.core.mvc" xmlns="goui.ui.m" xmlns:table="sap.ui.table"
xmlns:html="http:www.w3.org/1999/xhtml">
   <mvc:content>
	 <Panel id="panel#">
	   <content>
	     <Input value="{/Field1}"/>
		 <Input value="{/Field1}"/>
		 <Label text="{/Field1}"/>
		 <Panel id="pnl2#">

		 </Panel>
	  </content>
	 </Panel>
   </mvc:content>
</mvc:View>
`

func viewDef() (io.Reader, error) {
	return strings.NewReader(view01), nil
}

var (
	MD_Tick = base.Extend(base.MD_EventProvider, "epTick",
		base.MDSettingEventDef{base.EventDef{Name: "tick"}},
	)
	createInput = m.MD_Input.GetClass()
	createLabel = m.MD_Label.GetClass()
	createPanel = m.MD_Panel.GetClass()
)

func main() {
	mvc.RegisterController("simpleController", simpleController)
	d1 := &Data1{
		Field1: "show me",
	}
	ep := MD_Tick.GetClass().New("tick").(base.EventProvider)
	c := make(chan struct{}, 0)
	vr, _ := viewDef()
	mdl := model.NewStructModel(nil)
	mdl = core.GetCore().SetModel("", mdl)
	viewCreator := mvc.NewXMLView("v1", vr)
	view := viewCreator()
	core.GetCore().StartApplication("appl", view)
	log.Printf("########################Set1 %v\n", d1)
	mdl.SetObject(d1)
	d1 = &Data1{
		Field1: "show me2",
	}

	mdl.SetObject(d1)

	changeIt := view.ElementById("pnl2")
	content := changeIt.Aggregation("content")
	dynCtrls := make([]core.Control, 50)
	for i := 0; i < len(dynCtrls); i++ {
		id := "Row#" + strconv.Itoa(i)
		txt := "Row #" + strconv.Itoa(i)
		dynLabel := createLabel.New(id).(m.Label)
		dynLabel.Property("text").Set(txt)
		dynCtrls[i] = dynLabel
	}
	ticker := time.NewTicker(100 * time.Millisecond)
	offset := 2
	direction := 1
	tick := func(e com.Event) {
		offset += direction
		if offset > len(dynCtrls) {
			direction = -1
		} else if offset < 0 {
			direction = 1
		} else {
			data := make([]interface{}, offset)
			for i, k := range dynCtrls[0:offset] {
				data[i] = k
			}
			content.Set(data)
		}
	}

	ep.AttachEvent("tick", nil, com.MakeHandler(tick))
	for {
		select {
		case <-ticker.C:
			ep.FireEvent("tick", nil)
		}
	}

	<-c
}
