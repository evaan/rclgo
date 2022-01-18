/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	geometry_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/geometry_msgs/msg"
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/point_cloud.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/PointCloud", PointCloudTypeSupport)
	typemap.RegisterMessage("sensor_msgs/msg/PointCloud", PointCloudTypeSupport)
}

// Do not create instances of this type directly. Always use NewPointCloud
// function instead.
type PointCloud struct {
	Header std_msgs_msg.Header `yaml:"header"`// Time of sensor data acquisition, coordinate frame ID.
	Points []geometry_msgs_msg.Point32 `yaml:"points"`// Array of 3d points. Each Point32 should be interpreted as a 3d pointin the frame given in the header.
	Channels []ChannelFloat32 `yaml:"channels"`// Each channel should have the same number of elements as points array,and the data in each channel should correspond 1:1 with each point.Channel names in common practice are listed in ChannelFloat32.msg.
}

// NewPointCloud creates a new PointCloud with default values.
func NewPointCloud() *PointCloud {
	self := PointCloud{}
	self.SetDefaults()
	return &self
}

func (t *PointCloud) Clone() *PointCloud {
	c := &PointCloud{}
	c.Header = *t.Header.Clone()
	if t.Points != nil {
		c.Points = make([]geometry_msgs_msg.Point32, len(t.Points))
		geometry_msgs_msg.ClonePoint32Slice(c.Points, t.Points)
	}
	if t.Channels != nil {
		c.Channels = make([]ChannelFloat32, len(t.Channels))
		CloneChannelFloat32Slice(c.Channels, t.Channels)
	}
	return c
}

func (t *PointCloud) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PointCloud) SetDefaults() {
	t.Header.SetDefaults()
	t.Points = nil
	t.Channels = nil
}

func (t *PointCloud) GetTypeSupport() types.MessageTypeSupport {
	return PointCloudTypeSupport
}

// PointCloudPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type PointCloudPublisher struct {
	*rclgo.Publisher
}

// NewPointCloudPublisher creates and returns a new publisher for the
// PointCloud
func NewPointCloudPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*PointCloudPublisher, error) {
	pub, err := node.NewPublisher(topic_name, PointCloudTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PointCloudPublisher{pub}, nil
}

func (p *PointCloudPublisher) Publish(msg *PointCloud) error {
	return p.Publisher.Publish(msg)
}

// PointCloudSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type PointCloudSubscription struct {
	*rclgo.Subscription
}

// PointCloudSubscriptionCallback type is used to provide a subscription
// handler function for a PointCloudSubscription.
type PointCloudSubscriptionCallback func(msg *PointCloud, info *rclgo.RmwMessageInfo, err error)

// NewPointCloudSubscription creates and returns a new subscription for the
// PointCloud
func NewPointCloudSubscription(node *rclgo.Node, topic_name string, subscriptionCallback PointCloudSubscriptionCallback) (*PointCloudSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg PointCloud
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, PointCloudTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &PointCloudSubscription{sub}, nil
}

func (s *PointCloudSubscription) TakeMessage(out *PointCloud) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePointCloudSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePointCloudSlice(dst, src []PointCloud) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PointCloudTypeSupport types.MessageTypeSupport = _PointCloudTypeSupport{}

type _PointCloudTypeSupport struct{}

func (t _PointCloudTypeSupport) New() types.Message {
	return NewPointCloud()
}

func (t _PointCloudTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__PointCloud
	return (unsafe.Pointer)(C.sensor_msgs__msg__PointCloud__create())
}

func (t _PointCloudTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__PointCloud__destroy((*C.sensor_msgs__msg__PointCloud)(pointer_to_free))
}

func (t _PointCloudTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PointCloud)
	mem := (*C.sensor_msgs__msg__PointCloud)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	geometry_msgs_msg.Point32__Sequence_to_C((*geometry_msgs_msg.CPoint32__Sequence)(unsafe.Pointer(&mem.points)), m.Points)
	ChannelFloat32__Sequence_to_C(&mem.channels, m.Channels)
}

func (t _PointCloudTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PointCloud)
	mem := (*C.sensor_msgs__msg__PointCloud)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	geometry_msgs_msg.Point32__Sequence_to_Go(&m.Points, *(*geometry_msgs_msg.CPoint32__Sequence)(unsafe.Pointer(&mem.points)))
	ChannelFloat32__Sequence_to_Go(&m.Channels, mem.channels)
}

func (t _PointCloudTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__PointCloud())
}

type CPointCloud = C.sensor_msgs__msg__PointCloud
type CPointCloud__Sequence = C.sensor_msgs__msg__PointCloud__Sequence

func PointCloud__Sequence_to_Go(goSlice *[]PointCloud, cSlice CPointCloud__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PointCloud, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		PointCloudTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func PointCloud__Sequence_to_C(cSlice *CPointCloud__Sequence, goSlice []PointCloud) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__PointCloud)(C.malloc(C.sizeof_struct_sensor_msgs__msg__PointCloud * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		PointCloudTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func PointCloud__Array_to_Go(goSlice []PointCloud, cSlice []CPointCloud) {
	for i := 0; i < len(cSlice); i++ {
		PointCloudTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PointCloud__Array_to_C(cSlice []CPointCloud, goSlice []PointCloud) {
	for i := 0; i < len(goSlice); i++ {
		PointCloudTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
