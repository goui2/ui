package main

import (
	"io"
	"log"
	"strings"

	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
	"github.com/goui2/ui/core/mvc"
	_ "github.com/goui2/ui/m"
	"github.com/goui2/ui/model"
)

type Data struct {
	Input string
	Items []DataItem
}
type DataItem struct {
	Field1 string
	Field2 string
	Field3 string
}

var (
	data = &Data{
		Input: "text",
		Items: []DataItem{
			{"a1", "b1", "c1"},
			{"a2", "b2", "c2"},
			{"a3", "b3", "c3"},
		},
	}
)

const view01 = `
<mvc:View controllerName="simpleController" xmlns:core="goui.core"
xmlns:mvc="goui.ui.core.mvc" xmlns="goui.ui.m" xmlns:table="sap.ui.table"
xmlns:html="http:www.w3.org/1999/xhtml">
   <mvc:content>
	 <Panel>
	   <content>
			<Label text="{/Input}"/>
			<Input value="{/Input}" />
			<Button text="Run" onPress="doRun" />
			<List id="list#" items="{/Items}" mode="MultiSelect">
				<items>
					<CustomListItem>
						<content>	 
							<Input value="{Field1}"/>
							<Input value="{Field2}"/>
							<Input value="{Field3}"/>
							<!--
							-->
						</content>
					</CustomListItem>
				</items>
		    </List>
	   </content>
	 </Panel>
   </mvc:content>
</mvc:View>
`

type controller struct {
	view mvc.View
}

func (c *controller) Init() {
	mdl, _ := core.GetCore().GetModel("")
	mdl.SetObject(data)

}
func (c *controller) Handler(name string) (com.EventHandler, bool) {
	log.Println("handler for", name)
	switch name {
	case "doRun":
		handle := func(e com.Event) {
			di := DataItem{}
			data.Items = append(data.Items, di)
			mdl, _ := core.GetCore().GetModel("")
			mdl.SetObject(data)
			log.Println("press")
		}
		return com.MakeHandler(handle), true
	default:
		return nil, false
	}
}

func (c *controller) Destroy() {}

func simpleController(v mvc.View) mvc.Controller {

	return &controller{view: v}
}

func viewDef() (io.Reader, error) {
	return strings.NewReader(view01), nil
}

func main() {
	vr, _ := viewDef()
	viewCreator := mvc.NewXMLView("v1", vr)
	mdl := model.NewStructModel(nil)
	mdl = core.GetCore().SetModel("", mdl)
	mvc.RegisterController("simpleController", simpleController)
	view := viewCreator()

	core.GetCore().StartApplication("appl", view)

	c := make(chan struct{}, 0)
	<-c
}
