package base

type MDSetting struct {
	Name string
}

func (mds MDSetting) Class() string {
	return mds.Name
}

type MDSettingExtend struct {
	MDSetting
	New func(s string, m ...MetaDataSetting) (MetaData, func(MetaData) MetaData)
}

type MDSettingCallback struct {
	Level    int
	Callback Callback
}

func (s MDSettingCallback) Mtype() SettingType {
	return Combine
}

type MDSettingParent struct {
	MetaData MetaData
}

func (s MDSettingParent) Mtype() SettingType {
	return Overwrite
}

type MDSettingConstructor ConstructorBuilder

func (s MDSettingConstructor) Mtype() SettingType {
	return Overwrite
}

type MDSettingAbstract bool

func (s MDSettingAbstract) Mtype() SettingType {
	return Overwrite
}

type MDSettingFinal bool

func (s MDSettingFinal) Mtype() SettingType {
	return Overwrite
}

type MDSettingEventDef struct {
	EventDef
}

func (s MDSettingEventDef) Mtype() SettingType {
	return Required
}

type MDSettingPropertyDef struct {
	PropertyDef
}

func (s MDSettingPropertyDef) Mtype() SettingType {
	return Required
}

type MDSettingAggregationDef struct {
	AggregationDef
}

func (s MDSettingAggregationDef) Mtype() SettingType {
	return Required
}
