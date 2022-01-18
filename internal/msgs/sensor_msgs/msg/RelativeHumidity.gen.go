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
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/relative_humidity.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/RelativeHumidity", RelativeHumidityTypeSupport)
	typemap.RegisterMessage("sensor_msgs/msg/RelativeHumidity", RelativeHumidityTypeSupport)
}

// Do not create instances of this type directly. Always use NewRelativeHumidity
// function instead.
type RelativeHumidity struct {
	Header std_msgs_msg.Header `yaml:"header"`// timestamp of the measurement
	RelativeHumidity float64 `yaml:"relative_humidity"`// Expression of the relative humidity
	Variance float64 `yaml:"variance"`// 0 is interpreted as variance unknown
}

// NewRelativeHumidity creates a new RelativeHumidity with default values.
func NewRelativeHumidity() *RelativeHumidity {
	self := RelativeHumidity{}
	self.SetDefaults()
	return &self
}

func (t *RelativeHumidity) Clone() *RelativeHumidity {
	c := &RelativeHumidity{}
	c.Header = *t.Header.Clone()
	c.RelativeHumidity = t.RelativeHumidity
	c.Variance = t.Variance
	return c
}

func (t *RelativeHumidity) CloneMsg() types.Message {
	return t.Clone()
}

func (t *RelativeHumidity) SetDefaults() {
	t.Header.SetDefaults()
	t.RelativeHumidity = 0
	t.Variance = 0
}

func (t *RelativeHumidity) GetTypeSupport() types.MessageTypeSupport {
	return RelativeHumidityTypeSupport
}

// RelativeHumidityPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type RelativeHumidityPublisher struct {
	*rclgo.Publisher
}

// NewRelativeHumidityPublisher creates and returns a new publisher for the
// RelativeHumidity
func NewRelativeHumidityPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*RelativeHumidityPublisher, error) {
	pub, err := node.NewPublisher(topic_name, RelativeHumidityTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &RelativeHumidityPublisher{pub}, nil
}

func (p *RelativeHumidityPublisher) Publish(msg *RelativeHumidity) error {
	return p.Publisher.Publish(msg)
}

// RelativeHumiditySubscription wraps rclgo.Subscription to provide type safe helper
// functions
type RelativeHumiditySubscription struct {
	*rclgo.Subscription
}

// RelativeHumiditySubscriptionCallback type is used to provide a subscription
// handler function for a RelativeHumiditySubscription.
type RelativeHumiditySubscriptionCallback func(msg *RelativeHumidity, info *rclgo.RmwMessageInfo, err error)

// NewRelativeHumiditySubscription creates and returns a new subscription for the
// RelativeHumidity
func NewRelativeHumiditySubscription(node *rclgo.Node, topic_name string, subscriptionCallback RelativeHumiditySubscriptionCallback) (*RelativeHumiditySubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg RelativeHumidity
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, RelativeHumidityTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &RelativeHumiditySubscription{sub}, nil
}

func (s *RelativeHumiditySubscription) TakeMessage(out *RelativeHumidity) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneRelativeHumiditySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneRelativeHumiditySlice(dst, src []RelativeHumidity) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var RelativeHumidityTypeSupport types.MessageTypeSupport = _RelativeHumidityTypeSupport{}

type _RelativeHumidityTypeSupport struct{}

func (t _RelativeHumidityTypeSupport) New() types.Message {
	return NewRelativeHumidity()
}

func (t _RelativeHumidityTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__RelativeHumidity
	return (unsafe.Pointer)(C.sensor_msgs__msg__RelativeHumidity__create())
}

func (t _RelativeHumidityTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__RelativeHumidity__destroy((*C.sensor_msgs__msg__RelativeHumidity)(pointer_to_free))
}

func (t _RelativeHumidityTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*RelativeHumidity)
	mem := (*C.sensor_msgs__msg__RelativeHumidity)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.relative_humidity = C.double(m.RelativeHumidity)
	mem.variance = C.double(m.Variance)
}

func (t _RelativeHumidityTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*RelativeHumidity)
	mem := (*C.sensor_msgs__msg__RelativeHumidity)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.RelativeHumidity = float64(mem.relative_humidity)
	m.Variance = float64(mem.variance)
}

func (t _RelativeHumidityTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__RelativeHumidity())
}

type CRelativeHumidity = C.sensor_msgs__msg__RelativeHumidity
type CRelativeHumidity__Sequence = C.sensor_msgs__msg__RelativeHumidity__Sequence

func RelativeHumidity__Sequence_to_Go(goSlice *[]RelativeHumidity, cSlice CRelativeHumidity__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RelativeHumidity, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		RelativeHumidityTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func RelativeHumidity__Sequence_to_C(cSlice *CRelativeHumidity__Sequence, goSlice []RelativeHumidity) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__RelativeHumidity)(C.malloc(C.sizeof_struct_sensor_msgs__msg__RelativeHumidity * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		RelativeHumidityTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func RelativeHumidity__Array_to_Go(goSlice []RelativeHumidity, cSlice []CRelativeHumidity) {
	for i := 0; i < len(cSlice); i++ {
		RelativeHumidityTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func RelativeHumidity__Array_to_C(cSlice []CRelativeHumidity, goSlice []RelativeHumidity) {
	for i := 0; i < len(goSlice); i++ {
		RelativeHumidityTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
