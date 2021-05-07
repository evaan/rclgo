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
#include <px4_msgs/msg/estimator_selector_status.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/EstimatorSelectorStatus", &EstimatorSelectorStatus{})
}

// Do not create instances of this type directly. Always use NewEstimatorSelectorStatus
// function instead.
type EstimatorSelectorStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	PrimaryInstance uint8 `yaml:"primary_instance"`
	InstancesAvailable uint8 `yaml:"instances_available"`
	InstanceChangedCount uint32 `yaml:"instance_changed_count"`
	LastInstanceChange uint64 `yaml:"last_instance_change"`
	AccelDeviceId uint32 `yaml:"accel_device_id"`
	BaroDeviceId uint32 `yaml:"baro_device_id"`
	GyroDeviceId uint32 `yaml:"gyro_device_id"`
	MagDeviceId uint32 `yaml:"mag_device_id"`
	CombinedTestRatio [9]float32 `yaml:"combined_test_ratio"`
	RelativeTestRatio [9]float32 `yaml:"relative_test_ratio"`
	Healthy [9]bool `yaml:"healthy"`
	AccumulatedGyroError [4]float32 `yaml:"accumulated_gyro_error"`
	AccumulatedAccelError [4]float32 `yaml:"accumulated_accel_error"`
	GyroFaultDetected bool `yaml:"gyro_fault_detected"`
	AccelFaultDetected bool `yaml:"accel_fault_detected"`
}

// NewEstimatorSelectorStatus creates a new EstimatorSelectorStatus with default values.
func NewEstimatorSelectorStatus() *EstimatorSelectorStatus {
	self := EstimatorSelectorStatus{}
	self.SetDefaults(nil)
	return &self
}

func (t *EstimatorSelectorStatus) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *EstimatorSelectorStatus) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EstimatorSelectorStatus())
}
func (t *EstimatorSelectorStatus) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EstimatorSelectorStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__EstimatorSelectorStatus__create())
}
func (t *EstimatorSelectorStatus) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EstimatorSelectorStatus__destroy((*C.px4_msgs__msg__EstimatorSelectorStatus)(pointer_to_free))
}
func (t *EstimatorSelectorStatus) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__EstimatorSelectorStatus)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.primary_instance = C.uint8_t(t.PrimaryInstance)
	mem.instances_available = C.uint8_t(t.InstancesAvailable)
	mem.instance_changed_count = C.uint32_t(t.InstanceChangedCount)
	mem.last_instance_change = C.uint64_t(t.LastInstanceChange)
	mem.accel_device_id = C.uint32_t(t.AccelDeviceId)
	mem.baro_device_id = C.uint32_t(t.BaroDeviceId)
	mem.gyro_device_id = C.uint32_t(t.GyroDeviceId)
	mem.mag_device_id = C.uint32_t(t.MagDeviceId)
	cSlice_combined_test_ratio := mem.combined_test_ratio[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_combined_test_ratio)), t.CombinedTestRatio[:])
	cSlice_relative_test_ratio := mem.relative_test_ratio[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_relative_test_ratio)), t.RelativeTestRatio[:])
	cSlice_healthy := mem.healthy[:]
	rosidl_runtime_c.Bool__Array_to_C(*(*[]rosidl_runtime_c.CBool)(unsafe.Pointer(&cSlice_healthy)), t.Healthy[:])
	cSlice_accumulated_gyro_error := mem.accumulated_gyro_error[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_accumulated_gyro_error)), t.AccumulatedGyroError[:])
	cSlice_accumulated_accel_error := mem.accumulated_accel_error[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_accumulated_accel_error)), t.AccumulatedAccelError[:])
	mem.gyro_fault_detected = C.bool(t.GyroFaultDetected)
	mem.accel_fault_detected = C.bool(t.AccelFaultDetected)
	return unsafe.Pointer(mem)
}
func (t *EstimatorSelectorStatus) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__EstimatorSelectorStatus)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.PrimaryInstance = uint8(mem.primary_instance)
	t.InstancesAvailable = uint8(mem.instances_available)
	t.InstanceChangedCount = uint32(mem.instance_changed_count)
	t.LastInstanceChange = uint64(mem.last_instance_change)
	t.AccelDeviceId = uint32(mem.accel_device_id)
	t.BaroDeviceId = uint32(mem.baro_device_id)
	t.GyroDeviceId = uint32(mem.gyro_device_id)
	t.MagDeviceId = uint32(mem.mag_device_id)
	cSlice_combined_test_ratio := mem.combined_test_ratio[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.CombinedTestRatio[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_combined_test_ratio)))
	cSlice_relative_test_ratio := mem.relative_test_ratio[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.RelativeTestRatio[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_relative_test_ratio)))
	cSlice_healthy := mem.healthy[:]
	rosidl_runtime_c.Bool__Array_to_Go(t.Healthy[:], *(*[]rosidl_runtime_c.CBool)(unsafe.Pointer(&cSlice_healthy)))
	cSlice_accumulated_gyro_error := mem.accumulated_gyro_error[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.AccumulatedGyroError[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_accumulated_gyro_error)))
	cSlice_accumulated_accel_error := mem.accumulated_accel_error[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.AccumulatedAccelError[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_accumulated_accel_error)))
	t.GyroFaultDetected = bool(mem.gyro_fault_detected)
	t.AccelFaultDetected = bool(mem.accel_fault_detected)
}
func (t *EstimatorSelectorStatus) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CEstimatorSelectorStatus = C.px4_msgs__msg__EstimatorSelectorStatus
type CEstimatorSelectorStatus__Sequence = C.px4_msgs__msg__EstimatorSelectorStatus__Sequence

func EstimatorSelectorStatus__Sequence_to_Go(goSlice *[]EstimatorSelectorStatus, cSlice CEstimatorSelectorStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorSelectorStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EstimatorSelectorStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorSelectorStatus * uintptr(i)),
		))
		(*goSlice)[i] = EstimatorSelectorStatus{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func EstimatorSelectorStatus__Sequence_to_C(cSlice *CEstimatorSelectorStatus__Sequence, goSlice []EstimatorSelectorStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EstimatorSelectorStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EstimatorSelectorStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EstimatorSelectorStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorSelectorStatus * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__EstimatorSelectorStatus)(v.AsCStruct())
	}
}
func EstimatorSelectorStatus__Array_to_Go(goSlice []EstimatorSelectorStatus, cSlice []CEstimatorSelectorStatus) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorSelectorStatus__Array_to_C(cSlice []CEstimatorSelectorStatus, goSlice []EstimatorSelectorStatus) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__EstimatorSelectorStatus)(goSlice[i].AsCStruct())
	}
}


