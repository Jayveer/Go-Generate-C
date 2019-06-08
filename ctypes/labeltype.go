package ctypes

type LabelType struct {
	Name string
}

func (l *LabelType) SetName(name string) {
	l.Name = name
}

func (l *LabelType) ToC() string {
	return l.Name + ":"
}
