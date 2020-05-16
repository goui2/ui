package com

import (
	"fmt"
	"regexp"
	"strings"
)

type BindingType string

const (
	BindingType_Property BindingType = "Property"
	BindingType_Context              = "Context"
	BindingType_List                 = "List"
)

type BindingDescription struct {
	active bool
	Model  string
	Path   string
	Type   BindingType
}

var emptyBindingDescription BindingDescription

func trim(s string) string { return strings.Trim(s, " \t\n\r") }

func NewBindingDescription(m, p string) BindingDescription {
	return BindingDescription{
		active: true,
		Model:  trim(m),
		Path:   trim(p),
	}
}

func ParseBindingString(s string) (BindingDescription, bool) {
	re := regexp.MustCompile(`\{(([^>]*)>)?([^>]+)\}`)
	m := re.FindAllStringSubmatch(s, -1)
	if len(m) > 0 && s == m[0][0] {
		return NewBindingDescription(m[0][2], m[0][3]), true
	} else {
		return emptyBindingDescription, false
	}
}

func (bd BindingDescription) IsNil() bool {
	return !bd.active
}

func (bd BindingDescription) IsAbsolute() bool {
	if bd.Path != "" && bd.Path[0] == '/' {
		return true
	} else {
		return false
	}
}

func (bd BindingDescription) String() string {
	if bd.active {
		return fmt.Sprintf("BD:%s(%s)", bd.Model, bd.Path)
	} else {
		return "BD<nil>"
	}
}

func (bd BindingDescription) NotEquals(d BindingDescription) bool {
	if bd.active != d.active {
		return true
	} else if bd.Model != d.Model {
		return true
	} else if bd.Path != d.Path {
		return true
	} else {
		return false
	}
}
