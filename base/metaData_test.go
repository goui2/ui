package base

import (
	"strconv"
	"testing"
	"time"

	"github.com/goui2/ui/com"
)

func TestMetadataMDObject(t *testing.T) {
	if MD_Object.GetClass() == nil {
		t.Error("invalid MD_Object")
	}
	obj := MD_Object.GetClass()("object1")
	if obj.Metadata() == nil {
		t.Error("object miss metadata")
	}
	if obj.Metadata().GetName() != "goui.base.Object" {
		t.Error("object has wrong name ")
	}
}

type eventHandlerTest1 struct {
	received int
}

func (e *eventHandlerTest1) Handle(com.Event) {
	e.received++
}

func TestMetadataMDEventProvider(t *testing.T) {
	var ehAllProviders eventHandlerTest1
	var ehCallMeProvider eventHandlerTest1
	epCallMeMD := Extend(MD_EventProvider, "ep1", MDSettingEventDef{EventDef{Name: "callMe"}})
	epCallMe1 := epCallMeMD.GetClass().New("i1").(EventProvider)
	epCallMe1.AttachEvent("callMe", nil, &ehAllProviders)
	epCallMe1.AttachEvent("callMe", nil, &ehCallMeProvider)
	epCallMe1.FireEvent("callMe", nil)

	epCallMeMDr, _ := LoadMetaData("ep1")
	epCallMeR1 := epCallMeMDr.GetClass().New("i2").(EventProvider)
	epCallMeR1.AttachEvent("callMe", nil, &ehAllProviders)
	epCallMeR1.AttachEvent("callMe", nil, &ehCallMeProvider)
	epCallMeR1.FireEvent("callMe", nil)

	time.Sleep(40 * time.Millisecond)
	if ehAllProviders.received != 2 {
		t.Error("eventProvider test failed: exp 3 act: " + strconv.Itoa(ehAllProviders.received))
	}
	if ehCallMeProvider.received != 2 {
		t.Error("eventProvider test failed: exp 2 act: " + strconv.Itoa(ehCallMeProvider.received))
	}
}

func TestMetadataMDEventProvider02(t *testing.T) {

	var ehAllProviders eventHandlerTest1
	var ehCallMeProvider eventHandlerTest1
	epCallMeMD := Extend(MD_EventProvider, "ep1", MDSettingEventDef{EventDef{Name: "callMe"}})
	epCallMeToMD := Extend(epCallMeMD, "ep12", MDSettingEventDef{EventDef{Name: "callMeTo"}})
	epCallMe := epCallMeToMD.GetClass().New("callMeTo").(EventProvider)
	epCallMe.AttachEvent("callMe", nil, &ehAllProviders)
	epCallMe.AttachEvent("callMeTo", nil, &ehCallMeProvider)

}

func TestManagedObjectMetaData(t *testing.T) {
	mdMO := MD_ManagedObject
	mdMOProp := Extend(mdMO, "mdMOProp", MDSettingPropertyDef{PropertyDef{"prop1", PropTypeString, false, "value1"}})
	prop1 := mdMOProp.GetClass().New("prop1").(ManagedObject)
	val := prop1.Property("prop1").Get().(string)
	if val != "value1" {
		t.Error("managedObject test failed exp: value1 act: " + val)
	}

	mdMOAggr := Extend(mdMO, "mdMOAggr", MDSettingAggregationDef{AggregationDef{"aggr1", TypeCheckString, "1..n"}})
	aggr1Obj := mdMOAggr.GetClass().New("aggr1").(ManagedObject)
	aggr := aggr1Obj.Aggregation("aggr1")
	if aggr == nil {
		t.Error("managedObject test failed aggregation is nil")
	}

	mdMOL2 := Extend(mdMO, "mdMOL2",
		MDSettingPropertyDef{PropertyDef{"prop1", PropTypeString, false, "value1"}},
		MDSettingAggregationDef{AggregationDef{"aggr1", TypeCheckString, "1..n"}},
	).(ManagedObjectMetaData)
	mdMOL3 := Extend(mdMOL2, "mdMOL3",
		MDSettingPropertyDef{PropertyDef{"prop2", PropTypeString, false, "value1"}},
		MDSettingAggregationDef{AggregationDef{"aggr2", TypeCheckString, "1..n"}},
	).(ManagedObjectMetaData)
	if _, ok := mdMOL3.Property("prop1"); !ok {
		t.Error("missing prop1")
	}
	if _, ok := mdMOL3.Property("prop2"); !ok {
		t.Error("missing prop2")
	}
	mdMOL3.Aggregation("aggr1")
	mdMOL3.Aggregation("aggr2")
}

