package ctypes

import (
	"strings"
)

type FuncType struct {
	inline bool
	extern bool
	Type   BaseType //return type
	Name   string
	Params []MemberType
	Body   CBlock
}

func (f *FuncType) iterateContents() string {
	var sb strings.Builder
	for i, p := range f.Params {
		sep := ", "
		if i+1 == len(f.Params) {
			sep = ""
		}
		sb.WriteString(p.ToC() + sep)
	}
	return sb.String()
}

func (f *FuncType) checkQualifier() string {
	var qualifier strings.Builder

	if f.extern {
		qualifier.WriteString("extern ")
	}

	if f.inline {
		qualifier.WriteString("inline ")
	}

	return qualifier.String()
}

func (f *FuncType) SetName(name string) {
	f.Name = name
}

func (f *FuncType) SetType(t BaseType) {
	f.Type = t
}

func (f *FuncType) SetInline(state bool) {
	f.inline = state
}

func (f *FuncType) SetExtern(state bool) {
	f.extern = state
}

func (f *FuncType) ToC() string {
	return (f.checkQualifier() + f.Type.ToC() + " " + f.Name + "(" + f.iterateContents() + ")" + " " + f.Body.ToC())
}
