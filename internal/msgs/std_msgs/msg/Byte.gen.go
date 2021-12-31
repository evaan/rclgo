/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package std_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/byte.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/Byte", ByteTypeSupport)
}

// Do not create instances of this type directly. Always use NewByte
// function instead.
type Byte struct {
	Data byte `yaml:"data"`
}

// NewByte creates a new Byte with default values.
func NewByte() *Byte {
	self := Byte{}
	self.SetDefaults()
	return &self
}

func (t *Byte) Clone() *Byte {
	c := &Byte{}
	c.Data = t.Data
	return c
}

func (t *Byte) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Byte) SetDefaults() {
	t.Data = 0
}

// BytePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type BytePublisher struct {
	*rclgo.Publisher
}

// NewBytePublisher creates and returns a new publisher for the
// Byte
func NewBytePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*BytePublisher, error) {
	pub, err := node.NewPublisher(topic_name, ByteTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &BytePublisher{pub}, nil
}

func (p *BytePublisher) Publish(msg *Byte) error {
	return p.Publisher.Publish(msg)
}

// ByteSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ByteSubscription struct {
	*rclgo.Subscription
}

// ByteSubscriptionCallback type is used to provide a subscription
// handler function for a ByteSubscription.
type ByteSubscriptionCallback func(msg *Byte, info *rclgo.RmwMessageInfo, err error)

// NewByteSubscription creates and returns a new subscription for the
// Byte
func NewByteSubscription(node *rclgo.Node, topic_name string, subscriptionCallback ByteSubscriptionCallback) (*ByteSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Byte
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ByteTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &ByteSubscription{sub}, nil
}

func (s *ByteSubscription) TakeMessage(out *Byte) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneByteSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneByteSlice(dst, src []Byte) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ByteTypeSupport types.MessageTypeSupport = _ByteTypeSupport{}

type _ByteTypeSupport struct{}

func (t _ByteTypeSupport) New() types.Message {
	return NewByte()
}

func (t _ByteTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__Byte
	return (unsafe.Pointer)(C.std_msgs__msg__Byte__create())
}

func (t _ByteTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__Byte__destroy((*C.std_msgs__msg__Byte)(pointer_to_free))
}

func (t _ByteTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Byte)
	mem := (*C.std_msgs__msg__Byte)(dst)
	mem.data = C.uint8_t(m.Data)
}

func (t _ByteTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Byte)
	mem := (*C.std_msgs__msg__Byte)(ros2_message_buffer)
	m.Data = byte(mem.data)
}

func (t _ByteTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__Byte())
}

type CByte = C.std_msgs__msg__Byte
type CByte__Sequence = C.std_msgs__msg__Byte__Sequence

func Byte__Sequence_to_Go(goSlice *[]Byte, cSlice CByte__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Byte, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__Byte__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Byte * uintptr(i)),
		))
		ByteTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Byte__Sequence_to_C(cSlice *CByte__Sequence, goSlice []Byte) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__Byte)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__Byte * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__Byte)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Byte * uintptr(i)),
		))
		ByteTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Byte__Array_to_Go(goSlice []Byte, cSlice []CByte) {
	for i := 0; i < len(cSlice); i++ {
		ByteTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Byte__Array_to_C(cSlice []CByte, goSlice []Byte) {
	for i := 0; i < len(goSlice); i++ {
		ByteTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
