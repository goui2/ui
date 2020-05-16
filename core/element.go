package core

import (
	"errors"
	"strconv"
	"strings"
	"syscall/js"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/com"
	"github.com/goui2/ui/model"
)

type Element interface {
	base.ManagedObject
	Id() string
	Renderer() func(rm RenderManager, item Element)
	DomRef() (js.Value, bool)
	OnBeforeRendering()
	OnAfterRendering(v js.Value)
	FindDomChild(child Element) (js.Value, error)
	FindDomById(id string) (js.Value, error)
	Context() com.Context
	SetParent(p base.Object)
	Parent() Element
	ElementById(id string) Element
	BindObject(bd com.BindingDescription)
	BindAggregation(n string, b com.BindingDescription, of base.ObjectFactory)
	BindProperty(n string, b com.BindingDescription)
	String() string
}

func constructorElement(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &element{}

		renderer := md.(ElementMetaData).Renderer()

		parentSettings := base.SelectInstanceSettings(s, func(i base.InstanceSetting) bool {
			switch x := i.(type) {
			case ISRenderer:
				if mo.renderer == nil {
					mo.renderer = x.Renderer
				}
				return false
			}
			return true
		})
		parentSettings = base.AdjustSelfSetting(mo, md, parentSettings)

		if idx := strings.Index(id, "#"); idx > -1 {
			mo.id = id[0:idx]
		} else {
			mo.id = id + nextId()
		}

		mo.renderer = checkRenderer(mo.renderer, renderer)
		mo.ManagedObject = parent.New(mo.id, parentSettings...).(base.ManagedObject)

		return mo
	}
}

var defaultRenderer = func(rm RenderManager, item Element) {}

func checkRenderer(r func(rm RenderManager, item Element), mr func(rm RenderManager, item Element)) func(rm RenderManager, item Element) {
	if r == nil {
		if mr == nil {
			return defaultRenderer
		} else {
			return mr
		}
	} else {
		return r
	}
}

var ids int = 0

func nextId() string {
	ids++
	return strconv.Itoa(ids)
}

type element struct {
	base.ManagedObject
	id           string
	domRef       js.Value
	domRefExists bool
	renderer     func(rm RenderManager, item Element)
	binding      com.Binding
	//	bindingDescription com.BindingDescription
	bdSaved      com.BindingDescription
	bdCalculated com.BindingDescription
	parent       Element
	propBinding  map[string]*elementBindProperty
	aggrBinding  map[string]*elementBindAggregation
}

type elementBindProperty struct {
	name               string
	bindingDescription com.BindingDescription
	changed            bool
}

type elementBindAggregation struct {
	name               string
	bindingDescription com.BindingDescription
	objectFactory      base.ObjectFactory
	handler            com.EventHandler
	changed            bool
	oldBinding         *elementBindAggregation
	itemControls       []Control
}

func (e *element) Id() string { return e.id }
func (e *element) Renderer() func(rm RenderManager, item Element) {
	return e.renderer
}
func (e *element) DomRef() (js.Value, bool) { return e.domRef, e.domRefExists }
func (e *element) OnAfterRendering(d js.Value) {
	e.domRef = d
	e.domRefExists = true
}
func (e *element) OnBeforeRendering() {}
func (e *element) FindDomChild(child Element) (v js.Value, err error) {
	return e.FindDomById(child.Id())
}

func (e *element) FindDomById(id string) (v js.Value, err error) {
	if e.domRefExists {
		v = e.domRef.Call("querySelector", "#"+id)
		if testValue(v) {
			err = errors.New("no child with id: #" + id)
		}
	} else {
		err = errors.New("no domRef exists")
	}
	return
}

func testValue(v js.Value) bool { return true }

func (e *element) Context() com.Context {
	if e.binding != nil {
		return e.binding.GetContext()
	} else {
		return nil
	}
}

func (e *element) BindObject(b com.BindingDescription) {
	//log.Println("BindObject( ", e.String(), ") - ", b)
	e.bdSaved = b
	rebind := func() {
		if mdl, ok := GetCore().GetModel(e.bdCalculated.Model); ok {
			e.binding = mdl.BindContext(e.bdCalculated.Path)
		} else {
			e.binding = nil
		}
		e.updateForAggr(func(n string, dep Element) {
			dep.BindObject(e.bdCalculated)
		})
		e.updateObjectBinding()
	}

	if !b.IsNil() && b.IsAbsolute() {
		e.bdCalculated = b
		rebind()
	} else if e.parent != nil {
		tmpBD := BindingDescriptionOfContext(e.parent.Context())
		newBD := CalcBindingDescription(tmpBD, b)
		if !newBD.IsNil() && newBD.NotEquals(e.bdCalculated) {
			e.bdCalculated = newBD
			rebind()
		} else if newBD.IsNil() {
			e.binding = nil
			e.updateForAggr(func(n string, dep Element) {
				dep.BindObject(com.BindingDescription{})
			})
			e.removeObjectBinding()
		}
	}

}

func (e *element) BindProperty(n string, b com.BindingDescription) {
	//log.Println("BindProperty( ", e.String(), ") ", n, " - ", b)
	if e.propBinding == nil {
		e.propBinding = make(map[string]*elementBindProperty)
	}
	if pb, ok := e.propBinding[n]; ok {
		pb.bindingDescription = b
		pb.changed = true
		e.updatePropertyBinding(pb)
	} else {
		pb = &elementBindProperty{
			name:               n,
			bindingDescription: b,
			changed:            true,
		}
		e.propBinding[n] = pb
		e.updatePropertyBinding(pb)
	}
}

