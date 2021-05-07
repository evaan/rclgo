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
#include <px4_msgs/msg/vehicle_control_mode.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleControlMode", &VehicleControlMode{})
}

// Do not create instances of this type directly. Always use NewVehicleControlMode
// function instead.
type VehicleControlMode struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	FlagArmed bool `yaml:"flag_armed"`// synonym for actuator_armed.armed
	FlagExternalManualOverrideOk bool `yaml:"flag_external_manual_override_ok"`// external override non-fatal for system. Only true for fixed wing
	FlagControlManualEnabled bool `yaml:"flag_control_manual_enabled"`// true if manual input is mixed in
	FlagControlAutoEnabled bool `yaml:"flag_control_auto_enabled"`// true if onboard autopilot should act
	FlagControlOffboardEnabled bool `yaml:"flag_control_offboard_enabled"`// true if offboard control should be used
	FlagControlRatesEnabled bool `yaml:"flag_control_rates_enabled"`// true if rates are stabilized
	FlagControlAttitudeEnabled bool `yaml:"flag_control_attitude_enabled"`// true if attitude stabilization is mixed in
	FlagControlAccelerationEnabled bool `yaml:"flag_control_acceleration_enabled"`// true if acceleration is controlled
	FlagControlVelocityEnabled bool `yaml:"flag_control_velocity_enabled"`// true if horizontal velocity (implies direction) is controlled
	FlagControlPositionEnabled bool `yaml:"flag_control_position_enabled"`// true if position is controlled
	FlagControlAltitudeEnabled bool `yaml:"flag_control_altitude_enabled"`// true if altitude is controlled
	FlagControlClimbRateEnabled bool `yaml:"flag_control_climb_rate_enabled"`// true if climb rate is controlled
	FlagControlTerminationEnabled bool `yaml:"flag_control_termination_enabled"`// true if flighttermination is enabled
}

// NewVehicleControlMode creates a new VehicleControlMode with default values.
func NewVehicleControlMode() *VehicleControlMode {
	self := VehicleControlMode{}
	self.SetDefaults(nil)
	return &self
}

func (t *VehicleControlMode) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *VehicleControlMode) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleControlMode())
}
func (t *VehicleControlMode) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleControlMode
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleControlMode__create())
}
func (t *VehicleControlMode) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleControlMode__destroy((*C.px4_msgs__msg__VehicleControlMode)(pointer_to_free))
}
func (t *VehicleControlMode) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__VehicleControlMode)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.flag_armed = C.bool(t.FlagArmed)
	mem.flag_external_manual_override_ok = C.bool(t.FlagExternalManualOverrideOk)
	mem.flag_control_manual_enabled = C.bool(t.FlagControlManualEnabled)
	mem.flag_control_auto_enabled = C.bool(t.FlagControlAutoEnabled)
	mem.flag_control_offboard_enabled = C.bool(t.FlagControlOffboardEnabled)
	mem.flag_control_rates_enabled = C.bool(t.FlagControlRatesEnabled)
	mem.flag_control_attitude_enabled = C.bool(t.FlagControlAttitudeEnabled)
	mem.flag_control_acceleration_enabled = C.bool(t.FlagControlAccelerationEnabled)
	mem.flag_control_velocity_enabled = C.bool(t.FlagControlVelocityEnabled)
	mem.flag_control_position_enabled = C.bool(t.FlagControlPositionEnabled)
	mem.flag_control_altitude_enabled = C.bool(t.FlagControlAltitudeEnabled)
	mem.flag_control_climb_rate_enabled = C.bool(t.FlagControlClimbRateEnabled)
	mem.flag_control_termination_enabled = C.bool(t.FlagControlTerminationEnabled)
	return unsafe.Pointer(mem)
}
func (t *VehicleControlMode) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__VehicleControlMode)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.FlagArmed = bool(mem.flag_armed)
	t.FlagExternalManualOverrideOk = bool(mem.flag_external_manual_override_ok)
	t.FlagControlManualEnabled = bool(mem.flag_control_manual_enabled)
	t.FlagControlAutoEnabled = bool(mem.flag_control_auto_enabled)
	t.FlagControlOffboardEnabled = bool(mem.flag_control_offboard_enabled)
	t.FlagControlRatesEnabled = bool(mem.flag_control_rates_enabled)
	t.FlagControlAttitudeEnabled = bool(mem.flag_control_attitude_enabled)
	t.FlagControlAccelerationEnabled = bool(mem.flag_control_acceleration_enabled)
	t.FlagControlVelocityEnabled = bool(mem.flag_control_velocity_enabled)
	t.FlagControlPositionEnabled = bool(mem.flag_control_position_enabled)
	t.FlagControlAltitudeEnabled = bool(mem.flag_control_altitude_enabled)
	t.FlagControlClimbRateEnabled = bool(mem.flag_control_climb_rate_enabled)
	t.FlagControlTerminationEnabled = bool(mem.flag_control_termination_enabled)
}
func (t *VehicleControlMode) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CVehicleControlMode = C.px4_msgs__msg__VehicleControlMode
type CVehicleControlMode__Sequence = C.px4_msgs__msg__VehicleControlMode__Sequence

func VehicleControlMode__Sequence_to_Go(goSlice *[]VehicleControlMode, cSlice CVehicleControlMode__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleControlMode, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleControlMode__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleControlMode * uintptr(i)),
		))
		(*goSlice)[i] = VehicleControlMode{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func VehicleControlMode__Sequence_to_C(cSlice *CVehicleControlMode__Sequence, goSlice []VehicleControlMode) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleControlMode)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleControlMode * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleControlMode)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleControlMode * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__VehicleControlMode)(v.AsCStruct())
	}
}
func VehicleControlMode__Array_to_Go(goSlice []VehicleControlMode, cSlice []CVehicleControlMode) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleControlMode__Array_to_C(cSlice []CVehicleControlMode, goSlice []VehicleControlMode) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__VehicleControlMode)(goSlice[i].AsCStruct())
	}
}


