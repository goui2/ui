package base

type ManagedObjectMetaData interface {
	EventProviderMetadata
	AllProperties() []PropertyDef
	AllAggregations() []AggregationDef
	Property(s string) (PropertyDef, bool)
	Aggregation(s string) (AggregationDef, bool)
	DefaultAggregation() (AggregationDef, bool)
}

type managedObjectMetaData struct {
	EventProviderMetadata
	properties             []PropertyDef
	aggregations           []AggregationDef
	propertiesMap          map[string]PropertyDef
	aggregationsMap        map[string]AggregationDef
	defaultAggregationName string
}

func (m managedObjectMetaData) AllAggregations() []AggregationDef {
	l := make([]AggregationDef, 0)
	l = append(l, m.aggregations...)
	if p := m.GetParent(); p != nil {
		if mp, ok := p.(ManagedObjectMetaData); ok {
			l = append(l, mp.AllAggregations()...)
		}
	}
	return l
}
func (m managedObjectMetaData) AllProperties() []PropertyDef {
	a := make([]PropertyDef, 0)
	a = append(a, m.properties...)
	if p := m.GetParent(); p != nil {
		if mp, ok := p.(ManagedObjectMetaData); ok {
			a = append(a, mp.AllProperties()...)
		}
	}
	return a
}

func (m managedObjectMetaData) Property(s string) (p PropertyDef, ok bool) {
	p, ok = m.propertiesMap[s]
	if !ok {
		if pa := m.GetParent(); pa != nil {
			if mp, ok1 := pa.(ManagedObjectMetaData); ok1 {
				p, ok = mp.Property(s)
			} else {
				ok = false
			}
		}
	}
	return
}

func (m managedObjectMetaData) Aggregation(s string) (a AggregationDef, ok bool) {
	a, ok = m.aggregationsMap[s]
	if !ok {
		if p := m.GetParent(); p != nil {
			if mp, ok := p.(ManagedObjectMetaData); ok {
				return mp.Aggregation(s)
			}
		}
	}
	return
}

func (m managedObjectMetaData) DefaultAggregation() (a AggregationDef, ok bool) {
	a, ok = m.Aggregation(m.defaultAggregationName)
	return
}

func managedObjectMetaDataBuilder(s string, ms ...MetaDataSetting) []MetaDataSetting {
	epm := &managedObjectMetaData{
		properties:      make([]PropertyDef, 0),
		aggregations:    make([]AggregationDef, 0),
		propertiesMap:   make(map[string]PropertyDef),
		aggregationsMap: make(map[string]AggregationDef),
	}
	selector := func(i MetaDataSetting) bool {
		switch s := i.(type) {
		case MDSettingAggregationDef:
			epm.aggregations = append(epm.aggregations, s.AggregationDef)
			epm.aggregationsMap[s.AggregationDef.Name] = s.AggregationDef
			return true
		case MDSettingPropertyDef:
			epm.properties = append(epm.properties, s.PropertyDef)
			epm.propertiesMap[s.PropertyDef.Name] = s.PropertyDef
			return true
		}
		return true
	}
	parentSettings := SelectMetadataSettings(ms, selector)
	parentSettings = append(
		parentSettings,
		MDSettingCallback{
			2,
			func(m MetaData, prev MetaData) MetaData {
				epm.EventProviderMetadata = m.(EventProviderMetadata)
				if prevMD, ok := prev.(ManagedObjectMetaData); ok {
					for _, aDef := range prevMD.AllAggregations() {
						epm.aggregations = append(epm.aggregations, aDef)
						epm.aggregationsMap[aDef.Name] = aDef
					}
					for _, pDef := range prevMD.AllProperties() {
						epm.properties = append(epm.properties, pDef)
						epm.propertiesMap[pDef.Name] = pDef
					}
				}
				return epm
			},
		},
	)
	return parentSettings
}

/*
func newManagedObjectMetadata(m MetaData) MetaData {
	epm := &managedObjectMetaData{
		EventProviderMetadata: m.(EventProviderMetadata),
		properties:            make([]PropertyDef, 0),
		aggregations:          make([]AggregationDef, 0),
		propertiesMap:         make(map[string]PropertyDef),
		aggregationsMap:       make(map[string]AggregationDef),
	}
	return epm
}
*/
