/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package example_interfaces_action

/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <example_interfaces/action/fibonacci.h>
*/
import "C"

import (
	"context"
	"time"
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"

	action_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/action_msgs/msg"
	action_msgs_srv "github.com/tiiuae/rclgo/internal/msgs/action_msgs/srv"
)

func init() {
	typemap.RegisterAction("example_interfaces/Fibonacci", FibonacciTypeSupport)
}

type _FibonacciTypeSupport struct {}

func (s _FibonacciTypeSupport) Goal() types.MessageTypeSupport {
	return Fibonacci_GoalTypeSupport
}

func (s _FibonacciTypeSupport) SendGoal() types.ServiceTypeSupport {
	return Fibonacci_SendGoalTypeSupport
}

func (s _FibonacciTypeSupport) NewSendGoalResponse(accepted bool, stamp time.Duration) types.Message {
	msg := NewFibonacci_SendGoal_Response()
	msg.Accepted = accepted
	secs := stamp.Truncate(time.Second)
	msg.Stamp.Sec = int32(secs)
	msg.Stamp.Nanosec = uint32(stamp - secs)
	return msg
}

func (s _FibonacciTypeSupport) Result() types.MessageTypeSupport {
	return Fibonacci_ResultTypeSupport
}

func (s _FibonacciTypeSupport) GetResult() types.ServiceTypeSupport {
	return Fibonacci_GetResultTypeSupport
}

func (s _FibonacciTypeSupport) NewGetResultResponse(status int8, result types.Message) types.Message {
	msg := NewFibonacci_GetResult_Response()
	msg.Status = status
	if result == nil {
		msg.Result = *NewFibonacci_Result()
	} else {
		msg.Result = *result.(*Fibonacci_Result)
	}
	return msg
}

func (s _FibonacciTypeSupport) CancelGoal() types.ServiceTypeSupport {
	return action_msgs_srv.CancelGoalTypeSupport
}

func (s _FibonacciTypeSupport) Feedback() types.MessageTypeSupport {
	return Fibonacci_FeedbackTypeSupport
}

func (s _FibonacciTypeSupport) FeedbackMessage() types.MessageTypeSupport {
	return Fibonacci_FeedbackMessageTypeSupport
}

func (s _FibonacciTypeSupport) NewFeedbackMessage(goalID *types.GoalID, feedback types.Message) types.Message {
	msg := NewFibonacci_FeedbackMessage()
	msg.GoalID.Uuid = *goalID
	msg.Feedback = *feedback.(*Fibonacci_Feedback)
	return msg
}

func (s _FibonacciTypeSupport) GoalStatusArray() types.MessageTypeSupport {
	return action_msgs_msg.GoalStatusArrayTypeSupport
}

func (s _FibonacciTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_action_type_support_handle__example_interfaces__action__Fibonacci())
}

// Modifying this variable is undefined behavior.
var FibonacciTypeSupport types.ActionTypeSupport = _FibonacciTypeSupport{}

type FibonacciFeedbackSender struct {
	sender rclgo.FeedbackSender
}

func (s *FibonacciFeedbackSender) Send(msg *Fibonacci_Feedback) error {
	return s.sender.Send(msg)
}

type FibonacciGoalHandle struct{
	*rclgo.GoalHandle

	Description *Fibonacci_Goal
}

func (g *FibonacciGoalHandle) Accept() (*FibonacciFeedbackSender, error) {
	s, err := g.GoalHandle.Accept()
	if err != nil {
		return nil, err
	}
	return &FibonacciFeedbackSender{*s}, nil
}

type FibonacciAction interface {
	ExecuteGoal(context.Context, *FibonacciGoalHandle) (*Fibonacci_Result, error)
}

func NewFibonacciAction(
	executeGoal func(context.Context, *FibonacciGoalHandle) (*Fibonacci_Result, error),
) FibonacciAction {
	return _FibonacciFuncAction(executeGoal)
}

type _FibonacciFuncAction func(context.Context, *FibonacciGoalHandle) (*Fibonacci_Result, error)

