package ctypes

import (
	"strings"
)

type BaseType struct {
	volatile bool
	constant bool
	nPointer int
	Name     string
}

func (b *BaseType) checkQualifier() string {
	var prefix strings.Builder

	if b.constant {
		prefix.WriteString("const ")
	}

	if b.volatile {
		prefix.WriteString("volatile ")
	}

	return prefix.String()
}

func (b *BaseType) checkPointer() string {
	var p strings.Builder

	for i := 0; i < b.nPointer; i++ {
		p.WriteString("*")
	}

	return p.String()
}

func (b *BaseType) AddPointer() {
	b.nPointer++
}

func (b *BaseType) RemovePointer() {
	b.nPointer--
}

func (b *BaseType) SetName(name string) {
	b.Name = name
}

func (b *BaseType) SetConstant(state bool) {
	b.constant = state
}

func (b *BaseType) SetVolatile(state bool) {
	b.volatile = state
}

func (b *BaseType) ToC() string {
	return b.checkQualifier() + b.Name + b.checkPointer()
}
