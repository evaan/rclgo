/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_action
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	unique_identifier_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/unique_identifier_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/action/nested_message.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/NestedMessage_FeedbackMessage", NestedMessage_FeedbackMessageTypeSupport)
	typemap.RegisterMessage("test_msgs/action/NestedMessage_FeedbackMessage", NestedMessage_FeedbackMessageTypeSupport)
}

// Do not create instances of this type directly. Always use NewNestedMessage_FeedbackMessage
// function instead.
type NestedMessage_FeedbackMessage struct {
	GoalID unique_identifier_msgs_msg.UUID `yaml:"goal_id"`
	Feedback NestedMessage_Feedback `yaml:"feedback"`
}

// NewNestedMessage_FeedbackMessage creates a new NestedMessage_FeedbackMessage with default values.
func NewNestedMessage_FeedbackMessage() *NestedMessage_FeedbackMessage {
	self := NestedMessage_FeedbackMessage{}
	self.SetDefaults()
	return &self
}

func (t *NestedMessage_FeedbackMessage) Clone() *NestedMessage_FeedbackMessage {
	c := &NestedMessage_FeedbackMessage{}
	c.GoalID = *t.GoalID.Clone()
	c.Feedback = *t.Feedback.Clone()
	return c
}

func (t *NestedMessage_FeedbackMessage) CloneMsg() types.Message {
	return t.Clone()
}

func (t *NestedMessage_FeedbackMessage) SetDefaults() {
	t.GoalID.SetDefaults()
	t.Feedback.SetDefaults()
}

func (t *NestedMessage_FeedbackMessage) GetTypeSupport() types.MessageTypeSupport {
	return NestedMessage_FeedbackMessageTypeSupport
}
func (t *NestedMessage_FeedbackMessage) GetGoalID() *types.GoalID {
	return (*types.GoalID)(&t.GoalID.Uuid)
}

func (t *NestedMessage_FeedbackMessage) SetGoalID(id *types.GoalID) {
	t.GoalID.Uuid = *id
}

// NestedMessage_FeedbackMessagePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type NestedMessage_FeedbackMessagePublisher struct {
	*rclgo.Publisher
}

// NewNestedMessage_FeedbackMessagePublisher creates and returns a new publisher for the
// NestedMessage_FeedbackMessage
func NewNestedMessage_FeedbackMessagePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*NestedMessage_FeedbackMessagePublisher, error) {
	pub, err := node.NewPublisher(topic_name, NestedMessage_FeedbackMessageTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_FeedbackMessagePublisher{pub}, nil
}

func (p *NestedMessage_FeedbackMessagePublisher) Publish(msg *NestedMessage_FeedbackMessage) error {
	return p.Publisher.Publish(msg)
}

// NestedMessage_FeedbackMessageSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type NestedMessage_FeedbackMessageSubscription struct {
	*rclgo.Subscription
}

// NestedMessage_FeedbackMessageSubscriptionCallback type is used to provide a subscription
// handler function for a NestedMessage_FeedbackMessageSubscription.
type NestedMessage_FeedbackMessageSubscriptionCallback func(msg *NestedMessage_FeedbackMessage, info *rclgo.RmwMessageInfo, err error)

// NewNestedMessage_FeedbackMessageSubscription creates and returns a new subscription for the
// NestedMessage_FeedbackMessage
func NewNestedMessage_FeedbackMessageSubscription(node *rclgo.Node, topic_name string, subscriptionCallback NestedMessage_FeedbackMessageSubscriptionCallback) (*NestedMessage_FeedbackMessageSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg NestedMessage_FeedbackMessage
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, NestedMessage_FeedbackMessageTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_FeedbackMessageSubscription{sub}, nil
}

func (s *NestedMessage_FeedbackMessageSubscription) TakeMessage(out *NestedMessage_FeedbackMessage) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneNestedMessage_FeedbackMessageSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneNestedMessage_FeedbackMessageSlice(dst, src []NestedMessage_FeedbackMessage) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var NestedMessage_FeedbackMessageTypeSupport types.MessageTypeSupport = _NestedMessage_FeedbackMessageTypeSupport{}

type _NestedMessage_FeedbackMessageTypeSupport struct{}

func (t _NestedMessage_FeedbackMessageTypeSupport) New() types.Message {
	return NewNestedMessage_FeedbackMessage()
}

func (t _NestedMessage_FeedbackMessageTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__action__NestedMessage_FeedbackMessage
	return (unsafe.Pointer)(C.test_msgs__action__NestedMessage_FeedbackMessage__create())
}

func (t _NestedMessage_FeedbackMessageTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__action__NestedMessage_FeedbackMessage__destroy((*C.test_msgs__action__NestedMessage_FeedbackMessage)(pointer_to_free))
}

func (t _NestedMessage_FeedbackMessageTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*NestedMessage_FeedbackMessage)
	mem := (*C.test_msgs__action__NestedMessage_FeedbackMessage)(dst)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal_id), &m.GoalID)
	NestedMessage_FeedbackTypeSupport.AsCStruct(unsafe.Pointer(&mem.feedback), &m.Feedback)
}

func (t _NestedMessage_FeedbackMessageTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*NestedMessage_FeedbackMessage)
	mem := (*C.test_msgs__action__NestedMessage_FeedbackMessage)(ros2_message_buffer)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsGoStruct(&m.GoalID, unsafe.Pointer(&mem.goal_id))
	NestedMessage_FeedbackTypeSupport.AsGoStruct(&m.Feedback, unsafe.Pointer(&mem.feedback))
}

func (t _NestedMessage_FeedbackMessageTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__action__NestedMessage_FeedbackMessage())
}

type CNestedMessage_FeedbackMessage = C.test_msgs__action__NestedMessage_FeedbackMessage
type CNestedMessage_FeedbackMessage__Sequence = C.test_msgs__action__NestedMessage_FeedbackMessage__Sequence

func NestedMessage_FeedbackMessage__Sequence_to_Go(goSlice *[]NestedMessage_FeedbackMessage, cSlice CNestedMessage_FeedbackMessage__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]NestedMessage_FeedbackMessage, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		NestedMessage_FeedbackMessageTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func NestedMessage_FeedbackMessage__Sequence_to_C(cSlice *CNestedMessage_FeedbackMessage__Sequence, goSlice []NestedMessage_FeedbackMessage) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__action__NestedMessage_FeedbackMessage)(C.malloc(C.sizeof_struct_test_msgs__action__NestedMessage_FeedbackMessage * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		NestedMessage_FeedbackMessageTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func NestedMessage_FeedbackMessage__Array_to_Go(goSlice []NestedMessage_FeedbackMessage, cSlice []CNestedMessage_FeedbackMessage) {
	for i := 0; i < len(cSlice); i++ {
		NestedMessage_FeedbackMessageTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func NestedMessage_FeedbackMessage__Array_to_C(cSlice []CNestedMessage_FeedbackMessage, goSlice []NestedMessage_FeedbackMessage) {
	for i := 0; i < len(goSlice); i++ {
		NestedMessage_FeedbackMessageTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
