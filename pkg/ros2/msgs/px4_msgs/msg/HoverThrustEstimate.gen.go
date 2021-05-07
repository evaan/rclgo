/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <px4_msgs/msg/hover_thrust_estimate.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/HoverThrustEstimate", &HoverThrustEstimate{})
}

// Do not create instances of this type directly. Always use NewHoverThrustEstimate
// function instead.
type HoverThrustEstimate struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// time of corresponding sensor data last used for this estimate
	HoverThrust float32 `yaml:"hover_thrust"`// estimated hover thrust [0.1, 0.9]
	HoverThrustVar float32 `yaml:"hover_thrust_var"`// estimated hover thrust variance
	AccelInnov float32 `yaml:"accel_innov"`// innovation of the last acceleration fusion
	AccelInnovVar float32 `yaml:"accel_innov_var"`// innovation variance of the last acceleration fusion
	AccelInnovTestRatio float32 `yaml:"accel_innov_test_ratio"`// normalized innovation squared test ratio
	AccelNoiseVar float32 `yaml:"accel_noise_var"`// vertical acceleration noise variance estimated form innovation residual
	Valid bool `yaml:"valid"`
}

// NewHoverThrustEstimate creates a new HoverThrustEstimate with default values.
func NewHoverThrustEstimate() *HoverThrustEstimate {
	self := HoverThrustEstimate{}
	self.SetDefaults(nil)
	return &self
}

func (t *HoverThrustEstimate) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *HoverThrustEstimate) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__HoverThrustEstimate())
}
func (t *HoverThrustEstimate) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__HoverThrustEstimate
	return (unsafe.Pointer)(C.px4_msgs__msg__HoverThrustEstimate__create())
}
func (t *HoverThrustEstimate) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__HoverThrustEstimate__destroy((*C.px4_msgs__msg__HoverThrustEstimate)(pointer_to_free))
}
func (t *HoverThrustEstimate) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__HoverThrustEstimate)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_sample = C.uint64_t(t.TimestampSample)
	mem.hover_thrust = C.float(t.HoverThrust)
	mem.hover_thrust_var = C.float(t.HoverThrustVar)
	mem.accel_innov = C.float(t.AccelInnov)
	mem.accel_innov_var = C.float(t.AccelInnovVar)
	mem.accel_innov_test_ratio = C.float(t.AccelInnovTestRatio)
	mem.accel_noise_var = C.float(t.AccelNoiseVar)
	mem.valid = C.bool(t.Valid)
	return unsafe.Pointer(mem)
}
func (t *HoverThrustEstimate) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__HoverThrustEstimate)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampSample = uint64(mem.timestamp_sample)
	t.HoverThrust = float32(mem.hover_thrust)
	t.HoverThrustVar = float32(mem.hover_thrust_var)
	t.AccelInnov = float32(mem.accel_innov)
	t.AccelInnovVar = float32(mem.accel_innov_var)
	t.AccelInnovTestRatio = float32(mem.accel_innov_test_ratio)
	t.AccelNoiseVar = float32(mem.accel_noise_var)
	t.Valid = bool(mem.valid)
}
func (t *HoverThrustEstimate) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CHoverThrustEstimate = C.px4_msgs__msg__HoverThrustEstimate
type CHoverThrustEstimate__Sequence = C.px4_msgs__msg__HoverThrustEstimate__Sequence

func HoverThrustEstimate__Sequence_to_Go(goSlice *[]HoverThrustEstimate, cSlice CHoverThrustEstimate__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]HoverThrustEstimate, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__HoverThrustEstimate__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__HoverThrustEstimate * uintptr(i)),
		))
		(*goSlice)[i] = HoverThrustEstimate{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func HoverThrustEstimate__Sequence_to_C(cSlice *CHoverThrustEstimate__Sequence, goSlice []HoverThrustEstimate) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__HoverThrustEstimate)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__HoverThrustEstimate * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__HoverThrustEstimate)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__HoverThrustEstimate * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__HoverThrustEstimate)(v.AsCStruct())
	}
}
func HoverThrustEstimate__Array_to_Go(goSlice []HoverThrustEstimate, cSlice []CHoverThrustEstimate) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func HoverThrustEstimate__Array_to_C(cSlice []CHoverThrustEstimate, goSlice []HoverThrustEstimate) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__HoverThrustEstimate)(goSlice[i].AsCStruct())
	}
}


