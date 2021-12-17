/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <example_interfaces/srv/add_two_ints.h>
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
	typemap.RegisterService("example_interfaces/AddTwoInts", AddTwoIntsTypeSupport)
}

type _AddTwoIntsTypeSupport struct {}

func (s _AddTwoIntsTypeSupport) Request() types.MessageTypeSupport {
	return AddTwoInts_RequestTypeSupport
}

func (s _AddTwoIntsTypeSupport) Response() types.MessageTypeSupport {
	return AddTwoInts_ResponseTypeSupport
}

func (s _AddTwoIntsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__example_interfaces__srv__AddTwoInts())
}

// Modifying this variable is undefined behavior.
var AddTwoIntsTypeSupport types.ServiceTypeSupport = _AddTwoIntsTypeSupport{}

// AddTwoIntsClient wraps rclgo.Client to provide type safe helper
// functions
type AddTwoIntsClient struct {
	*rclgo.Client
}

// NewAddTwoIntsClient creates and returns a new client for the
// AddTwoInts
func NewAddTwoIntsClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*AddTwoIntsClient, error) {
	client, err := node.NewClient(serviceName, AddTwoIntsTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &AddTwoIntsClient{client}, nil
}

func (s *AddTwoIntsClient) Send(ctx context.Context, req *AddTwoInts_Request) (*AddTwoInts_Response, *rclgo.RmwServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*AddTwoInts_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type AddTwoIntsServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s AddTwoIntsServiceResponseSender) SendResponse(resp *AddTwoInts_Response) error {
	return s.sender.SendResponse(resp)
}

type AddTwoIntsServiceRequestHandler func(*rclgo.RmwServiceInfo, *AddTwoInts_Request, AddTwoIntsServiceResponseSender)

// AddTwoIntsService wraps rclgo.Service to provide type safe helper
// functions
type AddTwoIntsService struct {
	*rclgo.Service
}

// NewAddTwoIntsService creates and returns a new service for the
// AddTwoInts
func NewAddTwoIntsService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler AddTwoIntsServiceRequestHandler) (*AddTwoIntsService, error) {
	h := func(rmw *rclgo.RmwServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*AddTwoInts_Request)
		responseSender := AddTwoIntsServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, AddTwoIntsTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &AddTwoIntsService{service}, nil
}

