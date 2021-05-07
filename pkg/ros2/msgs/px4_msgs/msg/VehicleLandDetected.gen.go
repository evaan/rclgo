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
#include <px4_msgs/msg/vehicle_land_detected.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleLandDetected", &VehicleLandDetected{})
}

// Do not create instances of this type directly. Always use NewVehicleLandDetected
// function instead.
type VehicleLandDetected struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	AltMax float32 `yaml:"alt_max"`// maximum altitude in [m] that can be reached
	Freefall bool `yaml:"freefall"`// true if vehicle is currently in free-fall
	GroundContact bool `yaml:"ground_contact"`// true if vehicle has ground contact but is not landed (1. stage)
	MaybeLanded bool `yaml:"maybe_landed"`// true if the vehicle might have landed (2. stage)
	Landed bool `yaml:"landed"`// true if vehicle is currently landed on the ground (3. stage)
	InGroundEffect bool `yaml:"in_ground_effect"`// indicates if from the perspective of the landing detector the vehicle might be in ground effect (baro). This flag will become true if the vehicle is not moving horizontally and is descending (crude assumption that user is landing).
}

// NewVehicleLandDetected creates a new VehicleLandDetected with default values.
func NewVehicleLandDetected() *VehicleLandDetected {
	self := VehicleLandDetected{}
	self.SetDefaults(nil)
	return &self
}

func (t *VehicleLandDetected) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *VehicleLandDetected) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleLandDetected())
}
func (t *VehicleLandDetected) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleLandDetected
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleLandDetected__create())
}
func (t *VehicleLandDetected) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleLandDetected__destroy((*C.px4_msgs__msg__VehicleLandDetected)(pointer_to_free))
}
func (t *VehicleLandDetected) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__VehicleLandDetected)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.alt_max = C.float(t.AltMax)
	mem.freefall = C.bool(t.Freefall)
	mem.ground_contact = C.bool(t.GroundContact)
	mem.maybe_landed = C.bool(t.MaybeLanded)
	mem.landed = C.bool(t.Landed)
	mem.in_ground_effect = C.bool(t.InGroundEffect)
	return unsafe.Pointer(mem)
}
func (t *VehicleLandDetected) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__VehicleLandDetected)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.AltMax = float32(mem.alt_max)
	t.Freefall = bool(mem.freefall)
	t.GroundContact = bool(mem.ground_contact)
	t.MaybeLanded = bool(mem.maybe_landed)
	t.Landed = bool(mem.landed)
	t.InGroundEffect = bool(mem.in_ground_effect)
}
func (t *VehicleLandDetected) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CVehicleLandDetected = C.px4_msgs__msg__VehicleLandDetected
type CVehicleLandDetected__Sequence = C.px4_msgs__msg__VehicleLandDetected__Sequence

func VehicleLandDetected__Sequence_to_Go(goSlice *[]VehicleLandDetected, cSlice CVehicleLandDetected__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleLandDetected, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleLandDetected__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleLandDetected * uintptr(i)),
		))
		(*goSlice)[i] = VehicleLandDetected{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func VehicleLandDetected__Sequence_to_C(cSlice *CVehicleLandDetected__Sequence, goSlice []VehicleLandDetected) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleLandDetected)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleLandDetected * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleLandDetected)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleLandDetected * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__VehicleLandDetected)(v.AsCStruct())
	}
}
func VehicleLandDetected__Array_to_Go(goSlice []VehicleLandDetected, cSlice []CVehicleLandDetected) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleLandDetected__Array_to_C(cSlice []CVehicleLandDetected, goSlice []VehicleLandDetected) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__VehicleLandDetected)(goSlice[i].AsCStruct())
	}
}


