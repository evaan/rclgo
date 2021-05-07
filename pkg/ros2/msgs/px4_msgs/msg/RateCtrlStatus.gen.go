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
#include <px4_msgs/msg/rate_ctrl_status.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/RateCtrlStatus", &RateCtrlStatus{})
}

// Do not create instances of this type directly. Always use NewRateCtrlStatus
// function instead.
type RateCtrlStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	RollspeedInteg float32 `yaml:"rollspeed_integ"`// rate controller integrator status
	PitchspeedInteg float32 `yaml:"pitchspeed_integ"`// rate controller integrator status
	YawspeedInteg float32 `yaml:"yawspeed_integ"`// rate controller integrator status
	AdditionalInteg1 float32 `yaml:"additional_integ1"`// FW: wheel rate integrator (optional). rate controller integrator status
}

// NewRateCtrlStatus creates a new RateCtrlStatus with default values.
func NewRateCtrlStatus() *RateCtrlStatus {
	self := RateCtrlStatus{}
	self.SetDefaults(nil)
	return &self
}

func (t *RateCtrlStatus) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *RateCtrlStatus) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__RateCtrlStatus())
}
func (t *RateCtrlStatus) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__RateCtrlStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__RateCtrlStatus__create())
}
func (t *RateCtrlStatus) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__RateCtrlStatus__destroy((*C.px4_msgs__msg__RateCtrlStatus)(pointer_to_free))
}
func (t *RateCtrlStatus) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__RateCtrlStatus)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.rollspeed_integ = C.float(t.RollspeedInteg)
	mem.pitchspeed_integ = C.float(t.PitchspeedInteg)
	mem.yawspeed_integ = C.float(t.YawspeedInteg)
	mem.additional_integ1 = C.float(t.AdditionalInteg1)
	return unsafe.Pointer(mem)
}
func (t *RateCtrlStatus) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__RateCtrlStatus)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.RollspeedInteg = float32(mem.rollspeed_integ)
	t.PitchspeedInteg = float32(mem.pitchspeed_integ)
	t.YawspeedInteg = float32(mem.yawspeed_integ)
	t.AdditionalInteg1 = float32(mem.additional_integ1)
}
func (t *RateCtrlStatus) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CRateCtrlStatus = C.px4_msgs__msg__RateCtrlStatus
type CRateCtrlStatus__Sequence = C.px4_msgs__msg__RateCtrlStatus__Sequence

func RateCtrlStatus__Sequence_to_Go(goSlice *[]RateCtrlStatus, cSlice CRateCtrlStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RateCtrlStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__RateCtrlStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__RateCtrlStatus * uintptr(i)),
		))
		(*goSlice)[i] = RateCtrlStatus{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func RateCtrlStatus__Sequence_to_C(cSlice *CRateCtrlStatus__Sequence, goSlice []RateCtrlStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__RateCtrlStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__RateCtrlStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__RateCtrlStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__RateCtrlStatus * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__RateCtrlStatus)(v.AsCStruct())
	}
}
func RateCtrlStatus__Array_to_Go(goSlice []RateCtrlStatus, cSlice []CRateCtrlStatus) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func RateCtrlStatus__Array_to_C(cSlice []CRateCtrlStatus, goSlice []RateCtrlStatus) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__RateCtrlStatus)(goSlice[i].AsCStruct())
	}
}


