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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/int16.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Int16", Int16TypeSupport)
}

// Do not create instances of this type directly. Always use NewInt16
// function instead.
type Int16 struct {
	Data int16 `yaml:"data"`// This is an example message of using a primitive datatype, int16.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewInt16 creates a new Int16 with default values.
func NewInt16() *Int16 {
	self := Int16{}
	self.SetDefaults()
	return &self
}

func (t *Int16) Clone() *Int16 {
	c := &Int16{}
	c.Data = t.Data
	return c
}

func (t *Int16) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Int16) SetDefaults() {
	t.Data = 0
}

// Int16Publisher wraps rclgo.Publisher to provide type safe helper
// functions
type Int16Publisher struct {
	*rclgo.Publisher
}

// NewInt16Publisher creates and returns a new publisher for the
// Int16
func NewInt16Publisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Int16Publisher, error) {
	pub, err := node.NewPublisher(topic_name, Int16TypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Int16Publisher{pub}, nil
}

func (p *Int16Publisher) Publish(msg *Int16) error {
	return p.Publisher.Publish(msg)
}

// Int16Subscription wraps rclgo.Subscription to provide type safe helper
// functions
type Int16Subscription struct {
	*rclgo.Subscription
}

// NewInt16Subscription creates and returns a new subscription for the
// Int16
func NewInt16Subscription(node *rclgo.Node, topic_name string, subscriptionCallback rclgo.SubscriptionCallback) (*Int16Subscription, error) {
	sub, err := node.NewSubscription(topic_name, Int16TypeSupport, subscriptionCallback)
	if err != nil {
		return nil, err
	}
	return &Int16Subscription{sub}, nil
}

func (s *Int16Subscription) TakeMessage(out *Int16) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneInt16Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInt16Slice(dst, src []Int16) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Int16TypeSupport types.MessageTypeSupport = _Int16TypeSupport{}

type _Int16TypeSupport struct{}

func (t _Int16TypeSupport) New() types.Message {
	return NewInt16()
}

func (t _Int16TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Int16
	return (unsafe.Pointer)(C.example_interfaces__msg__Int16__create())
}

func (t _Int16TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Int16__destroy((*C.example_interfaces__msg__Int16)(pointer_to_free))
}

func (t _Int16TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Int16)
	mem := (*C.example_interfaces__msg__Int16)(dst)
	mem.data = C.int16_t(m.Data)
}

func (t _Int16TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Int16)
	mem := (*C.example_interfaces__msg__Int16)(ros2_message_buffer)
	m.Data = int16(mem.data)
}

func (t _Int16TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Int16())
}

type CInt16 = C.example_interfaces__msg__Int16
type CInt16__Sequence = C.example_interfaces__msg__Int16__Sequence

func Int16__Sequence_to_Go(goSlice *[]Int16, cSlice CInt16__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Int16, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__Int16__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Int16 * uintptr(i)),
		))
		Int16TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Int16__Sequence_to_C(cSlice *CInt16__Sequence, goSlice []Int16) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Int16)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__Int16 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__Int16)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Int16 * uintptr(i)),
		))
		Int16TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Int16__Array_to_Go(goSlice []Int16, cSlice []CInt16) {
	for i := 0; i < len(cSlice); i++ {
		Int16TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Int16__Array_to_C(cSlice []CInt16, goSlice []Int16) {
	for i := 0; i < len(goSlice); i++ {
		Int16TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