func (e *element) BindAggregation(n string, b com.BindingDescription, of base.ObjectFactory) {
	//log.Println("BindAggregation(", e.String(), ") :", n, " - ", b)
	if e.aggrBinding == nil {
		e.aggrBinding = make(map[string]*elementBindAggregation)
	}
	if ag, ok := e.aggrBinding[n]; ok {
		ag.changed = true
		ag.bindingDescription = b
		ag.objectFactory = of
		e.updateAggregationBinding(ag)
	} else {
		ag := &elementBindAggregation{
			name:               n,
			bindingDescription: b,
			objectFactory:      of,
			changed:            true,
		}
		e.aggrBinding[n] = ag
		e.updateAggregationBinding(ag)
	}
}

func (e *element) Parent() Element {
	return e.parent
}

func (e *element) SetParent(p base.Object) {
	if pe, ok := p.(Element); ok {
		//		//log.Printf("%s.SetParent(%s) Step1\n", e, pe)
		ctx := pe.Context()
		bd := BindingDescriptionOfContext(ctx)
		e.bdCalculated = CalcBindingDescription(bd, e.bdSaved)
		if !e.bdCalculated.IsNil() {
			if mdl, ok := GetCore().GetModel(e.bdCalculated.Model); ok {
				e.binding = mdl.BindContext(e.bdCalculated.Path)
			} else {
				e.binding = nil
			}
			e.updateForAggr(func(n string, dep Element) {
				dep.BindObject(e.bdCalculated)
			})
			e.updateObjectBinding()
		}
	}
}

func (e *element) updateForAggr(f func(n string, e Element)) {
	for _, ad := range e.Metadata().(base.ManagedObjectMetaData).AllAggregations() {
		for _, c := range e.Aggregation(ad.Name).Get() {
			if ctrl, ok := c.(Element); ok {
				f(ad.Name, ctrl)
			}
		}
	}
}

func (e *element) updateObjectBinding() {
	//log.Println(e, "updateObjectBinding")
	e.updatePropertyBindings()
	e.updateAggregationBindings()
}

func (e *element) removeObjectBinding() {

}

func (e *element) updatePropertyBindings() {
	for _, pb := range e.propBinding {
		e.updatePropertyBinding(pb)
	}
}

func (e *element) updatePropertyBinding(pb *elementBindProperty) {
	newBD := CalcBindingDescription(e.bdCalculated, pb.bindingDescription)
	//log.Println(e.String(), ".updatePB: ", e.bdCalculated.String(), " with ", pb.bindingDescription.String(), " => ", newBD.String())
	if !newBD.IsNil() && newBD.IsAbsolute() {
		if mdl, ok := GetCore().GetModel(newBD.Model); ok {
			b := mdl.BindProperty(newBD.Path)
			e.Property(pb.name).Bind(b)
		}
	}
}

func (e *element) updateAggregationBindings() {
	for _, ab := range e.aggrBinding {
		e.updateAggregationBinding(ab)
	}
}

func (e *element) updateAggregationBinding(ab *elementBindAggregation) {

	newBD := CalcBindingDescription(e.bdCalculated, ab.bindingDescription)
	//log.Println(e.String(), ".updateAB: ", e.bdCalculated.String(), " with ", ab.bindingDescription.String(), " => ", newBD.String())
	var b com.ListBinding
	if !newBD.IsNil() && newBD.IsAbsolute() {
		if mdl, ok := GetCore().GetModel(newBD.Model); ok {
			b = mdl.BindList(newBD.Path)
		}
	}

	aggr := e.Aggregation(ab.name)

	if b != nil {
		handlerFunc := func(e com.Event) {
			ctx := b.GetContext()
			length := ctx.GetValue("").ArraySize()
			mdl := ctx.GetModel().(model.Model)
			path := model.NewPath(ctx.GetPath())
			ab.itemControls = make([]Control, 0)
			startOff := 0
			if ab.oldBinding != nil {
				ab.itemControls = append(ab.itemControls, ab.oldBinding.itemControls...)
				startOff = len(ab.oldBinding.itemControls)
			}
			for i := startOff; i < length; i++ {
				itemPath := path.Append(model.NewPath(strconv.Itoa(i)))
				itemBind := mdl.BindContext(itemPath.String())
				itemCtrl := ab.objectFactory(itemBind).(Control)
				//				itemCtrl.BindObject(com.NewBindingDescription("", strconv.Itoa(i)))
				ab.itemControls = append(ab.itemControls, itemCtrl)
			}
			intA := make([]interface{}, length)
			for i := 0; i < length; i++ {
				intA[i] = ab.itemControls[i]
			}
			aggr.Set(intA)
		}
		ab.handler = com.MakeHandler(handlerFunc)
		b.AttachChange(nil, ab.handler)
	}
}

func (e *element) String() string {
	return e.id + "<" + e.Metadata().GetName() + ">"
}

func (e *element) ElementById(id string) Element {
	if e.Id() == id {
		return e
	} else {
		for _, aggrdef := range e.Metadata().(ElementMetaData).AllAggregations() {
			for _, cin := range e.Aggregation(aggrdef.Name).Get() {
				if child, ok := cin.(Element); ok {
					if child.Id() == id {
						return child
					} else {
						elem := child.ElementById(id)
						if elem != nil {
							return elem
						}
					}
				}
			}
		}
	}
	return nil
}