type Dummy1 interface {
	EventProvider
	DoDummy1()
}
type Dummy2 interface {
	Dummy1
	DoDummy2()
}

type Dummy1MD interface {
	EventProviderMetadata
}

type Dummy2MD interface {
	Dummy1MD
}

type dummy1 struct {
	EventProvider
}

type dummy1MD struct {
	EventProviderMetadata
}

type dummy2MD struct {
	Dummy1MD
}

func (d *dummy1) DoDummy1() {

}

type dummy2 struct {
	Dummy1
}

func (d *dummy2) DoDummy2() {

}

func dummy1Constructor(m MetaData, emb Constructor) Constructor {
	return func(n string, settings ...InstanceSetting) Object {
		return &dummy1{
			EventProvider: emb(n, settings...).(EventProvider),
		}
	}
}

func dummy2Constructor(m MetaData, emb Constructor) Constructor {
	return func(n string, settings ...InstanceSetting) Object {
		return &dummy2{
			Dummy1: emb(n, settings...).(Dummy1),
		}
	}
}

func dummy1MetadataBuilder(s string, ms ...MetaDataSetting) []MetaDataSetting {
	newEPM := &dummy1MD{}
	parentSettings := append(
		ms,
		MDSettingCallback{2,
			func(m MetaData, pre MetaData) MetaData {
				newEPM.EventProviderMetadata = m.(EventProviderMetadata)
				return newEPM
			},
		},
	)
	return parentSettings
}
func dummy2MetadataBuilder(s string, ms ...MetaDataSetting) []MetaDataSetting {
	newEPM := &dummy2MD{}
	parentSettings := append(
		ms,
		MDSettingCallback{3,
			func(m MetaData, pre MetaData) MetaData {
				newEPM.Dummy1MD = m.(Dummy1MD)
				return newEPM
			},
		},
	)
	return parentSettings
}

func TestMDWithoutMetaDataType(t *testing.T) {
	var (
		MD_Dummy1 = Extend(MD_EventProvider, "dummy1",
			MDSettingEventDef{EventDef{Name: "dummy1"}},
			//			MetaModelBuilder(dummy1MetadataBuilder),
			MDSettingConstructor(dummy1Constructor),
		).(EventProviderMetadata)
		MD_Dummy2 = Extend(MD_Dummy1, "dummy2",
			MDSettingEventDef{EventDef{Name: "dummy2"}},
			//			MetaModelBuilder(dummy2MetadataBuilder),
			MDSettingConstructor(dummy2Constructor),
		).(EventProviderMetadata)
	)
	if _, ok := MD_Dummy2.Event("dummy1"); !ok {
		t.Error("missing event dummy1")
	}
	if _, ok := MD_Dummy2.Event("dummy2"); !ok {
		t.Error("missing event dummy2")
	}
	MD_Dummy1.GetClass().New("dummy1_01").(Dummy1).HasListener("dummy1")
	MD_Dummy2.GetClass().New("dummy2_01").(Dummy2).HasListener("dummy2")
}

func TestManagedObjectMetaData02(t *testing.T) {
	mdMO := MD_ManagedObject
	mdMO01 := Extend(mdMO, "mdMO01",
		MDSettingPropertyDef{PropertyDef{"prop1", PropTypeString, false, "value1"}},
		MDSettingAggregationDef{AggregationDef{"aggr1", TypeCheckString, "1..n"}},
	).(ManagedObjectMetaData)
	mdMO02 := Extend(mdMO01, "mdMO02",
		MDSettingPropertyDef{PropertyDef{"prop2", PropTypeString, false, "value1"}},
		MDSettingAggregationDef{AggregationDef{"aggr2", TypeCheckString, "1..n"}},
	).(ManagedObjectMetaData)
	if _, ok := mdMO02.Property("prop1"); !ok {
		t.Error("missing property prop1")
	}
	if _, ok := mdMO02.Property("prop2"); !ok {
		t.Error("missing property prop2")
	}
	if _, ok := mdMO02.Property("propUnknown"); ok {
		t.Error("property to much propUnknown")
	}
	if _, ok := mdMO01.Property("prop1"); !ok {
		t.Error("missing property prop1")
	}
	if _, ok := mdMO01.Property("prop2"); ok {
		t.Error("property to much prop2")
	}
}
