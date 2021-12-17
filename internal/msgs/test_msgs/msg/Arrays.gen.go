/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package test_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltest_msgs__rosidl_typesupport_c -ltest_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/msg/arrays.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/Arrays", ArraysTypeSupport)
}

// Do not create instances of this type directly. Always use NewArrays
// function instead.
type Arrays struct {
	BoolValues [3]bool `yaml:"bool_values"`// Arrays of different types
	ByteValues [3]byte `yaml:"byte_values"`// Arrays of different types
	CharValues [3]byte `yaml:"char_values"`// Arrays of different types
	Float32Values [3]float32 `yaml:"float32_values"`// Arrays of different types
	Float64Values [3]float64 `yaml:"float64_values"`// Arrays of different types
	Int8Values [3]int8 `yaml:"int8_values"`// Arrays of different types
	Uint8Values [3]uint8 `yaml:"uint8_values"`// Arrays of different types
	Int16Values [3]int16 `yaml:"int16_values"`// Arrays of different types
	Uint16Values [3]uint16 `yaml:"uint16_values"`// Arrays of different types
	Int32Values [3]int32 `yaml:"int32_values"`// Arrays of different types
	Uint32Values [3]uint32 `yaml:"uint32_values"`// Arrays of different types
	Int64Values [3]int64 `yaml:"int64_values"`// Arrays of different types
	Uint64Values [3]uint64 `yaml:"uint64_values"`// Arrays of different types
	StringValues [3]string `yaml:"string_values"`// Arrays of different types
	BasicTypesValues [3]BasicTypes `yaml:"basic_types_values"`// Arrays of different types
	ConstantsValues [3]Constants `yaml:"constants_values"`// Arrays of different types
	DefaultsValues [3]Defaults `yaml:"defaults_values"`// Arrays of different types
	BoolValuesDefault [3]bool `yaml:"bool_values_default"`// Arrays of different types
	ByteValuesDefault [3]byte `yaml:"byte_values_default"`// Arrays of different types
	CharValuesDefault [3]byte `yaml:"char_values_default"`// Arrays of different types
	Float32ValuesDefault [3]float32 `yaml:"float32_values_default"`// Arrays of different types
	Float64ValuesDefault [3]float64 `yaml:"float64_values_default"`// Arrays of different types
	Int8ValuesDefault [3]int8 `yaml:"int8_values_default"`// Arrays of different types
	Uint8ValuesDefault [3]uint8 `yaml:"uint8_values_default"`// Arrays of different types
	Int16ValuesDefault [3]int16 `yaml:"int16_values_default"`// Arrays of different types
	Uint16ValuesDefault [3]uint16 `yaml:"uint16_values_default"`// Arrays of different types
	Int32ValuesDefault [3]int32 `yaml:"int32_values_default"`// Arrays of different types
	Uint32ValuesDefault [3]uint32 `yaml:"uint32_values_default"`// Arrays of different types
	Int64ValuesDefault [3]int64 `yaml:"int64_values_default"`// Arrays of different types
	Uint64ValuesDefault [3]uint64 `yaml:"uint64_values_default"`// Arrays of different types
	StringValuesDefault [3]string `yaml:"string_values_default"`// Arrays of different types
	AlignmentCheck int32 `yaml:"alignment_check"`// Arrays of different typesRegression test: check alignment of basic field after an array field is correct
}

// NewArrays creates a new Arrays with default values.
func NewArrays() *Arrays {
	self := Arrays{}
	self.SetDefaults()
	return &self
}

