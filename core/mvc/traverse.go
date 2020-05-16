package mvc

import (
	"log"
	"strings"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/core"
	"github.com/goui2/ui/core/mvc/brick"
)

type ViewCreator func() View

type viewModification func(anc core.Control, vf ViewFactory)
type viewSetup []viewModification

type eventBinding struct {
	provider    base.EventProvider
	name        string
	handlerName string
}

type viewFactory struct {
	view     View
	bindings []func()
	events   []eventBinding
}

func (vf *viewFactory) ViewId(id string) string {
	return vf.view.Id() + "_" + id
}
func (vf *viewFactory) addBinding(f func()) {
	vf.bindings = append(vf.bindings, f)
}

func (vf *viewFactory) addEvent(event string, handler string, ep base.EventProvider) {
	vf.events = append(vf.events, eventBinding{ep, event, handler})
}

type childFactory func(vf ViewFactory) core.Control

func traverseChildBrick(brk *brick.ViewBrick) childFactory {
	bmd, ok := base.LoadMetaData(brk.ClassName())
	if !ok {
		panic("unknown class: " + brk.ClassName())
	}
	md := bmd.(core.ControlMetaData)
	id := ""
	for _, attr := range brk.Attr {
		if attr.Name.Local == "id" {
			id = attr.Value
			break
		}
	}
	setups := traverseBrick(brk, md)
	return func(vf ViewFactory) core.Control {
		ctrl := core.NewControlMD(md, vf.ViewId(id))
		for _, setup := range setups {
			setup(ctrl, vf)
		}
		return ctrl
	}
}

func eventName(n string) string {
	p := strings.ToLower(string(n[2:3]))
	return p + string(n[3:])
}

func traverseBrick(brk *brick.ViewBrick, md core.ControlMetaData) viewSetup {
	inits := make(viewSetup, 0)
	aggrBinding := make(map[string]com.BindingDescription)

	for _, attr := range brk.Attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if _, ok := md.Property(attrName); ok {
			bd, valid := com.ParseBindingString(attrValue)
			inits = append(inits, func(anc core.Control, vf ViewFactory) {
				if valid {
					anc.BindProperty(attrName, bd)
				} else {
					anc.Property(attrName).Set(attrValue)
				}
			})
		} else if _, ok := md.Aggregation(attrName); ok {
			if bd, ok := com.ParseBindingString(attrValue); ok {
				aggrBinding[attrName] = bd
			}
		} else if strings.HasPrefix(attrName, "on") {
			en := eventName(attrName)
			if en == "press" {
				log.Println("event", en, md.GetName())
				log.Println("events", md.AllEvents())
			}
			if _, ok := md.Event(en); ok {
				f := func(anc core.Control, vf ViewFactory) {
					vfs := vf.(*viewFactory)
					vfs.addEvent(en, attrValue, anc)
				}
				inits = append(inits, f)
			}
		}
	}

	for _, cld := range brk.Children {
		if _, ok := md.Aggregation(cld.Name.Local); ok {
			childSetup := make([]childFactory, 0)
			for _, aggrCld := range cld.Children {
				childSetup = append(childSetup, traverseChildBrick(aggrCld))
			}
			if bd, ok := aggrBinding[cld.Name.Local]; ok {
				inits = append(inits, func(anc core.Control, vf ViewFactory) {
					anc.BindAggregation(cld.Name.Local, bd, func(b com.Binding) base.ManagedObject {
						if len(childSetup) > 0 {
							ctrl := childSetup[0](vf)
							ctrl.BindObject(com.NewBindingDescription("", b.GetPath()))
							return ctrl
						} else {
							return nil
						}
					})
				})
			} else {
				inits = append(inits, func(anc core.Control, vf ViewFactory) {
					for _, setup := range childSetup {
						ctrl := setup(vf)
						anc.Aggregation(cld.Name.Local).Add(ctrl)
					}
				})
			}
		}
	}
	return inits
}

func attrByName(name string, brk *brick.ViewBrick) string {
	for _, attr := range brk.Attr {
		if attr.Name.Local == name {
			return attr.Value
		}
	}
	return ""
}

func traverseViewBrick(brk *brick.ViewBrick, md core.ControlMetaData) ViewCreator {
	viewId := "view"
	inits := traverseBrick(brk, md)
	controllerName := attrByName("controllerName", brk)

	return func() View {
		if cf, ok := LoadController(controllerName); ok {

			view := md.GetClass().New(viewId).(View)
			vf := &viewFactory{view, make([]func(), 0), make([]eventBinding, 0)}
			for _, i := range inits {
				i(view, vf)
			}
			for _, f := range vf.bindings {
				f()
			}
			ctrl := cf(view)
			for _, eb := range vf.events {
				if h, ok := ctrl.Handler(eb.handlerName); ok {
					eb.provider.AttachEvent(eb.name, nil, h)
				}
			}
			ctrl.Init()
			return view
		} else {
			panic("controller unknown: '" + controllerName + "'")
		}
	}
}
