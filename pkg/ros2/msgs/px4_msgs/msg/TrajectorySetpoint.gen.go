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
#include <px4_msgs/msg/trajectory_setpoint.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/TrajectorySetpoint", &TrajectorySetpoint{})
}

// Do not create instances of this type directly. Always use NewTrajectorySetpoint
// function instead.
type TrajectorySetpoint struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	X float32 `yaml:"x"`// in meters NED
	Y float32 `yaml:"y"`// in meters NED
	Z float32 `yaml:"z"`// in meters NED
	Yaw float32 `yaml:"yaw"`// in radians NED -PI..+PI
	Yawspeed float32 `yaml:"yawspeed"`// in radians/sec
	Vx float32 `yaml:"vx"`// in meters/sec
	Vy float32 `yaml:"vy"`// in meters/sec
	Vz float32 `yaml:"vz"`// in meters/sec
	Acceleration [3]float32 `yaml:"acceleration"`// in meters/sec^2
	Jerk [3]float32 `yaml:"jerk"`// in meters/sec^3
	Thrust [3]float32 `yaml:"thrust"`// normalized thrust vector in NED
}

// NewTrajectorySetpoint creates a new TrajectorySetpoint with default values.
func NewTrajectorySetpoint() *TrajectorySetpoint {
	self := TrajectorySetpoint{}
	self.SetDefaults(nil)
	return &self
}

func (t *TrajectorySetpoint) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *TrajectorySetpoint) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__TrajectorySetpoint())
}
func (t *TrajectorySetpoint) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__TrajectorySetpoint
	return (unsafe.Pointer)(C.px4_msgs__msg__TrajectorySetpoint__create())
}
func (t *TrajectorySetpoint) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__TrajectorySetpoint__destroy((*C.px4_msgs__msg__TrajectorySetpoint)(pointer_to_free))
}
func (t *TrajectorySetpoint) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__TrajectorySetpoint)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.x = C.float(t.X)
	mem.y = C.float(t.Y)
	mem.z = C.float(t.Z)
	mem.yaw = C.float(t.Yaw)
	mem.yawspeed = C.float(t.Yawspeed)
	mem.vx = C.float(t.Vx)
	mem.vy = C.float(t.Vy)
	mem.vz = C.float(t.Vz)
	cSlice_acceleration := mem.acceleration[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_acceleration)), t.Acceleration[:])
	cSlice_jerk := mem.jerk[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_jerk)), t.Jerk[:])
	cSlice_thrust := mem.thrust[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_thrust)), t.Thrust[:])
	return unsafe.Pointer(mem)
}
func (t *TrajectorySetpoint) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__TrajectorySetpoint)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.X = float32(mem.x)
	t.Y = float32(mem.y)
	t.Z = float32(mem.z)
	t.Yaw = float32(mem.yaw)
	t.Yawspeed = float32(mem.yawspeed)
	t.Vx = float32(mem.vx)
	t.Vy = float32(mem.vy)
	t.Vz = float32(mem.vz)
	cSlice_acceleration := mem.acceleration[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Acceleration[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_acceleration)))
	cSlice_jerk := mem.jerk[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Jerk[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_jerk)))
	cSlice_thrust := mem.thrust[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Thrust[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_thrust)))
}
func (t *TrajectorySetpoint) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CTrajectorySetpoint = C.px4_msgs__msg__TrajectorySetpoint
type CTrajectorySetpoint__Sequence = C.px4_msgs__msg__TrajectorySetpoint__Sequence

func TrajectorySetpoint__Sequence_to_Go(goSlice *[]TrajectorySetpoint, cSlice CTrajectorySetpoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TrajectorySetpoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__TrajectorySetpoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TrajectorySetpoint * uintptr(i)),
		))
		(*goSlice)[i] = TrajectorySetpoint{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func TrajectorySetpoint__Sequence_to_C(cSlice *CTrajectorySetpoint__Sequence, goSlice []TrajectorySetpoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__TrajectorySetpoint)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__TrajectorySetpoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__TrajectorySetpoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TrajectorySetpoint * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__TrajectorySetpoint)(v.AsCStruct())
	}
}
func TrajectorySetpoint__Array_to_Go(goSlice []TrajectorySetpoint, cSlice []CTrajectorySetpoint) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func TrajectorySetpoint__Array_to_C(cSlice []CTrajectorySetpoint, goSlice []TrajectorySetpoint) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__TrajectorySetpoint)(goSlice[i].AsCStruct())
	}
}


