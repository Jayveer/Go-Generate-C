package ctypes

import (
	"strconv"
	"strings"
)

//ArrayType represents a Single/Multidimensional C Array
type ArrayType struct {
	Type      CType
	Name      string
	SubRanges []int
}

func (a *ArrayType) iterateContents() string {
	var sb strings.Builder
	for _, r := range a.SubRanges {
		sb.WriteString("[" + strconv.Itoa(r) + "]")
	}
	return sb.String()
}

func (a *ArrayType) SetName(name string) {
	a.Name = name
}

func (a *ArrayType) ToC() string {
	return (a.Type.ToC() + " " + a.Name + a.iterateContents())
}
