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
#include <px4_msgs/msg/vehicle_attitude_setpoint.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleAttitudeSetpoint", &VehicleAttitudeSetpoint{})
}
const (
	VehicleAttitudeSetpoint_FLAPS_OFF uint8 = 0// no flaps
	VehicleAttitudeSetpoint_FLAPS_LAND uint8 = 1// landing config flaps
	VehicleAttitudeSetpoint_FLAPS_TAKEOFF uint8 = 2// take-off config flaps
)

// Do not create instances of this type directly. Always use NewVehicleAttitudeSetpoint
// function instead.
type VehicleAttitudeSetpoint struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	RollBody float32 `yaml:"roll_body"`// body angle in NED frame (can be NaN for FW)
	PitchBody float32 `yaml:"pitch_body"`// body angle in NED frame (can be NaN for FW)
	YawBody float32 `yaml:"yaw_body"`// body angle in NED frame (can be NaN for FW)
	YawSpMoveRate float32 `yaml:"yaw_sp_move_rate"`// rad/s (commanded by user)
	QD [4]float32 `yaml:"q_d"`// Desired quaternion for quaternion control. For quaternion-based attitude control
	ThrustBody [3]float32 `yaml:"thrust_body"`// Normalized thrust command in body NED frame [-1,1]. For clarification: For multicopters thrust_body[0] and thrust[1] are usually 0 and thrust[2] is the negative throttle demand.For fixed wings thrust_x is the throttle demand and thrust_y, thrust_z will usually be zero.
	RollResetIntegral bool `yaml:"roll_reset_integral"`// Reset roll integral part (navigation logic change)
	PitchResetIntegral bool `yaml:"pitch_reset_integral"`// Reset pitch integral part (navigation logic change)
	YawResetIntegral bool `yaml:"yaw_reset_integral"`// Reset yaw integral part (navigation logic change)
	FwControlYaw bool `yaml:"fw_control_yaw"`// control heading with rudder (used for auto takeoff on runway)
	ApplyFlaps uint8 `yaml:"apply_flaps"`// flap config specifier
}

// NewVehicleAttitudeSetpoint creates a new VehicleAttitudeSetpoint with default values.
func NewVehicleAttitudeSetpoint() *VehicleAttitudeSetpoint {
	self := VehicleAttitudeSetpoint{}
	self.SetDefaults(nil)
	return &self
}

func (t *VehicleAttitudeSetpoint) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *VehicleAttitudeSetpoint) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleAttitudeSetpoint())
}
func (t *VehicleAttitudeSetpoint) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleAttitudeSetpoint
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleAttitudeSetpoint__create())
}
func (t *VehicleAttitudeSetpoint) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleAttitudeSetpoint__destroy((*C.px4_msgs__msg__VehicleAttitudeSetpoint)(pointer_to_free))
}
func (t *VehicleAttitudeSetpoint) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__VehicleAttitudeSetpoint)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.roll_body = C.float(t.RollBody)
	mem.pitch_body = C.float(t.PitchBody)
	mem.yaw_body = C.float(t.YawBody)
	mem.yaw_sp_move_rate = C.float(t.YawSpMoveRate)
	cSlice_q_d := mem.q_d[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_q_d)), t.QD[:])
	cSlice_thrust_body := mem.thrust_body[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_thrust_body)), t.ThrustBody[:])
	mem.roll_reset_integral = C.bool(t.RollResetIntegral)
	mem.pitch_reset_integral = C.bool(t.PitchResetIntegral)
	mem.yaw_reset_integral = C.bool(t.YawResetIntegral)
	mem.fw_control_yaw = C.bool(t.FwControlYaw)
	mem.apply_flaps = C.uint8_t(t.ApplyFlaps)
	return unsafe.Pointer(mem)
}
func (t *VehicleAttitudeSetpoint) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__VehicleAttitudeSetpoint)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.RollBody = float32(mem.roll_body)
	t.PitchBody = float32(mem.pitch_body)
	t.YawBody = float32(mem.yaw_body)
	t.YawSpMoveRate = float32(mem.yaw_sp_move_rate)
	cSlice_q_d := mem.q_d[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.QD[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_q_d)))
	cSlice_thrust_body := mem.thrust_body[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.ThrustBody[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_thrust_body)))
	t.RollResetIntegral = bool(mem.roll_reset_integral)
	t.PitchResetIntegral = bool(mem.pitch_reset_integral)
	t.YawResetIntegral = bool(mem.yaw_reset_integral)
	t.FwControlYaw = bool(mem.fw_control_yaw)
	t.ApplyFlaps = uint8(mem.apply_flaps)
}
func (t *VehicleAttitudeSetpoint) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CVehicleAttitudeSetpoint = C.px4_msgs__msg__VehicleAttitudeSetpoint
type CVehicleAttitudeSetpoint__Sequence = C.px4_msgs__msg__VehicleAttitudeSetpoint__Sequence

func VehicleAttitudeSetpoint__Sequence_to_Go(goSlice *[]VehicleAttitudeSetpoint, cSlice CVehicleAttitudeSetpoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleAttitudeSetpoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleAttitudeSetpoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAttitudeSetpoint * uintptr(i)),
		))
		(*goSlice)[i] = VehicleAttitudeSetpoint{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func VehicleAttitudeSetpoint__Sequence_to_C(cSlice *CVehicleAttitudeSetpoint__Sequence, goSlice []VehicleAttitudeSetpoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleAttitudeSetpoint)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleAttitudeSetpoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleAttitudeSetpoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAttitudeSetpoint * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__VehicleAttitudeSetpoint)(v.AsCStruct())
	}
}
func VehicleAttitudeSetpoint__Array_to_Go(goSlice []VehicleAttitudeSetpoint, cSlice []CVehicleAttitudeSetpoint) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleAttitudeSetpoint__Array_to_C(cSlice []CVehicleAttitudeSetpoint, goSlice []VehicleAttitudeSetpoint) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__VehicleAttitudeSetpoint)(goSlice[i].AsCStruct())
	}
}


