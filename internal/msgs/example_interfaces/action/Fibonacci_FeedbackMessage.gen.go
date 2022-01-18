/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package example_interfaces_action
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	unique_identifier_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/unique_identifier_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/action/fibonacci.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Fibonacci_FeedbackMessage", Fibonacci_FeedbackMessageTypeSupport)
	typemap.RegisterMessage("example_interfaces/action/Fibonacci_FeedbackMessage", Fibonacci_FeedbackMessageTypeSupport)
}

// Do not create instances of this type directly. Always use NewFibonacci_FeedbackMessage
// function instead.
type Fibonacci_FeedbackMessage struct {
	GoalID unique_identifier_msgs_msg.UUID `yaml:"goal_id"`
	Feedback Fibonacci_Feedback `yaml:"feedback"`
}

// NewFibonacci_FeedbackMessage creates a new Fibonacci_FeedbackMessage with default values.
func NewFibonacci_FeedbackMessage() *Fibonacci_FeedbackMessage {
	self := Fibonacci_FeedbackMessage{}
	self.SetDefaults()
	return &self
}

func (t *Fibonacci_FeedbackMessage) Clone() *Fibonacci_FeedbackMessage {
	c := &Fibonacci_FeedbackMessage{}
	c.GoalID = *t.GoalID.Clone()
	c.Feedback = *t.Feedback.Clone()
	return c
}

func (t *Fibonacci_FeedbackMessage) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Fibonacci_FeedbackMessage) SetDefaults() {
	t.GoalID.SetDefaults()
	t.Feedback.SetDefaults()
}

func (t *Fibonacci_FeedbackMessage) GetTypeSupport() types.MessageTypeSupport {
	return Fibonacci_FeedbackMessageTypeSupport
}
func (t *Fibonacci_FeedbackMessage) GetGoalID() *types.GoalID {
	return (*types.GoalID)(&t.GoalID.Uuid)
}

func (t *Fibonacci_FeedbackMessage) SetGoalID(id *types.GoalID) {
	t.GoalID.Uuid = *id
}

// Fibonacci_FeedbackMessagePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Fibonacci_FeedbackMessagePublisher struct {
	*rclgo.Publisher
}

// NewFibonacci_FeedbackMessagePublisher creates and returns a new publisher for the
// Fibonacci_FeedbackMessage
func NewFibonacci_FeedbackMessagePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Fibonacci_FeedbackMessagePublisher, error) {
	pub, err := node.NewPublisher(topic_name, Fibonacci_FeedbackMessageTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_FeedbackMessagePublisher{pub}, nil
}

func (p *Fibonacci_FeedbackMessagePublisher) Publish(msg *Fibonacci_FeedbackMessage) error {
	return p.Publisher.Publish(msg)
}

// Fibonacci_FeedbackMessageSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Fibonacci_FeedbackMessageSubscription struct {
	*rclgo.Subscription
}

// Fibonacci_FeedbackMessageSubscriptionCallback type is used to provide a subscription
// handler function for a Fibonacci_FeedbackMessageSubscription.
type Fibonacci_FeedbackMessageSubscriptionCallback func(msg *Fibonacci_FeedbackMessage, info *rclgo.RmwMessageInfo, err error)

// NewFibonacci_FeedbackMessageSubscription creates and returns a new subscription for the
// Fibonacci_FeedbackMessage
func NewFibonacci_FeedbackMessageSubscription(node *rclgo.Node, topic_name string, subscriptionCallback Fibonacci_FeedbackMessageSubscriptionCallback) (*Fibonacci_FeedbackMessageSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Fibonacci_FeedbackMessage
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Fibonacci_FeedbackMessageTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_FeedbackMessageSubscription{sub}, nil
}

func (s *Fibonacci_FeedbackMessageSubscription) TakeMessage(out *Fibonacci_FeedbackMessage) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFibonacci_FeedbackMessageSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFibonacci_FeedbackMessageSlice(dst, src []Fibonacci_FeedbackMessage) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Fibonacci_FeedbackMessageTypeSupport types.MessageTypeSupport = _Fibonacci_FeedbackMessageTypeSupport{}

type _Fibonacci_FeedbackMessageTypeSupport struct{}

func (t _Fibonacci_FeedbackMessageTypeSupport) New() types.Message {
	return NewFibonacci_FeedbackMessage()
}

func (t _Fibonacci_FeedbackMessageTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__action__Fibonacci_FeedbackMessage
	return (unsafe.Pointer)(C.example_interfaces__action__Fibonacci_FeedbackMessage__create())
}

func (t _Fibonacci_FeedbackMessageTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__action__Fibonacci_FeedbackMessage__destroy((*C.example_interfaces__action__Fibonacci_FeedbackMessage)(pointer_to_free))
}

func (t _Fibonacci_FeedbackMessageTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Fibonacci_FeedbackMessage)
	mem := (*C.example_interfaces__action__Fibonacci_FeedbackMessage)(dst)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal_id), &m.GoalID)
	Fibonacci_FeedbackTypeSupport.AsCStruct(unsafe.Pointer(&mem.feedback), &m.Feedback)
}

func (t _Fibonacci_FeedbackMessageTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Fibonacci_FeedbackMessage)
	mem := (*C.example_interfaces__action__Fibonacci_FeedbackMessage)(ros2_message_buffer)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsGoStruct(&m.GoalID, unsafe.Pointer(&mem.goal_id))
	Fibonacci_FeedbackTypeSupport.AsGoStruct(&m.Feedback, unsafe.Pointer(&mem.feedback))
}

func (t _Fibonacci_FeedbackMessageTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__action__Fibonacci_FeedbackMessage())
}

type CFibonacci_FeedbackMessage = C.example_interfaces__action__Fibonacci_FeedbackMessage
type CFibonacci_FeedbackMessage__Sequence = C.example_interfaces__action__Fibonacci_FeedbackMessage__Sequence

func Fibonacci_FeedbackMessage__Sequence_to_Go(goSlice *[]Fibonacci_FeedbackMessage, cSlice CFibonacci_FeedbackMessage__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Fibonacci_FeedbackMessage, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Fibonacci_FeedbackMessageTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Fibonacci_FeedbackMessage__Sequence_to_C(cSlice *CFibonacci_FeedbackMessage__Sequence, goSlice []Fibonacci_FeedbackMessage) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__action__Fibonacci_FeedbackMessage)(C.malloc(C.sizeof_struct_example_interfaces__action__Fibonacci_FeedbackMessage * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Fibonacci_FeedbackMessageTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Fibonacci_FeedbackMessage__Array_to_Go(goSlice []Fibonacci_FeedbackMessage, cSlice []CFibonacci_FeedbackMessage) {
	for i := 0; i < len(cSlice); i++ {
		Fibonacci_FeedbackMessageTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Fibonacci_FeedbackMessage__Array_to_C(cSlice []CFibonacci_FeedbackMessage, goSlice []Fibonacci_FeedbackMessage) {
	for i := 0; i < len(goSlice); i++ {
		Fibonacci_FeedbackMessageTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
