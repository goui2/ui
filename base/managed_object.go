package base

type ManagedObject interface {
	EventProvider
	Property(n string) Property
	Aggregation(n string) Aggregation
	DefaultAggregation() string
	//	BindObject(com.Binding)
}

type managedObject struct {
	EventProvider
	props              map[string]*property
	aggrs              map[string]*aggregation
	defaultAggregation string
	className          string
}

func constructorManagedObject(md MetaData, emb Constructor) Constructor {
	return func(id string, s ...InstanceSetting) Object {
		moMD := md.(ManagedObjectMetaData)
		mo := &managedObject{
			props: make(map[string]*property),
			aggrs: make(map[string]*aggregation),
		}

		parentSettings := AdjustSelfSetting(mo, md, s)
		mo.EventProvider = emb(id, parentSettings...).(EventProvider)

		mo.initProperties(moMD.AllProperties())
		mo.initAggregations(moMD.AllAggregations())
		return mo
	}
}

func (mo *managedObject) initProperties(pds []PropertyDef) {
	for _, pd := range pds {
		p := property{
			EventProvider: MD_EventProviderProp.GetClass()("").(EventProvider), //NewEventProvider(EventProviderSettings{Events: []string{"changed", "updateHTML"}}),
			mO:            mo,
			name:          pd.Name,
			value:         pd.DefaultValue,
			typeCheck:     checkPropTypeCheck(pd.Type),
		}
		mo.props[pd.Name] = &p
	}
}

func (mo *managedObject) initAggregations(ags []AggregationDef) {
	for _, ad := range ags {
		a := aggregation{
			EventProvider: MD_EventProviderAggr.GetClass().New("").(EventProvider),
			mO:            mo,
			name:          ad.Name,
			typeCheck:     checkTypeCheck(ad.Type),
			values:        make([]interface{}, 0),
		}
		mo.aggrs[ad.Name] = &a
	}
}

var defaultTypeCheck = func(v interface{}) bool { return true }

func checkPropTypeCheck(f func(v interface{}) (interface{}, bool)) func(v interface{}) (interface{}, bool) {
	if f == nil {
		return PropTypeString
	} else {
		return f
	}
}

func checkTypeCheck(f func(v interface{}) bool) func(v interface{}) bool {
	if f == nil {
		return defaultTypeCheck
	} else {
		return f
	}
}

func (mo *managedObject) Property(n string) Property {
	if p, ok := mo.props[n]; ok {
		return p
	} else {
		panic("unknown property: " + n)
	}
}
func (mo *managedObject) Aggregation(n string) Aggregation {
	if a, ok := mo.aggrs[n]; ok {
		return a
	} else {
		panic("unknown aggregation: " + n)
	}
}

/*
func (mo *managedObject) BindObject(b com.Binding) {
	panic("BindObject must be implemented")
}
*/

func (mo *managedObject) Destroy() {
	for _, ag := range mo.aggrs {
		for _, v := range ag.values {
			if cld, ok := v.(ManagedObject); ok {
				cld.Destroy()
			}
		}
		ag.values = make([]interface{}, 0)
	}
}

func (mo *managedObject) DefaultAggregation() string {
	return mo.defaultAggregation
}
