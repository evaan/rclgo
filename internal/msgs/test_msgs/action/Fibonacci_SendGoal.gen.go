/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_action

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <test_msgs/action/fibonacci.h>
*/
import "C"

import (
	"context"
	"errors"
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
)

func init() {
	typemap.RegisterService("test_msgs/Fibonacci_SendGoal", Fibonacci_SendGoalTypeSupport)
	typemap.RegisterService("test_msgs/action/Fibonacci_SendGoal", Fibonacci_SendGoalTypeSupport)
}

type _Fibonacci_SendGoalTypeSupport struct {}

func (s _Fibonacci_SendGoalTypeSupport) Request() types.MessageTypeSupport {
	return Fibonacci_SendGoal_RequestTypeSupport
}

func (s _Fibonacci_SendGoalTypeSupport) Response() types.MessageTypeSupport {
	return Fibonacci_SendGoal_ResponseTypeSupport
}

func (s _Fibonacci_SendGoalTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__test_msgs__action__Fibonacci_SendGoal())
}

// Modifying this variable is undefined behavior.
var Fibonacci_SendGoalTypeSupport types.ServiceTypeSupport = _Fibonacci_SendGoalTypeSupport{}

// Fibonacci_SendGoalClient wraps rclgo.Client to provide type safe helper
// functions
type Fibonacci_SendGoalClient struct {
	*rclgo.Client
}

// NewFibonacci_SendGoalClient creates and returns a new client for the
// Fibonacci_SendGoal
func NewFibonacci_SendGoalClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*Fibonacci_SendGoalClient, error) {
	client, err := node.NewClient(serviceName, Fibonacci_SendGoalTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_SendGoalClient{client}, nil
}

func (s *Fibonacci_SendGoalClient) Send(ctx context.Context, req *Fibonacci_SendGoal_Request) (*Fibonacci_SendGoal_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*Fibonacci_SendGoal_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type Fibonacci_SendGoalServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s Fibonacci_SendGoalServiceResponseSender) SendResponse(resp *Fibonacci_SendGoal_Response) error {
	return s.sender.SendResponse(resp)
}

type Fibonacci_SendGoalServiceRequestHandler func(*rclgo.ServiceInfo, *Fibonacci_SendGoal_Request, Fibonacci_SendGoalServiceResponseSender)

// Fibonacci_SendGoalService wraps rclgo.Service to provide type safe helper
// functions
type Fibonacci_SendGoalService struct {
	*rclgo.Service
}

// NewFibonacci_SendGoalService creates and returns a new service for the
// Fibonacci_SendGoal
func NewFibonacci_SendGoalService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler Fibonacci_SendGoalServiceRequestHandler) (*Fibonacci_SendGoalService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*Fibonacci_SendGoal_Request)
		responseSender := Fibonacci_SendGoalServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, Fibonacci_SendGoalTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_SendGoalService{service}, nil
}