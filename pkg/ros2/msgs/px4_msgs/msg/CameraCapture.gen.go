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
#include <px4_msgs/msg/camera_capture.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/CameraCapture", &CameraCapture{})
}

// Do not create instances of this type directly. Always use NewCameraCapture
// function instead.
type CameraCapture struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampUtc uint64 `yaml:"timestamp_utc"`// Capture time in UTC / GPS time
	Seq uint32 `yaml:"seq"`// Image sequence number
	Lat float64 `yaml:"lat"`// Latitude in degrees (WGS84)
	Lon float64 `yaml:"lon"`// Longitude in degrees (WGS84)
	Alt float32 `yaml:"alt"`// Altitude (AMSL)
	GroundDistance float32 `yaml:"ground_distance"`// Altitude above ground (meters)
	Q [4]float32 `yaml:"q"`// Attitude of the camera, zero rotation is facing towards front of vehicle
	Result int8 `yaml:"result"`// 1 for success, 0 for failure, -1 if camera does not provide feedback
}

// NewCameraCapture creates a new CameraCapture with default values.
func NewCameraCapture() *CameraCapture {
	self := CameraCapture{}
	self.SetDefaults(nil)
	return &self
}

func (t *CameraCapture) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *CameraCapture) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__CameraCapture())
}
func (t *CameraCapture) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__CameraCapture
	return (unsafe.Pointer)(C.px4_msgs__msg__CameraCapture__create())
}
func (t *CameraCapture) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__CameraCapture__destroy((*C.px4_msgs__msg__CameraCapture)(pointer_to_free))
}
func (t *CameraCapture) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__CameraCapture)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_utc = C.uint64_t(t.TimestampUtc)
	mem.seq = C.uint32_t(t.Seq)
	mem.lat = C.double(t.Lat)
	mem.lon = C.double(t.Lon)
	mem.alt = C.float(t.Alt)
	mem.ground_distance = C.float(t.GroundDistance)
	cSlice_q := mem.q[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_q)), t.Q[:])
	mem.result = C.int8_t(t.Result)
	return unsafe.Pointer(mem)
}
func (t *CameraCapture) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__CameraCapture)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampUtc = uint64(mem.timestamp_utc)
	t.Seq = uint32(mem.seq)
	t.Lat = float64(mem.lat)
	t.Lon = float64(mem.lon)
	t.Alt = float32(mem.alt)
	t.GroundDistance = float32(mem.ground_distance)
	cSlice_q := mem.q[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Q[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_q)))
	t.Result = int8(mem.result)
}
func (t *CameraCapture) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CCameraCapture = C.px4_msgs__msg__CameraCapture
type CCameraCapture__Sequence = C.px4_msgs__msg__CameraCapture__Sequence

func CameraCapture__Sequence_to_Go(goSlice *[]CameraCapture, cSlice CCameraCapture__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CameraCapture, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__CameraCapture__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CameraCapture * uintptr(i)),
		))
		(*goSlice)[i] = CameraCapture{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func CameraCapture__Sequence_to_C(cSlice *CCameraCapture__Sequence, goSlice []CameraCapture) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__CameraCapture)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__CameraCapture * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__CameraCapture)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CameraCapture * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__CameraCapture)(v.AsCStruct())
	}
}
func CameraCapture__Array_to_Go(goSlice []CameraCapture, cSlice []CCameraCapture) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func CameraCapture__Array_to_C(cSlice []CCameraCapture, goSlice []CameraCapture) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__CameraCapture)(goSlice[i].AsCStruct())
	}
}


