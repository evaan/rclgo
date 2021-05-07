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
#include <px4_msgs/msg/vehicle_trajectory_waypoint_desired.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleTrajectoryWaypointDesired", &VehicleTrajectoryWaypointDesired{})
}
const (
	VehicleTrajectoryWaypointDesired_MAV_TRAJECTORY_REPRESENTATION_WAYPOINTS uint8 = 0
	VehicleTrajectoryWaypointDesired_POINT_0 uint8 = 0
	VehicleTrajectoryWaypointDesired_POINT_1 uint8 = 1
	VehicleTrajectoryWaypointDesired_POINT_2 uint8 = 2
	VehicleTrajectoryWaypointDesired_POINT_3 uint8 = 3
	VehicleTrajectoryWaypointDesired_POINT_4 uint8 = 4
	VehicleTrajectoryWaypointDesired_NUMBER_POINTS uint8 = 5
)

// Do not create instances of this type directly. Always use NewVehicleTrajectoryWaypointDesired
// function instead.
type VehicleTrajectoryWaypointDesired struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Type uint8 `yaml:"type"`// Type from MAV_TRAJECTORY_REPRESENTATION enum.
	Waypoints [5]TrajectoryWaypoint `yaml:"waypoints"`
}

// NewVehicleTrajectoryWaypointDesired creates a new VehicleTrajectoryWaypointDesired with default values.
func NewVehicleTrajectoryWaypointDesired() *VehicleTrajectoryWaypointDesired {
	self := VehicleTrajectoryWaypointDesired{}
	self.SetDefaults(nil)
	return &self
}

func (t *VehicleTrajectoryWaypointDesired) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Waypoints[0].SetDefaults(nil)
	t.Waypoints[1].SetDefaults(nil)
	t.Waypoints[2].SetDefaults(nil)
	t.Waypoints[3].SetDefaults(nil)
	t.Waypoints[4].SetDefaults(nil)
	
	return t
}

func (t *VehicleTrajectoryWaypointDesired) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleTrajectoryWaypointDesired())
}
func (t *VehicleTrajectoryWaypointDesired) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleTrajectoryWaypointDesired
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleTrajectoryWaypointDesired__create())
}
func (t *VehicleTrajectoryWaypointDesired) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleTrajectoryWaypointDesired__destroy((*C.px4_msgs__msg__VehicleTrajectoryWaypointDesired)(pointer_to_free))
}
func (t *VehicleTrajectoryWaypointDesired) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__VehicleTrajectoryWaypointDesired)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem._type = C.uint8_t(t.Type)
	TrajectoryWaypoint__Array_to_C(mem.waypoints[:], t.Waypoints[:])
	return unsafe.Pointer(mem)
}
func (t *VehicleTrajectoryWaypointDesired) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__VehicleTrajectoryWaypointDesired)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.Type = uint8(mem._type)
	TrajectoryWaypoint__Array_to_Go(t.Waypoints[:], mem.waypoints[:])
}
func (t *VehicleTrajectoryWaypointDesired) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CVehicleTrajectoryWaypointDesired = C.px4_msgs__msg__VehicleTrajectoryWaypointDesired
type CVehicleTrajectoryWaypointDesired__Sequence = C.px4_msgs__msg__VehicleTrajectoryWaypointDesired__Sequence

func VehicleTrajectoryWaypointDesired__Sequence_to_Go(goSlice *[]VehicleTrajectoryWaypointDesired, cSlice CVehicleTrajectoryWaypointDesired__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleTrajectoryWaypointDesired, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleTrajectoryWaypointDesired__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleTrajectoryWaypointDesired * uintptr(i)),
		))
		(*goSlice)[i] = VehicleTrajectoryWaypointDesired{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func VehicleTrajectoryWaypointDesired__Sequence_to_C(cSlice *CVehicleTrajectoryWaypointDesired__Sequence, goSlice []VehicleTrajectoryWaypointDesired) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleTrajectoryWaypointDesired)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleTrajectoryWaypointDesired * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleTrajectoryWaypointDesired)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleTrajectoryWaypointDesired * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__VehicleTrajectoryWaypointDesired)(v.AsCStruct())
	}
}
func VehicleTrajectoryWaypointDesired__Array_to_Go(goSlice []VehicleTrajectoryWaypointDesired, cSlice []CVehicleTrajectoryWaypointDesired) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleTrajectoryWaypointDesired__Array_to_C(cSlice []CVehicleTrajectoryWaypointDesired, goSlice []VehicleTrajectoryWaypointDesired) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__VehicleTrajectoryWaypointDesired)(goSlice[i].AsCStruct())
	}
}


