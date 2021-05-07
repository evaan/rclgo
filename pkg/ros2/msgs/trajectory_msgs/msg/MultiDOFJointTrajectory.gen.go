/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package trajectory_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	std_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltrajectory_msgs__rosidl_typesupport_c -ltrajectory_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <trajectory_msgs/msg/multi_dof_joint_trajectory.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("trajectory_msgs/MultiDOFJointTrajectory", &MultiDOFJointTrajectory{})
}

// Do not create instances of this type directly. Always use NewMultiDOFJointTrajectory
// function instead.
type MultiDOFJointTrajectory struct {
	Header std_msgs.Header `yaml:"header"`// The header is used to specify the coordinate frame and the reference time for the trajectory durations
	JointNames []rosidl_runtime_c.String `yaml:"joint_names"`
	Points []MultiDOFJointTrajectoryPoint `yaml:"points"`
}

// NewMultiDOFJointTrajectory creates a new MultiDOFJointTrajectory with default values.
func NewMultiDOFJointTrajectory() *MultiDOFJointTrajectory {
	self := MultiDOFJointTrajectory{}
	self.SetDefaults(nil)
	return &self
}

func (t *MultiDOFJointTrajectory) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	
	return t
}

func (t *MultiDOFJointTrajectory) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__trajectory_msgs__msg__MultiDOFJointTrajectory())
}
func (t *MultiDOFJointTrajectory) PrepareMemory() unsafe.Pointer { //returns *C.trajectory_msgs__msg__MultiDOFJointTrajectory
	return (unsafe.Pointer)(C.trajectory_msgs__msg__MultiDOFJointTrajectory__create())
}
func (t *MultiDOFJointTrajectory) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.trajectory_msgs__msg__MultiDOFJointTrajectory__destroy((*C.trajectory_msgs__msg__MultiDOFJointTrajectory)(pointer_to_free))
}
func (t *MultiDOFJointTrajectory) AsCStruct() unsafe.Pointer {
	mem := (*C.trajectory_msgs__msg__MultiDOFJointTrajectory)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	rosidl_runtime_c.String__Sequence_to_C((*rosidl_runtime_c.CString__Sequence)(unsafe.Pointer(&mem.joint_names)), t.JointNames)
	MultiDOFJointTrajectoryPoint__Sequence_to_C(&mem.points, t.Points)
	return unsafe.Pointer(mem)
}
func (t *MultiDOFJointTrajectory) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.trajectory_msgs__msg__MultiDOFJointTrajectory)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	rosidl_runtime_c.String__Sequence_to_Go(&t.JointNames, *(*rosidl_runtime_c.CString__Sequence)(unsafe.Pointer(&mem.joint_names)))
	MultiDOFJointTrajectoryPoint__Sequence_to_Go(&t.Points, mem.points)
}
func (t *MultiDOFJointTrajectory) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CMultiDOFJointTrajectory = C.trajectory_msgs__msg__MultiDOFJointTrajectory
type CMultiDOFJointTrajectory__Sequence = C.trajectory_msgs__msg__MultiDOFJointTrajectory__Sequence

func MultiDOFJointTrajectory__Sequence_to_Go(goSlice *[]MultiDOFJointTrajectory, cSlice CMultiDOFJointTrajectory__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MultiDOFJointTrajectory, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.trajectory_msgs__msg__MultiDOFJointTrajectory__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_trajectory_msgs__msg__MultiDOFJointTrajectory * uintptr(i)),
		))
		(*goSlice)[i] = MultiDOFJointTrajectory{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func MultiDOFJointTrajectory__Sequence_to_C(cSlice *CMultiDOFJointTrajectory__Sequence, goSlice []MultiDOFJointTrajectory) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.trajectory_msgs__msg__MultiDOFJointTrajectory)(C.malloc((C.size_t)(C.sizeof_struct_trajectory_msgs__msg__MultiDOFJointTrajectory * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.trajectory_msgs__msg__MultiDOFJointTrajectory)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_trajectory_msgs__msg__MultiDOFJointTrajectory * uintptr(i)),
		))
		*cIdx = *(*C.trajectory_msgs__msg__MultiDOFJointTrajectory)(v.AsCStruct())
	}
}
func MultiDOFJointTrajectory__Array_to_Go(goSlice []MultiDOFJointTrajectory, cSlice []CMultiDOFJointTrajectory) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func MultiDOFJointTrajectory__Array_to_C(cSlice []CMultiDOFJointTrajectory, goSlice []MultiDOFJointTrajectory) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.trajectory_msgs__msg__MultiDOFJointTrajectory)(goSlice[i].AsCStruct())
	}
}


