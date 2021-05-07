/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package rcl_interfaces
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <rcl_interfaces/msg/parameter_descriptor.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("rcl_interfaces/ParameterDescriptor", &ParameterDescriptor{})
}

// Do not create instances of this type directly. Always use NewParameterDescriptor
// function instead.
type ParameterDescriptor struct {
	Name rosidl_runtime_c.String `yaml:"name"`// The name of the parameter.
	Type uint8 `yaml:"type"`// Enum values are defined in the `ParameterType.msg` message.
	Description rosidl_runtime_c.String `yaml:"description"`// Description of the parameter, visible from introspection tools.
	AdditionalConstraints rosidl_runtime_c.String `yaml:"additional_constraints"`// Plain English description of additional constraints which cannot be expressedwith the available constraints, e.g. "only prime numbers".By convention, this should only be used to clarify constraints which cannotbe completely expressed with the parameter constraints below.
	ReadOnly bool `yaml:"read_only"`// If 'true' then the value cannot change after it has been initialized.
	FloatingPointRange []FloatingPointRange `yaml:"floating_point_range"`// FloatingPointRange consists of a from_value, a to_value, and a step.
	IntegerRange []IntegerRange `yaml:"integer_range"`// IntegerRange consists of a from_value, a to_value, and a step.
}

// NewParameterDescriptor creates a new ParameterDescriptor with default values.
func NewParameterDescriptor() *ParameterDescriptor {
	self := ParameterDescriptor{}
	self.SetDefaults(nil)
	return &self
}

func (t *ParameterDescriptor) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Name.SetDefaults("")
	t.Description.SetDefaults("")
	t.AdditionalConstraints.SetDefaults("")
	t.ReadOnly = false
	
	return t
}

func (t *ParameterDescriptor) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__msg__ParameterDescriptor())
}
func (t *ParameterDescriptor) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__msg__ParameterDescriptor
	return (unsafe.Pointer)(C.rcl_interfaces__msg__ParameterDescriptor__create())
}
func (t *ParameterDescriptor) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__msg__ParameterDescriptor__destroy((*C.rcl_interfaces__msg__ParameterDescriptor)(pointer_to_free))
}
func (t *ParameterDescriptor) AsCStruct() unsafe.Pointer {
	mem := (*C.rcl_interfaces__msg__ParameterDescriptor)(t.PrepareMemory())
	mem.name = *(*C.rosidl_runtime_c__String)(t.Name.AsCStruct())
	mem._type = C.uint8_t(t.Type)
	mem.description = *(*C.rosidl_runtime_c__String)(t.Description.AsCStruct())
	mem.additional_constraints = *(*C.rosidl_runtime_c__String)(t.AdditionalConstraints.AsCStruct())
	mem.read_only = C.bool(t.ReadOnly)
	FloatingPointRange__Sequence_to_C(&mem.floating_point_range, t.FloatingPointRange)
	IntegerRange__Sequence_to_C(&mem.integer_range, t.IntegerRange)
	return unsafe.Pointer(mem)
}
func (t *ParameterDescriptor) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.rcl_interfaces__msg__ParameterDescriptor)(ros2_message_buffer)
	t.Name.AsGoStruct(unsafe.Pointer(&mem.name))
	t.Type = uint8(mem._type)
	t.Description.AsGoStruct(unsafe.Pointer(&mem.description))
	t.AdditionalConstraints.AsGoStruct(unsafe.Pointer(&mem.additional_constraints))
	t.ReadOnly = bool(mem.read_only)
	FloatingPointRange__Sequence_to_Go(&t.FloatingPointRange, mem.floating_point_range)
	IntegerRange__Sequence_to_Go(&t.IntegerRange, mem.integer_range)
}
func (t *ParameterDescriptor) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CParameterDescriptor = C.rcl_interfaces__msg__ParameterDescriptor
type CParameterDescriptor__Sequence = C.rcl_interfaces__msg__ParameterDescriptor__Sequence

func ParameterDescriptor__Sequence_to_Go(goSlice *[]ParameterDescriptor, cSlice CParameterDescriptor__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ParameterDescriptor, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rcl_interfaces__msg__ParameterDescriptor__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__msg__ParameterDescriptor * uintptr(i)),
		))
		(*goSlice)[i] = ParameterDescriptor{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func ParameterDescriptor__Sequence_to_C(cSlice *CParameterDescriptor__Sequence, goSlice []ParameterDescriptor) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rcl_interfaces__msg__ParameterDescriptor)(C.malloc((C.size_t)(C.sizeof_struct_rcl_interfaces__msg__ParameterDescriptor * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rcl_interfaces__msg__ParameterDescriptor)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__msg__ParameterDescriptor * uintptr(i)),
		))
		*cIdx = *(*C.rcl_interfaces__msg__ParameterDescriptor)(v.AsCStruct())
	}
}
func ParameterDescriptor__Array_to_Go(goSlice []ParameterDescriptor, cSlice []CParameterDescriptor) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func ParameterDescriptor__Array_to_C(cSlice []CParameterDescriptor, goSlice []ParameterDescriptor) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.rcl_interfaces__msg__ParameterDescriptor)(goSlice[i].AsCStruct())
	}
}


