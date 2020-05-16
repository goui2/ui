package model

import (
	"regexp"
	"strconv"
)

type path struct {
	name   string
	offset int
}

type Path []path

func (p Path) Len() uint8           { return uint8(len(p)) }
func (p Path) isValid(i uint8) bool { return 0 <= i && i < p.Len() }
func (p Path) Name(i uint8) string {
	if p.isValid(i) {
		return p[int(i)].name
	}
	panic("invalid index " + strconv.Itoa(int(i)))
}
func (p Path) Offset(i uint8) int {
	if p.isValid(i) {
		return p[int(i)].offset
	}
	panic("invalid index " + strconv.Itoa(int(i)))
}

func (p Path) IsAbsolut() bool {
	if p.isValid(0) {
		return p[0].name == "/"
	} else {
		return false
	}
}

func (p Path) String() string {
	s := ""
	d := ""
	r := 0
	for i, x := range p {
		s += d + x.name
		if i == 0 && x.name == "/" {
			r = 1
		}
		if r == i {
			d = "/"
		}

	}
	return s
}

func (p Path) Append(a Path) Path {
	if a.IsAbsolut() {
		panic("absolut path can't be appended")
	}
	np := make(Path, 0)
	np = append(np, p...)
	np = append(np, a...)
	return np
}

func NewPath(p string) Path {
	re := regexp.MustCompile(`(/?)((\d+)|([\w_][\w_\d_]*))`)
	parts := re.FindAllStringSubmatch(p, -1)
	pth := make(Path, 0)
	if p == "/" {
		pth = append(pth, path{"/", -1})
	} else {
		for i, ps := range parts {
			if i == 0 && ps[1] == "/" {
				pth = append(pth, path{"/", -1})
			}
			if ps[4] != "" {
				pth = append(pth, path{ps[4], -1})
			} else if ps[3] != "" {
				x, _ := strconv.Atoi(ps[3])
				pth = append(pth, path{ps[3], x})
			}
		}
	}
	return pth
}
