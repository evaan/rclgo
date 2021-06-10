/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	sensor_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/sensor_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/srv/set_camera_info.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/SetCameraInfo_Request", SetCameraInfo_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetCameraInfo_Request
// function instead.
type SetCameraInfo_Request struct {
	CameraInfo sensor_msgs_msg.CameraInfo `yaml:"camera_info"`// The camera_info to store
}

// NewSetCameraInfo_Request creates a new SetCameraInfo_Request with default values.
func NewSetCameraInfo_Request() *SetCameraInfo_Request {
	self := SetCameraInfo_Request{}
	self.SetDefaults()
	return &self
}

func (t *SetCameraInfo_Request) Clone() *SetCameraInfo_Request {
	c := &SetCameraInfo_Request{}
	c.CameraInfo = *t.CameraInfo.Clone()
	return c
}

func (t *SetCameraInfo_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetCameraInfo_Request) SetDefaults() {
	t.CameraInfo.SetDefaults()
}

// CloneSetCameraInfo_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetCameraInfo_RequestSlice(dst, src []SetCameraInfo_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetCameraInfo_RequestTypeSupport types.MessageTypeSupport = _SetCameraInfo_RequestTypeSupport{}

type _SetCameraInfo_RequestTypeSupport struct{}

func (t _SetCameraInfo_RequestTypeSupport) New() types.Message {
	return NewSetCameraInfo_Request()
}

func (t _SetCameraInfo_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__srv__SetCameraInfo_Request
	return (unsafe.Pointer)(C.sensor_msgs__srv__SetCameraInfo_Request__create())
}

func (t _SetCameraInfo_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__srv__SetCameraInfo_Request__destroy((*C.sensor_msgs__srv__SetCameraInfo_Request)(pointer_to_free))
}

func (t _SetCameraInfo_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetCameraInfo_Request)
	mem := (*C.sensor_msgs__srv__SetCameraInfo_Request)(dst)
	sensor_msgs_msg.CameraInfoTypeSupport.AsCStruct(unsafe.Pointer(&mem.camera_info), &m.CameraInfo)
}

func (t _SetCameraInfo_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetCameraInfo_Request)
	mem := (*C.sensor_msgs__srv__SetCameraInfo_Request)(ros2_message_buffer)
	sensor_msgs_msg.CameraInfoTypeSupport.AsGoStruct(&m.CameraInfo, unsafe.Pointer(&mem.camera_info))
}

func (t _SetCameraInfo_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__srv__SetCameraInfo_Request())
}

type CSetCameraInfo_Request = C.sensor_msgs__srv__SetCameraInfo_Request
type CSetCameraInfo_Request__Sequence = C.sensor_msgs__srv__SetCameraInfo_Request__Sequence

func SetCameraInfo_Request__Sequence_to_Go(goSlice *[]SetCameraInfo_Request, cSlice CSetCameraInfo_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetCameraInfo_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__srv__SetCameraInfo_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__srv__SetCameraInfo_Request * uintptr(i)),
		))
		SetCameraInfo_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SetCameraInfo_Request__Sequence_to_C(cSlice *CSetCameraInfo_Request__Sequence, goSlice []SetCameraInfo_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__srv__SetCameraInfo_Request)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__srv__SetCameraInfo_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__srv__SetCameraInfo_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__srv__SetCameraInfo_Request * uintptr(i)),
		))
		SetCameraInfo_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SetCameraInfo_Request__Array_to_Go(goSlice []SetCameraInfo_Request, cSlice []CSetCameraInfo_Request) {
	for i := 0; i < len(cSlice); i++ {
		SetCameraInfo_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetCameraInfo_Request__Array_to_C(cSlice []CSetCameraInfo_Request, goSlice []SetCameraInfo_Request) {
	for i := 0; i < len(goSlice); i++ {
		SetCameraInfo_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}