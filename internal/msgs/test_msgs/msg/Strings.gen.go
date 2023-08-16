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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/msg/strings.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/Strings", StringsTypeSupport)
	typemap.RegisterMessage("test_msgs/msg/Strings", StringsTypeSupport)
}
const (
	Strings_STRING_CONST string = "Hello world!"
)

// Do not create instances of this type directly. Always use NewStrings
// function instead.
type Strings struct {
	StringValue string `yaml:"string_value"`
	StringValueDefault1 string `yaml:"string_value_default1"`
	StringValueDefault2 string `yaml:"string_value_default2"`
	StringValueDefault3 string `yaml:"string_value_default3"`
	StringValueDefault4 string `yaml:"string_value_default4"`
	StringValueDefault5 string `yaml:"string_value_default5"`
	BoundedStringValue string `yaml:"bounded_string_value"`
	BoundedStringValueDefault1 string `yaml:"bounded_string_value_default1"`
	BoundedStringValueDefault2 string `yaml:"bounded_string_value_default2"`
	BoundedStringValueDefault3 string `yaml:"bounded_string_value_default3"`
	BoundedStringValueDefault4 string `yaml:"bounded_string_value_default4"`
	BoundedStringValueDefault5 string `yaml:"bounded_string_value_default5"`
}

// NewStrings creates a new Strings with default values.
func NewStrings() *Strings {
	self := Strings{}
	self.SetDefaults()
	return &self
}

func (t *Strings) Clone() *Strings {
	c := &Strings{}
	c.StringValue = t.StringValue
	c.StringValueDefault1 = t.StringValueDefault1
	c.StringValueDefault2 = t.StringValueDefault2
	c.StringValueDefault3 = t.StringValueDefault3
	c.StringValueDefault4 = t.StringValueDefault4
	c.StringValueDefault5 = t.StringValueDefault5
	c.BoundedStringValue = t.BoundedStringValue
	c.BoundedStringValueDefault1 = t.BoundedStringValueDefault1
	c.BoundedStringValueDefault2 = t.BoundedStringValueDefault2
	c.BoundedStringValueDefault3 = t.BoundedStringValueDefault3
	c.BoundedStringValueDefault4 = t.BoundedStringValueDefault4
	c.BoundedStringValueDefault5 = t.BoundedStringValueDefault5
	return c
}

func (t *Strings) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Strings) SetDefaults() {
	t.StringValue = ""
	t.StringValueDefault1 = "Hello world!"
	t.StringValueDefault2 = "Hello'world!"
	t.StringValueDefault3 = "Hello\"world!"
	t.StringValueDefault4 = "Hello'world!"
	t.StringValueDefault5 = "Hello\"world!"
	t.BoundedStringValue = ""
	t.BoundedStringValueDefault1 = "Hello world!"
	t.BoundedStringValueDefault2 = "Hello'world!"
	t.BoundedStringValueDefault3 = "Hello\"world!"
	t.BoundedStringValueDefault4 = "Hello'world!"
	t.BoundedStringValueDefault5 = "Hello\"world!"
}

func (t *Strings) GetTypeSupport() types.MessageTypeSupport {
	return StringsTypeSupport
}

// StringsPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type StringsPublisher struct {
	*rclgo.Publisher
}

// NewStringsPublisher creates and returns a new publisher for the
// Strings
func NewStringsPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*StringsPublisher, error) {
	pub, err := node.NewPublisher(topic_name, StringsTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &StringsPublisher{pub}, nil
}

func (p *StringsPublisher) Publish(msg *Strings) error {
	return p.Publisher.Publish(msg)
}

// StringsSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type StringsSubscription struct {
	*rclgo.Subscription
}

// StringsSubscriptionCallback type is used to provide a subscription
// handler function for a StringsSubscription.
type StringsSubscriptionCallback func(msg *Strings, info *rclgo.MessageInfo, err error)

