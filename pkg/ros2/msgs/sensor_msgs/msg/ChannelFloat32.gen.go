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
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <sensor_msgs/msg/channel_float32.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("sensor_msgs/ChannelFloat32", &ChannelFloat32{})
}

// Do not create instances of this type directly. Always use NewChannelFloat32
// function instead.
type ChannelFloat32 struct {
	Name rosidl_runtime_c.String `yaml:"name"`// The channel name should give semantics of the channel (e.g."intensity" instead of "value").
	Values []float32 `yaml:"values"`// The values array should be 1-1 with the elements of the associatedPointCloud.
}

// NewChannelFloat32 creates a new ChannelFloat32 with default values.
func NewChannelFloat32() *ChannelFloat32 {
	self := ChannelFloat32{}
	self.SetDefaults(nil)
	return &self
}

func (t *ChannelFloat32) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Name.SetDefaults("")
	
	return t
}

func (t *ChannelFloat32) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__ChannelFloat32())
}
func (t *ChannelFloat32) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__ChannelFloat32
	return (unsafe.Pointer)(C.sensor_msgs__msg__ChannelFloat32__create())
}
func (t *ChannelFloat32) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__ChannelFloat32__destroy((*C.sensor_msgs__msg__ChannelFloat32)(pointer_to_free))
}
func (t *ChannelFloat32) AsCStruct() unsafe.Pointer {
	mem := (*C.sensor_msgs__msg__ChannelFloat32)(t.PrepareMemory())
	mem.name = *(*C.rosidl_runtime_c__String)(t.Name.AsCStruct())
	rosidl_runtime_c.Float32__Sequence_to_C((*rosidl_runtime_c.CFloat32__Sequence)(unsafe.Pointer(&mem.values)), t.Values)
	return unsafe.Pointer(mem)
}
func (t *ChannelFloat32) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.sensor_msgs__msg__ChannelFloat32)(ros2_message_buffer)
	t.Name.AsGoStruct(unsafe.Pointer(&mem.name))
	rosidl_runtime_c.Float32__Sequence_to_Go(&t.Values, *(*rosidl_runtime_c.CFloat32__Sequence)(unsafe.Pointer(&mem.values)))
}
func (t *ChannelFloat32) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CChannelFloat32 = C.sensor_msgs__msg__ChannelFloat32
type CChannelFloat32__Sequence = C.sensor_msgs__msg__ChannelFloat32__Sequence

func ChannelFloat32__Sequence_to_Go(goSlice *[]ChannelFloat32, cSlice CChannelFloat32__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ChannelFloat32, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__ChannelFloat32__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__ChannelFloat32 * uintptr(i)),
		))
		(*goSlice)[i] = ChannelFloat32{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func ChannelFloat32__Sequence_to_C(cSlice *CChannelFloat32__Sequence, goSlice []ChannelFloat32) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__ChannelFloat32)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__ChannelFloat32 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__ChannelFloat32)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__ChannelFloat32 * uintptr(i)),
		))
		*cIdx = *(*C.sensor_msgs__msg__ChannelFloat32)(v.AsCStruct())
	}
}
func ChannelFloat32__Array_to_Go(goSlice []ChannelFloat32, cSlice []CChannelFloat32) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func ChannelFloat32__Array_to_C(cSlice []CChannelFloat32, goSlice []ChannelFloat32) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.sensor_msgs__msg__ChannelFloat32)(goSlice[i].AsCStruct())
	}
}


