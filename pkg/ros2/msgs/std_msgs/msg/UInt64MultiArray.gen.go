/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package std_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <std_msgs/msg/u_int64_multi_array.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("std_msgs/UInt64MultiArray", &UInt64MultiArray{})
}

// Do not create instances of this type directly. Always use NewUInt64MultiArray
// function instead.
type UInt64MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []uint64 `yaml:"data"`// array of data
}

// NewUInt64MultiArray creates a new UInt64MultiArray with default values.
func NewUInt64MultiArray() *UInt64MultiArray {
	self := UInt64MultiArray{}
	self.SetDefaults(nil)
	return &self
}

func (t *UInt64MultiArray) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Layout.SetDefaults(nil)
	
	return t
}

func (t *UInt64MultiArray) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__UInt64MultiArray())
}
func (t *UInt64MultiArray) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__UInt64MultiArray
	return (unsafe.Pointer)(C.std_msgs__msg__UInt64MultiArray__create())
}
func (t *UInt64MultiArray) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__UInt64MultiArray__destroy((*C.std_msgs__msg__UInt64MultiArray)(pointer_to_free))
}
func (t *UInt64MultiArray) AsCStruct() unsafe.Pointer {
	mem := (*C.std_msgs__msg__UInt64MultiArray)(t.PrepareMemory())
	mem.layout = *(*C.std_msgs__msg__MultiArrayLayout)(t.Layout.AsCStruct())
	rosidl_runtime_c.Uint64__Sequence_to_C((*rosidl_runtime_c.CUint64__Sequence)(unsafe.Pointer(&mem.data)), t.Data)
	return unsafe.Pointer(mem)
}
func (t *UInt64MultiArray) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.std_msgs__msg__UInt64MultiArray)(ros2_message_buffer)
	t.Layout.AsGoStruct(unsafe.Pointer(&mem.layout))
	rosidl_runtime_c.Uint64__Sequence_to_Go(&t.Data, *(*rosidl_runtime_c.CUint64__Sequence)(unsafe.Pointer(&mem.data)))
}
func (t *UInt64MultiArray) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CUInt64MultiArray = C.std_msgs__msg__UInt64MultiArray
type CUInt64MultiArray__Sequence = C.std_msgs__msg__UInt64MultiArray__Sequence

func UInt64MultiArray__Sequence_to_Go(goSlice *[]UInt64MultiArray, cSlice CUInt64MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt64MultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__UInt64MultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt64MultiArray * uintptr(i)),
		))
		(*goSlice)[i] = UInt64MultiArray{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func UInt64MultiArray__Sequence_to_C(cSlice *CUInt64MultiArray__Sequence, goSlice []UInt64MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__UInt64MultiArray)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__UInt64MultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__UInt64MultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt64MultiArray * uintptr(i)),
		))
		*cIdx = *(*C.std_msgs__msg__UInt64MultiArray)(v.AsCStruct())
	}
}
func UInt64MultiArray__Array_to_Go(goSlice []UInt64MultiArray, cSlice []CUInt64MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func UInt64MultiArray__Array_to_C(cSlice []CUInt64MultiArray, goSlice []UInt64MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.std_msgs__msg__UInt64MultiArray)(goSlice[i].AsCStruct())
	}
}


