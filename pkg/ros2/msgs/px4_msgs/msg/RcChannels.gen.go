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
#include <px4_msgs/msg/rc_channels.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/RcChannels", &RcChannels{})
}
const (
	RcChannels_FUNCTION_THROTTLE uint8 = 0
	RcChannels_FUNCTION_ROLL uint8 = 1
	RcChannels_FUNCTION_PITCH uint8 = 2
	RcChannels_FUNCTION_YAW uint8 = 3
	RcChannels_FUNCTION_MODE uint8 = 4
	RcChannels_FUNCTION_RETURN uint8 = 5
	RcChannels_FUNCTION_POSCTL uint8 = 6
	RcChannels_FUNCTION_LOITER uint8 = 7
	RcChannels_FUNCTION_OFFBOARD uint8 = 8
	RcChannels_FUNCTION_ACRO uint8 = 9
	RcChannels_FUNCTION_FLAPS uint8 = 10
	RcChannels_FUNCTION_AUX_1 uint8 = 11
	RcChannels_FUNCTION_AUX_2 uint8 = 12
	RcChannels_FUNCTION_AUX_3 uint8 = 13
	RcChannels_FUNCTION_AUX_4 uint8 = 14
	RcChannels_FUNCTION_AUX_5 uint8 = 15
	RcChannels_FUNCTION_PARAM_1 uint8 = 16
	RcChannels_FUNCTION_PARAM_2 uint8 = 17
	RcChannels_FUNCTION_PARAM_3_5 uint8 = 18
	RcChannels_FUNCTION_KILLSWITCH uint8 = 19
	RcChannels_FUNCTION_TRANSITION uint8 = 20
	RcChannels_FUNCTION_GEAR uint8 = 21
	RcChannels_FUNCTION_ARMSWITCH uint8 = 22
	RcChannels_FUNCTION_STAB uint8 = 23
	RcChannels_FUNCTION_AUX_6 uint8 = 24
	RcChannels_FUNCTION_MAN uint8 = 25
)

// Do not create instances of this type directly. Always use NewRcChannels
// function instead.
type RcChannels struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampLastValid uint64 `yaml:"timestamp_last_valid"`// Timestamp of last valid RC signal
	Channels [18]float32 `yaml:"channels"`// Scaled to -1..1 (throttle: 0..1)
	ChannelCount uint8 `yaml:"channel_count"`// Number of valid channels
	Function [26]int8 `yaml:"function"`// Functions mapping
	Rssi uint8 `yaml:"rssi"`// Receive signal strength index
	SignalLost bool `yaml:"signal_lost"`// Control signal lost, should be checked together with topic timeout
	FrameDropCount uint32 `yaml:"frame_drop_count"`// Number of dropped frames
}

// NewRcChannels creates a new RcChannels with default values.
func NewRcChannels() *RcChannels {
	self := RcChannels{}
	self.SetDefaults(nil)
	return &self
}

func (t *RcChannels) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *RcChannels) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__RcChannels())
}
func (t *RcChannels) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__RcChannels
	return (unsafe.Pointer)(C.px4_msgs__msg__RcChannels__create())
}
func (t *RcChannels) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__RcChannels__destroy((*C.px4_msgs__msg__RcChannels)(pointer_to_free))
}
func (t *RcChannels) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__RcChannels)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_last_valid = C.uint64_t(t.TimestampLastValid)
	cSlice_channels := mem.channels[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_channels)), t.Channels[:])
	mem.channel_count = C.uint8_t(t.ChannelCount)
	cSlice_function := mem.function[:]
	rosidl_runtime_c.Int8__Array_to_C(*(*[]rosidl_runtime_c.CInt8)(unsafe.Pointer(&cSlice_function)), t.Function[:])
	mem.rssi = C.uint8_t(t.Rssi)
	mem.signal_lost = C.bool(t.SignalLost)
	mem.frame_drop_count = C.uint32_t(t.FrameDropCount)
	return unsafe.Pointer(mem)
}
func (t *RcChannels) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__RcChannels)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampLastValid = uint64(mem.timestamp_last_valid)
	cSlice_channels := mem.channels[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Channels[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_channels)))
	t.ChannelCount = uint8(mem.channel_count)
	cSlice_function := mem.function[:]
	rosidl_runtime_c.Int8__Array_to_Go(t.Function[:], *(*[]rosidl_runtime_c.CInt8)(unsafe.Pointer(&cSlice_function)))
	t.Rssi = uint8(mem.rssi)
	t.SignalLost = bool(mem.signal_lost)
	t.FrameDropCount = uint32(mem.frame_drop_count)
}
func (t *RcChannels) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CRcChannels = C.px4_msgs__msg__RcChannels
type CRcChannels__Sequence = C.px4_msgs__msg__RcChannels__Sequence

func RcChannels__Sequence_to_Go(goSlice *[]RcChannels, cSlice CRcChannels__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RcChannels, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__RcChannels__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__RcChannels * uintptr(i)),
		))
		(*goSlice)[i] = RcChannels{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func RcChannels__Sequence_to_C(cSlice *CRcChannels__Sequence, goSlice []RcChannels) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__RcChannels)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__RcChannels * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__RcChannels)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__RcChannels * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__RcChannels)(v.AsCStruct())
	}
}
func RcChannels__Array_to_Go(goSlice []RcChannels, cSlice []CRcChannels) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func RcChannels__Array_to_C(cSlice []CRcChannels, goSlice []RcChannels) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__RcChannels)(goSlice[i].AsCStruct())
	}
}


