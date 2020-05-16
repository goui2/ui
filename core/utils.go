package core

import (
	"strings"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/model"
)

func NewControl(n, id string, s ...base.InstanceSetting) Control {
	if md, ok := base.LoadMetaData(n); ok {
		return NewControlMD(md.(ControlMetaData), id, s...)
	} else {
		panic("metadata not exists. class " + n)
	}
}

func NewControlMD(md ControlMetaData, id string, s ...base.InstanceSetting) Control {
	return md.GetClass().New(id, s...).(Control)
}

func BindingDescriptionOfBinding(b com.Binding) com.BindingDescription {
	if b == nil {
		return com.BindingDescription{}
	} else {
		return BindingDescriptionOfContext(b.GetContext())
	}
}

func BindingDescriptionOfContext(c com.Context) com.BindingDescription {
	if c == nil {
		return com.BindingDescription{}
	} else if mn, ok := GetCore().GetModelName(c.GetModel().(model.Model)); ok {
		return com.NewBindingDescription(mn, c.GetPath())
	} else {
		return com.BindingDescription{}
	}

}

func CalcBindingDescription(baseBD com.BindingDescription, itemBD com.BindingDescription) com.BindingDescription {
	if itemBD.IsNil() {
		return baseBD
	} else if itemBD.IsAbsolute() {
		return itemBD
	} else if baseBD.IsNil() {
		return baseBD
	} else if itemBD.Model != "" {
		panic("CalcBindingDescription item binding description is not absolute and contains model name: " + itemBD.String())
	} else {
		return com.NewBindingDescription(baseBD.Model, pathAppend(baseBD.Path, itemBD.Path))
	}
}

func pathAppend(p1, p2 string) string {
	join := p1 + "/" + p2
	path := ""
	for path = strings.ReplaceAll(join, "//", "/"); path != join; {
		join = path
		path = strings.ReplaceAll(join, "//", "/")
	}
	//	log.Println("Path: "+path, " parts: ", p1, p2)
	return path
}
