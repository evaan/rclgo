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
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <px4_msgs/msg/estimator_local_position.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/EstimatorLocalPosition", &EstimatorLocalPosition{})
}
const (
	EstimatorLocalPosition_DIST_BOTTOM_SENSOR_NONE uint8 = 0// Distance to surface
	EstimatorLocalPosition_DIST_BOTTOM_SENSOR_RANGE uint8 = 1// (1 << 0) a range sensor is used to estimate dist_bottom field. Distance to surface
	EstimatorLocalPosition_DIST_BOTTOM_SENSOR_FLOW uint8 = 2// (1 << 1) a flow sensor is used to estimate dist_bottom field (mostly fixed-wing use case). Distance to surface
)

// Do not create instances of this type directly. Always use NewEstimatorLocalPosition
// function instead.
type EstimatorLocalPosition struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	XyValid bool `yaml:"xy_valid"`// true if x and y are valid
	ZValid bool `yaml:"z_valid"`// true if z is valid
	VXyValid bool `yaml:"v_xy_valid"`// true if vy and vy are valid
	VZValid bool `yaml:"v_z_valid"`// true if vz is valid
	X float32 `yaml:"x"`// North position in NED earth-fixed frame, (metres). Position in local NED frame
	Y float32 `yaml:"y"`// East position in NED earth-fixed frame, (metres). Position in local NED frame
	Z float32 `yaml:"z"`// Down position (negative altitude) in NED earth-fixed frame, (metres). Position in local NED frame
	DeltaXy [2]float32 `yaml:"delta_xy"`// Position reset delta
	XyResetCounter uint8 `yaml:"xy_reset_counter"`// Position reset delta
	DeltaZ float32 `yaml:"delta_z"`
	ZResetCounter uint8 `yaml:"z_reset_counter"`
	Vx float32 `yaml:"vx"`// North velocity in NED earth-fixed frame, (metres/sec). Velocity in NED frame
	Vy float32 `yaml:"vy"`// East velocity in NED earth-fixed frame, (metres/sec). Velocity in NED frame
	Vz float32 `yaml:"vz"`// Down velocity in NED earth-fixed frame, (metres/sec). Velocity in NED frame
	ZDeriv float32 `yaml:"z_deriv"`// Down position time derivative in NED earth-fixed frame, (metres/sec). Velocity in NED frame
	DeltaVxy [2]float32 `yaml:"delta_vxy"`// Velocity reset delta
	VxyResetCounter uint8 `yaml:"vxy_reset_counter"`// Velocity reset delta
	DeltaVz float32 `yaml:"delta_vz"`
	VzResetCounter uint8 `yaml:"vz_reset_counter"`
	Ax float32 `yaml:"ax"`// North velocity derivative in NED earth-fixed frame, (metres/sec^2). Acceleration in NED frame
	Ay float32 `yaml:"ay"`// East velocity derivative in NED earth-fixed frame, (metres/sec^2). Acceleration in NED frame
	Az float32 `yaml:"az"`// Down velocity derivative in NED earth-fixed frame, (metres/sec^2). Acceleration in NED frame
	Heading float32 `yaml:"heading"`// Euler yaw angle transforming the tangent plane relative to NED earth-fixed frame, -PI..+PI,  (radians)
	DeltaHeading float32 `yaml:"delta_heading"`
	HeadingResetCounter uint8 `yaml:"heading_reset_counter"`
	XyGlobal bool `yaml:"xy_global"`// true if position (x, y) has a valid global reference (ref_lat, ref_lon). Position of reference point (local NED frame origin) in global (GPS / WGS84) frame
	ZGlobal bool `yaml:"z_global"`// true if z has a valid global reference (ref_alt). Position of reference point (local NED frame origin) in global (GPS / WGS84) frame
	RefTimestamp uint64 `yaml:"ref_timestamp"`// Time when reference position was set since system start, (microseconds). Position of reference point (local NED frame origin) in global (GPS / WGS84) frame
	RefLat float64 `yaml:"ref_lat"`// Reference point latitude, (degrees). Position of reference point (local NED frame origin) in global (GPS / WGS84) frame
	RefLon float64 `yaml:"ref_lon"`// Reference point longitude, (degrees). Position of reference point (local NED frame origin) in global (GPS / WGS84) frame
	RefAlt float32 `yaml:"ref_alt"`// Reference altitude AMSL, (metres). Position of reference point (local NED frame origin) in global (GPS / WGS84) frame
	DistBottom float32 `yaml:"dist_bottom"`// Distance from from bottom surface to ground, (metres). Distance to surface
	DistBottomValid bool `yaml:"dist_bottom_valid"`// true if distance to bottom surface is valid. Distance to surface
	DistBottomSensorBitfield uint8 `yaml:"dist_bottom_sensor_bitfield"`// bitfield indicating what type of sensor is used to estimate dist_bottom. Distance to surface
	Eph float32 `yaml:"eph"`// Standard deviation of horizontal position error, (metres)
	Epv float32 `yaml:"epv"`// Standard deviation of vertical position error, (metres)
	Evh float32 `yaml:"evh"`// Standard deviation of horizontal velocity error, (metres/sec)
	Evv float32 `yaml:"evv"`// Standard deviation of horizontal velocity error, (metres/sec)
	VxyMax float32 `yaml:"vxy_max"`// maximum horizontal speed - set to 0 when limiting not required (meters/sec). estimator specified vehicle limits
	VzMax float32 `yaml:"vz_max"`// maximum vertical speed - set to 0 when limiting not required (meters/sec). estimator specified vehicle limits
	HaglMin float32 `yaml:"hagl_min"`// minimum height above ground level - set to 0 when limiting not required (meters). estimator specified vehicle limits
	HaglMax float32 `yaml:"hagl_max"`// maximum height above ground level - set to 0 when limiting not required (meters). estimator specified vehicle limits
}

