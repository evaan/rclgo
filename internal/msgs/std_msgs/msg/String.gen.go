/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package std_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/string.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/String", StringTypeSupport)
	typemap.RegisterMessage("std_msgs/msg/String", StringTypeSupport)
}

// Do not create instances of this type directly. Always use NewString
// function instead.
type String struct {
	Data string `yaml:"data"`
}

// NewString creates a new String with default values.
func NewString() *String {
	self := String{}
	self.SetDefaults()
	return &self
}

func (t *String) Clone() *String {
	c := &String{}
	c.Data = t.Data
	return c
}

func (t *String) CloneMsg() types.Message {
	return t.Clone()
}

func (t *String) SetDefaults() {
	t.Data = ""
}

func (t *String) GetTypeSupport() types.MessageTypeSupport {
	return StringTypeSupport
}

// StringPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type StringPublisher struct {
	*rclgo.Publisher
}

// NewStringPublisher creates and returns a new publisher for the
// String
func NewStringPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*StringPublisher, error) {
	pub, err := node.NewPublisher(topic_name, StringTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &StringPublisher{pub}, nil
}

func (p *StringPublisher) Publish(msg *String) error {
	return p.Publisher.Publish(msg)
}

// StringSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type StringSubscription struct {
	*rclgo.Subscription
}

// StringSubscriptionCallback type is used to provide a subscription
// handler function for a StringSubscription.
type StringSubscriptionCallback func(msg *String, info *rclgo.RmwMessageInfo, err error)

// NewStringSubscription creates and returns a new subscription for the
// String
func NewStringSubscription(node *rclgo.Node, topic_name string, subscriptionCallback StringSubscriptionCallback) (*StringSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg String
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, StringTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &StringSubscription{sub}, nil
}

func (s *StringSubscription) TakeMessage(out *String) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneStringSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneStringSlice(dst, src []String) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var StringTypeSupport types.MessageTypeSupport = _StringTypeSupport{}

type _StringTypeSupport struct{}

func (t _StringTypeSupport) New() types.Message {
	return NewString()
}

func (t _StringTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__String
	return (unsafe.Pointer)(C.std_msgs__msg__String__create())
}

func (t _StringTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__String__destroy((*C.std_msgs__msg__String)(pointer_to_free))
}

func (t _StringTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*String)
	mem := (*C.std_msgs__msg__String)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.data), m.Data)
}

func (t _StringTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*String)
	mem := (*C.std_msgs__msg__String)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.Data, unsafe.Pointer(&mem.data))
}

func (t _StringTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__String())
}

type CString = C.std_msgs__msg__String
type CString__Sequence = C.std_msgs__msg__String__Sequence

func String__Sequence_to_Go(goSlice *[]String, cSlice CString__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]String, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		StringTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func String__Sequence_to_C(cSlice *CString__Sequence, goSlice []String) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__String)(C.malloc(C.sizeof_struct_std_msgs__msg__String * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		StringTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func String__Array_to_Go(goSlice []String, cSlice []CString) {
	for i := 0; i < len(cSlice); i++ {
		StringTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func String__Array_to_C(cSlice []CString, goSlice []String) {
	for i := 0; i < len(goSlice); i++ {
		StringTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
