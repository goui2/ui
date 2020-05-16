package base

type Object interface {
	Metadata() MetaData
	Self() Object
}

type object struct {
	metaData MetaData
	self     Object
}

func (o *object) Metadata() MetaData {
	return o.metaData
}
func (o *object) Self() Object {
	return o.self
}
