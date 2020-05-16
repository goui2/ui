package m
import (
	"syscall/js"

)

type htmlElement struct {
	Html js.Value
}

func (h *htmlElement) FindElementById(id string) (*htmlElement,bool)  {
	element := h.Html.Call("querySelector", "#" + id)
	if element.IsNull() || element.IsUndefined() {
		return nil,false
	}else {
		return &htmlElement{element},true
	}
}