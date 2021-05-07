/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	std_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <sensor_msgs/msg/point_cloud2.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("sensor_msgs/PointCloud2", &PointCloud2{})
}

// Do not create instances of this type directly. Always use NewPointCloud2
// function instead.
type PointCloud2 struct {
	Header std_msgs.Header `yaml:"header"`// Time of sensor data acquisition, and the coordinate frame ID (for 3d points).
	Height uint32 `yaml:"height"`// 2D structure of the point cloud. If the cloud is unordered, height is1 and width is the length of the point cloud.
	Width uint32 `yaml:"width"`// 2D structure of the point cloud. If the cloud is unordered, height is1 and width is the length of the point cloud.
	Fields []PointField `yaml:"fields"`// Describes the channels and their layout in the binary data blob.
	IsBigendian bool `yaml:"is_bigendian"`// Is this data bigendian?
	PointStep uint32 `yaml:"point_step"`// Length of a point in bytes
	RowStep uint32 `yaml:"row_step"`// Length of a row in bytes
	Data []uint8 `yaml:"data"`// Actual point data, size is (row_step*height)
	IsDense bool `yaml:"is_dense"`// True if there are no invalid points
}

// NewPointCloud2 creates a new PointCloud2 with default values.
func NewPointCloud2() *PointCloud2 {
	self := PointCloud2{}
	self.SetDefaults(nil)
	return &self
}

func (t *PointCloud2) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	
	return t
}

func (t *PointCloud2) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__PointCloud2())
}
func (t *PointCloud2) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__PointCloud2
	return (unsafe.Pointer)(C.sensor_msgs__msg__PointCloud2__create())
}
func (t *PointCloud2) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__PointCloud2__destroy((*C.sensor_msgs__msg__PointCloud2)(pointer_to_free))
}
func (t *PointCloud2) AsCStruct() unsafe.Pointer {
	mem := (*C.sensor_msgs__msg__PointCloud2)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	mem.height = C.uint32_t(t.Height)
	mem.width = C.uint32_t(t.Width)
	PointField__Sequence_to_C(&mem.fields, t.Fields)
	mem.is_bigendian = C.bool(t.IsBigendian)
	mem.point_step = C.uint32_t(t.PointStep)
	mem.row_step = C.uint32_t(t.RowStep)
	rosidl_runtime_c.Uint8__Sequence_to_C((*rosidl_runtime_c.CUint8__Sequence)(unsafe.Pointer(&mem.data)), t.Data)
	mem.is_dense = C.bool(t.IsDense)
	return unsafe.Pointer(mem)
}
func (t *PointCloud2) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.sensor_msgs__msg__PointCloud2)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	t.Height = uint32(mem.height)
	t.Width = uint32(mem.width)
	PointField__Sequence_to_Go(&t.Fields, mem.fields)
	t.IsBigendian = bool(mem.is_bigendian)
	t.PointStep = uint32(mem.point_step)
	t.RowStep = uint32(mem.row_step)
	rosidl_runtime_c.Uint8__Sequence_to_Go(&t.Data, *(*rosidl_runtime_c.CUint8__Sequence)(unsafe.Pointer(&mem.data)))
	t.IsDense = bool(mem.is_dense)
}
func (t *PointCloud2) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CPointCloud2 = C.sensor_msgs__msg__PointCloud2
type CPointCloud2__Sequence = C.sensor_msgs__msg__PointCloud2__Sequence

func PointCloud2__Sequence_to_Go(goSlice *[]PointCloud2, cSlice CPointCloud2__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PointCloud2, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__PointCloud2__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__PointCloud2 * uintptr(i)),
		))
		(*goSlice)[i] = PointCloud2{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func PointCloud2__Sequence_to_C(cSlice *CPointCloud2__Sequence, goSlice []PointCloud2) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__PointCloud2)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__PointCloud2 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__PointCloud2)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__PointCloud2 * uintptr(i)),
		))
		*cIdx = *(*C.sensor_msgs__msg__PointCloud2)(v.AsCStruct())
	}
}
func PointCloud2__Array_to_Go(goSlice []PointCloud2, cSlice []CPointCloud2) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func PointCloud2__Array_to_C(cSlice []CPointCloud2, goSlice []PointCloud2) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.sensor_msgs__msg__PointCloud2)(goSlice[i].AsCStruct())
	}
}


