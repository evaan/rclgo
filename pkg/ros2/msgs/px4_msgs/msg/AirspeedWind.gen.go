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
#include <px4_msgs/msg/airspeed_wind.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/AirspeedWind", &AirspeedWind{})
}
const (
	AirspeedWind_SOURCE_AS_BETA_ONLY uint8 = 0// wind estimate only based on synthetic sideslip fusion
	AirspeedWind_SOURCE_AS_SENSOR_1 uint8 = 1// combined synthetic sideslip and airspeed fusion (data from first airspeed sensor)
	AirspeedWind_SOURCE_AS_SENSOR_2 uint8 = 2// combined synthetic sideslip and airspeed fusion (data from second airspeed sensor)
	AirspeedWind_SOURCE_AS_SENSOR_3 uint8 = 3// combined synthetic sideslip and airspeed fusion (data from third airspeed sensor)
)

// Do not create instances of this type directly. Always use NewAirspeedWind
// function instead.
type AirspeedWind struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	WindspeedNorth float32 `yaml:"windspeed_north"`// Wind component in north / X direction (m/sec)
	WindspeedEast float32 `yaml:"windspeed_east"`// Wind component in east / Y direction (m/sec)
	VarianceNorth float32 `yaml:"variance_north"`// Wind estimate error variance in north / X direction (m/sec)**2 - set to zero (no uncertainty) if not estimated
	VarianceEast float32 `yaml:"variance_east"`// Wind estimate error variance in east / Y direction (m/sec)**2 - set to zero (no uncertainty) if not estimated
	TasInnov float32 `yaml:"tas_innov"`// True airspeed innovation
	TasInnovVar float32 `yaml:"tas_innov_var"`// True airspeed innovation variance
	TasScale float32 `yaml:"tas_scale"`// Estimated true airspeed scale factor
	BetaInnov float32 `yaml:"beta_innov"`// Sideslip measurement innovation
	BetaInnovVar float32 `yaml:"beta_innov_var"`// Sideslip measurement innovation variance
	Source uint8 `yaml:"source"`// source of wind estimate
}

// NewAirspeedWind creates a new AirspeedWind with default values.
func NewAirspeedWind() *AirspeedWind {
	self := AirspeedWind{}
	self.SetDefaults(nil)
	return &self
}

func (t *AirspeedWind) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *AirspeedWind) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__AirspeedWind())
}
func (t *AirspeedWind) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__AirspeedWind
	return (unsafe.Pointer)(C.px4_msgs__msg__AirspeedWind__create())
}
func (t *AirspeedWind) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__AirspeedWind__destroy((*C.px4_msgs__msg__AirspeedWind)(pointer_to_free))
}
func (t *AirspeedWind) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__AirspeedWind)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_sample = C.uint64_t(t.TimestampSample)
	mem.windspeed_north = C.float(t.WindspeedNorth)
	mem.windspeed_east = C.float(t.WindspeedEast)
	mem.variance_north = C.float(t.VarianceNorth)
	mem.variance_east = C.float(t.VarianceEast)
	mem.tas_innov = C.float(t.TasInnov)
	mem.tas_innov_var = C.float(t.TasInnovVar)
	mem.tas_scale = C.float(t.TasScale)
	mem.beta_innov = C.float(t.BetaInnov)
	mem.beta_innov_var = C.float(t.BetaInnovVar)
	mem.source = C.uint8_t(t.Source)
	return unsafe.Pointer(mem)
}
func (t *AirspeedWind) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__AirspeedWind)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampSample = uint64(mem.timestamp_sample)
	t.WindspeedNorth = float32(mem.windspeed_north)
	t.WindspeedEast = float32(mem.windspeed_east)
	t.VarianceNorth = float32(mem.variance_north)
	t.VarianceEast = float32(mem.variance_east)
	t.TasInnov = float32(mem.tas_innov)
	t.TasInnovVar = float32(mem.tas_innov_var)
	t.TasScale = float32(mem.tas_scale)
	t.BetaInnov = float32(mem.beta_innov)
	t.BetaInnovVar = float32(mem.beta_innov_var)
	t.Source = uint8(mem.source)
}
func (t *AirspeedWind) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CAirspeedWind = C.px4_msgs__msg__AirspeedWind
type CAirspeedWind__Sequence = C.px4_msgs__msg__AirspeedWind__Sequence

func AirspeedWind__Sequence_to_Go(goSlice *[]AirspeedWind, cSlice CAirspeedWind__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]AirspeedWind, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__AirspeedWind__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__AirspeedWind * uintptr(i)),
		))
		(*goSlice)[i] = AirspeedWind{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func AirspeedWind__Sequence_to_C(cSlice *CAirspeedWind__Sequence, goSlice []AirspeedWind) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__AirspeedWind)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__AirspeedWind * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__AirspeedWind)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__AirspeedWind * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__AirspeedWind)(v.AsCStruct())
	}
}
func AirspeedWind__Array_to_Go(goSlice []AirspeedWind, cSlice []CAirspeedWind) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func AirspeedWind__Array_to_C(cSlice []CAirspeedWind, goSlice []AirspeedWind) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__AirspeedWind)(goSlice[i].AsCStruct())
	}
}


