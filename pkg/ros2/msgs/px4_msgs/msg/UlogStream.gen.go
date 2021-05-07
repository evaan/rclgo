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
#include <px4_msgs/msg/ulog_stream.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/UlogStream", &UlogStream{})
}
const (
	UlogStream_FLAGS_NEED_ACK uint8 = 1// if set, this message requires to be acked.. flags bitmasks
	UlogStream_ORB_QUEUE_LENGTH uint8 = 16// TODO: we might be able to reduce this if mavlink polled on the topic
)

// Do not create instances of this type directly. Always use NewUlogStream
// function instead.
type UlogStream struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Length uint8 `yaml:"length"`// length of data
	FirstMessageOffset uint8 `yaml:"first_message_offset"`// offset into data where first message starts. This
	MsgSequence uint16 `yaml:"msg_sequence"`// allows determine drops. can be used for recovery, when a previous message got lost
	Flags uint8 `yaml:"flags"`// see FLAGS_*. can be used for recovery, when a previous message got lost
	Data [249]uint8 `yaml:"data"`// ulog data. can be used for recovery, when a previous message got lost
}

// NewUlogStream creates a new UlogStream with default values.
func NewUlogStream() *UlogStream {
	self := UlogStream{}
	self.SetDefaults(nil)
	return &self
}

func (t *UlogStream) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *UlogStream) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__UlogStream())
}
func (t *UlogStream) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__UlogStream
	return (unsafe.Pointer)(C.px4_msgs__msg__UlogStream__create())
}
func (t *UlogStream) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__UlogStream__destroy((*C.px4_msgs__msg__UlogStream)(pointer_to_free))
}
func (t *UlogStream) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__UlogStream)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.length = C.uint8_t(t.Length)
	mem.first_message_offset = C.uint8_t(t.FirstMessageOffset)
	mem.msg_sequence = C.uint16_t(t.MsgSequence)
	mem.flags = C.uint8_t(t.Flags)
	cSlice_data := mem.data[:]
	rosidl_runtime_c.Uint8__Array_to_C(*(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_data)), t.Data[:])
	return unsafe.Pointer(mem)
}
func (t *UlogStream) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__UlogStream)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.Length = uint8(mem.length)
	t.FirstMessageOffset = uint8(mem.first_message_offset)
	t.MsgSequence = uint16(mem.msg_sequence)
	t.Flags = uint8(mem.flags)
	cSlice_data := mem.data[:]
	rosidl_runtime_c.Uint8__Array_to_Go(t.Data[:], *(*[]rosidl_runtime_c.CUint8)(unsafe.Pointer(&cSlice_data)))
}
func (t *UlogStream) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CUlogStream = C.px4_msgs__msg__UlogStream
type CUlogStream__Sequence = C.px4_msgs__msg__UlogStream__Sequence

func UlogStream__Sequence_to_Go(goSlice *[]UlogStream, cSlice CUlogStream__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UlogStream, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__UlogStream__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__UlogStream * uintptr(i)),
		))
		(*goSlice)[i] = UlogStream{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func UlogStream__Sequence_to_C(cSlice *CUlogStream__Sequence, goSlice []UlogStream) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__UlogStream)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__UlogStream * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__UlogStream)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__UlogStream * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__UlogStream)(v.AsCStruct())
	}
}
func UlogStream__Array_to_Go(goSlice []UlogStream, cSlice []CUlogStream) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func UlogStream__Array_to_C(cSlice []CUlogStream, goSlice []UlogStream) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__UlogStream)(goSlice[i].AsCStruct())
	}
}


