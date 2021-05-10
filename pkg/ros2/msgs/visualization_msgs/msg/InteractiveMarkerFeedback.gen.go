/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package visualization_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	geometry_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/geometry_msgs/msg"
	std_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <visualization_msgs/msg/interactive_marker_feedback.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("visualization_msgs/InteractiveMarkerFeedback", &InteractiveMarkerFeedback{})
}
const (
	InteractiveMarkerFeedback_KEEP_ALIVE uint8 = 0// Type of the eventKEEP_ALIVE: sent while dragging to keep up control of the markerMENU_SELECT: a menu entry has been selectedBUTTON_CLICK: a button control has been clickedPOSE_UPDATE: the pose has been changed using one of the controls
	InteractiveMarkerFeedback_POSE_UPDATE uint8 = 1// Type of the eventKEEP_ALIVE: sent while dragging to keep up control of the markerMENU_SELECT: a menu entry has been selectedBUTTON_CLICK: a button control has been clickedPOSE_UPDATE: the pose has been changed using one of the controls
	InteractiveMarkerFeedback_MENU_SELECT uint8 = 2// Type of the eventKEEP_ALIVE: sent while dragging to keep up control of the markerMENU_SELECT: a menu entry has been selectedBUTTON_CLICK: a button control has been clickedPOSE_UPDATE: the pose has been changed using one of the controls
	InteractiveMarkerFeedback_BUTTON_CLICK uint8 = 3// Type of the eventKEEP_ALIVE: sent while dragging to keep up control of the markerMENU_SELECT: a menu entry has been selectedBUTTON_CLICK: a button control has been clickedPOSE_UPDATE: the pose has been changed using one of the controls
	InteractiveMarkerFeedback_MOUSE_DOWN uint8 = 4
	InteractiveMarkerFeedback_MOUSE_UP uint8 = 5
)

// Do not create instances of this type directly. Always use NewInteractiveMarkerFeedback
// function instead.
type InteractiveMarkerFeedback struct {
	Header std_msgs.Header `yaml:"header"`// Time/frame info.
	ClientId rosidl_runtime_c.String `yaml:"client_id"`// Identifying string. Must be unique in the topic namespace.
	MarkerName rosidl_runtime_c.String `yaml:"marker_name"`// Specifies which interactive marker and control this message refers to
	ControlName rosidl_runtime_c.String `yaml:"control_name"`// Specifies which interactive marker and control this message refers to
	EventType uint8 `yaml:"event_type"`
	Pose geometry_msgs.Pose `yaml:"pose"`// Current pose of the markerNote: Has to be valid for all feedback types.
	MenuEntryId uint32 `yaml:"menu_entry_id"`// Contains the ID of the selected menu entryOnly valid for MENU_SELECT events.
	MousePoint geometry_msgs.Point `yaml:"mouse_point"`// If event_type is BUTTON_CLICK, MOUSE_DOWN, or MOUSE_UP, mouse_pointmay contain the 3 dimensional position of the event on thecontrol.  If it does, mouse_point_valid will be true.  mouse_pointwill be relative to the frame listed in the header.
	MousePointValid bool `yaml:"mouse_point_valid"`// If event_type is BUTTON_CLICK, MOUSE_DOWN, or MOUSE_UP, mouse_pointmay contain the 3 dimensional position of the event on thecontrol.  If it does, mouse_point_valid will be true.  mouse_pointwill be relative to the frame listed in the header.
}

// NewInteractiveMarkerFeedback creates a new InteractiveMarkerFeedback with default values.
func NewInteractiveMarkerFeedback() *InteractiveMarkerFeedback {
	self := InteractiveMarkerFeedback{}
	self.SetDefaults(nil)
	return &self
}

func (t *InteractiveMarkerFeedback) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	t.ClientId.SetDefaults("")
	t.MarkerName.SetDefaults("")
	t.ControlName.SetDefaults("")
	t.Pose.SetDefaults(nil)
	t.MousePoint.SetDefaults(nil)
	
	return t
}

func (t *InteractiveMarkerFeedback) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__InteractiveMarkerFeedback())
}
func (t *InteractiveMarkerFeedback) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__InteractiveMarkerFeedback
	return (unsafe.Pointer)(C.visualization_msgs__msg__InteractiveMarkerFeedback__create())
}
func (t *InteractiveMarkerFeedback) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__InteractiveMarkerFeedback__destroy((*C.visualization_msgs__msg__InteractiveMarkerFeedback)(pointer_to_free))
}
func (t *InteractiveMarkerFeedback) AsCStruct() unsafe.Pointer {
	mem := (*C.visualization_msgs__msg__InteractiveMarkerFeedback)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	mem.client_id = *(*C.rosidl_runtime_c__String)(t.ClientId.AsCStruct())
	mem.marker_name = *(*C.rosidl_runtime_c__String)(t.MarkerName.AsCStruct())
	mem.control_name = *(*C.rosidl_runtime_c__String)(t.ControlName.AsCStruct())
	mem.event_type = C.uint8_t(t.EventType)
	mem.pose = *(*C.geometry_msgs__msg__Pose)(t.Pose.AsCStruct())
	mem.menu_entry_id = C.uint32_t(t.MenuEntryId)
	mem.mouse_point = *(*C.geometry_msgs__msg__Point)(t.MousePoint.AsCStruct())
	mem.mouse_point_valid = C.bool(t.MousePointValid)
	return unsafe.Pointer(mem)
}
func (t *InteractiveMarkerFeedback) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.visualization_msgs__msg__InteractiveMarkerFeedback)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	t.ClientId.AsGoStruct(unsafe.Pointer(&mem.client_id))
	t.MarkerName.AsGoStruct(unsafe.Pointer(&mem.marker_name))
	t.ControlName.AsGoStruct(unsafe.Pointer(&mem.control_name))
	t.EventType = uint8(mem.event_type)
	t.Pose.AsGoStruct(unsafe.Pointer(&mem.pose))
	t.MenuEntryId = uint32(mem.menu_entry_id)
	t.MousePoint.AsGoStruct(unsafe.Pointer(&mem.mouse_point))
	t.MousePointValid = bool(mem.mouse_point_valid)
}
func (t *InteractiveMarkerFeedback) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CInteractiveMarkerFeedback = C.visualization_msgs__msg__InteractiveMarkerFeedback
type CInteractiveMarkerFeedback__Sequence = C.visualization_msgs__msg__InteractiveMarkerFeedback__Sequence

func InteractiveMarkerFeedback__Sequence_to_Go(goSlice *[]InteractiveMarkerFeedback, cSlice CInteractiveMarkerFeedback__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]InteractiveMarkerFeedback, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarkerFeedback__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerFeedback * uintptr(i)),
		))
		(*goSlice)[i] = InteractiveMarkerFeedback{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func InteractiveMarkerFeedback__Sequence_to_C(cSlice *CInteractiveMarkerFeedback__Sequence, goSlice []InteractiveMarkerFeedback) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__InteractiveMarkerFeedback)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerFeedback * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarkerFeedback)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerFeedback * uintptr(i)),
		))
		*cIdx = *(*C.visualization_msgs__msg__InteractiveMarkerFeedback)(v.AsCStruct())
	}
}
func InteractiveMarkerFeedback__Array_to_Go(goSlice []InteractiveMarkerFeedback, cSlice []CInteractiveMarkerFeedback) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func InteractiveMarkerFeedback__Array_to_C(cSlice []CInteractiveMarkerFeedback, goSlice []InteractiveMarkerFeedback) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.visualization_msgs__msg__InteractiveMarkerFeedback)(goSlice[i].AsCStruct())
	}
}

