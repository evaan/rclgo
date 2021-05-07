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
#include <px4_msgs/msg/heater_status.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/HeaterStatus", &HeaterStatus{})
}
const (
	HeaterStatus_MODE_GPIO uint8 = 1
	HeaterStatus_MODE_PX4IO uint8 = 2
)

// Do not create instances of this type directly. Always use NewHeaterStatus
// function instead.
type HeaterStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	DeviceId uint32 `yaml:"device_id"`
	HeaterOn bool `yaml:"heater_on"`
	TemperatureSensor float32 `yaml:"temperature_sensor"`
	TemperatureTarget float32 `yaml:"temperature_target"`
	ControllerPeriodUsec uint32 `yaml:"controller_period_usec"`
	ControllerTimeOnUsec uint32 `yaml:"controller_time_on_usec"`
	ProportionalValue float32 `yaml:"proportional_value"`
	IntegratorValue float32 `yaml:"integrator_value"`
	FeedForwardValue float32 `yaml:"feed_forward_value"`
	Mode uint8 `yaml:"mode"`
}

// NewHeaterStatus creates a new HeaterStatus with default values.
func NewHeaterStatus() *HeaterStatus {
	self := HeaterStatus{}
	self.SetDefaults(nil)
	return &self
}

func (t *HeaterStatus) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *HeaterStatus) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__HeaterStatus())
}
func (t *HeaterStatus) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__HeaterStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__HeaterStatus__create())
}
func (t *HeaterStatus) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__HeaterStatus__destroy((*C.px4_msgs__msg__HeaterStatus)(pointer_to_free))
}
func (t *HeaterStatus) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__HeaterStatus)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.device_id = C.uint32_t(t.DeviceId)
	mem.heater_on = C.bool(t.HeaterOn)
	mem.temperature_sensor = C.float(t.TemperatureSensor)
	mem.temperature_target = C.float(t.TemperatureTarget)
	mem.controller_period_usec = C.uint32_t(t.ControllerPeriodUsec)
	mem.controller_time_on_usec = C.uint32_t(t.ControllerTimeOnUsec)
	mem.proportional_value = C.float(t.ProportionalValue)
	mem.integrator_value = C.float(t.IntegratorValue)
	mem.feed_forward_value = C.float(t.FeedForwardValue)
	mem.mode = C.uint8_t(t.Mode)
	return unsafe.Pointer(mem)
}
func (t *HeaterStatus) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__HeaterStatus)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.DeviceId = uint32(mem.device_id)
	t.HeaterOn = bool(mem.heater_on)
	t.TemperatureSensor = float32(mem.temperature_sensor)
	t.TemperatureTarget = float32(mem.temperature_target)
	t.ControllerPeriodUsec = uint32(mem.controller_period_usec)
	t.ControllerTimeOnUsec = uint32(mem.controller_time_on_usec)
	t.ProportionalValue = float32(mem.proportional_value)
	t.IntegratorValue = float32(mem.integrator_value)
	t.FeedForwardValue = float32(mem.feed_forward_value)
	t.Mode = uint8(mem.mode)
}
func (t *HeaterStatus) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CHeaterStatus = C.px4_msgs__msg__HeaterStatus
type CHeaterStatus__Sequence = C.px4_msgs__msg__HeaterStatus__Sequence

func HeaterStatus__Sequence_to_Go(goSlice *[]HeaterStatus, cSlice CHeaterStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]HeaterStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__HeaterStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__HeaterStatus * uintptr(i)),
		))
		(*goSlice)[i] = HeaterStatus{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func HeaterStatus__Sequence_to_C(cSlice *CHeaterStatus__Sequence, goSlice []HeaterStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__HeaterStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__HeaterStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__HeaterStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__HeaterStatus * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__HeaterStatus)(v.AsCStruct())
	}
}
func HeaterStatus__Array_to_Go(goSlice []HeaterStatus, cSlice []CHeaterStatus) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func HeaterStatus__Array_to_C(cSlice []CHeaterStatus, goSlice []HeaterStatus) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__HeaterStatus)(goSlice[i].AsCStruct())
	}
}


