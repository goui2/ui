package model

import (
	"testing"
	"time"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
)

func init() {
	go func() {
		time.Sleep(200 * time.Second)
	}()
}

type bHandle struct {
	t      *testing.T
	amount int
}

func (bh *bHandle) Handle(e com.Event) {
	//	bh.t.Log("Handle")
	bh.amount++
}

type tmpData struct {
	Feld1 string
	Feld2 string
}

var (
	PD_value      = base.MDSettingPropertyDef{base.PropertyDef{Name: "value", Type: base.PropTypeString, Required: false, DefaultValue: ""}}
	Event_Changed = base.EventSettingFactor(base.EventDef{Name: "changed"})
	MD_EPChange   = base.Extend(base.MD_EventProvider, "epchange",
		Event_Changed.AsSet("epchange"),
	)
	MD_MOPropValue = base.Extend(base.MD_ManagedObject, "propValue",
		PD_value,
	)
)

func TestStructModel01(t *testing.T) {
	epChange := MD_EPChange.GetClass()("id1").(base.EventProvider)
	h := &bHandle{t, 0}
	tmp := &tmpData{
		Feld1: "abc",
		Feld2: "def",
	}
	sm := MD_StructModel.GetClass()("sm1").(Model)
	b := sm.BindProperty("/Feld1")
	b.AttachChange(nil, h)
	epChange.AttachEvent("changed", nil, b)
	sm.SetObject(tmp)
	epChange.FireEvent("changed", base.PropertyChanged{nil, "changed"})
	time.Sleep(150 * time.Millisecond)
	if tmp.Feld1 != "changed" {
		t.Error("field1 not changed")
	}
	t.Logf("%#v", tmp)
	if h.amount != 2 {
		t.Errorf("invalid handler access exp 2 act %d", h.amount)
	}
}

func TestStructModel02(t *testing.T) {
	mo := MD_MOPropValue.GetClass()("mo1").(base.ManagedObject)
	tmp := &tmpData{
		Feld1: "abc",
		Feld2: "def",
	}
	sm := MD_StructModel.GetClass()("sm2").(Model)
	b := sm.BindProperty("/Feld1")
	mo.Property("value").Bind(b)
	sm.SetObject(tmp)
	time.Sleep(50 * time.Millisecond)
	if mo.Property("value").Get() != tmp.Feld1 {
		t.Error("property hasn't field value")
	}
	mo.Property("value").Set("anders")
	time.Sleep(50 * time.Millisecond)
	if "anders" != tmp.Feld1 {
		t.Error("field hasn't property value")
	}
}

func TestStructModelContext(t *testing.T) {
	tmp := &tmpData{
		Feld1: "abc",
		Feld2: "def",
	}
	sm := MD_StructModel.GetClass()("sm2").(Model)
	sm.SetObject(tmp)
	b := sm.BindProperty("/Feld1")
	v := b.GetContext().GetObject("")
	if v != "abc" {
		t.Error("invalid binding")
	}
	bc := sm.BindContext("/")
	v = bc.GetContext().GetObject("Feld1")
	if v != "abc" {
		t.Error("invalid binding")
	}
	tmp.Feld1 = "new-value"
	v = bc.GetContext().GetObject("Feld1")
	if v != "new-value" {
		t.Error("invalid binding")
	}

}

type tmpData03 struct {
	Feld1 string
	Feld2 string
}

type TmpManagedObject interface {
	base.ManagedObject
}

type tmpManagedObject struct {
	base.ManagedObject
}

func (t *tmpManagedObject) GetSource() base.EventProvider { return t }

func NewTmpManagedObject(s string) TmpManagedObject {
	pd := base.MDSettingPropertyDef{base.PropertyDef{Name: s, Type: base.PropTypeString, Required: false, DefaultValue: ""}}
	md := base.Extend(base.MD_ManagedObject, s, pd)
	mo1 := md.GetClass()("mo" + s).(base.ManagedObject)
	return &tmpManagedObject{mo1}
}

func TestStructModel03(t *testing.T) {
	mo1 := NewTmpManagedObject("value1")
	mo2 := NewTmpManagedObject("text")
	mo3 := NewTmpManagedObject("value2")
	tmp := &tmpData03{
		Feld1: "abc",
		Feld2: "def",
	}
	sm := MD_StructModel.GetClass()("sm03").(StructModel)
	b := sm.BindProperty("/Feld1")
	mo1.Property("value1").Set("")
	mo2.Property("text").Set("")
	mo3.Property("value2").Set("")
	mo1.Property("value1").Bind(b)
	mo2.Property("text").Bind(b)
	mo3.Property("value2").Bind(b)
	sm.SetObject(tmp)
	time.Sleep(150 * time.Millisecond)
	if mo1.Property("value1").Get() != tmp.Feld1 {
		t.Errorf("value1 %s <> %s", mo1.Property("value1").Get(), tmp.Feld1)
	}
	mo1.Property("value1").Set("anders")
	time.Sleep(150 * time.Millisecond)
	if "anders" != tmp.Feld1 {
		t.Errorf("Field1 %s <> %s", mo1.Property("value1").Get(), tmp.Feld1)
	}

	if "anders" != mo2.Property("text").Get() {
		t.Errorf("property text hasn't field value %s", mo2.Property("text").Get())
	}
	//	time.Sleep(200 * time.Second)
}

func TestStructModel04(t *testing.T) {
	mo1 := NewTmpManagedObject("value1")
	mo2 := NewTmpManagedObject("text")
	mo3 := NewTmpManagedObject("value2")
	tmp := &tmpData03{
		Feld1: "abc",
		Feld2: "def",
	}
	sm := MD_StructModel.GetClass()("sm03").(StructModel)
	b1 := sm.BindProperty("/Feld1")
	b2 := sm.BindProperty("/Feld1")
	b3 := sm.BindProperty("/Feld1")
	mo1.Property("value1").Set("")
	mo2.Property("text").Set("")
	mo3.Property("value2").Set("")
	mo1.Property("value1").Bind(b1)
	mo2.Property("text").Bind(b2)
	mo3.Property("value2").Bind(b3)
	sm.SetObject(tmp)
	time.Sleep(150 * time.Millisecond)
	if mo1.Property("value1").Get() != tmp.Feld1 {
		t.Errorf("value1 %s <> %s", mo1.Property("value1").Get(), tmp.Feld1)
	}
	mo1.Property("value1").Set("anders")
	time.Sleep(150 * time.Millisecond)
	if "anders" != tmp.Feld1 {
		t.Errorf("Field1 %s <> %s", mo1.Property("value1").Get(), tmp.Feld1)
	}

	if "anders" != mo2.Property("text").Get() {
		t.Errorf("property text hasn't field value %s", mo2.Property("text").Get())
	}
	//	time.Sleep(200 * time.Second)
}
