package main

import (
	"log"
	"strconv"
	"time"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
	"github.com/goui2/ui/m"
	"github.com/goui2/ui/model"
)

type Data1 struct {
	Field1 string
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
	d1 := &Data1{
		Field1: "show me",
	}
	ep := MD_Tick.GetClass().New("tick").(base.EventProvider)
	c := make(chan struct{}, 0)
	m1 := model.NewStructModel(d1)
	b1 := m1.BindProperty("/Field1")

	p1 := createPanel.New("panel").(m.Panel)
	mc := p1.Aggregation("content")
	lb1 := createLabel.New("lbl1").(m.Label)
	lb1.Property("text").Set("Label 1")
	mc.Add(lb1)

	l1 := createInput.New("input1").(m.Input)
	l1.Property("name").Set("input1")
	l1.Property("value").Bind(b1)
	mc.Add(l1)
	log.Println("ID " + l1.Id())
	l1 = createInput.New("input2").(m.Input)
	l1.Property("name").Set("input2")
	l1.Property("value").Bind(b1)
	mc.Add(l1)
	l2 := createLabel.New("lbl2").(m.Label)
	l2.Property("text").Bind(b1)
	mc.Add(l2)
	p2 := createPanel.New("panel2").(m.Panel)
	mc.Add(p2)
	mc = p2.Aggregation("content")

	core.GetCore().StartApplication("appl", p1)
	d1 = &Data1{
		Field1: "show me2",
	}

	m1.SetObject(d1)
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
			mc.Set(data)
		}
	}

	/*
		data := make([]interface{}, offset)
		for i, k := range dynCtrls[0:offset] {
			data[i] = k
		}
		mc.Set(data)
	*/
	ep.AttachEvent("tick", nil, com.MakeHandler(tick))
	for {
		select {
		case <-ticker.C:
			ep.FireEvent("tick", nil)
		}
	}
	<-c
}
