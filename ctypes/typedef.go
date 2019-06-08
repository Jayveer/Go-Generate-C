package ctypes

import "strings"

type TypeDefType struct {
	Type     CType
	Name     string
	nPointer int
}

func (t *TypeDefType) checkPointer() string {
	var p strings.Builder

	for i := 0; i < t.nPointer; i++ {
		p.WriteString("*")
	}

	return p.String()
}

func (t *TypeDefType) AddPointer() {
	t.nPointer++
}

func (t *TypeDefType) RemovePointer() {
	t.nPointer--
}

func (t *TypeDefType) SetName(name string) {
	t.Name = name
}

func (t *TypeDefType) SetType(ct CType) {
	t.Type = ct
}

func (t *TypeDefType) ToC() string {
	return ("typedef " + t.Type.ToC() + " " + t.Name + t.checkPointer())
}