// NewStringsSubscription creates and returns a new subscription for the
// Strings
func NewStringsSubscription(node *rclgo.Node, topic_name string, subscriptionCallback StringsSubscriptionCallback) (*StringsSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Strings
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, StringsTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &StringsSubscription{sub}, nil
}

func (s *StringsSubscription) TakeMessage(out *Strings) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneStringsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneStringsSlice(dst, src []Strings) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var StringsTypeSupport types.MessageTypeSupport = _StringsTypeSupport{}

type _StringsTypeSupport struct{}

func (t _StringsTypeSupport) New() types.Message {
	return NewStrings()
}

func (t _StringsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__Strings
	return (unsafe.Pointer)(C.test_msgs__msg__Strings__create())
}

func (t _StringsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__Strings__destroy((*C.test_msgs__msg__Strings)(pointer_to_free))
}

func (t _StringsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Strings)
	mem := (*C.test_msgs__msg__Strings)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.string_value), m.StringValue)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.string_value_default1), m.StringValueDefault1)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.string_value_default2), m.StringValueDefault2)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.string_value_default3), m.StringValueDefault3)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.string_value_default4), m.StringValueDefault4)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.string_value_default5), m.StringValueDefault5)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.bounded_string_value), m.BoundedStringValue)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.bounded_string_value_default1), m.BoundedStringValueDefault1)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.bounded_string_value_default2), m.BoundedStringValueDefault2)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.bounded_string_value_default3), m.BoundedStringValueDefault3)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.bounded_string_value_default4), m.BoundedStringValueDefault4)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.bounded_string_value_default5), m.BoundedStringValueDefault5)
}

func (t _StringsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Strings)
	mem := (*C.test_msgs__msg__Strings)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.StringValue, unsafe.Pointer(&mem.string_value))
	primitives.StringAsGoStruct(&m.StringValueDefault1, unsafe.Pointer(&mem.string_value_default1))
	primitives.StringAsGoStruct(&m.StringValueDefault2, unsafe.Pointer(&mem.string_value_default2))
	primitives.StringAsGoStruct(&m.StringValueDefault3, unsafe.Pointer(&mem.string_value_default3))
	primitives.StringAsGoStruct(&m.StringValueDefault4, unsafe.Pointer(&mem.string_value_default4))
	primitives.StringAsGoStruct(&m.StringValueDefault5, unsafe.Pointer(&mem.string_value_default5))
	primitives.StringAsGoStruct(&m.BoundedStringValue, unsafe.Pointer(&mem.bounded_string_value))
	primitives.StringAsGoStruct(&m.BoundedStringValueDefault1, unsafe.Pointer(&mem.bounded_string_value_default1))
	primitives.StringAsGoStruct(&m.BoundedStringValueDefault2, unsafe.Pointer(&mem.bounded_string_value_default2))
	primitives.StringAsGoStruct(&m.BoundedStringValueDefault3, unsafe.Pointer(&mem.bounded_string_value_default3))
	primitives.StringAsGoStruct(&m.BoundedStringValueDefault4, unsafe.Pointer(&mem.bounded_string_value_default4))
	primitives.StringAsGoStruct(&m.BoundedStringValueDefault5, unsafe.Pointer(&mem.bounded_string_value_default5))
}

func (t _StringsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__Strings())
}

type CStrings = C.test_msgs__msg__Strings
type CStrings__Sequence = C.test_msgs__msg__Strings__Sequence

func Strings__Sequence_to_Go(goSlice *[]Strings, cSlice CStrings__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Strings, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		StringsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Strings__Sequence_to_C(cSlice *CStrings__Sequence, goSlice []Strings) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.test_msgs__msg__Strings)(C.malloc(C.sizeof_struct_test_msgs__msg__Strings * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		StringsTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Strings__Array_to_Go(goSlice []Strings, cSlice []CStrings) {
	for i := 0; i < len(cSlice); i++ {
		StringsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Strings__Array_to_C(cSlice []CStrings, goSlice []Strings) {
	for i := 0; i < len(goSlice); i++ {
		StringsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
