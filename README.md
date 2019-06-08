
# Go Generate C


Go Generate C is a Go library which allows printable C code to be constructed via Go methods. It is still a work in progress.
 
## Installing

### Prerequisites

* [Git](https://git-scm.com)


Simply clone this repository into your Go dependencies folder and include the two packages within your project.

```go
import (
	"github.com/jayveerk/goGenerateC/cgenerator"
	"github.com/jayveerk/goGenerateC/ctypes"
)
```
## Example Usage

An Example.go file exists in the root of this project that demonstrates how the library can be used. The created variables remain consistent throughout the example.

```go
//creating a typedef
aBaseType := cgenerator.CreateBaseType("unsigned long")
aBaseType.AddPointer()
aTypeDef := cgenerator.CreateTypeDef(aBaseType, "uint_32t")
cgenerator.PrintC(aTypeDef)
```
```c
output: typedef unsigned long* uint_32t;
```
----

```go
//creating an array
aBaseType = cgenerator.CreateBaseType("char")
aBaseType.SetConstant(true)
anArray := cgenerator.CreateArrayType(aBaseType, "mark", []int{8, 256})
cgenerator.PrintC(anArray)
```
```c
output: const char mark[8][256]
```
----
```go
//creating a member
aMember := cgenerator.CreateMemberType(aBaseType, "test")
cgenerator.PrintC(aMember)
```
```c
output: const char test;
```
----
```go
//creating a struct
aBlock := cgenerator.CreateBlock([]ctypes.CType{aTypeDef, anArray, aMember})
aStruct := cgenerator.CreateStructType("_A_Header", aBlock)
cgenerator.PrintC(aStruct)
```
```c
output:
struct _A_Header {
	typedef unsigned long* uint_32t;
	const char mark[8][256];
	const char test;
}; 
```
----
```go
//struct typedef
anotherTypeDef := cgenerator.CreateTypeDef(aStruct, "A_Header")
cgenerator.PrintC(anotherTypeDef)
```
```c
output:
typedef struct _A_Header {
	typedef unsigned long* uint_32t;
	const char mark[8][256];
	const char test;
} A_Header;
```
----
```go
//creating an enum
anEnum := cgenerator.CreateEnumType("STATE", []string{"WAITING", "WORKING"})
cgenerator.PrintC(anEnum)
```
```c
output:
enum STATE {
	WAITING,
	WORKING
}; 
```
----
```go
//create function
anotherMember := cgenerator.CreateMemberType(aBaseType, "another")
params :=  &[]ctypes.MemberType{*aMember, *anotherMember}
aFunction := cgenerator.CreateFuncType(aBaseType, "FN_Open", params, aBlock)
aFunction.SetInline(true)
aFunction.SetExtern(true)
cgenerator.PrintC(aFunction)
```
```c
output:
extern inline const char FN_Open(const char test, const char another) {
	typedef unsigned long* uint_32t;
	const char mark[8][256];
	const char test;
}; 
```
## License
[MIT](LICENSE.md)