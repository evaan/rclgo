/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <example_interfaces/msg/char.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("example_interfaces/Char", &Char{})
}

// Do not create instances of this type directly. Always use NewChar
// function instead.
type Char struct {
	Data byte `yaml:"data"`// This is an example message of using a primitive datatype, char.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewChar creates a new Char with default values.
func NewChar() *Char {
	self := Char{}
	self.SetDefaults(nil)
	return &self
}

func (t *Char) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *Char) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Char())
}
func (t *Char) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Char
	return (unsafe.Pointer)(C.example_interfaces__msg__Char__create())
}
func (t *Char) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Char__destroy((*C.example_interfaces__msg__Char)(pointer_to_free))
}
func (t *Char) AsCStruct() unsafe.Pointer {
	mem := (*C.example_interfaces__msg__Char)(t.PrepareMemory())
	mem.data = C.uchar(t.Data)
	return unsafe.Pointer(mem)
}
func (t *Char) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.example_interfaces__msg__Char)(ros2_message_buffer)
	t.Data = byte(mem.data)
}
func (t *Char) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CChar = C.example_interfaces__msg__Char
type CChar__Sequence = C.example_interfaces__msg__Char__Sequence

func Char__Sequence_to_Go(goSlice *[]Char, cSlice CChar__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Char, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__Char__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Char * uintptr(i)),
		))
		(*goSlice)[i] = Char{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Char__Sequence_to_C(cSlice *CChar__Sequence, goSlice []Char) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Char)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__Char * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__Char)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Char * uintptr(i)),
		))
		*cIdx = *(*C.example_interfaces__msg__Char)(v.AsCStruct())
	}
}
func Char__Array_to_Go(goSlice []Char, cSlice []CChar) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Char__Array_to_C(cSlice []CChar, goSlice []Char) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.example_interfaces__msg__Char)(goSlice[i].AsCStruct())
	}
}


