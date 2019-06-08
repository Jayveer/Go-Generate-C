package ctypes

type MemberType struct {
	Type CType
	Name string
}

func (m *MemberType) SetName(name string) {
	m.Name = name
}

func (m *MemberType) ToC() string {

	if mt, ok := m.Type.(*ArrayType); ok {
		mt.SetName(m.Name)
		return m.Type.ToC()
	}

	return m.Type.ToC() + " " + m.Name
}
