package model

type eventData struct {
	v interface{}
}

func (e eventData) Param() interface{} {
	return e.v
}
