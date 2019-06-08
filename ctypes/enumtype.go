package ctypes

import (
	"strings"
)

type EnumType struct {
	Name   string
	Values []string
}

func (e *EnumType) iterateContents() string {
	var sb strings.Builder
	for i, ev := range e.Values {
		sep := ","
		if i+1 == len(e.Values) {
			sep = ""
		}
		sb.WriteString("\t" + ev + sep + "\n")
	}
	return sb.String()
}

func (e *EnumType) SetName(name string) {
	e.Name = name
}

func (e *EnumType) ToC() string {
	return "enum" + " " + e.Name + " " + "{\n" + e.iterateContents() + "}"
}
