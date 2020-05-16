package brick

import "encoding/xml"

type ViewBrickType uint

const (
	ViewUndefined ViewBrickType = iota + 1
	ViewElement
	ViewAggregation
	ViewProperty
	ViewRoot
)

type ViewBrick struct {
	Name     xml.Name
	Attr     []ViewAttr
	Type     ViewBrickType
	Children []*ViewBrick
}

func (v ViewBrick) ClassName() string {
	return v.Name.Space + "." + v.Name.Local
}

type ViewAttr struct {
	Name  xml.Name
	Value string
	Type  ViewBrickType
}

func (vb *ViewBrick) Add(c *ViewBrick) {
	if c == nil {
		return
	}
	if vb.Children == nil {
		vb.Children = make([]*ViewBrick, 0)
	}
	vb.Children = append(vb.Children, c)
}
