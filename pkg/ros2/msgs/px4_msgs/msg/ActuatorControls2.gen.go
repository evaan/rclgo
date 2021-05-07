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
#include <px4_msgs/msg/actuator_controls2.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/ActuatorControls2", &ActuatorControls2{})
}
const (
	ActuatorControls2_NUM_ACTUATOR_CONTROLS uint8 = 8
	ActuatorControls2_NUM_ACTUATOR_CONTROL_GROUPS uint8 = 6
	ActuatorControls2_INDEX_ROLL uint8 = 0
	ActuatorControls2_INDEX_PITCH uint8 = 1
	ActuatorControls2_INDEX_YAW uint8 = 2
	ActuatorControls2_INDEX_THROTTLE uint8 = 3
	ActuatorControls2_INDEX_FLAPS uint8 = 4
	ActuatorControls2_INDEX_SPOILERS uint8 = 5
	ActuatorControls2_INDEX_AIRBRAKES uint8 = 6
	ActuatorControls2_INDEX_LANDING_GEAR uint8 = 7
	ActuatorControls2_INDEX_GIMBAL_SHUTTER uint8 = 3
	ActuatorControls2_INDEX_CAMERA_ZOOM uint8 = 4
	ActuatorControls2_GROUP_INDEX_ATTITUDE uint8 = 0
	ActuatorControls2_GROUP_INDEX_ATTITUDE_ALTERNATE uint8 = 1
	ActuatorControls2_GROUP_INDEX_GIMBAL uint8 = 2
	ActuatorControls2_GROUP_INDEX_MANUAL_PASSTHROUGH uint8 = 3
	ActuatorControls2_GROUP_INDEX_ALLOCATED_PART1 uint8 = 4
	ActuatorControls2_GROUP_INDEX_ALLOCATED_PART2 uint8 = 5
	ActuatorControls2_GROUP_INDEX_PAYLOAD uint8 = 6
)

// Do not create instances of this type directly. Always use NewActuatorControls2
// function instead.
type ActuatorControls2 struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp the data this control response is based on was sampled
	Control [8]float32 `yaml:"control"`
}

// NewActuatorControls2 creates a new ActuatorControls2 with default values.
func NewActuatorControls2() *ActuatorControls2 {
	self := ActuatorControls2{}
	self.SetDefaults(nil)
	return &self
}

func (t *ActuatorControls2) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *ActuatorControls2) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__ActuatorControls2())
}
func (t *ActuatorControls2) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__ActuatorControls2
	return (unsafe.Pointer)(C.px4_msgs__msg__ActuatorControls2__create())
}
func (t *ActuatorControls2) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__ActuatorControls2__destroy((*C.px4_msgs__msg__ActuatorControls2)(pointer_to_free))
}
func (t *ActuatorControls2) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__ActuatorControls2)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_sample = C.uint64_t(t.TimestampSample)
	cSlice_control := mem.control[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_control)), t.Control[:])
	return unsafe.Pointer(mem)
}
func (t *ActuatorControls2) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__ActuatorControls2)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_control := mem.control[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Control[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_control)))
}
func (t *ActuatorControls2) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CActuatorControls2 = C.px4_msgs__msg__ActuatorControls2
type CActuatorControls2__Sequence = C.px4_msgs__msg__ActuatorControls2__Sequence

func ActuatorControls2__Sequence_to_Go(goSlice *[]ActuatorControls2, cSlice CActuatorControls2__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ActuatorControls2, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__ActuatorControls2__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControls2 * uintptr(i)),
		))
		(*goSlice)[i] = ActuatorControls2{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func ActuatorControls2__Sequence_to_C(cSlice *CActuatorControls2__Sequence, goSlice []ActuatorControls2) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__ActuatorControls2)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__ActuatorControls2 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__ActuatorControls2)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControls2 * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__ActuatorControls2)(v.AsCStruct())
	}
}
func ActuatorControls2__Array_to_Go(goSlice []ActuatorControls2, cSlice []CActuatorControls2) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func ActuatorControls2__Array_to_C(cSlice []CActuatorControls2, goSlice []ActuatorControls2) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__ActuatorControls2)(goSlice[i].AsCStruct())
	}
}


