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
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/color_rgba.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/ColorRGBA", ColorRGBATypeSupport)
	typemap.RegisterMessage("std_msgs/msg/ColorRGBA", ColorRGBATypeSupport)
}

// Do not create instances of this type directly. Always use NewColorRGBA
// function instead.
type ColorRGBA struct {
	R float32 `yaml:"r"`
	G float32 `yaml:"g"`
	B float32 `yaml:"b"`
	A float32 `yaml:"a"`
}

// NewColorRGBA creates a new ColorRGBA with default values.
func NewColorRGBA() *ColorRGBA {
	self := ColorRGBA{}
	self.SetDefaults()
	return &self
}

func (t *ColorRGBA) Clone() *ColorRGBA {
	c := &ColorRGBA{}
	c.R = t.R
	c.G = t.G
	c.B = t.B
	c.A = t.A
	return c
}

func (t *ColorRGBA) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ColorRGBA) SetDefaults() {
	t.R = 0
	t.G = 0
	t.B = 0
	t.A = 0
}

func (t *ColorRGBA) GetTypeSupport() types.MessageTypeSupport {
	return ColorRGBATypeSupport
}

// ColorRGBAPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ColorRGBAPublisher struct {
	*rclgo.Publisher
}

// NewColorRGBAPublisher creates and returns a new publisher for the
// ColorRGBA
func NewColorRGBAPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ColorRGBAPublisher, error) {
	pub, err := node.NewPublisher(topic_name, ColorRGBATypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ColorRGBAPublisher{pub}, nil
}

func (p *ColorRGBAPublisher) Publish(msg *ColorRGBA) error {
	return p.Publisher.Publish(msg)
}

// ColorRGBASubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ColorRGBASubscription struct {
	*rclgo.Subscription
}

// ColorRGBASubscriptionCallback type is used to provide a subscription
// handler function for a ColorRGBASubscription.
type ColorRGBASubscriptionCallback func(msg *ColorRGBA, info *rclgo.MessageInfo, err error)

// NewColorRGBASubscription creates and returns a new subscription for the
// ColorRGBA
func NewColorRGBASubscription(node *rclgo.Node, topic_name string, subscriptionCallback ColorRGBASubscriptionCallback) (*ColorRGBASubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg ColorRGBA
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ColorRGBATypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &ColorRGBASubscription{sub}, nil
}

func (s *ColorRGBASubscription) TakeMessage(out *ColorRGBA) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneColorRGBASlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneColorRGBASlice(dst, src []ColorRGBA) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ColorRGBATypeSupport types.MessageTypeSupport = _ColorRGBATypeSupport{}

type _ColorRGBATypeSupport struct{}

func (t _ColorRGBATypeSupport) New() types.Message {
	return NewColorRGBA()
}

func (t _ColorRGBATypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__ColorRGBA
	return (unsafe.Pointer)(C.std_msgs__msg__ColorRGBA__create())
}

func (t _ColorRGBATypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__ColorRGBA__destroy((*C.std_msgs__msg__ColorRGBA)(pointer_to_free))
}

func (t _ColorRGBATypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ColorRGBA)
	mem := (*C.std_msgs__msg__ColorRGBA)(dst)
	mem.r = C.float(m.R)
	mem.g = C.float(m.G)
	mem.b = C.float(m.B)
	mem.a = C.float(m.A)
}

func (t _ColorRGBATypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ColorRGBA)
	mem := (*C.std_msgs__msg__ColorRGBA)(ros2_message_buffer)
	m.R = float32(mem.r)
	m.G = float32(mem.g)
	m.B = float32(mem.b)
	m.A = float32(mem.a)
}

func (t _ColorRGBATypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__ColorRGBA())
}

type CColorRGBA = C.std_msgs__msg__ColorRGBA
type CColorRGBA__Sequence = C.std_msgs__msg__ColorRGBA__Sequence

func ColorRGBA__Sequence_to_Go(goSlice *[]ColorRGBA, cSlice CColorRGBA__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ColorRGBA, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ColorRGBATypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func ColorRGBA__Sequence_to_C(cSlice *CColorRGBA__Sequence, goSlice []ColorRGBA) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.std_msgs__msg__ColorRGBA)(C.malloc(C.sizeof_struct_std_msgs__msg__ColorRGBA * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ColorRGBATypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func ColorRGBA__Array_to_Go(goSlice []ColorRGBA, cSlice []CColorRGBA) {
	for i := 0; i < len(cSlice); i++ {
		ColorRGBATypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ColorRGBA__Array_to_C(cSlice []CColorRGBA, goSlice []ColorRGBA) {
	for i := 0; i < len(goSlice); i++ {
		ColorRGBATypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
