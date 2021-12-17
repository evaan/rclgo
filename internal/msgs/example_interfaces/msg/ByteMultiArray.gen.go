/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/byte_multi_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/ByteMultiArray", ByteMultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewByteMultiArray
// function instead.
type ByteMultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []byte `yaml:"data"`// array of data
}

// NewByteMultiArray creates a new ByteMultiArray with default values.
func NewByteMultiArray() *ByteMultiArray {
	self := ByteMultiArray{}
	self.SetDefaults()
	return &self
}

func (t *ByteMultiArray) Clone() *ByteMultiArray {
	c := &ByteMultiArray{}
	c.Layout = *t.Layout.Clone()
	if t.Data != nil {
		c.Data = make([]byte, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *ByteMultiArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ByteMultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	t.Data = nil
}

// ByteMultiArrayPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ByteMultiArrayPublisher struct {
	*rclgo.Publisher
}

// NewByteMultiArrayPublisher creates and returns a new publisher for the
// ByteMultiArray
func NewByteMultiArrayPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ByteMultiArrayPublisher, error) {
	pub, err := node.NewPublisher(topic_name, ByteMultiArrayTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ByteMultiArrayPublisher{pub}, nil
}

func (p *ByteMultiArrayPublisher) Publish(msg *ByteMultiArray) error {
	return p.Publisher.Publish(msg)
}

// ByteMultiArraySubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ByteMultiArraySubscription struct {
	*rclgo.Subscription
}

// NewByteMultiArraySubscription creates and returns a new subscription for the
// ByteMultiArray
func NewByteMultiArraySubscription(node *rclgo.Node, topic_name string, subscriptionCallback rclgo.SubscriptionCallback) (*ByteMultiArraySubscription, error) {
	sub, err := node.NewSubscription(topic_name, ByteMultiArrayTypeSupport, subscriptionCallback)
	if err != nil {
		return nil, err
	}
	return &ByteMultiArraySubscription{sub}, nil
}

func (s *ByteMultiArraySubscription) TakeMessage(out *ByteMultiArray) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneByteMultiArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneByteMultiArraySlice(dst, src []ByteMultiArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ByteMultiArrayTypeSupport types.MessageTypeSupport = _ByteMultiArrayTypeSupport{}

type _ByteMultiArrayTypeSupport struct{}

func (t _ByteMultiArrayTypeSupport) New() types.Message {
	return NewByteMultiArray()
}

func (t _ByteMultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__ByteMultiArray
	return (unsafe.Pointer)(C.example_interfaces__msg__ByteMultiArray__create())
}

func (t _ByteMultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__ByteMultiArray__destroy((*C.example_interfaces__msg__ByteMultiArray)(pointer_to_free))
}

func (t _ByteMultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ByteMultiArray)
	mem := (*C.example_interfaces__msg__ByteMultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	primitives.Byte__Sequence_to_C((*primitives.CByte__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _ByteMultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ByteMultiArray)
	mem := (*C.example_interfaces__msg__ByteMultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	primitives.Byte__Sequence_to_Go(&m.Data, *(*primitives.CByte__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _ByteMultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__ByteMultiArray())
}

type CByteMultiArray = C.example_interfaces__msg__ByteMultiArray
type CByteMultiArray__Sequence = C.example_interfaces__msg__ByteMultiArray__Sequence

func ByteMultiArray__Sequence_to_Go(goSlice *[]ByteMultiArray, cSlice CByteMultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ByteMultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__ByteMultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__ByteMultiArray * uintptr(i)),
		))
		ByteMultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ByteMultiArray__Sequence_to_C(cSlice *CByteMultiArray__Sequence, goSlice []ByteMultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__ByteMultiArray)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__ByteMultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__ByteMultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__ByteMultiArray * uintptr(i)),
		))
		ByteMultiArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ByteMultiArray__Array_to_Go(goSlice []ByteMultiArray, cSlice []CByteMultiArray) {
	for i := 0; i < len(cSlice); i++ {
		ByteMultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ByteMultiArray__Array_to_C(cSlice []CByteMultiArray, goSlice []ByteMultiArray) {
	for i := 0; i < len(goSlice); i++ {
		ByteMultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