// NewEstimatorLocalPosition creates a new EstimatorLocalPosition with default values.
func NewEstimatorLocalPosition() *EstimatorLocalPosition {
	self := EstimatorLocalPosition{}
	self.SetDefaults(nil)
	return &self
}

func (t *EstimatorLocalPosition) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *EstimatorLocalPosition) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EstimatorLocalPosition())
}
func (t *EstimatorLocalPosition) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EstimatorLocalPosition
	return (unsafe.Pointer)(C.px4_msgs__msg__EstimatorLocalPosition__create())
}
func (t *EstimatorLocalPosition) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EstimatorLocalPosition__destroy((*C.px4_msgs__msg__EstimatorLocalPosition)(pointer_to_free))
}
func (t *EstimatorLocalPosition) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__EstimatorLocalPosition)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_sample = C.uint64_t(t.TimestampSample)
	mem.xy_valid = C.bool(t.XyValid)
	mem.z_valid = C.bool(t.ZValid)
	mem.v_xy_valid = C.bool(t.VXyValid)
	mem.v_z_valid = C.bool(t.VZValid)
	mem.x = C.float(t.X)
	mem.y = C.float(t.Y)
	mem.z = C.float(t.Z)
	cSlice_delta_xy := mem.delta_xy[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_delta_xy)), t.DeltaXy[:])
	mem.xy_reset_counter = C.uint8_t(t.XyResetCounter)
	mem.delta_z = C.float(t.DeltaZ)
	mem.z_reset_counter = C.uint8_t(t.ZResetCounter)
	mem.vx = C.float(t.Vx)
	mem.vy = C.float(t.Vy)
	mem.vz = C.float(t.Vz)
	mem.z_deriv = C.float(t.ZDeriv)
	cSlice_delta_vxy := mem.delta_vxy[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_delta_vxy)), t.DeltaVxy[:])
	mem.vxy_reset_counter = C.uint8_t(t.VxyResetCounter)
	mem.delta_vz = C.float(t.DeltaVz)
	mem.vz_reset_counter = C.uint8_t(t.VzResetCounter)
	mem.ax = C.float(t.Ax)
	mem.ay = C.float(t.Ay)
	mem.az = C.float(t.Az)
	mem.heading = C.float(t.Heading)
	mem.delta_heading = C.float(t.DeltaHeading)
	mem.heading_reset_counter = C.uint8_t(t.HeadingResetCounter)
	mem.xy_global = C.bool(t.XyGlobal)
	mem.z_global = C.bool(t.ZGlobal)
	mem.ref_timestamp = C.uint64_t(t.RefTimestamp)
	mem.ref_lat = C.double(t.RefLat)
	mem.ref_lon = C.double(t.RefLon)
	mem.ref_alt = C.float(t.RefAlt)
	mem.dist_bottom = C.float(t.DistBottom)
	mem.dist_bottom_valid = C.bool(t.DistBottomValid)
	mem.dist_bottom_sensor_bitfield = C.uint8_t(t.DistBottomSensorBitfield)
	mem.eph = C.float(t.Eph)
	mem.epv = C.float(t.Epv)
	mem.evh = C.float(t.Evh)
	mem.evv = C.float(t.Evv)
	mem.vxy_max = C.float(t.VxyMax)
	mem.vz_max = C.float(t.VzMax)
	mem.hagl_min = C.float(t.HaglMin)
	mem.hagl_max = C.float(t.HaglMax)
	return unsafe.Pointer(mem)
}
func (t *EstimatorLocalPosition) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__EstimatorLocalPosition)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampSample = uint64(mem.timestamp_sample)
	t.XyValid = bool(mem.xy_valid)
	t.ZValid = bool(mem.z_valid)
	t.VXyValid = bool(mem.v_xy_valid)
	t.VZValid = bool(mem.v_z_valid)
	t.X = float32(mem.x)
	t.Y = float32(mem.y)
	t.Z = float32(mem.z)
	cSlice_delta_xy := mem.delta_xy[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.DeltaXy[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_delta_xy)))
	t.XyResetCounter = uint8(mem.xy_reset_counter)
	t.DeltaZ = float32(mem.delta_z)
	t.ZResetCounter = uint8(mem.z_reset_counter)
	t.Vx = float32(mem.vx)
	t.Vy = float32(mem.vy)
	t.Vz = float32(mem.vz)
	t.ZDeriv = float32(mem.z_deriv)
	cSlice_delta_vxy := mem.delta_vxy[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.DeltaVxy[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_delta_vxy)))
	t.VxyResetCounter = uint8(mem.vxy_reset_counter)
	t.DeltaVz = float32(mem.delta_vz)
	t.VzResetCounter = uint8(mem.vz_reset_counter)
	t.Ax = float32(mem.ax)
	t.Ay = float32(mem.ay)
	t.Az = float32(mem.az)
	t.Heading = float32(mem.heading)
	t.DeltaHeading = float32(mem.delta_heading)
	t.HeadingResetCounter = uint8(mem.heading_reset_counter)
	t.XyGlobal = bool(mem.xy_global)
	t.ZGlobal = bool(mem.z_global)
	t.RefTimestamp = uint64(mem.ref_timestamp)
	t.RefLat = float64(mem.ref_lat)
	t.RefLon = float64(mem.ref_lon)
	t.RefAlt = float32(mem.ref_alt)
	t.DistBottom = float32(mem.dist_bottom)
	t.DistBottomValid = bool(mem.dist_bottom_valid)
	t.DistBottomSensorBitfield = uint8(mem.dist_bottom_sensor_bitfield)
	t.Eph = float32(mem.eph)
	t.Epv = float32(mem.epv)
	t.Evh = float32(mem.evh)
	t.Evv = float32(mem.evv)
	t.VxyMax = float32(mem.vxy_max)
	t.VzMax = float32(mem.vz_max)
	t.HaglMin = float32(mem.hagl_min)
	t.HaglMax = float32(mem.hagl_max)
}
func (t *EstimatorLocalPosition) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CEstimatorLocalPosition = C.px4_msgs__msg__EstimatorLocalPosition
type CEstimatorLocalPosition__Sequence = C.px4_msgs__msg__EstimatorLocalPosition__Sequence

func EstimatorLocalPosition__Sequence_to_Go(goSlice *[]EstimatorLocalPosition, cSlice CEstimatorLocalPosition__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorLocalPosition, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EstimatorLocalPosition__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorLocalPosition * uintptr(i)),
		))
		(*goSlice)[i] = EstimatorLocalPosition{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func EstimatorLocalPosition__Sequence_to_C(cSlice *CEstimatorLocalPosition__Sequence, goSlice []EstimatorLocalPosition) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EstimatorLocalPosition)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EstimatorLocalPosition * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EstimatorLocalPosition)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorLocalPosition * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__EstimatorLocalPosition)(v.AsCStruct())
	}
}
func EstimatorLocalPosition__Array_to_Go(goSlice []EstimatorLocalPosition, cSlice []CEstimatorLocalPosition) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorLocalPosition__Array_to_C(cSlice []CEstimatorLocalPosition, goSlice []EstimatorLocalPosition) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__EstimatorLocalPosition)(goSlice[i].AsCStruct())
	}
}


