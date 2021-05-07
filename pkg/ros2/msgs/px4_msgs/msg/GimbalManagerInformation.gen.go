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
#include <px4_msgs/msg/gimbal_manager_information.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/GimbalManagerInformation", &GimbalManagerInformation{})
}
const (
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_HAS_RETRACT uint32 = 1
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_HAS_NEUTRAL uint32 = 2
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_AXIS uint32 = 4
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_FOLLOW uint32 = 8
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_LOCK uint32 = 16
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_AXIS uint32 = 32
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_FOLLOW uint32 = 64
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_LOCK uint32 = 128
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_AXIS uint32 = 256
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_FOLLOW uint32 = 512
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_LOCK uint32 = 1024
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_INFINITE_YAW uint32 = 2048
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_LOCAL uint32 = 65536
	GimbalManagerInformation_GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_GLOBAL uint32 = 131072
)

// Do not create instances of this type directly. Always use NewGimbalManagerInformation
// function instead.
type GimbalManagerInformation struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	CapFlags uint32 `yaml:"cap_flags"`
	GimbalDeviceId uint8 `yaml:"gimbal_device_id"`
	RollMin float32 `yaml:"roll_min"`// [rad]
	RollMax float32 `yaml:"roll_max"`// [rad]
	PitchMin float32 `yaml:"pitch_min"`// [rad]
	PitchMax float32 `yaml:"pitch_max"`// [rad]
	YawMin float32 `yaml:"yaw_min"`// [rad]
	YawMax float32 `yaml:"yaw_max"`// [rad]
}

// NewGimbalManagerInformation creates a new GimbalManagerInformation with default values.
func NewGimbalManagerInformation() *GimbalManagerInformation {
	self := GimbalManagerInformation{}
	self.SetDefaults(nil)
	return &self
}

func (t *GimbalManagerInformation) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *GimbalManagerInformation) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__GimbalManagerInformation())
}
func (t *GimbalManagerInformation) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__GimbalManagerInformation
	return (unsafe.Pointer)(C.px4_msgs__msg__GimbalManagerInformation__create())
}
func (t *GimbalManagerInformation) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__GimbalManagerInformation__destroy((*C.px4_msgs__msg__GimbalManagerInformation)(pointer_to_free))
}
func (t *GimbalManagerInformation) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__GimbalManagerInformation)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.cap_flags = C.uint32_t(t.CapFlags)
	mem.gimbal_device_id = C.uint8_t(t.GimbalDeviceId)
	mem.roll_min = C.float(t.RollMin)
	mem.roll_max = C.float(t.RollMax)
	mem.pitch_min = C.float(t.PitchMin)
	mem.pitch_max = C.float(t.PitchMax)
	mem.yaw_min = C.float(t.YawMin)
	mem.yaw_max = C.float(t.YawMax)
	return unsafe.Pointer(mem)
}
func (t *GimbalManagerInformation) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__GimbalManagerInformation)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.CapFlags = uint32(mem.cap_flags)
	t.GimbalDeviceId = uint8(mem.gimbal_device_id)
	t.RollMin = float32(mem.roll_min)
	t.RollMax = float32(mem.roll_max)
	t.PitchMin = float32(mem.pitch_min)
	t.PitchMax = float32(mem.pitch_max)
	t.YawMin = float32(mem.yaw_min)
	t.YawMax = float32(mem.yaw_max)
}
func (t *GimbalManagerInformation) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CGimbalManagerInformation = C.px4_msgs__msg__GimbalManagerInformation
type CGimbalManagerInformation__Sequence = C.px4_msgs__msg__GimbalManagerInformation__Sequence

func GimbalManagerInformation__Sequence_to_Go(goSlice *[]GimbalManagerInformation, cSlice CGimbalManagerInformation__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GimbalManagerInformation, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__GimbalManagerInformation__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GimbalManagerInformation * uintptr(i)),
		))
		(*goSlice)[i] = GimbalManagerInformation{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func GimbalManagerInformation__Sequence_to_C(cSlice *CGimbalManagerInformation__Sequence, goSlice []GimbalManagerInformation) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__GimbalManagerInformation)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__GimbalManagerInformation * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__GimbalManagerInformation)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GimbalManagerInformation * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__GimbalManagerInformation)(v.AsCStruct())
	}
}
func GimbalManagerInformation__Array_to_Go(goSlice []GimbalManagerInformation, cSlice []CGimbalManagerInformation) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func GimbalManagerInformation__Array_to_C(cSlice []CGimbalManagerInformation, goSlice []GimbalManagerInformation) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__GimbalManagerInformation)(goSlice[i].AsCStruct())
	}
}


