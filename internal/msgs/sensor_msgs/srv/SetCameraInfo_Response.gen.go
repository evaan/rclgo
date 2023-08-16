/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package sensor_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/srv/set_camera_info.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/SetCameraInfo_Response", SetCameraInfo_ResponseTypeSupport)
	typemap.RegisterMessage("sensor_msgs/srv/SetCameraInfo_Response", SetCameraInfo_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetCameraInfo_Response
// function instead.
type SetCameraInfo_Response struct {
	Success bool `yaml:"success"`// True if the call succeeded
	StatusMessage string `yaml:"status_message"`// Used to give details about success
}

// NewSetCameraInfo_Response creates a new SetCameraInfo_Response with default values.
func NewSetCameraInfo_Response() *SetCameraInfo_Response {
	self := SetCameraInfo_Response{}
	self.SetDefaults()
	return &self
}

func (t *SetCameraInfo_Response) Clone() *SetCameraInfo_Response {
	c := &SetCameraInfo_Response{}
	c.Success = t.Success
	c.StatusMessage = t.StatusMessage
	return c
}

func (t *SetCameraInfo_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetCameraInfo_Response) SetDefaults() {
	t.Success = false
	t.StatusMessage = ""
}

func (t *SetCameraInfo_Response) GetTypeSupport() types.MessageTypeSupport {
	return SetCameraInfo_ResponseTypeSupport
}

// SetCameraInfo_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type SetCameraInfo_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewSetCameraInfo_ResponsePublisher creates and returns a new publisher for the
// SetCameraInfo_Response
func NewSetCameraInfo_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*SetCameraInfo_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, SetCameraInfo_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SetCameraInfo_ResponsePublisher{pub}, nil
}

func (p *SetCameraInfo_ResponsePublisher) Publish(msg *SetCameraInfo_Response) error {
	return p.Publisher.Publish(msg)
}

// SetCameraInfo_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type SetCameraInfo_ResponseSubscription struct {
	*rclgo.Subscription
}

// SetCameraInfo_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a SetCameraInfo_ResponseSubscription.
type SetCameraInfo_ResponseSubscriptionCallback func(msg *SetCameraInfo_Response, info *rclgo.MessageInfo, err error)

// NewSetCameraInfo_ResponseSubscription creates and returns a new subscription for the
// SetCameraInfo_Response
func NewSetCameraInfo_ResponseSubscription(node *rclgo.Node, topic_name string, subscriptionCallback SetCameraInfo_ResponseSubscriptionCallback) (*SetCameraInfo_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg SetCameraInfo_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, SetCameraInfo_ResponseTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &SetCameraInfo_ResponseSubscription{sub}, nil
}

func (s *SetCameraInfo_ResponseSubscription) TakeMessage(out *SetCameraInfo_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneSetCameraInfo_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetCameraInfo_ResponseSlice(dst, src []SetCameraInfo_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetCameraInfo_ResponseTypeSupport types.MessageTypeSupport = _SetCameraInfo_ResponseTypeSupport{}

type _SetCameraInfo_ResponseTypeSupport struct{}

func (t _SetCameraInfo_ResponseTypeSupport) New() types.Message {
	return NewSetCameraInfo_Response()
}

func (t _SetCameraInfo_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__srv__SetCameraInfo_Response
	return (unsafe.Pointer)(C.sensor_msgs__srv__SetCameraInfo_Response__create())
}

func (t _SetCameraInfo_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__srv__SetCameraInfo_Response__destroy((*C.sensor_msgs__srv__SetCameraInfo_Response)(pointer_to_free))
}

func (t _SetCameraInfo_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetCameraInfo_Response)
	mem := (*C.sensor_msgs__srv__SetCameraInfo_Response)(dst)
	mem.success = C.bool(m.Success)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.status_message), m.StatusMessage)
}

func (t _SetCameraInfo_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetCameraInfo_Response)
	mem := (*C.sensor_msgs__srv__SetCameraInfo_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	primitives.StringAsGoStruct(&m.StatusMessage, unsafe.Pointer(&mem.status_message))
}

func (t _SetCameraInfo_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__srv__SetCameraInfo_Response())
}

type CSetCameraInfo_Response = C.sensor_msgs__srv__SetCameraInfo_Response
type CSetCameraInfo_Response__Sequence = C.sensor_msgs__srv__SetCameraInfo_Response__Sequence

func SetCameraInfo_Response__Sequence_to_Go(goSlice *[]SetCameraInfo_Response, cSlice CSetCameraInfo_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetCameraInfo_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		SetCameraInfo_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func SetCameraInfo_Response__Sequence_to_C(cSlice *CSetCameraInfo_Response__Sequence, goSlice []SetCameraInfo_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.sensor_msgs__srv__SetCameraInfo_Response)(C.malloc(C.sizeof_struct_sensor_msgs__srv__SetCameraInfo_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		SetCameraInfo_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func SetCameraInfo_Response__Array_to_Go(goSlice []SetCameraInfo_Response, cSlice []CSetCameraInfo_Response) {
	for i := 0; i < len(cSlice); i++ {
		SetCameraInfo_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetCameraInfo_Response__Array_to_C(cSlice []CSetCameraInfo_Response, goSlice []SetCameraInfo_Response) {
	for i := 0; i < len(goSlice); i++ {
		SetCameraInfo_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
