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
	builtin_interfaces "github.com/tiiuae/rclgo/pkg/ros2/msgs/builtin_interfaces/msg"
	geometry_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/geometry_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltrajectory_msgs__rosidl_typesupport_c -ltrajectory_msgs__rosidl_generator_c
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <trajectory_msgs/msg/multi_dof_joint_trajectory_point.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("trajectory_msgs/MultiDOFJointTrajectoryPoint", &MultiDOFJointTrajectoryPoint{})
}

// Do not create instances of this type directly. Always use NewMultiDOFJointTrajectoryPoint
// function instead.
type MultiDOFJointTrajectoryPoint struct {
	Transforms []geometry_msgs.Transform `yaml:"transforms"`// Each multi-dof joint can specify a transform (up to 6 DOF).
	Velocities []geometry_msgs.Twist `yaml:"velocities"`// There can be a velocity specified for the origin of the joint.
	Accelerations []geometry_msgs.Twist `yaml:"accelerations"`// There can be an acceleration specified for the origin of the joint.
	TimeFromStart builtin_interfaces.Duration `yaml:"time_from_start"`// Desired time from the trajectory start to arrive at this trajectory point.
}

// NewMultiDOFJointTrajectoryPoint creates a new MultiDOFJointTrajectoryPoint with default values.
func NewMultiDOFJointTrajectoryPoint() *MultiDOFJointTrajectoryPoint {
	self := MultiDOFJointTrajectoryPoint{}
	self.SetDefaults(nil)
	return &self
}

func (t *MultiDOFJointTrajectoryPoint) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.TimeFromStart.SetDefaults(nil)
	
	return t
}

func (t *MultiDOFJointTrajectoryPoint) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__trajectory_msgs__msg__MultiDOFJointTrajectoryPoint())
}
func (t *MultiDOFJointTrajectoryPoint) PrepareMemory() unsafe.Pointer { //returns *C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint
	return (unsafe.Pointer)(C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint__create())
}
func (t *MultiDOFJointTrajectoryPoint) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint__destroy((*C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint)(pointer_to_free))
}
func (t *MultiDOFJointTrajectoryPoint) AsCStruct() unsafe.Pointer {
	mem := (*C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint)(t.PrepareMemory())
	geometry_msgs.Transform__Sequence_to_C((*geometry_msgs.CTransform__Sequence)(unsafe.Pointer(&mem.transforms)), t.Transforms)
	geometry_msgs.Twist__Sequence_to_C((*geometry_msgs.CTwist__Sequence)(unsafe.Pointer(&mem.velocities)), t.Velocities)
	geometry_msgs.Twist__Sequence_to_C((*geometry_msgs.CTwist__Sequence)(unsafe.Pointer(&mem.accelerations)), t.Accelerations)
	mem.time_from_start = *(*C.builtin_interfaces__msg__Duration)(t.TimeFromStart.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *MultiDOFJointTrajectoryPoint) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint)(ros2_message_buffer)
	geometry_msgs.Transform__Sequence_to_Go(&t.Transforms, *(*geometry_msgs.CTransform__Sequence)(unsafe.Pointer(&mem.transforms)))
	geometry_msgs.Twist__Sequence_to_Go(&t.Velocities, *(*geometry_msgs.CTwist__Sequence)(unsafe.Pointer(&mem.velocities)))
	geometry_msgs.Twist__Sequence_to_Go(&t.Accelerations, *(*geometry_msgs.CTwist__Sequence)(unsafe.Pointer(&mem.accelerations)))
	t.TimeFromStart.AsGoStruct(unsafe.Pointer(&mem.time_from_start))
}
func (t *MultiDOFJointTrajectoryPoint) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CMultiDOFJointTrajectoryPoint = C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint
type CMultiDOFJointTrajectoryPoint__Sequence = C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint__Sequence

func MultiDOFJointTrajectoryPoint__Sequence_to_Go(goSlice *[]MultiDOFJointTrajectoryPoint, cSlice CMultiDOFJointTrajectoryPoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MultiDOFJointTrajectoryPoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_trajectory_msgs__msg__MultiDOFJointTrajectoryPoint * uintptr(i)),
		))
		(*goSlice)[i] = MultiDOFJointTrajectoryPoint{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func MultiDOFJointTrajectoryPoint__Sequence_to_C(cSlice *CMultiDOFJointTrajectoryPoint__Sequence, goSlice []MultiDOFJointTrajectoryPoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint)(C.malloc((C.size_t)(C.sizeof_struct_trajectory_msgs__msg__MultiDOFJointTrajectoryPoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_trajectory_msgs__msg__MultiDOFJointTrajectoryPoint * uintptr(i)),
		))
		*cIdx = *(*C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint)(v.AsCStruct())
	}
}
func MultiDOFJointTrajectoryPoint__Array_to_Go(goSlice []MultiDOFJointTrajectoryPoint, cSlice []CMultiDOFJointTrajectoryPoint) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func MultiDOFJointTrajectoryPoint__Array_to_C(cSlice []CMultiDOFJointTrajectoryPoint, goSlice []MultiDOFJointTrajectoryPoint) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.trajectory_msgs__msg__MultiDOFJointTrajectoryPoint)(goSlice[i].AsCStruct())
	}
}


