/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/msg/basic_types.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/BasicTypes", BasicTypesTypeSupport)
	typemap.RegisterMessage("test_msgs/msg/BasicTypes", BasicTypesTypeSupport)
}

// Do not create instances of this type directly. Always use NewBasicTypes
// function instead.
type BasicTypes struct {
	BoolValue bool `yaml:"bool_value"`
	ByteValue byte `yaml:"byte_value"`
	CharValue byte `yaml:"char_value"`
	Float32Value float32 `yaml:"float32_value"`
	Float64Value float64 `yaml:"float64_value"`
	Int8Value int8 `yaml:"int8_value"`
	Uint8Value uint8 `yaml:"uint8_value"`
	Int16Value int16 `yaml:"int16_value"`
	Uint16Value uint16 `yaml:"uint16_value"`
	Int32Value int32 `yaml:"int32_value"`
	Uint32Value uint32 `yaml:"uint32_value"`
	Int64Value int64 `yaml:"int64_value"`
	Uint64Value uint64 `yaml:"uint64_value"`
}

// NewBasicTypes creates a new BasicTypes with default values.
func NewBasicTypes() *BasicTypes {
	self := BasicTypes{}
	self.SetDefaults()
	return &self
}

func (t *BasicTypes) Clone() *BasicTypes {
	c := &BasicTypes{}
	c.BoolValue = t.BoolValue
	c.ByteValue = t.ByteValue
	c.CharValue = t.CharValue
	c.Float32Value = t.Float32Value
	c.Float64Value = t.Float64Value
	c.Int8Value = t.Int8Value
	c.Uint8Value = t.Uint8Value
	c.Int16Value = t.Int16Value
	c.Uint16Value = t.Uint16Value
	c.Int32Value = t.Int32Value
	c.Uint32Value = t.Uint32Value
	c.Int64Value = t.Int64Value
	c.Uint64Value = t.Uint64Value
	return c
}

func (t *BasicTypes) CloneMsg() types.Message {
	return t.Clone()
}

func (t *BasicTypes) SetDefaults() {
	t.BoolValue = false
	t.ByteValue = 0
	t.CharValue = 0
	t.Float32Value = 0
	t.Float64Value = 0
	t.Int8Value = 0
	t.Uint8Value = 0
	t.Int16Value = 0
	t.Uint16Value = 0
	t.Int32Value = 0
	t.Uint32Value = 0
	t.Int64Value = 0
	t.Uint64Value = 0
}

func (t *BasicTypes) GetTypeSupport() types.MessageTypeSupport {
	return BasicTypesTypeSupport
}

// BasicTypesPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type BasicTypesPublisher struct {
	*rclgo.Publisher
}

// NewBasicTypesPublisher creates and returns a new publisher for the
// BasicTypes
func NewBasicTypesPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*BasicTypesPublisher, error) {
	pub, err := node.NewPublisher(topic_name, BasicTypesTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &BasicTypesPublisher{pub}, nil
}

func (p *BasicTypesPublisher) Publish(msg *BasicTypes) error {
	return p.Publisher.Publish(msg)
}

// BasicTypesSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type BasicTypesSubscription struct {
	*rclgo.Subscription
}

// BasicTypesSubscriptionCallback type is used to provide a subscription
// handler function for a BasicTypesSubscription.
type BasicTypesSubscriptionCallback func(msg *BasicTypes, info *rclgo.RmwMessageInfo, err error)

// NewBasicTypesSubscription creates and returns a new subscription for the
// BasicTypes
func NewBasicTypesSubscription(node *rclgo.Node, topic_name string, subscriptionCallback BasicTypesSubscriptionCallback) (*BasicTypesSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg BasicTypes
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, BasicTypesTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &BasicTypesSubscription{sub}, nil
}

func (s *BasicTypesSubscription) TakeMessage(out *BasicTypes) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneBasicTypesSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneBasicTypesSlice(dst, src []BasicTypes) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var BasicTypesTypeSupport types.MessageTypeSupport = _BasicTypesTypeSupport{}

type _BasicTypesTypeSupport struct{}

func (t _BasicTypesTypeSupport) New() types.Message {
	return NewBasicTypes()
}

func (t _BasicTypesTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__BasicTypes
	return (unsafe.Pointer)(C.test_msgs__msg__BasicTypes__create())
}

func (t _BasicTypesTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__BasicTypes__destroy((*C.test_msgs__msg__BasicTypes)(pointer_to_free))
}

func (t _BasicTypesTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*BasicTypes)
	mem := (*C.test_msgs__msg__BasicTypes)(dst)
	mem.bool_value = C.bool(m.BoolValue)
	mem.byte_value = C.uint8_t(m.ByteValue)
	mem.char_value = C.uchar(m.CharValue)
	mem.float32_value = C.float(m.Float32Value)
	mem.float64_value = C.double(m.Float64Value)
	mem.int8_value = C.int8_t(m.Int8Value)
	mem.uint8_value = C.uint8_t(m.Uint8Value)
	mem.int16_value = C.int16_t(m.Int16Value)
	mem.uint16_value = C.uint16_t(m.Uint16Value)
	mem.int32_value = C.int32_t(m.Int32Value)
	mem.uint32_value = C.uint32_t(m.Uint32Value)
	mem.int64_value = C.int64_t(m.Int64Value)
	mem.uint64_value = C.uint64_t(m.Uint64Value)
}

func (t _BasicTypesTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*BasicTypes)
	mem := (*C.test_msgs__msg__BasicTypes)(ros2_message_buffer)
	m.BoolValue = bool(mem.bool_value)
	m.ByteValue = byte(mem.byte_value)
	m.CharValue = byte(mem.char_value)
	m.Float32Value = float32(mem.float32_value)
	m.Float64Value = float64(mem.float64_value)
	m.Int8Value = int8(mem.int8_value)
	m.Uint8Value = uint8(mem.uint8_value)
	m.Int16Value = int16(mem.int16_value)
	m.Uint16Value = uint16(mem.uint16_value)
	m.Int32Value = int32(mem.int32_value)
	m.Uint32Value = uint32(mem.uint32_value)
	m.Int64Value = int64(mem.int64_value)
	m.Uint64Value = uint64(mem.uint64_value)
}

func (t _BasicTypesTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__BasicTypes())
}

type CBasicTypes = C.test_msgs__msg__BasicTypes
type CBasicTypes__Sequence = C.test_msgs__msg__BasicTypes__Sequence

func BasicTypes__Sequence_to_Go(goSlice *[]BasicTypes, cSlice CBasicTypes__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]BasicTypes, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		BasicTypesTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func BasicTypes__Sequence_to_C(cSlice *CBasicTypes__Sequence, goSlice []BasicTypes) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__msg__BasicTypes)(C.malloc(C.sizeof_struct_test_msgs__msg__BasicTypes * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		BasicTypesTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func BasicTypes__Array_to_Go(goSlice []BasicTypes, cSlice []CBasicTypes) {
	for i := 0; i < len(cSlice); i++ {
		BasicTypesTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func BasicTypes__Array_to_C(cSlice []CBasicTypes, goSlice []BasicTypes) {
	for i := 0; i < len(goSlice); i++ {
		BasicTypesTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
