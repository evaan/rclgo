/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package test_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltest_msgs__rosidl_typesupport_c -ltest_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <test_msgs/msg/unbounded_sequences.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("test_msgs/UnboundedSequences", &UnboundedSequences{})
}

// Do not create instances of this type directly. Always use NewUnboundedSequences
// function instead.
type UnboundedSequences struct {
	BoolValues []bool `yaml:"bool_values"`// Unbounded sequences of different types
	ByteValues []byte `yaml:"byte_values"`// Unbounded sequences of different types
	CharValues []byte `yaml:"char_values"`// Unbounded sequences of different types
	Float32Values []float32 `yaml:"float32_values"`// Unbounded sequences of different types
	Float64Values []float64 `yaml:"float64_values"`// Unbounded sequences of different types
	Int8Values []int8 `yaml:"int8_values"`// Unbounded sequences of different types
	Uint8Values []uint8 `yaml:"uint8_values"`// Unbounded sequences of different types
	Int16Values []int16 `yaml:"int16_values"`// Unbounded sequences of different types
	Uint16Values []uint16 `yaml:"uint16_values"`// Unbounded sequences of different types
	Int32Values []int32 `yaml:"int32_values"`// Unbounded sequences of different types
	Uint32Values []uint32 `yaml:"uint32_values"`// Unbounded sequences of different types
	Int64Values []int64 `yaml:"int64_values"`// Unbounded sequences of different types
	Uint64Values []uint64 `yaml:"uint64_values"`// Unbounded sequences of different types
	StringValues []rosidl_runtime_c.String `yaml:"string_values"`// Unbounded sequences of different types
	BasicTypesValues []BasicTypes `yaml:"basic_types_values"`// Unbounded sequences of different types
	ConstantsValues []Constants `yaml:"constants_values"`// Unbounded sequences of different types
	DefaultsValues []Defaults `yaml:"defaults_values"`// Unbounded sequences of different types
	BoolValuesDefault []bool `yaml:"bool_values_default"`// Unbounded sequences of different types
	ByteValuesDefault []byte `yaml:"byte_values_default"`// Unbounded sequences of different types
	CharValuesDefault []byte `yaml:"char_values_default"`// Unbounded sequences of different types
	Float32ValuesDefault []float32 `yaml:"float32_values_default"`// Unbounded sequences of different types
	Float64ValuesDefault []float64 `yaml:"float64_values_default"`// Unbounded sequences of different types
	Int8ValuesDefault []int8 `yaml:"int8_values_default"`// Unbounded sequences of different types
	Uint8ValuesDefault []uint8 `yaml:"uint8_values_default"`// Unbounded sequences of different types
	Int16ValuesDefault []int16 `yaml:"int16_values_default"`// Unbounded sequences of different types
	Uint16ValuesDefault []uint16 `yaml:"uint16_values_default"`// Unbounded sequences of different types
	Int32ValuesDefault []int32 `yaml:"int32_values_default"`// Unbounded sequences of different types
	Uint32ValuesDefault []uint32 `yaml:"uint32_values_default"`// Unbounded sequences of different types
	Int64ValuesDefault []int64 `yaml:"int64_values_default"`// Unbounded sequences of different types
	Uint64ValuesDefault []uint64 `yaml:"uint64_values_default"`// Unbounded sequences of different types
	StringValuesDefault []rosidl_runtime_c.String `yaml:"string_values_default"`// Unbounded sequences of different types
	AlignmentCheck int32 `yaml:"alignment_check"`// Unbounded sequences of different typesRegression test: check alignment of basic field after a sequence field is correct
}

// NewUnboundedSequences creates a new UnboundedSequences with default values.
func NewUnboundedSequences() *UnboundedSequences {
	self := UnboundedSequences{}
	self.SetDefaults(nil)
	return &self
}

