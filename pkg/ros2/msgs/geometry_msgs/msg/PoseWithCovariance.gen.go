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
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <geometry_msgs/msg/pose_with_covariance.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("geometry_msgs/PoseWithCovariance", &PoseWithCovariance{})
}

// Do not create instances of this type directly. Always use NewPoseWithCovariance
// function instead.
type PoseWithCovariance struct {
	Pose Pose `yaml:"pose"`
	Covariance [36]float64 `yaml:"covariance"`// Row-major representation of the 6x6 covariance matrixThe orientation parameters use a fixed-axis representation.In order, the parameters are:(x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axis)
}

// NewPoseWithCovariance creates a new PoseWithCovariance with default values.
func NewPoseWithCovariance() *PoseWithCovariance {
	self := PoseWithCovariance{}
	self.SetDefaults(nil)
	return &self
}

func (t *PoseWithCovariance) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Pose.SetDefaults(nil)
	
	return t
}

func (t *PoseWithCovariance) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__PoseWithCovariance())
}
func (t *PoseWithCovariance) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__PoseWithCovariance
	return (unsafe.Pointer)(C.geometry_msgs__msg__PoseWithCovariance__create())
}
func (t *PoseWithCovariance) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__PoseWithCovariance__destroy((*C.geometry_msgs__msg__PoseWithCovariance)(pointer_to_free))
}
func (t *PoseWithCovariance) AsCStruct() unsafe.Pointer {
	mem := (*C.geometry_msgs__msg__PoseWithCovariance)(t.PrepareMemory())
	mem.pose = *(*C.geometry_msgs__msg__Pose)(t.Pose.AsCStruct())
	cSlice_covariance := mem.covariance[:]
	rosidl_runtime_c.Float64__Array_to_C(*(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_covariance)), t.Covariance[:])
	return unsafe.Pointer(mem)
}
func (t *PoseWithCovariance) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.geometry_msgs__msg__PoseWithCovariance)(ros2_message_buffer)
	t.Pose.AsGoStruct(unsafe.Pointer(&mem.pose))
	cSlice_covariance := mem.covariance[:]
	rosidl_runtime_c.Float64__Array_to_Go(t.Covariance[:], *(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_covariance)))
}
func (t *PoseWithCovariance) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CPoseWithCovariance = C.geometry_msgs__msg__PoseWithCovariance
type CPoseWithCovariance__Sequence = C.geometry_msgs__msg__PoseWithCovariance__Sequence

func PoseWithCovariance__Sequence_to_Go(goSlice *[]PoseWithCovariance, cSlice CPoseWithCovariance__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PoseWithCovariance, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__PoseWithCovariance__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__PoseWithCovariance * uintptr(i)),
		))
		(*goSlice)[i] = PoseWithCovariance{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func PoseWithCovariance__Sequence_to_C(cSlice *CPoseWithCovariance__Sequence, goSlice []PoseWithCovariance) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__PoseWithCovariance)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__PoseWithCovariance * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__PoseWithCovariance)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__PoseWithCovariance * uintptr(i)),
		))
		*cIdx = *(*C.geometry_msgs__msg__PoseWithCovariance)(v.AsCStruct())
	}
}
func PoseWithCovariance__Array_to_Go(goSlice []PoseWithCovariance, cSlice []CPoseWithCovariance) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func PoseWithCovariance__Array_to_C(cSlice []CPoseWithCovariance, goSlice []PoseWithCovariance) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.geometry_msgs__msg__PoseWithCovariance)(goSlice[i].AsCStruct())
	}
}


