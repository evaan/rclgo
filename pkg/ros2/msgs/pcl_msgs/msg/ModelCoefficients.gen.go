/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package pcl_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	std_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpcl_msgs__rosidl_typesupport_c -lpcl_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <pcl_msgs/msg/model_coefficients.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("pcl_msgs/ModelCoefficients", &ModelCoefficients{})
}

// Do not create instances of this type directly. Always use NewModelCoefficients
// function instead.
type ModelCoefficients struct {
	Header std_msgs.Header `yaml:"header"`
	Values []float32 `yaml:"values"`
}

// NewModelCoefficients creates a new ModelCoefficients with default values.
func NewModelCoefficients() *ModelCoefficients {
	self := ModelCoefficients{}
	self.SetDefaults(nil)
	return &self
}

func (t *ModelCoefficients) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	
	return t
}

func (t *ModelCoefficients) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__pcl_msgs__msg__ModelCoefficients())
}
func (t *ModelCoefficients) PrepareMemory() unsafe.Pointer { //returns *C.pcl_msgs__msg__ModelCoefficients
	return (unsafe.Pointer)(C.pcl_msgs__msg__ModelCoefficients__create())
}
func (t *ModelCoefficients) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.pcl_msgs__msg__ModelCoefficients__destroy((*C.pcl_msgs__msg__ModelCoefficients)(pointer_to_free))
}
func (t *ModelCoefficients) AsCStruct() unsafe.Pointer {
	mem := (*C.pcl_msgs__msg__ModelCoefficients)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	rosidl_runtime_c.Float32__Sequence_to_C((*rosidl_runtime_c.CFloat32__Sequence)(unsafe.Pointer(&mem.values)), t.Values)
	return unsafe.Pointer(mem)
}
func (t *ModelCoefficients) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.pcl_msgs__msg__ModelCoefficients)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	rosidl_runtime_c.Float32__Sequence_to_Go(&t.Values, *(*rosidl_runtime_c.CFloat32__Sequence)(unsafe.Pointer(&mem.values)))
}
func (t *ModelCoefficients) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CModelCoefficients = C.pcl_msgs__msg__ModelCoefficients
type CModelCoefficients__Sequence = C.pcl_msgs__msg__ModelCoefficients__Sequence

func ModelCoefficients__Sequence_to_Go(goSlice *[]ModelCoefficients, cSlice CModelCoefficients__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ModelCoefficients, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.pcl_msgs__msg__ModelCoefficients__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pcl_msgs__msg__ModelCoefficients * uintptr(i)),
		))
		(*goSlice)[i] = ModelCoefficients{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func ModelCoefficients__Sequence_to_C(cSlice *CModelCoefficients__Sequence, goSlice []ModelCoefficients) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.pcl_msgs__msg__ModelCoefficients)(C.malloc((C.size_t)(C.sizeof_struct_pcl_msgs__msg__ModelCoefficients * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.pcl_msgs__msg__ModelCoefficients)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pcl_msgs__msg__ModelCoefficients * uintptr(i)),
		))
		*cIdx = *(*C.pcl_msgs__msg__ModelCoefficients)(v.AsCStruct())
	}
}
func ModelCoefficients__Array_to_Go(goSlice []ModelCoefficients, cSlice []CModelCoefficients) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func ModelCoefficients__Array_to_C(cSlice []CModelCoefficients, goSlice []ModelCoefficients) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.pcl_msgs__msg__ModelCoefficients)(goSlice[i].AsCStruct())
	}
}


