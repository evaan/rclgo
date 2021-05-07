/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package std_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <std_msgs/msg/color_rgba.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("std_msgs/ColorRGBA", &ColorRGBA{})
}

// Do not create instances of this type directly. Always use NewColorRGBA
// function instead.
type ColorRGBA struct {
	R float32 `yaml:"r"`
	G float32 `yaml:"g"`
	B float32 `yaml:"b"`
	A float32 `yaml:"a"`
}

// NewColorRGBA creates a new ColorRGBA with default values.
func NewColorRGBA() *ColorRGBA {
	self := ColorRGBA{}
	self.SetDefaults(nil)
	return &self
}

func (t *ColorRGBA) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *ColorRGBA) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__ColorRGBA())
}
func (t *ColorRGBA) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__ColorRGBA
	return (unsafe.Pointer)(C.std_msgs__msg__ColorRGBA__create())
}
func (t *ColorRGBA) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__ColorRGBA__destroy((*C.std_msgs__msg__ColorRGBA)(pointer_to_free))
}
func (t *ColorRGBA) AsCStruct() unsafe.Pointer {
	mem := (*C.std_msgs__msg__ColorRGBA)(t.PrepareMemory())
	mem.r = C.float(t.R)
	mem.g = C.float(t.G)
	mem.b = C.float(t.B)
	mem.a = C.float(t.A)
	return unsafe.Pointer(mem)
}
func (t *ColorRGBA) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.std_msgs__msg__ColorRGBA)(ros2_message_buffer)
	t.R = float32(mem.r)
	t.G = float32(mem.g)
	t.B = float32(mem.b)
	t.A = float32(mem.a)
}
func (t *ColorRGBA) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CColorRGBA = C.std_msgs__msg__ColorRGBA
type CColorRGBA__Sequence = C.std_msgs__msg__ColorRGBA__Sequence

func ColorRGBA__Sequence_to_Go(goSlice *[]ColorRGBA, cSlice CColorRGBA__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ColorRGBA, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__ColorRGBA__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__ColorRGBA * uintptr(i)),
		))
		(*goSlice)[i] = ColorRGBA{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func ColorRGBA__Sequence_to_C(cSlice *CColorRGBA__Sequence, goSlice []ColorRGBA) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__ColorRGBA)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__ColorRGBA * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__ColorRGBA)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__ColorRGBA * uintptr(i)),
		))
		*cIdx = *(*C.std_msgs__msg__ColorRGBA)(v.AsCStruct())
	}
}
func ColorRGBA__Array_to_Go(goSlice []ColorRGBA, cSlice []CColorRGBA) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func ColorRGBA__Array_to_C(cSlice []CColorRGBA, goSlice []ColorRGBA) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.std_msgs__msg__ColorRGBA)(goSlice[i].AsCStruct())
	}
}


