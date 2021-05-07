/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package actionlib_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	std_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lactionlib_msgs__rosidl_typesupport_c -lactionlib_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <actionlib_msgs/msg/goal_status_array.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("actionlib_msgs/GoalStatusArray", &GoalStatusArray{})
}

// Do not create instances of this type directly. Always use NewGoalStatusArray
// function instead.
type GoalStatusArray struct {
	Header std_msgs.Header `yaml:"header"`// Stores the statuses for goals that are currently being trackedby an action server
	StatusList []GoalStatus `yaml:"status_list"`// Stores the statuses for goals that are currently being trackedby an action server
}

// NewGoalStatusArray creates a new GoalStatusArray with default values.
func NewGoalStatusArray() *GoalStatusArray {
	self := GoalStatusArray{}
	self.SetDefaults(nil)
	return &self
}

func (t *GoalStatusArray) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	
	return t
}

func (t *GoalStatusArray) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__actionlib_msgs__msg__GoalStatusArray())
}
func (t *GoalStatusArray) PrepareMemory() unsafe.Pointer { //returns *C.actionlib_msgs__msg__GoalStatusArray
	return (unsafe.Pointer)(C.actionlib_msgs__msg__GoalStatusArray__create())
}
func (t *GoalStatusArray) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.actionlib_msgs__msg__GoalStatusArray__destroy((*C.actionlib_msgs__msg__GoalStatusArray)(pointer_to_free))
}
func (t *GoalStatusArray) AsCStruct() unsafe.Pointer {
	mem := (*C.actionlib_msgs__msg__GoalStatusArray)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	GoalStatus__Sequence_to_C(&mem.status_list, t.StatusList)
	return unsafe.Pointer(mem)
}
func (t *GoalStatusArray) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.actionlib_msgs__msg__GoalStatusArray)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	GoalStatus__Sequence_to_Go(&t.StatusList, mem.status_list)
}
func (t *GoalStatusArray) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CGoalStatusArray = C.actionlib_msgs__msg__GoalStatusArray
type CGoalStatusArray__Sequence = C.actionlib_msgs__msg__GoalStatusArray__Sequence

func GoalStatusArray__Sequence_to_Go(goSlice *[]GoalStatusArray, cSlice CGoalStatusArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GoalStatusArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.actionlib_msgs__msg__GoalStatusArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_actionlib_msgs__msg__GoalStatusArray * uintptr(i)),
		))
		(*goSlice)[i] = GoalStatusArray{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func GoalStatusArray__Sequence_to_C(cSlice *CGoalStatusArray__Sequence, goSlice []GoalStatusArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.actionlib_msgs__msg__GoalStatusArray)(C.malloc((C.size_t)(C.sizeof_struct_actionlib_msgs__msg__GoalStatusArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.actionlib_msgs__msg__GoalStatusArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_actionlib_msgs__msg__GoalStatusArray * uintptr(i)),
		))
		*cIdx = *(*C.actionlib_msgs__msg__GoalStatusArray)(v.AsCStruct())
	}
}
func GoalStatusArray__Array_to_Go(goSlice []GoalStatusArray, cSlice []CGoalStatusArray) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func GoalStatusArray__Array_to_C(cSlice []CGoalStatusArray, goSlice []GoalStatusArray) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.actionlib_msgs__msg__GoalStatusArray)(goSlice[i].AsCStruct())
	}
}