func (a _FibonacciFuncAction) ExecuteGoal(
	ctx context.Context, goal *FibonacciGoalHandle,
) (*Fibonacci_Result, error) {
	return a(ctx, goal)
}

type _FibonacciAction struct {
	action FibonacciAction
}

func (a _FibonacciAction) ExecuteGoal(ctx context.Context, handle *rclgo.GoalHandle) (types.Message, error) {
	return a.action.ExecuteGoal(ctx, &FibonacciGoalHandle{
		GoalHandle:  handle,
		Description: handle.Description.(*Fibonacci_Goal),
	})
}

func (a _FibonacciAction) TypeSupport() types.ActionTypeSupport {
	return FibonacciTypeSupport
}

type FibonacciServer struct{
	*rclgo.ActionServer
}

func NewFibonacciServer(node *rclgo.Node, name string, action FibonacciAction, opts *rclgo.ActionServerOptions) (*FibonacciServer, error) {
	server, err := node.NewActionServer(name, _FibonacciAction{action}, opts)
	if err != nil {
		return nil, err
	}
	return &FibonacciServer{server}, nil
}

type FibonacciFeedbackHandler func(context.Context, *Fibonacci_FeedbackMessage)

type FibonacciStatusHandler func(context.Context, *action_msgs_msg.GoalStatus)

type FibonacciClient struct{
	*rclgo.ActionClient
}

func NewFibonacciClient(node *rclgo.Node, name string, opts *rclgo.ActionClientOptions) (*FibonacciClient, error) {
	client, err := node.NewActionClient(name, FibonacciTypeSupport, opts)
	if err != nil {
		return nil, err
	}
	return &FibonacciClient{client}, nil
}

func (c *FibonacciClient) WatchGoal(ctx context.Context, goal *Fibonacci_Goal, onFeedback FibonacciFeedbackHandler) (*Fibonacci_GetResult_Response, error) {
	resp, err := c.ActionClient.WatchGoal(ctx, goal, func(ctx context.Context, msg types.Message) {
		onFeedback(ctx, msg.(*Fibonacci_FeedbackMessage))
	})
	return resp.(*Fibonacci_GetResult_Response), err
}

func (c *FibonacciClient) SendGoal(ctx context.Context, goal *Fibonacci_Goal) (*Fibonacci_SendGoal_Response, *types.GoalID, error) {
	resp, id, err := c.ActionClient.SendGoal(ctx, goal)
	return resp.(*Fibonacci_SendGoal_Response), id, err
}

func (c *FibonacciClient) SendGoalRequest(ctx context.Context, request *Fibonacci_SendGoal_Request) (*Fibonacci_SendGoal_Response, error) {
	resp, err := c.ActionClient.SendGoalRequest(ctx, request)
	return resp.(*Fibonacci_SendGoal_Response), err
}

func (c *FibonacciClient) GetResult(ctx context.Context, goalID *types.GoalID) (*Fibonacci_GetResult_Response, error) {
	resp, err := c.ActionClient.GetResult(ctx, goalID)
	return resp.(*Fibonacci_GetResult_Response), err
}

func (c *FibonacciClient) CancelGoal(ctx context.Context, request *action_msgs_srv.CancelGoal_Request) (*action_msgs_srv.CancelGoal_Response, error) {
	resp, err := c.ActionClient.CancelGoal(ctx, request)
	return resp.(*action_msgs_srv.CancelGoal_Response), err
}

func (c *FibonacciClient) WatchFeedback(ctx context.Context, goalID *types.GoalID, handler FibonacciFeedbackHandler) <-chan error {
	return c.ActionClient.WatchFeedback(ctx, goalID, func(ctx context.Context, msg types.Message) {
		handler(ctx, msg.(*Fibonacci_FeedbackMessage))
	})
}

func (c *FibonacciClient) WatchStatus(ctx context.Context, goalID *types.GoalID, handler FibonacciStatusHandler) <-chan error {
	return c.ActionClient.WatchStatus(ctx, goalID, func(ctx context.Context, msg types.Message) {
		handler(ctx, msg.(*action_msgs_msg.GoalStatus))
	})
}