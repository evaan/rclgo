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
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/pose_with_covariance_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/PoseWithCovarianceStamped", PoseWithCovarianceStampedTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/PoseWithCovarianceStamped", PoseWithCovarianceStampedTypeSupport)
}

// Do not create instances of this type directly. Always use NewPoseWithCovarianceStamped
// function instead.
type PoseWithCovarianceStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Pose PoseWithCovariance `yaml:"pose"`
}

// NewPoseWithCovarianceStamped creates a new PoseWithCovarianceStamped with default values.
func NewPoseWithCovarianceStamped() *PoseWithCovarianceStamped {
	self := PoseWithCovarianceStamped{}
	self.SetDefaults()
	return &self
}

func (t *PoseWithCovarianceStamped) Clone() *PoseWithCovarianceStamped {
	c := &PoseWithCovarianceStamped{}
	c.Header = *t.Header.Clone()
	c.Pose = *t.Pose.Clone()
	return c
}

func (t *PoseWithCovarianceStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PoseWithCovarianceStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Pose.SetDefaults()
}

func (t *PoseWithCovarianceStamped) GetTypeSupport() types.MessageTypeSupport {
	return PoseWithCovarianceStampedTypeSupport
}

// PoseWithCovarianceStampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type PoseWithCovarianceStampedPublisher struct {
	*rclgo.Publisher
}

// NewPoseWithCovarianceStampedPublisher creates and returns a new publisher for the
// PoseWithCovarianceStamped
func NewPoseWithCovarianceStampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*PoseWithCovarianceStampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, PoseWithCovarianceStampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PoseWithCovarianceStampedPublisher{pub}, nil
}

func (p *PoseWithCovarianceStampedPublisher) Publish(msg *PoseWithCovarianceStamped) error {
	return p.Publisher.Publish(msg)
}

// PoseWithCovarianceStampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type PoseWithCovarianceStampedSubscription struct {
	*rclgo.Subscription
}

// PoseWithCovarianceStampedSubscriptionCallback type is used to provide a subscription
// handler function for a PoseWithCovarianceStampedSubscription.
type PoseWithCovarianceStampedSubscriptionCallback func(msg *PoseWithCovarianceStamped, info *rclgo.RmwMessageInfo, err error)

// NewPoseWithCovarianceStampedSubscription creates and returns a new subscription for the
// PoseWithCovarianceStamped
func NewPoseWithCovarianceStampedSubscription(node *rclgo.Node, topic_name string, subscriptionCallback PoseWithCovarianceStampedSubscriptionCallback) (*PoseWithCovarianceStampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg PoseWithCovarianceStamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, PoseWithCovarianceStampedTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &PoseWithCovarianceStampedSubscription{sub}, nil
}

func (s *PoseWithCovarianceStampedSubscription) TakeMessage(out *PoseWithCovarianceStamped) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePoseWithCovarianceStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePoseWithCovarianceStampedSlice(dst, src []PoseWithCovarianceStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PoseWithCovarianceStampedTypeSupport types.MessageTypeSupport = _PoseWithCovarianceStampedTypeSupport{}

type _PoseWithCovarianceStampedTypeSupport struct{}

func (t _PoseWithCovarianceStampedTypeSupport) New() types.Message {
	return NewPoseWithCovarianceStamped()
}

func (t _PoseWithCovarianceStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__PoseWithCovarianceStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__PoseWithCovarianceStamped__create())
}

func (t _PoseWithCovarianceStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__PoseWithCovarianceStamped__destroy((*C.geometry_msgs__msg__PoseWithCovarianceStamped)(pointer_to_free))
}

func (t _PoseWithCovarianceStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PoseWithCovarianceStamped)
	mem := (*C.geometry_msgs__msg__PoseWithCovarianceStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	PoseWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(&mem.pose), &m.Pose)
}

func (t _PoseWithCovarianceStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PoseWithCovarianceStamped)
	mem := (*C.geometry_msgs__msg__PoseWithCovarianceStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	PoseWithCovarianceTypeSupport.AsGoStruct(&m.Pose, unsafe.Pointer(&mem.pose))
}

func (t _PoseWithCovarianceStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__PoseWithCovarianceStamped())
}

type CPoseWithCovarianceStamped = C.geometry_msgs__msg__PoseWithCovarianceStamped
type CPoseWithCovarianceStamped__Sequence = C.geometry_msgs__msg__PoseWithCovarianceStamped__Sequence

func PoseWithCovarianceStamped__Sequence_to_Go(goSlice *[]PoseWithCovarianceStamped, cSlice CPoseWithCovarianceStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PoseWithCovarianceStamped, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		PoseWithCovarianceStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func PoseWithCovarianceStamped__Sequence_to_C(cSlice *CPoseWithCovarianceStamped__Sequence, goSlice []PoseWithCovarianceStamped) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__PoseWithCovarianceStamped)(C.malloc(C.sizeof_struct_geometry_msgs__msg__PoseWithCovarianceStamped * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		PoseWithCovarianceStampedTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func PoseWithCovarianceStamped__Array_to_Go(goSlice []PoseWithCovarianceStamped, cSlice []CPoseWithCovarianceStamped) {
	for i := 0; i < len(cSlice); i++ {
		PoseWithCovarianceStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PoseWithCovarianceStamped__Array_to_C(cSlice []CPoseWithCovarianceStamped, goSlice []PoseWithCovarianceStamped) {
	for i := 0; i < len(goSlice); i++ {
		PoseWithCovarianceStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
