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
#include <px4_msgs/msg/sensor_gyro_fft.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/SensorGyroFft", &SensorGyroFft{})
}

// Do not create instances of this type directly. Always use NewSensorGyroFft
// function instead.
type SensorGyroFft struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`
	DeviceId uint32 `yaml:"device_id"`// unique device ID for the sensor that does not change between power cycles
	SensorSampleRateHz float32 `yaml:"sensor_sample_rate_hz"`
	ResolutionHz float32 `yaml:"resolution_hz"`
	PeakFrequenciesX [6]float32 `yaml:"peak_frequencies_x"`// x axis peak frequencies
	PeakFrequenciesY [6]float32 `yaml:"peak_frequencies_y"`// y axis peak frequencies
	PeakFrequenciesZ [6]float32 `yaml:"peak_frequencies_z"`// z axis peak frequencies
	PeakMagnitudeX [6]uint32 `yaml:"peak_magnitude_x"`// x axis peak frequencies magnitude
	PeakMagnitudeY [6]uint32 `yaml:"peak_magnitude_y"`// y axis peak frequencies magnitude
	PeakMagnitudeZ [6]uint32 `yaml:"peak_magnitude_z"`// z axis peak frequencies magnitude
}

// NewSensorGyroFft creates a new SensorGyroFft with default values.
func NewSensorGyroFft() *SensorGyroFft {
	self := SensorGyroFft{}
	self.SetDefaults(nil)
	return &self
}

func (t *SensorGyroFft) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *SensorGyroFft) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__SensorGyroFft())
}
func (t *SensorGyroFft) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__SensorGyroFft
	return (unsafe.Pointer)(C.px4_msgs__msg__SensorGyroFft__create())
}
func (t *SensorGyroFft) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__SensorGyroFft__destroy((*C.px4_msgs__msg__SensorGyroFft)(pointer_to_free))
}
func (t *SensorGyroFft) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__SensorGyroFft)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_sample = C.uint64_t(t.TimestampSample)
	mem.device_id = C.uint32_t(t.DeviceId)
	mem.sensor_sample_rate_hz = C.float(t.SensorSampleRateHz)
	mem.resolution_hz = C.float(t.ResolutionHz)
	cSlice_peak_frequencies_x := mem.peak_frequencies_x[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_peak_frequencies_x)), t.PeakFrequenciesX[:])
	cSlice_peak_frequencies_y := mem.peak_frequencies_y[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_peak_frequencies_y)), t.PeakFrequenciesY[:])
	cSlice_peak_frequencies_z := mem.peak_frequencies_z[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_peak_frequencies_z)), t.PeakFrequenciesZ[:])
	cSlice_peak_magnitude_x := mem.peak_magnitude_x[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_peak_magnitude_x)), t.PeakMagnitudeX[:])
	cSlice_peak_magnitude_y := mem.peak_magnitude_y[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_peak_magnitude_y)), t.PeakMagnitudeY[:])
	cSlice_peak_magnitude_z := mem.peak_magnitude_z[:]
	rosidl_runtime_c.Uint32__Array_to_C(*(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_peak_magnitude_z)), t.PeakMagnitudeZ[:])
	return unsafe.Pointer(mem)
}
func (t *SensorGyroFft) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__SensorGyroFft)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampSample = uint64(mem.timestamp_sample)
	t.DeviceId = uint32(mem.device_id)
	t.SensorSampleRateHz = float32(mem.sensor_sample_rate_hz)
	t.ResolutionHz = float32(mem.resolution_hz)
	cSlice_peak_frequencies_x := mem.peak_frequencies_x[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.PeakFrequenciesX[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_peak_frequencies_x)))
	cSlice_peak_frequencies_y := mem.peak_frequencies_y[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.PeakFrequenciesY[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_peak_frequencies_y)))
	cSlice_peak_frequencies_z := mem.peak_frequencies_z[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.PeakFrequenciesZ[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_peak_frequencies_z)))
	cSlice_peak_magnitude_x := mem.peak_magnitude_x[:]
	rosidl_runtime_c.Uint32__Array_to_Go(t.PeakMagnitudeX[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_peak_magnitude_x)))
	cSlice_peak_magnitude_y := mem.peak_magnitude_y[:]
	rosidl_runtime_c.Uint32__Array_to_Go(t.PeakMagnitudeY[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_peak_magnitude_y)))
	cSlice_peak_magnitude_z := mem.peak_magnitude_z[:]
	rosidl_runtime_c.Uint32__Array_to_Go(t.PeakMagnitudeZ[:], *(*[]rosidl_runtime_c.CUint32)(unsafe.Pointer(&cSlice_peak_magnitude_z)))
}
func (t *SensorGyroFft) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CSensorGyroFft = C.px4_msgs__msg__SensorGyroFft
type CSensorGyroFft__Sequence = C.px4_msgs__msg__SensorGyroFft__Sequence

func SensorGyroFft__Sequence_to_Go(goSlice *[]SensorGyroFft, cSlice CSensorGyroFft__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SensorGyroFft, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__SensorGyroFft__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorGyroFft * uintptr(i)),
		))
		(*goSlice)[i] = SensorGyroFft{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func SensorGyroFft__Sequence_to_C(cSlice *CSensorGyroFft__Sequence, goSlice []SensorGyroFft) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__SensorGyroFft)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__SensorGyroFft * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__SensorGyroFft)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorGyroFft * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__SensorGyroFft)(v.AsCStruct())
	}
}
func SensorGyroFft__Array_to_Go(goSlice []SensorGyroFft, cSlice []CSensorGyroFft) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func SensorGyroFft__Array_to_C(cSlice []CSensorGyroFft, goSlice []SensorGyroFft) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__SensorGyroFft)(goSlice[i].AsCStruct())
	}
}


