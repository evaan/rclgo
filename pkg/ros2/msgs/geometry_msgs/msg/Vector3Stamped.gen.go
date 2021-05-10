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
#include <geometry_msgs/msg/vector3_stamped.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("geometry_msgs/Vector3Stamped", &Vector3Stamped{})
}

// Do not create instances of this type directly. Always use NewVector3Stamped
// function instead.
type Vector3Stamped struct {
	Header std_msgs.Header `yaml:"header"`
	Vector Vector3 `yaml:"vector"`
}

// NewVector3Stamped creates a new Vector3Stamped with default values.
func NewVector3Stamped() *Vector3Stamped {
	self := Vector3Stamped{}
	self.SetDefaults(nil)
	return &self
}

func (t *Vector3Stamped) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	t.Vector.SetDefaults(nil)
	
	return t
}

func (t *Vector3Stamped) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Vector3Stamped())
}
func (t *Vector3Stamped) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Vector3Stamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__Vector3Stamped__create())
}
func (t *Vector3Stamped) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Vector3Stamped__destroy((*C.geometry_msgs__msg__Vector3Stamped)(pointer_to_free))
}
func (t *Vector3Stamped) AsCStruct() unsafe.Pointer {
	mem := (*C.geometry_msgs__msg__Vector3Stamped)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	mem.vector = *(*C.geometry_msgs__msg__Vector3)(t.Vector.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *Vector3Stamped) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.geometry_msgs__msg__Vector3Stamped)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	t.Vector.AsGoStruct(unsafe.Pointer(&mem.vector))
}
func (t *Vector3Stamped) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CVector3Stamped = C.geometry_msgs__msg__Vector3Stamped
type CVector3Stamped__Sequence = C.geometry_msgs__msg__Vector3Stamped__Sequence

func Vector3Stamped__Sequence_to_Go(goSlice *[]Vector3Stamped, cSlice CVector3Stamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Vector3Stamped, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Vector3Stamped__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Vector3Stamped * uintptr(i)),
		))
		(*goSlice)[i] = Vector3Stamped{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Vector3Stamped__Sequence_to_C(cSlice *CVector3Stamped__Sequence, goSlice []Vector3Stamped) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Vector3Stamped)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Vector3Stamped * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Vector3Stamped)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Vector3Stamped * uintptr(i)),
		))
		*cIdx = *(*C.geometry_msgs__msg__Vector3Stamped)(v.AsCStruct())
	}
}
func Vector3Stamped__Array_to_Go(goSlice []Vector3Stamped, cSlice []CVector3Stamped) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Vector3Stamped__Array_to_C(cSlice []CVector3Stamped, goSlice []Vector3Stamped) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.geometry_msgs__msg__Vector3Stamped)(goSlice[i].AsCStruct())
	}
}

