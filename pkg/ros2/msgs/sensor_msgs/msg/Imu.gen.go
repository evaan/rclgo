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
	geometry_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/geometry_msgs/msg"
	std_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <sensor_msgs/msg/imu.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("sensor_msgs/Imu", &Imu{})
}

// Do not create instances of this type directly. Always use NewImu
// function instead.
type Imu struct {
	Header std_msgs.Header `yaml:"header"`
	Orientation geometry_msgs.Quaternion `yaml:"orientation"`
	OrientationCovariance [9]float64 `yaml:"orientation_covariance"`// Row major about x, y, z axes
	AngularVelocity geometry_msgs.Vector3 `yaml:"angular_velocity"`
	AngularVelocityCovariance [9]float64 `yaml:"angular_velocity_covariance"`// Row major about x, y, z axes
	LinearAcceleration geometry_msgs.Vector3 `yaml:"linear_acceleration"`
	LinearAccelerationCovariance [9]float64 `yaml:"linear_acceleration_covariance"`// Row major x, y z
}

// NewImu creates a new Imu with default values.
func NewImu() *Imu {
	self := Imu{}
	self.SetDefaults(nil)
	return &self
}

func (t *Imu) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	t.Orientation.SetDefaults(nil)
	t.AngularVelocity.SetDefaults(nil)
	t.LinearAcceleration.SetDefaults(nil)
	
	return t
}

func (t *Imu) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__Imu())
}
func (t *Imu) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__Imu
	return (unsafe.Pointer)(C.sensor_msgs__msg__Imu__create())
}
func (t *Imu) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__Imu__destroy((*C.sensor_msgs__msg__Imu)(pointer_to_free))
}
func (t *Imu) AsCStruct() unsafe.Pointer {
	mem := (*C.sensor_msgs__msg__Imu)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	mem.orientation = *(*C.geometry_msgs__msg__Quaternion)(t.Orientation.AsCStruct())
	cSlice_orientation_covariance := mem.orientation_covariance[:]
	rosidl_runtime_c.Float64__Array_to_C(*(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_orientation_covariance)), t.OrientationCovariance[:])
	mem.angular_velocity = *(*C.geometry_msgs__msg__Vector3)(t.AngularVelocity.AsCStruct())
	cSlice_angular_velocity_covariance := mem.angular_velocity_covariance[:]
	rosidl_runtime_c.Float64__Array_to_C(*(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_angular_velocity_covariance)), t.AngularVelocityCovariance[:])
	mem.linear_acceleration = *(*C.geometry_msgs__msg__Vector3)(t.LinearAcceleration.AsCStruct())
	cSlice_linear_acceleration_covariance := mem.linear_acceleration_covariance[:]
	rosidl_runtime_c.Float64__Array_to_C(*(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_linear_acceleration_covariance)), t.LinearAccelerationCovariance[:])
	return unsafe.Pointer(mem)
}
func (t *Imu) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.sensor_msgs__msg__Imu)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	t.Orientation.AsGoStruct(unsafe.Pointer(&mem.orientation))
	cSlice_orientation_covariance := mem.orientation_covariance[:]
	rosidl_runtime_c.Float64__Array_to_Go(t.OrientationCovariance[:], *(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_orientation_covariance)))
	t.AngularVelocity.AsGoStruct(unsafe.Pointer(&mem.angular_velocity))
	cSlice_angular_velocity_covariance := mem.angular_velocity_covariance[:]
	rosidl_runtime_c.Float64__Array_to_Go(t.AngularVelocityCovariance[:], *(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_angular_velocity_covariance)))
	t.LinearAcceleration.AsGoStruct(unsafe.Pointer(&mem.linear_acceleration))
	cSlice_linear_acceleration_covariance := mem.linear_acceleration_covariance[:]
	rosidl_runtime_c.Float64__Array_to_Go(t.LinearAccelerationCovariance[:], *(*[]rosidl_runtime_c.CFloat64)(unsafe.Pointer(&cSlice_linear_acceleration_covariance)))
}
func (t *Imu) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CImu = C.sensor_msgs__msg__Imu
type CImu__Sequence = C.sensor_msgs__msg__Imu__Sequence

func Imu__Sequence_to_Go(goSlice *[]Imu, cSlice CImu__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Imu, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__Imu__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__Imu * uintptr(i)),
		))
		(*goSlice)[i] = Imu{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Imu__Sequence_to_C(cSlice *CImu__Sequence, goSlice []Imu) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__Imu)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__Imu * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__Imu)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__Imu * uintptr(i)),
		))
		*cIdx = *(*C.sensor_msgs__msg__Imu)(v.AsCStruct())
	}
}
func Imu__Array_to_Go(goSlice []Imu, cSlice []CImu) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Imu__Array_to_C(cSlice []CImu, goSlice []Imu) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.sensor_msgs__msg__Imu)(goSlice[i].AsCStruct())
	}
}


