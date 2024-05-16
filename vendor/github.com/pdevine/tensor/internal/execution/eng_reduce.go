// Code generated by genlib2. DO NOT EDIT.

package execution

import (
	"reflect"
	"unsafe"

	"github.com/pkg/errors"
	"github.com/pdevine/tensor/internal/storage"
)

func (e E) ReduceFirst(t reflect.Type, data *storage.Header, retVal *storage.Header, split int, size int, fn interface{}) (err error) {
	switch t {
	case Bool:
		dt := data.Bools()
		rt := retVal.Bools()
		switch f := fn.(type) {
		case func([]bool, []bool):
			reduceFirstB(dt, rt, split, size, f)
		case func(bool, bool) bool:
			genericReduceFirstB(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Int:
		dt := data.Ints()
		rt := retVal.Ints()
		switch f := fn.(type) {
		case func([]int, []int):
			reduceFirstI(dt, rt, split, size, f)
		case func(int, int) int:
			genericReduceFirstI(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Int8:
		dt := data.Int8s()
		rt := retVal.Int8s()
		switch f := fn.(type) {
		case func([]int8, []int8):
			reduceFirstI8(dt, rt, split, size, f)
		case func(int8, int8) int8:
			genericReduceFirstI8(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Int16:
		dt := data.Int16s()
		rt := retVal.Int16s()
		switch f := fn.(type) {
		case func([]int16, []int16):
			reduceFirstI16(dt, rt, split, size, f)
		case func(int16, int16) int16:
			genericReduceFirstI16(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Int32:
		dt := data.Int32s()
		rt := retVal.Int32s()
		switch f := fn.(type) {
		case func([]int32, []int32):
			reduceFirstI32(dt, rt, split, size, f)
		case func(int32, int32) int32:
			genericReduceFirstI32(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Int64:
		dt := data.Int64s()
		rt := retVal.Int64s()
		switch f := fn.(type) {
		case func([]int64, []int64):
			reduceFirstI64(dt, rt, split, size, f)
		case func(int64, int64) int64:
			genericReduceFirstI64(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Uint:
		dt := data.Uints()
		rt := retVal.Uints()
		switch f := fn.(type) {
		case func([]uint, []uint):
			reduceFirstU(dt, rt, split, size, f)
		case func(uint, uint) uint:
			genericReduceFirstU(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Uint8:
		dt := data.Uint8s()
		rt := retVal.Uint8s()
		switch f := fn.(type) {
		case func([]uint8, []uint8):
			reduceFirstU8(dt, rt, split, size, f)
		case func(uint8, uint8) uint8:
			genericReduceFirstU8(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Uint16:
		dt := data.Uint16s()
		rt := retVal.Uint16s()
		switch f := fn.(type) {
		case func([]uint16, []uint16):
			reduceFirstU16(dt, rt, split, size, f)
		case func(uint16, uint16) uint16:
			genericReduceFirstU16(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Uint32:
		dt := data.Uint32s()
		rt := retVal.Uint32s()
		switch f := fn.(type) {
		case func([]uint32, []uint32):
			reduceFirstU32(dt, rt, split, size, f)
		case func(uint32, uint32) uint32:
			genericReduceFirstU32(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Uint64:
		dt := data.Uint64s()
		rt := retVal.Uint64s()
		switch f := fn.(type) {
		case func([]uint64, []uint64):
			reduceFirstU64(dt, rt, split, size, f)
		case func(uint64, uint64) uint64:
			genericReduceFirstU64(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Uintptr:
		dt := data.Uintptrs()
		rt := retVal.Uintptrs()
		switch f := fn.(type) {
		case func([]uintptr, []uintptr):
			reduceFirstUintptr(dt, rt, split, size, f)
		case func(uintptr, uintptr) uintptr:
			genericReduceFirstUintptr(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Float32:
		dt := data.Float32s()
		rt := retVal.Float32s()
		switch f := fn.(type) {
		case func([]float32, []float32):
			reduceFirstF32(dt, rt, split, size, f)
		case func(float32, float32) float32:
			genericReduceFirstF32(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Float64:
		dt := data.Float64s()
		rt := retVal.Float64s()
		switch f := fn.(type) {
		case func([]float64, []float64):
			reduceFirstF64(dt, rt, split, size, f)
		case func(float64, float64) float64:
			genericReduceFirstF64(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Complex64:
		dt := data.Complex64s()
		rt := retVal.Complex64s()
		switch f := fn.(type) {
		case func([]complex64, []complex64):
			reduceFirstC64(dt, rt, split, size, f)
		case func(complex64, complex64) complex64:
			genericReduceFirstC64(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Complex128:
		dt := data.Complex128s()
		rt := retVal.Complex128s()
		switch f := fn.(type) {
		case func([]complex128, []complex128):
			reduceFirstC128(dt, rt, split, size, f)
		case func(complex128, complex128) complex128:
			genericReduceFirstC128(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case String:
		dt := data.Strings()
		rt := retVal.Strings()
		switch f := fn.(type) {
		case func([]string, []string):
			reduceFirstStr(dt, rt, split, size, f)
		case func(string, string) string:
			genericReduceFirstStr(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case UnsafePointer:
		dt := data.UnsafePointers()
		rt := retVal.UnsafePointers()
		switch f := fn.(type) {
		case func([]unsafe.Pointer, []unsafe.Pointer):
			reduceFirstUnsafePointer(dt, rt, split, size, f)
		case func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer:
			genericReduceFirstUnsafePointer(dt, rt, split, size, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	default:
		return errors.Errorf("Unsupported type %v for ReduceFirst", t)
	}
}

func (e E) ReduceLast(t reflect.Type, data *storage.Header, retVal *storage.Header, dimSize int, defaultValue interface{}, fn interface{}) (err error) {
	var ok bool
	switch t {
	case Bool:
		var def bool

		if def, ok = defaultValue.(bool); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Bools()
		rt := retVal.Bools()
		switch f := fn.(type) {
		case func([]bool) bool:
			reduceLastB(dt, rt, dimSize, def, f)
		case func(bool, bool) bool:
			genericReduceLastB(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Int:
		var def int

		if def, ok = defaultValue.(int); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Ints()
		rt := retVal.Ints()
		switch f := fn.(type) {
		case func([]int) int:
			reduceLastI(dt, rt, dimSize, def, f)
		case func(int, int) int:
			genericReduceLastI(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Int8:
		var def int8

		if def, ok = defaultValue.(int8); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Int8s()
		rt := retVal.Int8s()
		switch f := fn.(type) {
		case func([]int8) int8:
			reduceLastI8(dt, rt, dimSize, def, f)
		case func(int8, int8) int8:
			genericReduceLastI8(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Int16:
		var def int16

		if def, ok = defaultValue.(int16); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Int16s()
		rt := retVal.Int16s()
		switch f := fn.(type) {
		case func([]int16) int16:
			reduceLastI16(dt, rt, dimSize, def, f)
		case func(int16, int16) int16:
			genericReduceLastI16(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Int32:
		var def int32

		if def, ok = defaultValue.(int32); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Int32s()
		rt := retVal.Int32s()
		switch f := fn.(type) {
		case func([]int32) int32:
			reduceLastI32(dt, rt, dimSize, def, f)
		case func(int32, int32) int32:
			genericReduceLastI32(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Int64:
		var def int64

		if def, ok = defaultValue.(int64); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Int64s()
		rt := retVal.Int64s()
		switch f := fn.(type) {
		case func([]int64) int64:
			reduceLastI64(dt, rt, dimSize, def, f)
		case func(int64, int64) int64:
			genericReduceLastI64(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Uint:
		var def uint

		if def, ok = defaultValue.(uint); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Uints()
		rt := retVal.Uints()
		switch f := fn.(type) {
		case func([]uint) uint:
			reduceLastU(dt, rt, dimSize, def, f)
		case func(uint, uint) uint:
			genericReduceLastU(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Uint8:
		var def uint8

		if def, ok = defaultValue.(uint8); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Uint8s()
		rt := retVal.Uint8s()
		switch f := fn.(type) {
		case func([]uint8) uint8:
			reduceLastU8(dt, rt, dimSize, def, f)
		case func(uint8, uint8) uint8:
			genericReduceLastU8(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Uint16:
		var def uint16

		if def, ok = defaultValue.(uint16); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Uint16s()
		rt := retVal.Uint16s()
		switch f := fn.(type) {
		case func([]uint16) uint16:
			reduceLastU16(dt, rt, dimSize, def, f)
		case func(uint16, uint16) uint16:
			genericReduceLastU16(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Uint32:
		var def uint32

		if def, ok = defaultValue.(uint32); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Uint32s()
		rt := retVal.Uint32s()
		switch f := fn.(type) {
		case func([]uint32) uint32:
			reduceLastU32(dt, rt, dimSize, def, f)
		case func(uint32, uint32) uint32:
			genericReduceLastU32(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Uint64:
		var def uint64

		if def, ok = defaultValue.(uint64); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Uint64s()
		rt := retVal.Uint64s()
		switch f := fn.(type) {
		case func([]uint64) uint64:
			reduceLastU64(dt, rt, dimSize, def, f)
		case func(uint64, uint64) uint64:
			genericReduceLastU64(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Uintptr:
		var def uintptr

		if def, ok = defaultValue.(uintptr); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Uintptrs()
		rt := retVal.Uintptrs()
		switch f := fn.(type) {
		case func([]uintptr) uintptr:
			reduceLastUintptr(dt, rt, dimSize, def, f)
		case func(uintptr, uintptr) uintptr:
			genericReduceLastUintptr(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Float32:
		var def float32

		if def, ok = defaultValue.(float32); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Float32s()
		rt := retVal.Float32s()
		switch f := fn.(type) {
		case func([]float32) float32:
			reduceLastF32(dt, rt, dimSize, def, f)
		case func(float32, float32) float32:
			genericReduceLastF32(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Float64:
		var def float64

		if def, ok = defaultValue.(float64); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Float64s()
		rt := retVal.Float64s()
		switch f := fn.(type) {
		case func([]float64) float64:
			reduceLastF64(dt, rt, dimSize, def, f)
		case func(float64, float64) float64:
			genericReduceLastF64(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Complex64:
		var def complex64

		if def, ok = defaultValue.(complex64); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Complex64s()
		rt := retVal.Complex64s()
		switch f := fn.(type) {
		case func([]complex64) complex64:
			reduceLastC64(dt, rt, dimSize, def, f)
		case func(complex64, complex64) complex64:
			genericReduceLastC64(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case Complex128:
		var def complex128

		if def, ok = defaultValue.(complex128); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Complex128s()
		rt := retVal.Complex128s()
		switch f := fn.(type) {
		case func([]complex128) complex128:
			reduceLastC128(dt, rt, dimSize, def, f)
		case func(complex128, complex128) complex128:
			genericReduceLastC128(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case String:
		var def string

		if def, ok = defaultValue.(string); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.Strings()
		rt := retVal.Strings()
		switch f := fn.(type) {
		case func([]string) string:
			reduceLastStr(dt, rt, dimSize, def, f)
		case func(string, string) string:
			genericReduceLastStr(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	case UnsafePointer:
		var def unsafe.Pointer

		if def, ok = defaultValue.(unsafe.Pointer); !ok {
			return errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		dt := data.UnsafePointers()
		rt := retVal.UnsafePointers()
		switch f := fn.(type) {
		case func([]unsafe.Pointer) unsafe.Pointer:
			reduceLastUnsafePointer(dt, rt, dimSize, def, f)
		case func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer:
			genericReduceLastUnsafePointer(dt, rt, dimSize, def, f)
		default:
			return errors.Errorf(reductionErrMsg, fn)
		}
		return nil
	default:
		return errors.Errorf("Unsupported type %v for ReduceLast", t)
	}
}

func (e E) ReduceDefault(t reflect.Type, data *storage.Header, retVal *storage.Header, dim0 int, dimSize int, outerStride int, stride int, expected int, fn interface{}) (err error) {
	var ok bool
	switch t {
	case Bool:
		var f func(bool, bool) bool
		if f, ok = fn.(func(bool, bool) bool); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Bools()
		rt := retVal.Bools()
		reduceDefaultB(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Int:
		var f func(int, int) int
		if f, ok = fn.(func(int, int) int); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Ints()
		rt := retVal.Ints()
		reduceDefaultI(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Int8:
		var f func(int8, int8) int8
		if f, ok = fn.(func(int8, int8) int8); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Int8s()
		rt := retVal.Int8s()
		reduceDefaultI8(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Int16:
		var f func(int16, int16) int16
		if f, ok = fn.(func(int16, int16) int16); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Int16s()
		rt := retVal.Int16s()
		reduceDefaultI16(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Int32:
		var f func(int32, int32) int32
		if f, ok = fn.(func(int32, int32) int32); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Int32s()
		rt := retVal.Int32s()
		reduceDefaultI32(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Int64:
		var f func(int64, int64) int64
		if f, ok = fn.(func(int64, int64) int64); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Int64s()
		rt := retVal.Int64s()
		reduceDefaultI64(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Uint:
		var f func(uint, uint) uint
		if f, ok = fn.(func(uint, uint) uint); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Uints()
		rt := retVal.Uints()
		reduceDefaultU(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Uint8:
		var f func(uint8, uint8) uint8
		if f, ok = fn.(func(uint8, uint8) uint8); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Uint8s()
		rt := retVal.Uint8s()
		reduceDefaultU8(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Uint16:
		var f func(uint16, uint16) uint16
		if f, ok = fn.(func(uint16, uint16) uint16); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Uint16s()
		rt := retVal.Uint16s()
		reduceDefaultU16(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Uint32:
		var f func(uint32, uint32) uint32
		if f, ok = fn.(func(uint32, uint32) uint32); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Uint32s()
		rt := retVal.Uint32s()
		reduceDefaultU32(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Uint64:
		var f func(uint64, uint64) uint64
		if f, ok = fn.(func(uint64, uint64) uint64); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Uint64s()
		rt := retVal.Uint64s()
		reduceDefaultU64(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Uintptr:
		var f func(uintptr, uintptr) uintptr
		if f, ok = fn.(func(uintptr, uintptr) uintptr); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Uintptrs()
		rt := retVal.Uintptrs()
		reduceDefaultUintptr(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Float32:
		var f func(float32, float32) float32
		if f, ok = fn.(func(float32, float32) float32); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Float32s()
		rt := retVal.Float32s()
		reduceDefaultF32(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Float64:
		var f func(float64, float64) float64
		if f, ok = fn.(func(float64, float64) float64); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Float64s()
		rt := retVal.Float64s()
		reduceDefaultF64(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Complex64:
		var f func(complex64, complex64) complex64
		if f, ok = fn.(func(complex64, complex64) complex64); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Complex64s()
		rt := retVal.Complex64s()
		reduceDefaultC64(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case Complex128:
		var f func(complex128, complex128) complex128
		if f, ok = fn.(func(complex128, complex128) complex128); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Complex128s()
		rt := retVal.Complex128s()
		reduceDefaultC128(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case String:
		var f func(string, string) string
		if f, ok = fn.(func(string, string) string); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.Strings()
		rt := retVal.Strings()
		reduceDefaultStr(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	case UnsafePointer:
		var f func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
		if f, ok = fn.(func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer); !ok {
			return errors.Errorf(reductionErrMsg, fn)
		}
		dt := data.UnsafePointers()
		rt := retVal.UnsafePointers()
		reduceDefaultUnsafePointer(dt, rt, dim0, dimSize, outerStride, stride, expected, f)
		return nil
	default:
		return errors.Errorf("Unsupported type %v for ReduceDefault", t)
	}
}

func (e E) Reduce(t reflect.Type, a *storage.Header, defaultValue interface{}, fn interface{}) (retVal interface{}, err error) {
	var ok bool
	switch t {
	case Bool:
		var f func(bool, bool) bool
		var def bool
		if f, ok = fn.(func(bool, bool) bool); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(bool); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceB(f, def, a.Bools()...)
		return
	case Int:
		var f func(int, int) int
		var def int
		if f, ok = fn.(func(int, int) int); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(int); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceI(f, def, a.Ints()...)
		return
	case Int8:
		var f func(int8, int8) int8
		var def int8
		if f, ok = fn.(func(int8, int8) int8); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(int8); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceI8(f, def, a.Int8s()...)
		return
	case Int16:
		var f func(int16, int16) int16
		var def int16
		if f, ok = fn.(func(int16, int16) int16); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(int16); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceI16(f, def, a.Int16s()...)
		return
	case Int32:
		var f func(int32, int32) int32
		var def int32
		if f, ok = fn.(func(int32, int32) int32); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(int32); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceI32(f, def, a.Int32s()...)
		return
	case Int64:
		var f func(int64, int64) int64
		var def int64
		if f, ok = fn.(func(int64, int64) int64); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(int64); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceI64(f, def, a.Int64s()...)
		return
	case Uint:
		var f func(uint, uint) uint
		var def uint
		if f, ok = fn.(func(uint, uint) uint); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(uint); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceU(f, def, a.Uints()...)
		return
	case Uint8:
		var f func(uint8, uint8) uint8
		var def uint8
		if f, ok = fn.(func(uint8, uint8) uint8); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(uint8); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceU8(f, def, a.Uint8s()...)
		return
	case Uint16:
		var f func(uint16, uint16) uint16
		var def uint16
		if f, ok = fn.(func(uint16, uint16) uint16); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(uint16); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceU16(f, def, a.Uint16s()...)
		return
	case Uint32:
		var f func(uint32, uint32) uint32
		var def uint32
		if f, ok = fn.(func(uint32, uint32) uint32); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(uint32); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceU32(f, def, a.Uint32s()...)
		return
	case Uint64:
		var f func(uint64, uint64) uint64
		var def uint64
		if f, ok = fn.(func(uint64, uint64) uint64); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(uint64); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceU64(f, def, a.Uint64s()...)
		return
	case Uintptr:
		var f func(uintptr, uintptr) uintptr
		var def uintptr
		if f, ok = fn.(func(uintptr, uintptr) uintptr); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(uintptr); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceUintptr(f, def, a.Uintptrs()...)
		return
	case Float32:
		var f func(float32, float32) float32
		var def float32
		if f, ok = fn.(func(float32, float32) float32); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(float32); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceF32(f, def, a.Float32s()...)
		return
	case Float64:
		var f func(float64, float64) float64
		var def float64
		if f, ok = fn.(func(float64, float64) float64); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(float64); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceF64(f, def, a.Float64s()...)
		return
	case Complex64:
		var f func(complex64, complex64) complex64
		var def complex64
		if f, ok = fn.(func(complex64, complex64) complex64); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(complex64); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceC64(f, def, a.Complex64s()...)
		return
	case Complex128:
		var f func(complex128, complex128) complex128
		var def complex128
		if f, ok = fn.(func(complex128, complex128) complex128); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(complex128); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceC128(f, def, a.Complex128s()...)
		return
	case String:
		var f func(string, string) string
		var def string
		if f, ok = fn.(func(string, string) string); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(string); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceStr(f, def, a.Strings()...)
		return
	case UnsafePointer:
		var f func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
		var def unsafe.Pointer
		if f, ok = fn.(func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer); !ok {
			return nil, errors.Errorf(reductionErrMsg, fn)
		}
		if def, ok = defaultValue.(unsafe.Pointer); !ok {
			return nil, errors.Errorf(defaultValueErrMsg, def, defaultValue, defaultValue)
		}
		retVal = ReduceUnsafePointer(f, def, a.UnsafePointers()...)
		return
	default:
		return nil, errors.Errorf("Unsupported type %v for Reduce", t)
	}
}
