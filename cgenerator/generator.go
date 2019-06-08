package cgenerator

import (
	"fmt"

	"github.com/jayveerk/goGenerateC/ctypes"
)

const unnamed = "\b"

func PrintC(ct ctypes.CType) {
	fmt.Println(ct.ToC() + ";\n")
}

func CreateCommentType(comment string) *ctypes.CommentType {
	return &ctypes.CommentType{Comment: comment}
}

func CreateBaseType(typeName string) *ctypes.BaseType {
	return &ctypes.BaseType{Name: typeName}
}

func CreateEnumType(name string, values []string) *ctypes.EnumType {
	return &ctypes.EnumType{Name: name, Values: values}
}

func CreateMemberType(baseType ctypes.CType, varName string) *ctypes.MemberType {
	return &ctypes.MemberType{Type: baseType, Name: varName}
}

func CreateArrayType(baseType ctypes.CType, varName string, subranges []int) *ctypes.ArrayType {
	return &ctypes.ArrayType{Type: baseType, Name: varName, SubRanges: subranges}
}

func CreateFuncType(baseType *ctypes.BaseType, varName string, params *[]ctypes.MemberType, body *ctypes.CBlock) *ctypes.FuncType {
	return &ctypes.FuncType{Type: *baseType, Name: varName, Params: *params, Body: *body}
}

func CreateLabelType(name string) *ctypes.LabelType {
	return &ctypes.LabelType{Name: name}
}

func CreateTypeDef(bt ctypes.CType, varName string) *ctypes.TypeDefType {
	return &ctypes.TypeDefType{Type: bt, Name: varName}
}

func CreateStructType(name string, body *ctypes.CBlock) *ctypes.StructType {
	s := &ctypes.StructType{Name: name, Body: *body}
	s.SetStruct()
	return s
}

func CreateUnionType(name string, body *ctypes.CBlock) *ctypes.StructType {
	u := &ctypes.StructType{Name: name, Body: *body}
	u.SetUnion()
	return u
}

func CreateBlock(contents []ctypes.CType) *ctypes.CBlock {
	blk := &ctypes.CBlock{}
	for _, c := range contents {
		blk.AddContent(c)
	}
	return blk
}
