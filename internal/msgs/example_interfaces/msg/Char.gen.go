/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package example_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/char.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Char", CharTypeSupport)
	typemap.RegisterMessage("example_interfaces/msg/Char", CharTypeSupport)
}

// Do not create instances of this type directly. Always use NewChar
// function instead.
type Char struct {
	Data byte `yaml:"data"`// This is an example message of using a primitive datatype, char.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewChar creates a new Char with default values.
func NewChar() *Char {
	self := Char{}
	self.SetDefaults()
	return &self
}

func (t *Char) Clone() *Char {
	c := &Char{}
	c.Data = t.Data
	return c
}

func (t *Char) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Char) SetDefaults() {
	t.Data = 0
}

func (t *Char) GetTypeSupport() types.MessageTypeSupport {
	return CharTypeSupport
}

// CharPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type CharPublisher struct {
	*rclgo.Publisher
}

// NewCharPublisher creates and returns a new publisher for the
// Char
func NewCharPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*CharPublisher, error) {
	pub, err := node.NewPublisher(topic_name, CharTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CharPublisher{pub}, nil
}

func (p *CharPublisher) Publish(msg *Char) error {
	return p.Publisher.Publish(msg)
}

// CharSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type CharSubscription struct {
	*rclgo.Subscription
}

// CharSubscriptionCallback type is used to provide a subscription
// handler function for a CharSubscription.
type CharSubscriptionCallback func(msg *Char, info *rclgo.RmwMessageInfo, err error)

// NewCharSubscription creates and returns a new subscription for the
// Char
func NewCharSubscription(node *rclgo.Node, topic_name string, subscriptionCallback CharSubscriptionCallback) (*CharSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Char
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, CharTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &CharSubscription{sub}, nil
}

func (s *CharSubscription) TakeMessage(out *Char) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneCharSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCharSlice(dst, src []Char) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CharTypeSupport types.MessageTypeSupport = _CharTypeSupport{}

type _CharTypeSupport struct{}

func (t _CharTypeSupport) New() types.Message {
	return NewChar()
}

func (t _CharTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Char
	return (unsafe.Pointer)(C.example_interfaces__msg__Char__create())
}

func (t _CharTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Char__destroy((*C.example_interfaces__msg__Char)(pointer_to_free))
}

func (t _CharTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Char)
	mem := (*C.example_interfaces__msg__Char)(dst)
	mem.data = C.uchar(m.Data)
}

func (t _CharTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Char)
	mem := (*C.example_interfaces__msg__Char)(ros2_message_buffer)
	m.Data = byte(mem.data)
}

func (t _CharTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Char())
}

type CChar = C.example_interfaces__msg__Char
type CChar__Sequence = C.example_interfaces__msg__Char__Sequence

func Char__Sequence_to_Go(goSlice *[]Char, cSlice CChar__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Char, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		CharTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Char__Sequence_to_C(cSlice *CChar__Sequence, goSlice []Char) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Char)(C.malloc(C.sizeof_struct_example_interfaces__msg__Char * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		CharTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Char__Array_to_Go(goSlice []Char, cSlice []CChar) {
	for i := 0; i < len(cSlice); i++ {
		CharTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Char__Array_to_C(cSlice []CChar, goSlice []Char) {
	for i := 0; i < len(goSlice); i++ {
		CharTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
