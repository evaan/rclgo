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
#include <px4_msgs/msg/task_stack_info.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/TaskStackInfo", &TaskStackInfo{})
}
const (
	TaskStackInfo_ORB_QUEUE_LENGTH uint8 = 2
)

// Do not create instances of this type directly. Always use NewTaskStackInfo
// function instead.
type TaskStackInfo struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	StackFree uint16 `yaml:"stack_free"`
	TaskName [24]byte `yaml:"task_name"`
}

// NewTaskStackInfo creates a new TaskStackInfo with default values.
func NewTaskStackInfo() *TaskStackInfo {
	self := TaskStackInfo{}
	self.SetDefaults(nil)
	return &self
}

func (t *TaskStackInfo) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *TaskStackInfo) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__TaskStackInfo())
}
func (t *TaskStackInfo) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__TaskStackInfo
	return (unsafe.Pointer)(C.px4_msgs__msg__TaskStackInfo__create())
}
func (t *TaskStackInfo) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__TaskStackInfo__destroy((*C.px4_msgs__msg__TaskStackInfo)(pointer_to_free))
}
func (t *TaskStackInfo) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__TaskStackInfo)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.stack_free = C.uint16_t(t.StackFree)
	cSlice_task_name := mem.task_name[:]
	rosidl_runtime_c.Char__Array_to_C(*(*[]rosidl_runtime_c.CChar)(unsafe.Pointer(&cSlice_task_name)), t.TaskName[:])
	return unsafe.Pointer(mem)
}
func (t *TaskStackInfo) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__TaskStackInfo)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.StackFree = uint16(mem.stack_free)
	cSlice_task_name := mem.task_name[:]
	rosidl_runtime_c.Char__Array_to_Go(t.TaskName[:], *(*[]rosidl_runtime_c.CChar)(unsafe.Pointer(&cSlice_task_name)))
}
func (t *TaskStackInfo) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CTaskStackInfo = C.px4_msgs__msg__TaskStackInfo
type CTaskStackInfo__Sequence = C.px4_msgs__msg__TaskStackInfo__Sequence

func TaskStackInfo__Sequence_to_Go(goSlice *[]TaskStackInfo, cSlice CTaskStackInfo__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TaskStackInfo, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__TaskStackInfo__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TaskStackInfo * uintptr(i)),
		))
		(*goSlice)[i] = TaskStackInfo{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func TaskStackInfo__Sequence_to_C(cSlice *CTaskStackInfo__Sequence, goSlice []TaskStackInfo) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__TaskStackInfo)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__TaskStackInfo * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__TaskStackInfo)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TaskStackInfo * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__TaskStackInfo)(v.AsCStruct())
	}
}
func TaskStackInfo__Array_to_Go(goSlice []TaskStackInfo, cSlice []CTaskStackInfo) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func TaskStackInfo__Array_to_C(cSlice []CTaskStackInfo, goSlice []TaskStackInfo) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__TaskStackInfo)(goSlice[i].AsCStruct())
	}
}


