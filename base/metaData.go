package base

import "reflect"

// The MetaData contains all information for creating a new object instance.
// The main point is a standarized constructor, which creates a instance of
// object or any derived sub object.
//
// MetaData can also be derived, so that it contains additional meta information
// for creating the corrosponding instance. Creating new MetaData can be done with
// Extend.
//
type MetaData interface {
	GetName() string
	// Give back the constructor for creating a new instance
	GetClass() Constructor
	GetParent() MetaData
	IsAbstract() bool
	IsFinal() bool
	// Checks whether the described class or one of its ancestor classes implements the given interface.
	IsInstanceOf(s string) bool
	//	Extend(s string, m ...MetaDataSetting) MetaData
	layer() *metaData
}

type Constructor func(id string, s ...InstanceSetting) Object

func (c Constructor) New(id string, s ...InstanceSetting) Object {
	return c(id, s...)
}

type MetaModelBuilder func(string, ...MetaDataSetting) []MetaDataSetting

func (_ MetaModelBuilder) Mtype() SettingType { return Unspecific }

type ConstructorBuilder func(m MetaData, embbed Constructor) Constructor

type Callback func(item MetaData, prev MetaData) MetaData

type metaData struct {
	name               string
	parent             MetaData
	abstract           bool
	final              bool
	constructor        Constructor
	constructorBuilder ConstructorBuilder

	previous *metaData
	builder  MetaModelBuilder
	config   []MetaDataSetting
	tail     MetaData
}

func (layer *metaData) prevLayer() *metaData {
	return layer.previous
}

func collectConstructorBuilder(m *metaData) []ConstructorBuilder {
	constBuilders := make([]ConstructorBuilder, 0)
	var lastBuilder uintptr

	for layer := m; layer != nil; layer = layer.prevLayer() {
		if layer.constructorBuilder != nil {
			act := reflect.ValueOf(layer.constructorBuilder).Pointer()
			if act != lastBuilder {
				constBuilders = append(constBuilders, layer.constructorBuilder)
				lastBuilder = act
			}
		}
	}

	return constBuilders
}

func Extend(m MetaData, n string, cfg ...MetaDataSetting) MetaData {
	newModel := &metaData{
		name:     n,
		previous: m.layer(),
	}
	var x *metaData
	for x = newModel.previous; x.builder == nil; x = x.previous {
	}
	newModel.parent = MetaData(x)
	newModel.builder, newModel.config = selectMetaModelBuilder(cfg)
	constBuilders := make([]ConstructorBuilder, 0)
	//	var lastConstBuilder uintptr
	params := []MetaDataSetting{}
	for layer := newModel; layer != nil; {
		params = append(layer.config, params...)
		if layer.builder != nil {
			params = layer.builder(n, params...)
			constBuilders = append(constBuilders, layer.constructorBuilder)
		}

		if layer.parent != nil {
			layer = layer.parent.layer()
		} else {
			layer = nil
		}
	}
	callbacks := make(map[int]Callback)
	maxLevel := 0
	for _, s := range params {
		switch ms := s.(type) {
		case MDSettingCallback:
			callbacks[ms.Level] = ms.Callback
			if ms.Level > maxLevel {
				maxLevel = ms.Level
			}
		}
	}
	var typeModel MetaData = newModel
	for idx := 0; idx <= maxLevel; idx++ {
		if cb, ok := callbacks[idx]; ok {
			typeModel = cb(typeModel, newModel.previous.tail)
		}
	}
	newModel.tail = typeModel

	constBuilders = collectConstructorBuilder(newModel)
	var prevBuilder Constructor = baseObjectConstructor
	for idx := len(constBuilders) - 1; idx > 0; idx-- {
		prevBuilder = constBuilders[idx](typeModel, prevBuilder)
	}
	newModel.constructor = newModel.constructorBuilder(typeModel, prevBuilder)
	registerMetaData(typeModel)
	return typeModel
}

func (m metaData) GetName() string {
	return m.name
}

func (m metaData) GetClass() Constructor {
	return m.constructor
}

func (m metaData) GetParent() MetaData {
	return m.parent
}

func (m metaData) IsAbstract() bool {
	return m.abstract
}

func (m metaData) IsFinal() bool {
	return m.final
}

func (m metaData) IsInstanceOf(s string) bool {
	return false
}

func (m *metaData) layer() *metaData {
	return m
}

func baseObjectConstructor(id string, s ...InstanceSetting) Object {
	obj := &object{metaData: MD_Object}
	for _, o := range s {
		if self, ok := o.(SelfSetting); ok {
			obj.self = self.Object
			obj.metaData = self.MetaData
			return obj
		}
	}
	panic("missing self")
}

func baseObjectBuilder(n string, settings ...MetaDataSetting) []MetaDataSetting {
	objectMD := &metaData{
		name: n,
	}
	selector := func(mds MetaDataSetting) bool {
		switch item := mds.(type) {
		case MDSettingAbstract:
			objectMD.abstract = bool(item)
			return false
		case MDSettingFinal:
			objectMD.final = bool(item)
			return false
		case MDSettingConstructor:
			objectMD.constructorBuilder = ConstructorBuilder(item)
		}
		return true
	}
	SelectMetadataSettings(settings, selector)

	return append(settings, MDSettingCallback{0, func(m MetaData, pre MetaData) MetaData {
		dtls := m.(*metaData)
		dtls.abstract = objectMD.abstract
		dtls.final = objectMD.final
		dtls.constructorBuilder = objectMD.constructorBuilder
		return m
	}})
}

func selectMetaModelBuilder(m []MetaDataSetting) (MetaModelBuilder, []MetaDataSetting) {
	result := make([]MetaDataSetting, 0, len(m))
	var metaBuilder MetaModelBuilder
	for _, s := range m {
		switch ms := s.(type) {
		case MetaModelBuilder:
			metaBuilder = ms
		default:
			result = append(result, s)
		}
	}
	return metaBuilder, result
}
