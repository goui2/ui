package main

import (
	"io"
	"log"
	"strings"
	"time"

	"github.com/goui2/ui/core"
	"github.com/goui2/ui/core/mvc"
	_ "github.com/goui2/ui/m"
	"github.com/goui2/ui/model"
)

type Data struct {
	Items []DataItem
}
type DataItem struct {
	Field1 string
	Field2 string
	Field3 string
}

var (
	data = &Data{
		Items: []DataItem{
			{"a1", "b1", "c1"},
			{"a2", "b2", "c2"},
			{"a3", "b3", "c3"},
		},
	}
)

const view01 = `
<mvc:View controllerName="testdata.complexsyntax" xmlns:core="goui.core"
xmlns:mvc="goui.ui.core.mvc" xmlns="goui.ui.m" xmlns:table="sap.ui.table"
xmlns:html="http:www.w3.org/1999/xhtml">
   <mvc:content>
	 <List id="list#" items="{/Items}">
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
   </mvc:content>
</mvc:View>
`

func viewDef() (io.Reader, error) {
	return strings.NewReader(view01), nil
}

func main() {
	go printData()
	vr, _ := viewDef()
	c := make(chan struct{}, 0)
	mdl := model.NewStructModel(nil)
	mdl = core.GetCore().SetModel("", mdl)
	//	mdl, _ = core.GetCore().GetModel("")
	viewCreator := mvc.NewXMLView("v1", vr)
	view := viewCreator()
	core.GetCore().StartApplication("appl", view)
	log.Printf("########################Set1 %v\n", data)
	mdl.SetObject(data)

	<-c
}

func printData() {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
		}
		log.Printf("change: %v", data)
	}
}