func (t *Arrays) Clone() *Arrays {
	c := &Arrays{}
	c.BoolValues = t.BoolValues
	c.ByteValues = t.ByteValues
	c.CharValues = t.CharValues
	c.Float32Values = t.Float32Values
	c.Float64Values = t.Float64Values
	c.Int8Values = t.Int8Values
	c.Uint8Values = t.Uint8Values
	c.Int16Values = t.Int16Values
	c.Uint16Values = t.Uint16Values
	c.Int32Values = t.Int32Values
	c.Uint32Values = t.Uint32Values
	c.Int64Values = t.Int64Values
	c.Uint64Values = t.Uint64Values
	c.StringValues = t.StringValues
	CloneBasicTypesSlice(c.BasicTypesValues[:], t.BasicTypesValues[:])
	CloneConstantsSlice(c.ConstantsValues[:], t.ConstantsValues[:])
	CloneDefaultsSlice(c.DefaultsValues[:], t.DefaultsValues[:])
	c.BoolValuesDefault = t.BoolValuesDefault
	c.ByteValuesDefault = t.ByteValuesDefault
	c.CharValuesDefault = t.CharValuesDefault
	c.Float32ValuesDefault = t.Float32ValuesDefault
	c.Float64ValuesDefault = t.Float64ValuesDefault
	c.Int8ValuesDefault = t.Int8ValuesDefault
	c.Uint8ValuesDefault = t.Uint8ValuesDefault
	c.Int16ValuesDefault = t.Int16ValuesDefault
	c.Uint16ValuesDefault = t.Uint16ValuesDefault
	c.Int32ValuesDefault = t.Int32ValuesDefault
	c.Uint32ValuesDefault = t.Uint32ValuesDefault
	c.Int64ValuesDefault = t.Int64ValuesDefault
	c.Uint64ValuesDefault = t.Uint64ValuesDefault
	c.StringValuesDefault = t.StringValuesDefault
	c.AlignmentCheck = t.AlignmentCheck
	return c
}

func (t *Arrays) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Arrays) SetDefaults() {
	t.BoolValues = [3]bool{}
	t.ByteValues = [3]byte{}
	t.CharValues = [3]byte{}
	t.Float32Values = [3]float32{}
	t.Float64Values = [3]float64{}
	t.Int8Values = [3]int8{}
	t.Uint8Values = [3]uint8{}
	t.Int16Values = [3]int16{}
	t.Uint16Values = [3]uint16{}
	t.Int32Values = [3]int32{}
	t.Uint32Values = [3]uint32{}
	t.Int64Values = [3]int64{}
	t.Uint64Values = [3]uint64{}
	t.StringValues = [3]string{}
	for i := range t.BasicTypesValues {
		t.BasicTypesValues[i].SetDefaults()
	}
	for i := range t.ConstantsValues {
		t.ConstantsValues[i].SetDefaults()
	}
	for i := range t.DefaultsValues {
		t.DefaultsValues[i].SetDefaults()
	}
	t.BoolValuesDefault = [3]bool{false,true,false}
	t.ByteValuesDefault = [3]byte{0,1,255}
	t.CharValuesDefault = [3]byte{0,1,127}
	t.Float32ValuesDefault = [3]float32{1.125,0.0,-1.125}
	t.Float64ValuesDefault = [3]float64{3.1415,0.0,-3.1415}
	t.Int8ValuesDefault = [3]int8{0,127,-128}
	t.Uint8ValuesDefault = [3]uint8{0,1,255}
	t.Int16ValuesDefault = [3]int16{0,32767,-32768}
	t.Uint16ValuesDefault = [3]uint16{0,1,65535}
	t.Int32ValuesDefault = [3]int32{0,2147483647,-2147483648}
	t.Uint32ValuesDefault = [3]uint32{0,1,4294967295}
	t.Int64ValuesDefault = [3]int64{0,9223372036854775807,-9223372036854775808}
	t.Uint64ValuesDefault = [3]uint64{0,1,18446744073709551615}
	t.StringValuesDefault = [3]string{"","max value","min value"}
	t.AlignmentCheck = 0
}

// ArraysPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ArraysPublisher struct {
	*rclgo.Publisher
}

// NewArraysPublisher creates and returns a new publisher for the
// Arrays
func NewArraysPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ArraysPublisher, error) {
	pub, err := node.NewPublisher(topic_name, ArraysTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ArraysPublisher{pub}, nil
}

func (p *ArraysPublisher) Publish(msg *Arrays) error {
	return p.Publisher.Publish(msg)
}

// ArraysSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ArraysSubscription struct {
	*rclgo.Subscription
}

// NewArraysSubscription creates and returns a new subscription for the
// Arrays
func NewArraysSubscription(node *rclgo.Node, topic_name string, subscriptionCallback rclgo.SubscriptionCallback) (*ArraysSubscription, error) {
	sub, err := node.NewSubscription(topic_name, ArraysTypeSupport, subscriptionCallback)
	if err != nil {
		return nil, err
	}
	return &ArraysSubscription{sub}, nil
}

func (s *ArraysSubscription) TakeMessage(out *Arrays) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneArraysSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneArraysSlice(dst, src []Arrays) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ArraysTypeSupport types.MessageTypeSupport = _ArraysTypeSupport{}

type _ArraysTypeSupport struct{}

func (t _ArraysTypeSupport) New() types.Message {
	return NewArrays()
}

