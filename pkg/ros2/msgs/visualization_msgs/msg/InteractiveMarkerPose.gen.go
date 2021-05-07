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
#include <visualization_msgs/msg/interactive_marker_pose.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("visualization_msgs/InteractiveMarkerPose", &InteractiveMarkerPose{})
}

// Do not create instances of this type directly. Always use NewInteractiveMarkerPose
// function instead.
type InteractiveMarkerPose struct {
	Header std_msgs.Header `yaml:"header"`// Time/frame info.
	Pose geometry_msgs.Pose `yaml:"pose"`// Initial pose. Also, defines the pivot point for rotations.
	Name rosidl_runtime_c.String `yaml:"name"`// Identifying string. Must be globally unique inthe topic that this message is sent through.
}

// NewInteractiveMarkerPose creates a new InteractiveMarkerPose with default values.
func NewInteractiveMarkerPose() *InteractiveMarkerPose {
	self := InteractiveMarkerPose{}
	self.SetDefaults(nil)
	return &self
}

func (t *InteractiveMarkerPose) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	t.Pose.SetDefaults(nil)
	t.Name.SetDefaults("")
	
	return t
}

func (t *InteractiveMarkerPose) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__InteractiveMarkerPose())
}
func (t *InteractiveMarkerPose) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__InteractiveMarkerPose
	return (unsafe.Pointer)(C.visualization_msgs__msg__InteractiveMarkerPose__create())
}
func (t *InteractiveMarkerPose) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__InteractiveMarkerPose__destroy((*C.visualization_msgs__msg__InteractiveMarkerPose)(pointer_to_free))
}
func (t *InteractiveMarkerPose) AsCStruct() unsafe.Pointer {
	mem := (*C.visualization_msgs__msg__InteractiveMarkerPose)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	mem.pose = *(*C.geometry_msgs__msg__Pose)(t.Pose.AsCStruct())
	mem.name = *(*C.rosidl_runtime_c__String)(t.Name.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *InteractiveMarkerPose) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.visualization_msgs__msg__InteractiveMarkerPose)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	t.Pose.AsGoStruct(unsafe.Pointer(&mem.pose))
	t.Name.AsGoStruct(unsafe.Pointer(&mem.name))
}
func (t *InteractiveMarkerPose) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CInteractiveMarkerPose = C.visualization_msgs__msg__InteractiveMarkerPose
type CInteractiveMarkerPose__Sequence = C.visualization_msgs__msg__InteractiveMarkerPose__Sequence

func InteractiveMarkerPose__Sequence_to_Go(goSlice *[]InteractiveMarkerPose, cSlice CInteractiveMarkerPose__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]InteractiveMarkerPose, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarkerPose__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerPose * uintptr(i)),
		))
		(*goSlice)[i] = InteractiveMarkerPose{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func InteractiveMarkerPose__Sequence_to_C(cSlice *CInteractiveMarkerPose__Sequence, goSlice []InteractiveMarkerPose) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__InteractiveMarkerPose)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerPose * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarkerPose)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerPose * uintptr(i)),
		))
		*cIdx = *(*C.visualization_msgs__msg__InteractiveMarkerPose)(v.AsCStruct())
	}
}
func InteractiveMarkerPose__Array_to_Go(goSlice []InteractiveMarkerPose, cSlice []CInteractiveMarkerPose) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func InteractiveMarkerPose__Array_to_C(cSlice []CInteractiveMarkerPose, goSlice []InteractiveMarkerPose) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.visualization_msgs__msg__InteractiveMarkerPose)(goSlice[i].AsCStruct())
	}
}


