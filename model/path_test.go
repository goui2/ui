package model

import (
	"log"
	"testing"
)

func TestPath01(t *testing.T) {
	pth := NewPath("as2_2/df/3/x")
	if pth.Len() != 4 {
		t.Error("invalid path parsing: " + pth.String())
	}
	log.Println(pth.String())
}
