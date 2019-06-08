package ctypes

//CPtrType is a type that can be referenced by addr
type CPtrType interface {
	AddPointer()
	RemovePointer()
	SetName(name string)
	ToC() string
}

//CType is a type in C
type CType interface {
	SetName(name string)
	ToC() string
}

//CPrimitive is something that can be expressed in C but isn't a type
type CPrimitive interface {
	ToC() string
}
