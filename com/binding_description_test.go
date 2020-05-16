package com

import "testing"

func TestBindingDescription01(t *testing.T) {
	cases := []struct {
		pattern, model, path string
		found                bool
	}{
		{"{m1>/a/b}", "m1", "/a/b", true},
		{"{>/a/b}", "", "/a/b", true},
		{"{/a/b}", "", "/a/b", true},
		{"/a/b", "", "", false},
		{" {m1>/a/b}", "", "", false},
		{"{m1>/a/b} ", "", "", false},
	}
	for _, item := range cases {
		bd, ok := ParseBindingString(item.pattern)
		if ok != item.found {
			t.Errorf("wrong parse result '%s'", item.pattern)
		}
		if bd.Model != item.model {
			t.Errorf("wrong model '%s' exp %s act %s", item.pattern, item.model, bd.Model)
		}
		if bd.Path != item.path {
			t.Errorf("wrong path '%s' exp %s act %s", item.pattern, item.path, bd.Path)
		}
	}
}
func TestBindingDescription02(t *testing.T) {
	bd := NewBindingDescription("", "/abc")
	if !bd.IsAbsolute() {
		t.Error("IsAbsolute false")
	}
	bd = NewBindingDescription("", "")
	if bd.IsAbsolute() {
		t.Error("IsAbsolute error")
	}
}
