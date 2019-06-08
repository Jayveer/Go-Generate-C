package ctypes

import (
	"strings"
)

//CBlock is n amount of CTypes, eg. lexical block
type CBlock struct {
	Contents []CPrimitive
}

func (cb *CBlock) iterateContents() string {
	var sb strings.Builder
	for _, cp := range cb.Contents {
		sb.WriteString("\t" + cp.ToC() + ";\n")
	}
	return sb.String()
}

func (cb *CBlock) AddContent(ct CType) {
	cb.Contents = append(cb.Contents, ct)
}

func (cb *CBlock) ToC() string {
	return "{\n" + cb.iterateContents() + "}"
}