func (t _ArraysTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__Arrays
	return (unsafe.Pointer)(C.test_msgs__msg__Arrays__create())
}

func (t _ArraysTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__Arrays__destroy((*C.test_msgs__msg__Arrays)(pointer_to_free))
}

func (t _ArraysTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Arrays)
	mem := (*C.test_msgs__msg__Arrays)(dst)
	cSlice_bool_values := mem.bool_values[:]
	primitives.Bool__Array_to_C(*(*[]primitives.CBool)(unsafe.Pointer(&cSlice_bool_values)), m.BoolValues[:])
	cSlice_byte_values := mem.byte_values[:]
	primitives.Byte__Array_to_C(*(*[]primitives.CByte)(unsafe.Pointer(&cSlice_byte_values)), m.ByteValues[:])
	cSlice_char_values := mem.char_values[:]
	primitives.Char__Array_to_C(*(*[]primitives.CChar)(unsafe.Pointer(&cSlice_char_values)), m.CharValues[:])
	cSlice_float32_values := mem.float32_values[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_float32_values)), m.Float32Values[:])
	cSlice_float64_values := mem.float64_values[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_float64_values)), m.Float64Values[:])
	cSlice_int8_values := mem.int8_values[:]
	primitives.Int8__Array_to_C(*(*[]primitives.CInt8)(unsafe.Pointer(&cSlice_int8_values)), m.Int8Values[:])
	cSlice_uint8_values := mem.uint8_values[:]
	primitives.Uint8__Array_to_C(*(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_uint8_values)), m.Uint8Values[:])
	cSlice_int16_values := mem.int16_values[:]
	primitives.Int16__Array_to_C(*(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_int16_values)), m.Int16Values[:])
	cSlice_uint16_values := mem.uint16_values[:]
	primitives.Uint16__Array_to_C(*(*[]primitives.CUint16)(unsafe.Pointer(&cSlice_uint16_values)), m.Uint16Values[:])
	cSlice_int32_values := mem.int32_values[:]
	primitives.Int32__Array_to_C(*(*[]primitives.CInt32)(unsafe.Pointer(&cSlice_int32_values)), m.Int32Values[:])
	cSlice_uint32_values := mem.uint32_values[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_uint32_values)), m.Uint32Values[:])
	cSlice_int64_values := mem.int64_values[:]
	primitives.Int64__Array_to_C(*(*[]primitives.CInt64)(unsafe.Pointer(&cSlice_int64_values)), m.Int64Values[:])
	cSlice_uint64_values := mem.uint64_values[:]
	primitives.Uint64__Array_to_C(*(*[]primitives.CUint64)(unsafe.Pointer(&cSlice_uint64_values)), m.Uint64Values[:])
	cSlice_string_values := mem.string_values[:]
	primitives.String__Array_to_C(*(*[]primitives.CString)(unsafe.Pointer(&cSlice_string_values)), m.StringValues[:])
	BasicTypes__Array_to_C(mem.basic_types_values[:], m.BasicTypesValues[:])
	Constants__Array_to_C(mem.constants_values[:], m.ConstantsValues[:])
	Defaults__Array_to_C(mem.defaults_values[:], m.DefaultsValues[:])
	cSlice_bool_values_default := mem.bool_values_default[:]
	primitives.Bool__Array_to_C(*(*[]primitives.CBool)(unsafe.Pointer(&cSlice_bool_values_default)), m.BoolValuesDefault[:])
	cSlice_byte_values_default := mem.byte_values_default[:]
	primitives.Byte__Array_to_C(*(*[]primitives.CByte)(unsafe.Pointer(&cSlice_byte_values_default)), m.ByteValuesDefault[:])
	cSlice_char_values_default := mem.char_values_default[:]
	primitives.Char__Array_to_C(*(*[]primitives.CChar)(unsafe.Pointer(&cSlice_char_values_default)), m.CharValuesDefault[:])
	cSlice_float32_values_default := mem.float32_values_default[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_float32_values_default)), m.Float32ValuesDefault[:])
	cSlice_float64_values_default := mem.float64_values_default[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_float64_values_default)), m.Float64ValuesDefault[:])
	cSlice_int8_values_default := mem.int8_values_default[:]
	primitives.Int8__Array_to_C(*(*[]primitives.CInt8)(unsafe.Pointer(&cSlice_int8_values_default)), m.Int8ValuesDefault[:])
	cSlice_uint8_values_default := mem.uint8_values_default[:]
	primitives.Uint8__Array_to_C(*(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_uint8_values_default)), m.Uint8ValuesDefault[:])
	cSlice_int16_values_default := mem.int16_values_default[:]
	primitives.Int16__Array_to_C(*(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_int16_values_default)), m.Int16ValuesDefault[:])
	cSlice_uint16_values_default := mem.uint16_values_default[:]
	primitives.Uint16__Array_to_C(*(*[]primitives.CUint16)(unsafe.Pointer(&cSlice_uint16_values_default)), m.Uint16ValuesDefault[:])
	cSlice_int32_values_default := mem.int32_values_default[:]
	primitives.Int32__Array_to_C(*(*[]primitives.CInt32)(unsafe.Pointer(&cSlice_int32_values_default)), m.Int32ValuesDefault[:])
	cSlice_uint32_values_default := mem.uint32_values_default[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_uint32_values_default)), m.Uint32ValuesDefault[:])
	cSlice_int64_values_default := mem.int64_values_default[:]
	primitives.Int64__Array_to_C(*(*[]primitives.CInt64)(unsafe.Pointer(&cSlice_int64_values_default)), m.Int64ValuesDefault[:])
	cSlice_uint64_values_default := mem.uint64_values_default[:]
	primitives.Uint64__Array_to_C(*(*[]primitives.CUint64)(unsafe.Pointer(&cSlice_uint64_values_default)), m.Uint64ValuesDefault[:])
	cSlice_string_values_default := mem.string_values_default[:]
	primitives.String__Array_to_C(*(*[]primitives.CString)(unsafe.Pointer(&cSlice_string_values_default)), m.StringValuesDefault[:])
	mem.alignment_check = C.int32_t(m.AlignmentCheck)
}

