package ctypes

import "strings"

type StructType struct {
	sType    string
	Name     string
	nPointer int
	Body     CBlock
}

func (s *StructType) checkPointer() string {
	var p strings.Builder

	for i := 0; i < s.nPointer; i++ {
		p.WriteString("*")
	}

	return p.String()
}

func (s *StructType) AddPointer() {
	s.nPointer++
}

func (s *StructType) RemovePointer() {
	s.nPointer--
}

func (s *StructType) SetName(name string) {
	s.Name = name
}

func (s *StructType) SetStruct() {
	s.sType = "struct"
}

func (s *StructType) SetUnion() {
	s.sType = "union"
}

func (s *StructType) ToC() string {
	return s.sType + " " + s.checkPointer() + s.Name + " " + s.Body.ToC()
}
