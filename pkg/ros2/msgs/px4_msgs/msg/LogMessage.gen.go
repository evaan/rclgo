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
#include <px4_msgs/msg/log_message.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/LogMessage", &LogMessage{})
}
const (
	LogMessage_ORB_QUEUE_LENGTH uint8 = 4
)

// Do not create instances of this type directly. Always use NewLogMessage
// function instead.
type LogMessage struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Severity uint8 `yaml:"severity"`// log level (same as in the linux kernel, starting with 0)
	Text [127]byte `yaml:"text"`
}

// NewLogMessage creates a new LogMessage with default values.
func NewLogMessage() *LogMessage {
	self := LogMessage{}
	self.SetDefaults(nil)
	return &self
}

func (t *LogMessage) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *LogMessage) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__LogMessage())
}
func (t *LogMessage) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__LogMessage
	return (unsafe.Pointer)(C.px4_msgs__msg__LogMessage__create())
}
func (t *LogMessage) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__LogMessage__destroy((*C.px4_msgs__msg__LogMessage)(pointer_to_free))
}
func (t *LogMessage) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__LogMessage)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.severity = C.uint8_t(t.Severity)
	cSlice_text := mem.text[:]
	rosidl_runtime_c.Char__Array_to_C(*(*[]rosidl_runtime_c.CChar)(unsafe.Pointer(&cSlice_text)), t.Text[:])
	return unsafe.Pointer(mem)
}
func (t *LogMessage) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__LogMessage)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.Severity = uint8(mem.severity)
	cSlice_text := mem.text[:]
	rosidl_runtime_c.Char__Array_to_Go(t.Text[:], *(*[]rosidl_runtime_c.CChar)(unsafe.Pointer(&cSlice_text)))
}
func (t *LogMessage) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CLogMessage = C.px4_msgs__msg__LogMessage
type CLogMessage__Sequence = C.px4_msgs__msg__LogMessage__Sequence

func LogMessage__Sequence_to_Go(goSlice *[]LogMessage, cSlice CLogMessage__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]LogMessage, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__LogMessage__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__LogMessage * uintptr(i)),
		))
		(*goSlice)[i] = LogMessage{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func LogMessage__Sequence_to_C(cSlice *CLogMessage__Sequence, goSlice []LogMessage) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__LogMessage)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__LogMessage * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__LogMessage)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__LogMessage * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__LogMessage)(v.AsCStruct())
	}
}
func LogMessage__Array_to_Go(goSlice []LogMessage, cSlice []CLogMessage) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func LogMessage__Array_to_C(cSlice []CLogMessage, goSlice []LogMessage) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__LogMessage)(goSlice[i].AsCStruct())
	}
}

