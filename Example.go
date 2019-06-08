package main

import (
	"github.com/jayveerk/goGenerateC/cgenerator"
	"github.com/jayveerk/goGenerateC/ctypes"
)

func main() {
	//creating typedef
	aBaseType := cgenerator.CreateBaseType("unsigned long")
	aBaseType.AddPointer()
	aTypeDef := cgenerator.CreateTypeDef(aBaseType, "uint_32t")
	cgenerator.PrintC(aTypeDef)
	//output: typedef unsigned long* uint_32t;

	//creating array
	aBaseType = cgenerator.CreateBaseType("char")
	aBaseType.SetConstant(true)
	anArray := cgenerator.CreateArrayType(aBaseType, "mark", []int{8, 256})
	cgenerator.PrintC(anArray)
	//output: const char mark[8][256]

	//creating member
	aMember := cgenerator.CreateMemberType(aBaseType, "test")
	cgenerator.PrintC(aMember)
	//output: const char test;

	//creating struct
	aBlock := cgenerator.CreateBlock([]ctypes.CType{aTypeDef, anArray, aMember})
	aStruct := cgenerator.CreateStructType("_A_Header", aBlock)
	cgenerator.PrintC(aStruct)
	/* output:
	struct _A_Header {
		typedef unsigned long* uint_32t;
		const char mark[8][256];
		const char test;
	}; */

	//struct typedef
	anotherTypeDef := cgenerator.CreateTypeDef(aStruct, "A_Header")
	cgenerator.PrintC(anotherTypeDef)
	/* output:
	typedef struct _A_Header {
		typedef unsigned long* uint_32t;
		const char mark[8][256];
		const char test;
	} A_Header; */

	//make enum
	anEnum := cgenerator.CreateEnumType("STATE", []string{"WAITING", "WORKING"})
	cgenerator.PrintC(anEnum)
	/*output:
	enum STATE {
		WAITING,
		WORKING
	}; */

	//create function
	anotherMember := cgenerator.CreateMemberType(aBaseType, "another")
	params := &[]ctypes.MemberType{*aMember, *anotherMember}
	aFunction := cgenerator.CreateFuncType(aBaseType, "FN_Open", params, aBlock)
	aFunction.SetInline(true)
	aFunction.SetExtern(true)
	cgenerator.PrintC(aFunction)
	/* output:
	extern inline const char FN_Open(const char test, const char another) {
		typedef unsigned long* uint_32t;
		const char mark[8][256];
		const char test;
	}; */
}
