/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package geometry_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	std_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <geometry_msgs/msg/quaternion_stamped.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("geometry_msgs/QuaternionStamped", &QuaternionStamped{})
}

// Do not create instances of this type directly. Always use NewQuaternionStamped
// function instead.
type QuaternionStamped struct {
	Header std_msgs.Header `yaml:"header"`
	Quaternion Quaternion `yaml:"quaternion"`
}

// NewQuaternionStamped creates a new QuaternionStamped with default values.
func NewQuaternionStamped() *QuaternionStamped {
	self := QuaternionStamped{}
	self.SetDefaults(nil)
	return &self
}

func (t *QuaternionStamped) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	t.Quaternion.SetDefaults(nil)
	
	return t
}

func (t *QuaternionStamped) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__QuaternionStamped())
}
func (t *QuaternionStamped) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__QuaternionStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__QuaternionStamped__create())
}
func (t *QuaternionStamped) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__QuaternionStamped__destroy((*C.geometry_msgs__msg__QuaternionStamped)(pointer_to_free))
}
func (t *QuaternionStamped) AsCStruct() unsafe.Pointer {
	mem := (*C.geometry_msgs__msg__QuaternionStamped)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	mem.quaternion = *(*C.geometry_msgs__msg__Quaternion)(t.Quaternion.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *QuaternionStamped) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.geometry_msgs__msg__QuaternionStamped)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	t.Quaternion.AsGoStruct(unsafe.Pointer(&mem.quaternion))
}
func (t *QuaternionStamped) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CQuaternionStamped = C.geometry_msgs__msg__QuaternionStamped
type CQuaternionStamped__Sequence = C.geometry_msgs__msg__QuaternionStamped__Sequence

func QuaternionStamped__Sequence_to_Go(goSlice *[]QuaternionStamped, cSlice CQuaternionStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]QuaternionStamped, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__QuaternionStamped__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__QuaternionStamped * uintptr(i)),
		))
		(*goSlice)[i] = QuaternionStamped{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func QuaternionStamped__Sequence_to_C(cSlice *CQuaternionStamped__Sequence, goSlice []QuaternionStamped) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__QuaternionStamped)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__QuaternionStamped * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__QuaternionStamped)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__QuaternionStamped * uintptr(i)),
		))
		*cIdx = *(*C.geometry_msgs__msg__QuaternionStamped)(v.AsCStruct())
	}
}
func QuaternionStamped__Array_to_Go(goSlice []QuaternionStamped, cSlice []CQuaternionStamped) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func QuaternionStamped__Array_to_C(cSlice []CQuaternionStamped, goSlice []QuaternionStamped) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.geometry_msgs__msg__QuaternionStamped)(goSlice[i].AsCStruct())
	}
}


