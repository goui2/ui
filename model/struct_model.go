package model

import (
	"reflect"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
)

type StructModel interface {
	Model
}

type structModel struct {
	Model
}

func unwrapData(ed interface{}) interface{} {
	switch d := ed.(type) {
	case com.NewValue:
		return d.NewValue()
	default:
		panic("should not be reached")
		//return ed
	}
}

func NewStructModel(v interface{}) StructModel {
	sm := MD_StructModel.GetClass().New("m#").(StructModel)
	sm.SetObject(v)
	return sm
}

func constructorStructModel(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &structModel{}
		mo.Model = parent(id, s...).(Model)
		return mo
	}
}

func (sm *structModel) GetSource() base.EventProvider { return sm }

func (sm *structModel) SetObject(v interface{}) {
	if v == nil {
		sm.Model.SetObject(v)
	} else {
		v1 := prepStructModel(v)
		sm.Model.SetObject(v1)
	}
}

func prepStructModel(v interface{}) (va reflect.Value) {
	var ok bool
	if va, ok = v.(reflect.Value); !ok {
		va = reflect.ValueOf(v)
	}
	if va.Kind() != reflect.Ptr {
		va = va.Addr()
	}
	if va.Kind() != reflect.Ptr {
		panic("v isn't reflect.Ptr")
	}
	va = va.Elem()
	if va.Kind() != reflect.Struct {
		panic("v isn't a struct{}")
	}
	return
}