func (t *UnboundedSequences) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.BoolValuesDefault = []bool{false, true, false}
	t.ByteValuesDefault = []byte{0, 1, 255}
	t.CharValuesDefault = []byte{0, 1, 127}
	t.Float32ValuesDefault = []float32{1.125, 0.0, -1.125}
	t.Float64ValuesDefault = []float64{3.1415, 0.0, -3.1415}
	t.Int8ValuesDefault = []int8{0, 127, -128}
	t.Uint8ValuesDefault = []uint8{0, 1, 255}
	t.Int16ValuesDefault = []int16{0, 32767, -32768}
	t.Uint16ValuesDefault = []uint16{0, 1, 65535}
	t.Int32ValuesDefault = []int32{0, 2147483647, -2147483648}
	t.Uint32ValuesDefault = []uint32{0, 1, 4294967295}
	t.Int64ValuesDefault = []int64{0, 9223372036854775807, -9223372036854775808}
	t.Uint64ValuesDefault = []uint64{0, 1, 18446744073709551615}
	t.StringValuesDefault = make([]rosidl_runtime_c.String, 3)
	t.StringValuesDefault[0].SetDefaults("")
	t.StringValuesDefault[1].SetDefaults("max value")
	t.StringValuesDefault[2].SetDefaults("min value")
	
	return t
}

