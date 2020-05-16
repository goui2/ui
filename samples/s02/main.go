package main

import (
	"log"
	"time"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
	"github.com/goui2/ui/m"
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

func main() {
	go printData()
	c := make(chan struct{}, 0)
	mdl := model.NewStructModel(nil)
	mdl = core.GetCore().SetModel("", mdl)
	ctx := mdl.BindContext("/")
	applCtrl := buildAppl(ctx.GetContext())
	core.GetCore().StartApplication("appl", applCtrl)
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

func buildAppl(ctx com.Context) core.Control {
	list := m.MD_List.GetClass().New("list",
		//		base.ObjectFactorySetting{Factory: createListItem},
		m.ListMode_MultiSelect,
	).(m.List)
	list.BindAggregation("items", com.NewBindingDescription("", "/Items"), createListItem)
	return list
}

func createListItem(b com.Binding) base.ManagedObject {
	newInput := m.MD_Input.GetClass()
	newItem := m.MD_CustomListItem.GetClass()
	item := newItem("it").(m.CustomerListItem)
	item.BindObject(com.NewBindingDescription("", b.GetContext().GetPath()))
	aggr := item.Aggregation("content")

	input1 := newInput("ip").(m.Input)
	aggr.Add(input1)
	input1.BindProperty("value", com.NewBindingDescription("", "Field1"))

	input2 := newInput("ip").(m.Input)
	aggr.Add(input2)
	input2.BindProperty("value", com.NewBindingDescription("", "Field2"))

	input3 := newInput("ip").(m.Input)
	aggr.Add(input3)
	input3.BindProperty("value", com.NewBindingDescription("", "Field3"))

	return item
}
