# abi

> 原文：[https://pkg.go.dev/internal/abi@go1.24.2](https://pkg.go.dev/internal/abi@go1.24.2)








  


  

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/abi/abi_amd64.go;l=7)

``` go
const (

	// RAX, RBX, RCX, RDI, RSI, R8, R9, R10, R11.
	IntArgRegs = 9

	// X0 -> X14.
	FloatArgRegs = 15

	// We use SSE2 registers which support 64-bit float operations.
	EffectiveFloatRegSize = 8
)
```

## 变量

This section is empty.

## 函数

#### func FuncPCABI0 

``` go
func FuncPCABI0(f any) uintptr
```

FuncPCABI0 returns the entry PC of the function f, which must be a direct reference of a function defined as ABI0. Otherwise it is a compile-time error.

Implemented as a compile intrinsic.

#### func FuncPCABIInternal 

``` go
func FuncPCABIInternal(f any) uintptr
```

FuncPCABIInternal returns the entry PC of the function f. If f is a direct reference of a function, it must be defined as ABIInternal. Otherwise it is a compile-time error. If f is not a direct reference of a defined function, it assumes that f is a func value. Otherwise the behavior is undefined.

Implemented as a compile intrinsic.

## 类型

### type IntArgRegBitmap 

``` go
type IntArgRegBitmap [(IntArgRegs + 7) / 8]uint8
```

IntArgRegBitmap is a bitmap large enough to hold one bit per integer argument/return register.

#### (*IntArgRegBitmap) Get 

``` go
func (b *IntArgRegBitmap) Get(i int) bool
```

Get returns whether the i'th bit of the bitmap is set.

nosplit because it's called in extremely sensitive contexts, like on the reflectcall return path.

#### (*IntArgRegBitmap) Set 

``` go
func (b *IntArgRegBitmap) Set(i int)
```

Set sets the i'th bit of the bitmap to 1.

### type RegArgs 

``` go
type RegArgs struct {
	// Values in these slots should be precisely the bit-by-bit
	// representation of how they would appear in a register.
	//
	// This means that on big endian arches, integer values should
	// be in the top bits of the slot. Floats are usually just
	// directly represented, but some architectures treat narrow
	// width floating point values specially (e.g. they're promoted
	// first, or they need to be NaN-boxed).
	Ints   [IntArgRegs]uintptr  // untyped integer registers
	Floats [FloatArgRegs]uint64 // untyped float registers

	// Ptrs is a space that duplicates Ints but with pointer type,
	// used to make pointers passed or returned  in registers
	// visible to the GC by making the type unsafe.Pointer.
	Ptrs [IntArgRegs]unsafe.Pointer

	// ReturnIsPtr is a bitmap that indicates which registers
	// contain or will contain pointers on the return path from
	// a reflectcall. The i'th bit indicates whether the i'th
	// register contains or will contain a valid Go pointer.
	ReturnIsPtr IntArgRegBitmap
}
```

RegArgs is a struct that has space for each argument and return value register on the current architecture.

Assembly code knows the layout of the first two fields of RegArgs.

RegArgs also contains additional space to hold pointers when it may not be safe to keep them only in the integer register space otherwise.

#### (*RegArgs) Dump  <- go1.17.6

``` go
func (r *RegArgs) Dump()
```

#### (*RegArgs) IntRegArgAddr  <- go1.18

``` go
func (r *RegArgs) IntRegArgAddr(reg int, argSize uintptr) unsafe.Pointer
```

IntRegArgAddr returns a pointer inside of r.Ints[reg] that is appropriately offset for an argument of size argSize.

argSize must be non-zero, fit in a register, and a power-of-two.

This method is a helper for dealing with the endianness of different CPU architectures, since sub-word-sized arguments in big endian architectures need to be "aligned" to the upper edge of the register to be interpreted by the CPU correctly.