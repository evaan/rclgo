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
#include <px4_msgs/msg/camera_trigger.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/CameraTrigger", &CameraTrigger{})
}

// Do not create instances of this type directly. Always use NewCameraTrigger
// function instead.
type CameraTrigger struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampUtc uint64 `yaml:"timestamp_utc"`// UTC timestamp
	Seq uint32 `yaml:"seq"`// Image sequence number
	Feedback bool `yaml:"feedback"`// Trigger feedback from camera
}

// NewCameraTrigger creates a new CameraTrigger with default values.
func NewCameraTrigger() *CameraTrigger {
	self := CameraTrigger{}
	self.SetDefaults(nil)
	return &self
}

func (t *CameraTrigger) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *CameraTrigger) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__CameraTrigger())
}
func (t *CameraTrigger) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__CameraTrigger
	return (unsafe.Pointer)(C.px4_msgs__msg__CameraTrigger__create())
}
func (t *CameraTrigger) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__CameraTrigger__destroy((*C.px4_msgs__msg__CameraTrigger)(pointer_to_free))
}
func (t *CameraTrigger) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__CameraTrigger)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_utc = C.uint64_t(t.TimestampUtc)
	mem.seq = C.uint32_t(t.Seq)
	mem.feedback = C.bool(t.Feedback)
	return unsafe.Pointer(mem)
}
func (t *CameraTrigger) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__CameraTrigger)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampUtc = uint64(mem.timestamp_utc)
	t.Seq = uint32(mem.seq)
	t.Feedback = bool(mem.feedback)
}
func (t *CameraTrigger) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CCameraTrigger = C.px4_msgs__msg__CameraTrigger
type CCameraTrigger__Sequence = C.px4_msgs__msg__CameraTrigger__Sequence

func CameraTrigger__Sequence_to_Go(goSlice *[]CameraTrigger, cSlice CCameraTrigger__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CameraTrigger, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__CameraTrigger__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CameraTrigger * uintptr(i)),
		))
		(*goSlice)[i] = CameraTrigger{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func CameraTrigger__Sequence_to_C(cSlice *CCameraTrigger__Sequence, goSlice []CameraTrigger) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__CameraTrigger)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__CameraTrigger * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__CameraTrigger)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CameraTrigger * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__CameraTrigger)(v.AsCStruct())
	}
}
func CameraTrigger__Array_to_Go(goSlice []CameraTrigger, cSlice []CCameraTrigger) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func CameraTrigger__Array_to_C(cSlice []CCameraTrigger, goSlice []CameraTrigger) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__CameraTrigger)(goSlice[i].AsCStruct())
	}
}