func (t _ArraysTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Arrays)
	mem := (*C.test_msgs__msg__Arrays)(ros2_message_buffer)
	cSlice_bool_values := mem.bool_values[:]
	primitives.Bool__Array_to_Go(m.BoolValues[:], *(*[]primitives.CBool)(unsafe.Pointer(&cSlice_bool_values)))
	cSlice_byte_values := mem.byte_values[:]
	primitives.Byte__Array_to_Go(m.ByteValues[:], *(*[]primitives.CByte)(unsafe.Pointer(&cSlice_byte_values)))
	cSlice_char_values := mem.char_values[:]
	primitives.Char__Array_to_Go(m.CharValues[:], *(*[]primitives.CChar)(unsafe.Pointer(&cSlice_char_values)))
	cSlice_float32_values := mem.float32_values[:]
	primitives.Float32__Array_to_Go(m.Float32Values[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_float32_values)))
	cSlice_float64_values := mem.float64_values[:]
	primitives.Float64__Array_to_Go(m.Float64Values[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_float64_values)))
	cSlice_int8_values := mem.int8_values[:]
	primitives.Int8__Array_to_Go(m.Int8Values[:], *(*[]primitives.CInt8)(unsafe.Pointer(&cSlice_int8_values)))
	cSlice_uint8_values := mem.uint8_values[:]
	primitives.Uint8__Array_to_Go(m.Uint8Values[:], *(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_uint8_values)))
	cSlice_int16_values := mem.int16_values[:]
	primitives.Int16__Array_to_Go(m.Int16Values[:], *(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_int16_values)))
	cSlice_uint16_values := mem.uint16_values[:]
	primitives.Uint16__Array_to_Go(m.Uint16Values[:], *(*[]primitives.CUint16)(unsafe.Pointer(&cSlice_uint16_values)))
	cSlice_int32_values := mem.int32_values[:]
	primitives.Int32__Array_to_Go(m.Int32Values[:], *(*[]primitives.CInt32)(unsafe.Pointer(&cSlice_int32_values)))
	cSlice_uint32_values := mem.uint32_values[:]
	primitives.Uint32__Array_to_Go(m.Uint32Values[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_uint32_values)))
	cSlice_int64_values := mem.int64_values[:]
	primitives.Int64__Array_to_Go(m.Int64Values[:], *(*[]primitives.CInt64)(unsafe.Pointer(&cSlice_int64_values)))
	cSlice_uint64_values := mem.uint64_values[:]
	primitives.Uint64__Array_to_Go(m.Uint64Values[:], *(*[]primitives.CUint64)(unsafe.Pointer(&cSlice_uint64_values)))
	cSlice_string_values := mem.string_values[:]
	primitives.String__Array_to_Go(m.StringValues[:], *(*[]primitives.CString)(unsafe.Pointer(&cSlice_string_values)))
	BasicTypes__Array_to_Go(m.BasicTypesValues[:], mem.basic_types_values[:])
	Constants__Array_to_Go(m.ConstantsValues[:], mem.constants_values[:])
	Defaults__Array_to_Go(m.DefaultsValues[:], mem.defaults_values[:])
	cSlice_bool_values_default := mem.bool_values_default[:]
	primitives.Bool__Array_to_Go(m.BoolValuesDefault[:], *(*[]primitives.CBool)(unsafe.Pointer(&cSlice_bool_values_default)))
	cSlice_byte_values_default := mem.byte_values_default[:]
	primitives.Byte__Array_to_Go(m.ByteValuesDefault[:], *(*[]primitives.CByte)(unsafe.Pointer(&cSlice_byte_values_default)))
	cSlice_char_values_default := mem.char_values_default[:]
	primitives.Char__Array_to_Go(m.CharValuesDefault[:], *(*[]primitives.CChar)(unsafe.Pointer(&cSlice_char_values_default)))
	cSlice_float32_values_default := mem.float32_values_default[:]
	primitives.Float32__Array_to_Go(m.Float32ValuesDefault[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_float32_values_default)))
	cSlice_float64_values_default := mem.float64_values_default[:]
	primitives.Float64__Array_to_Go(m.Float64ValuesDefault[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_float64_values_default)))
	cSlice_int8_values_default := mem.int8_values_default[:]
	primitives.Int8__Array_to_Go(m.Int8ValuesDefault[:], *(*[]primitives.CInt8)(unsafe.Pointer(&cSlice_int8_values_default)))
	cSlice_uint8_values_default := mem.uint8_values_default[:]
	primitives.Uint8__Array_to_Go(m.Uint8ValuesDefault[:], *(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_uint8_values_default)))
	cSlice_int16_values_default := mem.int16_values_default[:]
	primitives.Int16__Array_to_Go(m.Int16ValuesDefault[:], *(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_int16_values_default)))
	cSlice_uint16_values_default := mem.uint16_values_default[:]
	primitives.Uint16__Array_to_Go(m.Uint16ValuesDefault[:], *(*[]primitives.CUint16)(unsafe.Pointer(&cSlice_uint16_values_default)))
	cSlice_int32_values_default := mem.int32_values_default[:]
	primitives.Int32__Array_to_Go(m.Int32ValuesDefault[:], *(*[]primitives.CInt32)(unsafe.Pointer(&cSlice_int32_values_default)))
	cSlice_uint32_values_default := mem.uint32_values_default[:]
	primitives.Uint32__Array_to_Go(m.Uint32ValuesDefault[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_uint32_values_default)))
	cSlice_int64_values_default := mem.int64_values_default[:]
	primitives.Int64__Array_to_Go(m.Int64ValuesDefault[:], *(*[]primitives.CInt64)(unsafe.Pointer(&cSlice_int64_values_default)))
	cSlice_uint64_values_default := mem.uint64_values_default[:]
	primitives.Uint64__Array_to_Go(m.Uint64ValuesDefault[:], *(*[]primitives.CUint64)(unsafe.Pointer(&cSlice_uint64_values_default)))
	cSlice_string_values_default := mem.string_values_default[:]
	primitives.String__Array_to_Go(m.StringValuesDefault[:], *(*[]primitives.CString)(unsafe.Pointer(&cSlice_string_values_default)))
	m.AlignmentCheck = int32(mem.alignment_check)
}

func (t _ArraysTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__Arrays())
}

type CArrays = C.test_msgs__msg__Arrays
type CArrays__Sequence = C.test_msgs__msg__Arrays__Sequence

func Arrays__Sequence_to_Go(goSlice *[]Arrays, cSlice CArrays__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Arrays, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.test_msgs__msg__Arrays__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__Arrays * uintptr(i)),
		))
		ArraysTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Arrays__Sequence_to_C(cSlice *CArrays__Sequence, goSlice []Arrays) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__msg__Arrays)(C.malloc((C.size_t)(C.sizeof_struct_test_msgs__msg__Arrays * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.test_msgs__msg__Arrays)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__Arrays * uintptr(i)),
		))
		ArraysTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Arrays__Array_to_Go(goSlice []Arrays, cSlice []CArrays) {
	for i := 0; i < len(cSlice); i++ {
		ArraysTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Arrays__Array_to_C(cSlice []CArrays, goSlice []Arrays) {
	for i := 0; i < len(goSlice); i++ {
		ArraysTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
