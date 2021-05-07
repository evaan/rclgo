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
#include <geometry_msgs/msg/accel_with_covariance.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("geometry_msgs/AccelWithCovariance", &AccelWithCovariance{})
}

// Do not create instances of this type directly. Always use NewAccelWithCovariance
// function instead.
type AccelWithCovariance struct {
	Accel Accel `yaml:"accel"`
	Covariance [36]float64 `yaml:"covariance"`// Row-major representation of the 6x6 covariance matrixThe orientation parameters use a fixed-axis representation.In order, the parameters are:(x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axis)
}

// NewAccelWithCovariance creates a new AccelWithCovariance with default values.
func NewAccelWithCovariance() *AccelWithCovariance {
	self := AccelWithCovariance{}
	self.SetDefaults(nil)
	return &self
}

func (t *AccelWithCovariance) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Accel.SetDefaults(nil)
	
	return t
}

func (t *AccelWithCovariance) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__AccelWithCovariance())
}
func (t *AccelWithCovariance) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__AccelWithCovariance
	return (unsafe.Pointer)(C.geometry_msgs__msg__AccelWithCovariance__create())
}
func (t *AccelWithCovariance) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__AccelWithCovariance__destroy((*C.geometry_msgs__msg__AccelWithCovariance)(pointer_to_free))
}
func (t *AccelWithCovariance) AsCStruct() unsafe.Pointer {
	mem := (*C.geometry_msgs__msg__AccelWithCovariance)(t.PrepareMemory())
	mem.accel = *(*C.geometry_msgs__msg__Accel)(t.Accel.AsCStruct())
	cSlice_covariance := mem.covariance[:]
	rosidl_runtime_c.Float64__Array_to_C(*(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_covariance)), t.Covariance[:])
	return unsafe.Pointer(mem)
}
func (t *AccelWithCovariance) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.geometry_msgs__msg__AccelWithCovariance)(ros2_message_buffer)
	t.Accel.AsGoStruct(unsafe.Pointer(&mem.accel))
	cSlice_covariance := mem.covariance[:]
	rosidl_runtime_c.Float64__Array_to_Go(t.Covariance[:], *(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_covariance)))
}
func (t *AccelWithCovariance) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CAccelWithCovariance = C.geometry_msgs__msg__AccelWithCovariance
type CAccelWithCovariance__Sequence = C.geometry_msgs__msg__AccelWithCovariance__Sequence

func AccelWithCovariance__Sequence_to_Go(goSlice *[]AccelWithCovariance, cSlice CAccelWithCovariance__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]AccelWithCovariance, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__AccelWithCovariance__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__AccelWithCovariance * uintptr(i)),
		))
		(*goSlice)[i] = AccelWithCovariance{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func AccelWithCovariance__Sequence_to_C(cSlice *CAccelWithCovariance__Sequence, goSlice []AccelWithCovariance) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__AccelWithCovariance)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__AccelWithCovariance * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__AccelWithCovariance)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__AccelWithCovariance * uintptr(i)),
		))
		*cIdx = *(*C.geometry_msgs__msg__AccelWithCovariance)(v.AsCStruct())
	}
}
func AccelWithCovariance__Array_to_Go(goSlice []AccelWithCovariance, cSlice []CAccelWithCovariance) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func AccelWithCovariance__Array_to_C(cSlice []CAccelWithCovariance, goSlice []AccelWithCovariance) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.geometry_msgs__msg__AccelWithCovariance)(goSlice[i].AsCStruct())
	}
}


