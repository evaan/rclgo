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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <geometry_msgs/msg/transform.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("geometry_msgs/Transform", &Transform{})
}

// Do not create instances of this type directly. Always use NewTransform
// function instead.
type Transform struct {
	Translation Vector3 `yaml:"translation"`
	Rotation Quaternion `yaml:"rotation"`
}

// NewTransform creates a new Transform with default values.
func NewTransform() *Transform {
	self := Transform{}
	self.SetDefaults(nil)
	return &self
}

func (t *Transform) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Translation.SetDefaults(nil)
	t.Rotation.SetDefaults(nil)
	
	return t
}

func (t *Transform) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Transform())
}
func (t *Transform) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Transform
	return (unsafe.Pointer)(C.geometry_msgs__msg__Transform__create())
}
func (t *Transform) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Transform__destroy((*C.geometry_msgs__msg__Transform)(pointer_to_free))
}
func (t *Transform) AsCStruct() unsafe.Pointer {
	mem := (*C.geometry_msgs__msg__Transform)(t.PrepareMemory())
	mem.translation = *(*C.geometry_msgs__msg__Vector3)(t.Translation.AsCStruct())
	mem.rotation = *(*C.geometry_msgs__msg__Quaternion)(t.Rotation.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *Transform) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.geometry_msgs__msg__Transform)(ros2_message_buffer)
	t.Translation.AsGoStruct(unsafe.Pointer(&mem.translation))
	t.Rotation.AsGoStruct(unsafe.Pointer(&mem.rotation))
}
func (t *Transform) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CTransform = C.geometry_msgs__msg__Transform
type CTransform__Sequence = C.geometry_msgs__msg__Transform__Sequence

func Transform__Sequence_to_Go(goSlice *[]Transform, cSlice CTransform__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Transform, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Transform__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Transform * uintptr(i)),
		))
		(*goSlice)[i] = Transform{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Transform__Sequence_to_C(cSlice *CTransform__Sequence, goSlice []Transform) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Transform)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Transform * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Transform)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Transform * uintptr(i)),
		))
		*cIdx = *(*C.geometry_msgs__msg__Transform)(v.AsCStruct())
	}
}
func Transform__Array_to_Go(goSlice []Transform, cSlice []CTransform) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Transform__Array_to_C(cSlice []CTransform, goSlice []Transform) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.geometry_msgs__msg__Transform)(goSlice[i].AsCStruct())
	}
}


