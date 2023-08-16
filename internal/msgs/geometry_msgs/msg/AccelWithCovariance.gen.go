/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/accel_with_covariance.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/AccelWithCovariance", AccelWithCovarianceTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/AccelWithCovariance", AccelWithCovarianceTypeSupport)
}

// Do not create instances of this type directly. Always use NewAccelWithCovariance
// function instead.
type AccelWithCovariance struct {
	Accel Accel `yaml:"accel"`
	Covariance [36]float64 `yaml:"covariance"`// Row-major representation of the 6x6 covariance matrixThe orientation parameters use a fixed-axis representation.In order, the parameters are:(x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axis)
}

// NewAccelWithCovariance creates a new AccelWithCovariance with default values.
func NewAccelWithCovariance() *AccelWithCovariance {
	self := AccelWithCovariance{}
	self.SetDefaults()
	return &self
}

func (t *AccelWithCovariance) Clone() *AccelWithCovariance {
	c := &AccelWithCovariance{}
	c.Accel = *t.Accel.Clone()
	c.Covariance = t.Covariance
	return c
}

func (t *AccelWithCovariance) CloneMsg() types.Message {
	return t.Clone()
}

func (t *AccelWithCovariance) SetDefaults() {
	t.Accel.SetDefaults()
	t.Covariance = [36]float64{}
}

func (t *AccelWithCovariance) GetTypeSupport() types.MessageTypeSupport {
	return AccelWithCovarianceTypeSupport
}

// AccelWithCovariancePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type AccelWithCovariancePublisher struct {
	*rclgo.Publisher
}

// NewAccelWithCovariancePublisher creates and returns a new publisher for the
// AccelWithCovariance
func NewAccelWithCovariancePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*AccelWithCovariancePublisher, error) {
	pub, err := node.NewPublisher(topic_name, AccelWithCovarianceTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &AccelWithCovariancePublisher{pub}, nil
}

func (p *AccelWithCovariancePublisher) Publish(msg *AccelWithCovariance) error {
	return p.Publisher.Publish(msg)
}

// AccelWithCovarianceSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type AccelWithCovarianceSubscription struct {
	*rclgo.Subscription
}

// AccelWithCovarianceSubscriptionCallback type is used to provide a subscription
// handler function for a AccelWithCovarianceSubscription.
type AccelWithCovarianceSubscriptionCallback func(msg *AccelWithCovariance, info *rclgo.MessageInfo, err error)

// NewAccelWithCovarianceSubscription creates and returns a new subscription for the
// AccelWithCovariance
func NewAccelWithCovarianceSubscription(node *rclgo.Node, topic_name string, subscriptionCallback AccelWithCovarianceSubscriptionCallback) (*AccelWithCovarianceSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg AccelWithCovariance
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, AccelWithCovarianceTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &AccelWithCovarianceSubscription{sub}, nil
}

func (s *AccelWithCovarianceSubscription) TakeMessage(out *AccelWithCovariance) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneAccelWithCovarianceSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneAccelWithCovarianceSlice(dst, src []AccelWithCovariance) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var AccelWithCovarianceTypeSupport types.MessageTypeSupport = _AccelWithCovarianceTypeSupport{}

type _AccelWithCovarianceTypeSupport struct{}

func (t _AccelWithCovarianceTypeSupport) New() types.Message {
	return NewAccelWithCovariance()
}

func (t _AccelWithCovarianceTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__AccelWithCovariance
	return (unsafe.Pointer)(C.geometry_msgs__msg__AccelWithCovariance__create())
}

func (t _AccelWithCovarianceTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__AccelWithCovariance__destroy((*C.geometry_msgs__msg__AccelWithCovariance)(pointer_to_free))
}

func (t _AccelWithCovarianceTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*AccelWithCovariance)
	mem := (*C.geometry_msgs__msg__AccelWithCovariance)(dst)
	AccelTypeSupport.AsCStruct(unsafe.Pointer(&mem.accel), &m.Accel)
	cSlice_covariance := mem.covariance[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_covariance)), m.Covariance[:])
}

func (t _AccelWithCovarianceTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*AccelWithCovariance)
	mem := (*C.geometry_msgs__msg__AccelWithCovariance)(ros2_message_buffer)
	AccelTypeSupport.AsGoStruct(&m.Accel, unsafe.Pointer(&mem.accel))
	cSlice_covariance := mem.covariance[:]
	primitives.Float64__Array_to_Go(m.Covariance[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_covariance)))
}

func (t _AccelWithCovarianceTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__AccelWithCovariance())
}

type CAccelWithCovariance = C.geometry_msgs__msg__AccelWithCovariance
type CAccelWithCovariance__Sequence = C.geometry_msgs__msg__AccelWithCovariance__Sequence

func AccelWithCovariance__Sequence_to_Go(goSlice *[]AccelWithCovariance, cSlice CAccelWithCovariance__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]AccelWithCovariance, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		AccelWithCovarianceTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func AccelWithCovariance__Sequence_to_C(cSlice *CAccelWithCovariance__Sequence, goSlice []AccelWithCovariance) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__AccelWithCovariance)(C.malloc(C.sizeof_struct_geometry_msgs__msg__AccelWithCovariance * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		AccelWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func AccelWithCovariance__Array_to_Go(goSlice []AccelWithCovariance, cSlice []CAccelWithCovariance) {
	for i := 0; i < len(cSlice); i++ {
		AccelWithCovarianceTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func AccelWithCovariance__Array_to_C(cSlice []CAccelWithCovariance, goSlice []AccelWithCovariance) {
	for i := 0; i < len(goSlice); i++ {
		AccelWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
