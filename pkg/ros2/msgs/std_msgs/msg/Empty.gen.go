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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <std_msgs/msg/empty.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("std_msgs/Empty", &Empty{})
}

// Do not create instances of this type directly. Always use NewEmpty
// function instead.
type Empty struct {
}

// NewEmpty creates a new Empty with default values.
func NewEmpty() *Empty {
	self := Empty{}
	self.SetDefaults(nil)
	return &self
}

func (t *Empty) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *Empty) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__Empty())
}
func (t *Empty) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__Empty
	return (unsafe.Pointer)(C.std_msgs__msg__Empty__create())
}
func (t *Empty) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__Empty__destroy((*C.std_msgs__msg__Empty)(pointer_to_free))
}
func (t *Empty) AsCStruct() unsafe.Pointer {
	mem := (*C.std_msgs__msg__Empty)(t.PrepareMemory())
	return unsafe.Pointer(mem)
}
func (t *Empty) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	
}
func (t *Empty) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CEmpty = C.std_msgs__msg__Empty
type CEmpty__Sequence = C.std_msgs__msg__Empty__Sequence

func Empty__Sequence_to_Go(goSlice *[]Empty, cSlice CEmpty__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Empty, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__Empty__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Empty * uintptr(i)),
		))
		(*goSlice)[i] = Empty{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Empty__Sequence_to_C(cSlice *CEmpty__Sequence, goSlice []Empty) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__Empty)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__Empty * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__Empty)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Empty * uintptr(i)),
		))
		*cIdx = *(*C.std_msgs__msg__Empty)(v.AsCStruct())
	}
}
func Empty__Array_to_Go(goSlice []Empty, cSlice []CEmpty) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Empty__Array_to_C(cSlice []CEmpty, goSlice []Empty) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.std_msgs__msg__Empty)(goSlice[i].AsCStruct())
	}
}


