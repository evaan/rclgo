/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package builtin_interfaces
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <builtin_interfaces/msg/time.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("builtin_interfaces/Time", &Time{})
}

// Do not create instances of this type directly. Always use NewTime
// function instead.
type Time struct {
	Sec int32 `yaml:"sec"`// The seconds component, valid over all int32 values.
	Nanosec uint32 `yaml:"nanosec"`// The nanoseconds component, valid in the range [0, 10e9).
}

// NewTime creates a new Time with default values.
func NewTime() *Time {
	self := Time{}
	self.SetDefaults(nil)
	return &self
}

func (t *Time) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *Time) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__builtin_interfaces__msg__Time())
}
func (t *Time) PrepareMemory() unsafe.Pointer { //returns *C.builtin_interfaces__msg__Time
	return (unsafe.Pointer)(C.builtin_interfaces__msg__Time__create())
}
func (t *Time) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.builtin_interfaces__msg__Time__destroy((*C.builtin_interfaces__msg__Time)(pointer_to_free))
}
func (t *Time) AsCStruct() unsafe.Pointer {
	mem := (*C.builtin_interfaces__msg__Time)(t.PrepareMemory())
	mem.sec = C.int32_t(t.Sec)
	mem.nanosec = C.uint32_t(t.Nanosec)
	return unsafe.Pointer(mem)
}
func (t *Time) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.builtin_interfaces__msg__Time)(ros2_message_buffer)
	t.Sec = int32(mem.sec)
	t.Nanosec = uint32(mem.nanosec)
}
func (t *Time) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CTime = C.builtin_interfaces__msg__Time
type CTime__Sequence = C.builtin_interfaces__msg__Time__Sequence

func Time__Sequence_to_Go(goSlice *[]Time, cSlice CTime__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Time, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.builtin_interfaces__msg__Time__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_builtin_interfaces__msg__Time * uintptr(i)),
		))
		(*goSlice)[i] = Time{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Time__Sequence_to_C(cSlice *CTime__Sequence, goSlice []Time) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.builtin_interfaces__msg__Time)(C.malloc((C.size_t)(C.sizeof_struct_builtin_interfaces__msg__Time * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.builtin_interfaces__msg__Time)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_builtin_interfaces__msg__Time * uintptr(i)),
		))
		*cIdx = *(*C.builtin_interfaces__msg__Time)(v.AsCStruct())
	}
}
func Time__Array_to_Go(goSlice []Time, cSlice []CTime) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Time__Array_to_C(cSlice []CTime, goSlice []Time) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.builtin_interfaces__msg__Time)(goSlice[i].AsCStruct())
	}
}