func (t *UnboundedSequences) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__UnboundedSequences())
}
func (t *UnboundedSequences) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__UnboundedSequences
	return (unsafe.Pointer)(C.test_msgs__msg__UnboundedSequences__create())
}
func (t *UnboundedSequences) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__UnboundedSequences__destroy((*C.test_msgs__msg__UnboundedSequences)(pointer_to_free))
}
func (t *UnboundedSequences) AsCStruct() unsafe.Pointer {
	mem := (*C.test_msgs__msg__UnboundedSequences)(t.PrepareMemory())
	rosidl_runtime_c.Bool__Sequence_to_C((*rosidl_runtime_c.CBool__Sequence)(unsafe.Pointer(&mem.bool_values)), t.BoolValues)
	rosidl_runtime_c.Byte__Sequence_to_C((*rosidl_runtime_c.CByte__Sequence)(unsafe.Pointer(&mem.byte_values)), t.ByteValues)
	rosidl_runtime_c.Char__Sequence_to_C((*rosidl_runtime_c.CChar__Sequence)(unsafe.Pointer(&mem.char_values)), t.CharValues)
	rosidl_runtime_c.Float32__Sequence_to_C((*rosidl_runtime_c.CFloat32__Sequence)(unsafe.Pointer(&mem.float32_values)), t.Float32Values)
	rosidl_runtime_c.Float64__Sequence_to_C((*rosidl_runtime_c.CFloat64__Sequence)(unsafe.Pointer(&mem.float64_values)), t.Float64Values)
	rosidl_runtime_c.Int8__Sequence_to_C((*rosidl_runtime_c.CInt8__Sequence)(unsafe.Pointer(&mem.int8_values)), t.Int8Values)
	rosidl_runtime_c.Uint8__Sequence_to_C((*rosidl_runtime_c.CUint8__Sequence)(unsafe.Pointer(&mem.uint8_values)), t.Uint8Values)
	rosidl_runtime_c.Int16__Sequence_to_C((*rosidl_runtime_c.CInt16__Sequence)(unsafe.Pointer(&mem.int16_values)), t.Int16Values)
	rosidl_runtime_c.Uint16__Sequence_to_C((*rosidl_runtime_c.CUint16__Sequence)(unsafe.Pointer(&mem.uint16_values)), t.Uint16Values)
	rosidl_runtime_c.Int32__Sequence_to_C((*rosidl_runtime_c.CInt32__Sequence)(unsafe.Pointer(&mem.int32_values)), t.Int32Values)
	rosidl_runtime_c.Uint32__Sequence_to_C((*rosidl_runtime_c.CUint32__Sequence)(unsafe.Pointer(&mem.uint32_values)), t.Uint32Values)
	rosidl_runtime_c.Int64__Sequence_to_C((*rosidl_runtime_c.CInt64__Sequence)(unsafe.Pointer(&mem.int64_values)), t.Int64Values)
	rosidl_runtime_c.Uint64__Sequence_to_C((*rosidl_runtime_c.CUint64__Sequence)(unsafe.Pointer(&mem.uint64_values)), t.Uint64Values)
	rosidl_runtime_c.String__Sequence_to_C((*rosidl_runtime_c.CString__Sequence)(unsafe.Pointer(&mem.string_values)), t.StringValues)
	BasicTypes__Sequence_to_C(&mem.basic_types_values, t.BasicTypesValues)
	Constants__Sequence_to_C(&mem.constants_values, t.ConstantsValues)
	Defaults__Sequence_to_C(&mem.defaults_values, t.DefaultsValues)
	rosidl_runtime_c.Bool__Sequence_to_C((*rosidl_runtime_c.CBool__Sequence)(unsafe.Pointer(&mem.bool_values_default)), t.BoolValuesDefault)
	rosidl_runtime_c.Byte__Sequence_to_C((*rosidl_runtime_c.CByte__Sequence)(unsafe.Pointer(&mem.byte_values_default)), t.ByteValuesDefault)
	rosidl_runtime_c.Char__Sequence_to_C((*rosidl_runtime_c.CChar__Sequence)(unsafe.Pointer(&mem.char_values_default)), t.CharValuesDefault)
	rosidl_runtime_c.Float32__Sequence_to_C((*rosidl_runtime_c.CFloat32__Sequence)(unsafe.Pointer(&mem.float32_values_default)), t.Float32ValuesDefault)
	rosidl_runtime_c.Float64__Sequence_to_C((*rosidl_runtime_c.CFloat64__Sequence)(unsafe.Pointer(&mem.float64_values_default)), t.Float64ValuesDefault)
	rosidl_runtime_c.Int8__Sequence_to_C((*rosidl_runtime_c.CInt8__Sequence)(unsafe.Pointer(&mem.int8_values_default)), t.Int8ValuesDefault)
	rosidl_runtime_c.Uint8__Sequence_to_C((*rosidl_runtime_c.CUint8__Sequence)(unsafe.Pointer(&mem.uint8_values_default)), t.Uint8ValuesDefault)
	rosidl_runtime_c.Int16__Sequence_to_C((*rosidl_runtime_c.CInt16__Sequence)(unsafe.Pointer(&mem.int16_values_default)), t.Int16ValuesDefault)
	rosidl_runtime_c.Uint16__Sequence_to_C((*rosidl_runtime_c.CUint16__Sequence)(unsafe.Pointer(&mem.uint16_values_default)), t.Uint16ValuesDefault)
	rosidl_runtime_c.Int32__Sequence_to_C((*rosidl_runtime_c.CInt32__Sequence)(unsafe.Pointer(&mem.int32_values_default)), t.Int32ValuesDefault)
	rosidl_runtime_c.Uint32__Sequence_to_C((*rosidl_runtime_c.CUint32__Sequence)(unsafe.Pointer(&mem.uint32_values_default)), t.Uint32ValuesDefault)
	rosidl_runtime_c.Int64__Sequence_to_C((*rosidl_runtime_c.CInt64__Sequence)(unsafe.Pointer(&mem.int64_values_default)), t.Int64ValuesDefault)
	rosidl_runtime_c.Uint64__Sequence_to_C((*rosidl_runtime_c.CUint64__Sequence)(unsafe.Pointer(&mem.uint64_values_default)), t.Uint64ValuesDefault)
	rosidl_runtime_c.String__Sequence_to_C((*rosidl_runtime_c.CString__Sequence)(unsafe.Pointer(&mem.string_values_default)), t.StringValuesDefault)
	mem.alignment_check = C.int32_t(t.AlignmentCheck)
	return unsafe.Pointer(mem)
}
func (t *UnboundedSequences) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.test_msgs__msg__UnboundedSequences)(ros2_message_buffer)
	rosidl_runtime_c.Bool__Sequence_to_Go(&t.BoolValues, *(*rosidl_runtime_c.CBool__Sequence)(unsafe.Pointer(&mem.bool_values)))
	rosidl_runtime_c.Byte__Sequence_to_Go(&t.ByteValues, *(*rosidl_runtime_c.CByte__Sequence)(unsafe.Pointer(&mem.byte_values)))
	rosidl_runtime_c.Char__Sequence_to_Go(&t.CharValues, *(*rosidl_runtime_c.CChar__Sequence)(unsafe.Pointer(&mem.char_values)))
	rosidl_runtime_c.Float32__Sequence_to_Go(&t.Float32Values, *(*rosidl_runtime_c.CFloat32__Sequence)(unsafe.Pointer(&mem.float32_values)))
	rosidl_runtime_c.Float64__Sequence_to_Go(&t.Float64Values, *(*rosidl_runtime_c.CFloat64__Sequence)(unsafe.Pointer(&mem.float64_values)))
	rosidl_runtime_c.Int8__Sequence_to_Go(&t.Int8Values, *(*rosidl_runtime_c.CInt8__Sequence)(unsafe.Pointer(&mem.int8_values)))
	rosidl_runtime_c.Uint8__Sequence_to_Go(&t.Uint8Values, *(*rosidl_runtime_c.CUint8__Sequence)(unsafe.Pointer(&mem.uint8_values)))
	rosidl_runtime_c.Int16__Sequence_to_Go(&t.Int16Values, *(*rosidl_runtime_c.CInt16__Sequence)(unsafe.Pointer(&mem.int16_values)))
	rosidl_runtime_c.Uint16__Sequence_to_Go(&t.Uint16Values, *(*rosidl_runtime_c.CUint16__Sequence)(unsafe.Pointer(&mem.uint16_values)))
	rosidl_runtime_c.Int32__Sequence_to_Go(&t.Int32Values, *(*rosidl_runtime_c.CInt32__Sequence)(unsafe.Pointer(&mem.int32_values)))
	rosidl_runtime_c.Uint32__Sequence_to_Go(&t.Uint32Values, *(*rosidl_runtime_c.CUint32__Sequence)(unsafe.Pointer(&mem.uint32_values)))
	rosidl_runtime_c.Int64__Sequence_to_Go(&t.Int64Values, *(*rosidl_runtime_c.CInt64__Sequence)(unsafe.Pointer(&mem.int64_values)))
	rosidl_runtime_c.Uint64__Sequence_to_Go(&t.Uint64Values, *(*rosidl_runtime_c.CUint64__Sequence)(unsafe.Pointer(&mem.uint64_values)))
	rosidl_runtime_c.String__Sequence_to_Go(&t.StringValues, *(*rosidl_runtime_c.CString__Sequence)(unsafe.Pointer(&mem.string_values)))
	BasicTypes__Sequence_to_Go(&t.BasicTypesValues, mem.basic_types_values)
	Constants__Sequence_to_Go(&t.ConstantsValues, mem.constants_values)
	Defaults__Sequence_to_Go(&t.DefaultsValues, mem.defaults_values)
	rosidl_runtime_c.Bool__Sequence_to_Go(&t.BoolValuesDefault, *(*rosidl_runtime_c.CBool__Sequence)(unsafe.Pointer(&mem.bool_values_default)))
	rosidl_runtime_c.Byte__Sequence_to_Go(&t.ByteValuesDefault, *(*rosidl_runtime_c.CByte__Sequence)(unsafe.Pointer(&mem.byte_values_default)))
	rosidl_runtime_c.Char__Sequence_to_Go(&t.CharValuesDefault, *(*rosidl_runtime_c.CChar__Sequence)(unsafe.Pointer(&mem.char_values_default)))
	rosidl_runtime_c.Float32__Sequence_to_Go(&t.Float32ValuesDefault, *(*rosidl_runtime_c.CFloat32__Sequence)(unsafe.Pointer(&mem.float32_values_default)))
	rosidl_runtime_c.Float64__Sequence_to_Go(&t.Float64ValuesDefault, *(*rosidl_runtime_c.CFloat64__Sequence)(unsafe.Pointer(&mem.float64_values_default)))
	rosidl_runtime_c.Int8__Sequence_to_Go(&t.Int8ValuesDefault, *(*rosidl_runtime_c.CInt8__Sequence)(unsafe.Pointer(&mem.int8_values_default)))
	rosidl_runtime_c.Uint8__Sequence_to_Go(&t.Uint8ValuesDefault, *(*rosidl_runtime_c.CUint8__Sequence)(unsafe.Pointer(&mem.uint8_values_default)))
	rosidl_runtime_c.Int16__Sequence_to_Go(&t.Int16ValuesDefault, *(*rosidl_runtime_c.CInt16__Sequence)(unsafe.Pointer(&mem.int16_values_default)))
	rosidl_runtime_c.Uint16__Sequence_to_Go(&t.Uint16ValuesDefault, *(*rosidl_runtime_c.CUint16__Sequence)(unsafe.Pointer(&mem.uint16_values_default)))
	rosidl_runtime_c.Int32__Sequence_to_Go(&t.Int32ValuesDefault, *(*rosidl_runtime_c.CInt32__Sequence)(unsafe.Pointer(&mem.int32_values_default)))
	rosidl_runtime_c.Uint32__Sequence_to_Go(&t.Uint32ValuesDefault, *(*rosidl_runtime_c.CUint32__Sequence)(unsafe.Pointer(&mem.uint32_values_default)))
	rosidl_runtime_c.Int64__Sequence_to_Go(&t.Int64ValuesDefault, *(*rosidl_runtime_c.CInt64__Sequence)(unsafe.Pointer(&mem.int64_values_default)))
	rosidl_runtime_c.Uint64__Sequence_to_Go(&t.Uint64ValuesDefault, *(*rosidl_runtime_c.CUint64__Sequence)(unsafe.Pointer(&mem.uint64_values_default)))
	rosidl_runtime_c.String__Sequence_to_Go(&t.StringValuesDefault, *(*rosidl_runtime_c.CString__Sequence)(unsafe.Pointer(&mem.string_values_default)))
	t.AlignmentCheck = int32(mem.alignment_check)
}
func (t *UnboundedSequences) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CUnboundedSequences = C.test_msgs__msg__UnboundedSequences
type CUnboundedSequences__Sequence = C.test_msgs__msg__UnboundedSequences__Sequence

func UnboundedSequences__Sequence_to_Go(goSlice *[]UnboundedSequences, cSlice CUnboundedSequences__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UnboundedSequences, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.test_msgs__msg__UnboundedSequences__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__UnboundedSequences * uintptr(i)),
		))
		(*goSlice)[i] = UnboundedSequences{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func UnboundedSequences__Sequence_to_C(cSlice *CUnboundedSequences__Sequence, goSlice []UnboundedSequences) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__msg__UnboundedSequences)(C.malloc((C.size_t)(C.sizeof_struct_test_msgs__msg__UnboundedSequences * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.test_msgs__msg__UnboundedSequences)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__UnboundedSequences * uintptr(i)),
		))
		*cIdx = *(*C.test_msgs__msg__UnboundedSequences)(v.AsCStruct())
	}
}
func UnboundedSequences__Array_to_Go(goSlice []UnboundedSequences, cSlice []CUnboundedSequences) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func UnboundedSequences__Array_to_C(cSlice []CUnboundedSequences, goSlice []UnboundedSequences) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.test_msgs__msg__UnboundedSequences)(goSlice[i].AsCStruct())
	}
}

