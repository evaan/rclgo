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
#include <sensor_msgs/msg/joint_state.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("sensor_msgs/JointState", &JointState{})
}

// Do not create instances of this type directly. Always use NewJointState
// function instead.
type JointState struct {
	Header std_msgs.Header `yaml:"header"`
	Name []rosidl_runtime_c.String `yaml:"name"`
	Position []float64 `yaml:"position"`
	Velocity []float64 `yaml:"velocity"`
	Effort []float64 `yaml:"effort"`
}

// NewJointState creates a new JointState with default values.
func NewJointState() *JointState {
	self := JointState{}
	self.SetDefaults(nil)
	return &self
}

func (t *JointState) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	
	return t
}

func (t *JointState) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__JointState())
}
func (t *JointState) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__JointState
	return (unsafe.Pointer)(C.sensor_msgs__msg__JointState__create())
}
func (t *JointState) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__JointState__destroy((*C.sensor_msgs__msg__JointState)(pointer_to_free))
}
func (t *JointState) AsCStruct() unsafe.Pointer {
	mem := (*C.sensor_msgs__msg__JointState)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	rosidl_runtime_c.String__Sequence_to_C((*rosidl_runtime_c.CString__Sequence)(unsafe.Pointer(&mem.name)), t.Name)
	rosidl_runtime_c.Float64__Sequence_to_C((*rosidl_runtime_c.CFloat64__Sequence)(unsafe.Pointer(&mem.position)), t.Position)
	rosidl_runtime_c.Float64__Sequence_to_C((*rosidl_runtime_c.CFloat64__Sequence)(unsafe.Pointer(&mem.velocity)), t.Velocity)
	rosidl_runtime_c.Float64__Sequence_to_C((*rosidl_runtime_c.CFloat64__Sequence)(unsafe.Pointer(&mem.effort)), t.Effort)
	return unsafe.Pointer(mem)
}
func (t *JointState) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.sensor_msgs__msg__JointState)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	rosidl_runtime_c.String__Sequence_to_Go(&t.Name, *(*rosidl_runtime_c.CString__Sequence)(unsafe.Pointer(&mem.name)))
	rosidl_runtime_c.Float64__Sequence_to_Go(&t.Position, *(*rosidl_runtime_c.CFloat64__Sequence)(unsafe.Pointer(&mem.position)))
	rosidl_runtime_c.Float64__Sequence_to_Go(&t.Velocity, *(*rosidl_runtime_c.CFloat64__Sequence)(unsafe.Pointer(&mem.velocity)))
	rosidl_runtime_c.Float64__Sequence_to_Go(&t.Effort, *(*rosidl_runtime_c.CFloat64__Sequence)(unsafe.Pointer(&mem.effort)))
}
func (t *JointState) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CJointState = C.sensor_msgs__msg__JointState
type CJointState__Sequence = C.sensor_msgs__msg__JointState__Sequence

func JointState__Sequence_to_Go(goSlice *[]JointState, cSlice CJointState__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]JointState, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__JointState__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__JointState * uintptr(i)),
		))
		(*goSlice)[i] = JointState{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func JointState__Sequence_to_C(cSlice *CJointState__Sequence, goSlice []JointState) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__JointState)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__JointState * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__JointState)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__JointState * uintptr(i)),
		))
		*cIdx = *(*C.sensor_msgs__msg__JointState)(v.AsCStruct())
	}
}
func JointState__Array_to_Go(goSlice []JointState, cSlice []CJointState) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func JointState__Array_to_C(cSlice []CJointState, goSlice []JointState) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.sensor_msgs__msg__JointState)(goSlice[i].AsCStruct())
	}
}


