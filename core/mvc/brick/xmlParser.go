package brick

import (
	"container/list"
	"encoding/xml"
	"io"
	"strings"
)

func ParseXmlView(r io.Reader) *ViewBrick {
	d := xml.NewDecoder(r)
	l := list.New()
	l.PushFront(&ViewBrick{})

	for t, err := d.Token(); err == nil; t, err = d.Token() {
		switch xt := t.(type) {
		case xml.StartElement:
			parent := l.Front().Value.(*ViewBrick)
			vb := &ViewBrick{
				Name: xt.Name,
				Attr: mapAttr(xt),
			}
			parent.Add(vb)
			l.PushFront(vb)
		case xml.EndElement:
			f := l.Front()
			l.Remove(f)
		}
	}
	t := l.Front().Value.(*ViewBrick)
	return t.Children[0]
}

func mapAttr(e xml.StartElement) []ViewAttr {
	l := make([]ViewAttr, 0)
	for _, xa := range e.Attr {
		if strings.HasPrefix(xa.Name.Space, "xmlns") || xa.Name.Local == "xmlns" {
			continue
		}
		va := ViewAttr{
			Name:  xa.Name,
			Value: xa.Value,
		}
		l = append(l, va)
	}
	return l
}
