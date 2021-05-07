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
#include <px4_msgs/msg/estimator_wind.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/EstimatorWind", &EstimatorWind{})
}

// Do not create instances of this type directly. Always use NewEstimatorWind
// function instead.
type EstimatorWind struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	WindspeedNorth float32 `yaml:"windspeed_north"`// Wind component in north / X direction (m/sec)
	WindspeedEast float32 `yaml:"windspeed_east"`// Wind component in east / Y direction (m/sec)
	VarianceNorth float32 `yaml:"variance_north"`// Wind estimate error variance in north / X direction (m/sec)**2 - set to zero (no uncertainty) if not estimated
	VarianceEast float32 `yaml:"variance_east"`// Wind estimate error variance in east / Y direction (m/sec)**2 - set to zero (no uncertainty) if not estimated
	TasInnov float32 `yaml:"tas_innov"`// True airspeed innovation
	TasInnovVar float32 `yaml:"tas_innov_var"`// True airspeed innovation variance
	BetaInnov float32 `yaml:"beta_innov"`// Sideslip measurement innovation
	BetaInnovVar float32 `yaml:"beta_innov_var"`// Sideslip measurement innovation variance
}

// NewEstimatorWind creates a new EstimatorWind with default values.
func NewEstimatorWind() *EstimatorWind {
	self := EstimatorWind{}
	self.SetDefaults(nil)
	return &self
}

func (t *EstimatorWind) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *EstimatorWind) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EstimatorWind())
}
func (t *EstimatorWind) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EstimatorWind
	return (unsafe.Pointer)(C.px4_msgs__msg__EstimatorWind__create())
}
func (t *EstimatorWind) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EstimatorWind__destroy((*C.px4_msgs__msg__EstimatorWind)(pointer_to_free))
}
func (t *EstimatorWind) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__EstimatorWind)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_sample = C.uint64_t(t.TimestampSample)
	mem.windspeed_north = C.float(t.WindspeedNorth)
	mem.windspeed_east = C.float(t.WindspeedEast)
	mem.variance_north = C.float(t.VarianceNorth)
	mem.variance_east = C.float(t.VarianceEast)
	mem.tas_innov = C.float(t.TasInnov)
	mem.tas_innov_var = C.float(t.TasInnovVar)
	mem.beta_innov = C.float(t.BetaInnov)
	mem.beta_innov_var = C.float(t.BetaInnovVar)
	return unsafe.Pointer(mem)
}
func (t *EstimatorWind) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__EstimatorWind)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampSample = uint64(mem.timestamp_sample)
	t.WindspeedNorth = float32(mem.windspeed_north)
	t.WindspeedEast = float32(mem.windspeed_east)
	t.VarianceNorth = float32(mem.variance_north)
	t.VarianceEast = float32(mem.variance_east)
	t.TasInnov = float32(mem.tas_innov)
	t.TasInnovVar = float32(mem.tas_innov_var)
	t.BetaInnov = float32(mem.beta_innov)
	t.BetaInnovVar = float32(mem.beta_innov_var)
}
func (t *EstimatorWind) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CEstimatorWind = C.px4_msgs__msg__EstimatorWind
type CEstimatorWind__Sequence = C.px4_msgs__msg__EstimatorWind__Sequence

func EstimatorWind__Sequence_to_Go(goSlice *[]EstimatorWind, cSlice CEstimatorWind__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorWind, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EstimatorWind__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorWind * uintptr(i)),
		))
		(*goSlice)[i] = EstimatorWind{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func EstimatorWind__Sequence_to_C(cSlice *CEstimatorWind__Sequence, goSlice []EstimatorWind) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EstimatorWind)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EstimatorWind * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EstimatorWind)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorWind * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__EstimatorWind)(v.AsCStruct())
	}
}
func EstimatorWind__Array_to_Go(goSlice []EstimatorWind, cSlice []CEstimatorWind) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorWind__Array_to_C(cSlice []CEstimatorWind, goSlice []EstimatorWind) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__EstimatorWind)(goSlice[i].AsCStruct())
	}
}


