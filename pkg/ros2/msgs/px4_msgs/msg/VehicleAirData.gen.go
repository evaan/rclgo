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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <px4_msgs/msg/vehicle_air_data.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/VehicleAirData", &VehicleAirData{})
}

// Do not create instances of this type directly. Always use NewVehicleAirData
// function instead.
type VehicleAirData struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	BaroDeviceId uint32 `yaml:"baro_device_id"`// unique device ID for the selected barometer
	BaroAltMeter float32 `yaml:"baro_alt_meter"`// Altitude above MSL calculated from temperature compensated baro sensor data using an ISA corrected for sea level pressure SENS_BARO_QNH.
	BaroTempCelcius float32 `yaml:"baro_temp_celcius"`// Temperature in degrees Celsius
	BaroPressurePa float32 `yaml:"baro_pressure_pa"`// Absolute pressure in Pascals
	Rho float32 `yaml:"rho"`// air density
}

// NewVehicleAirData creates a new VehicleAirData with default values.
func NewVehicleAirData() *VehicleAirData {
	self := VehicleAirData{}
	self.SetDefaults(nil)
	return &self
}

func (t *VehicleAirData) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *VehicleAirData) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleAirData())
}
func (t *VehicleAirData) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleAirData
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleAirData__create())
}
func (t *VehicleAirData) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleAirData__destroy((*C.px4_msgs__msg__VehicleAirData)(pointer_to_free))
}
func (t *VehicleAirData) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__VehicleAirData)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.timestamp_sample = C.uint64_t(t.TimestampSample)
	mem.baro_device_id = C.uint32_t(t.BaroDeviceId)
	mem.baro_alt_meter = C.float(t.BaroAltMeter)
	mem.baro_temp_celcius = C.float(t.BaroTempCelcius)
	mem.baro_pressure_pa = C.float(t.BaroPressurePa)
	mem.rho = C.float(t.Rho)
	return unsafe.Pointer(mem)
}
func (t *VehicleAirData) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__VehicleAirData)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.TimestampSample = uint64(mem.timestamp_sample)
	t.BaroDeviceId = uint32(mem.baro_device_id)
	t.BaroAltMeter = float32(mem.baro_alt_meter)
	t.BaroTempCelcius = float32(mem.baro_temp_celcius)
	t.BaroPressurePa = float32(mem.baro_pressure_pa)
	t.Rho = float32(mem.rho)
}
func (t *VehicleAirData) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CVehicleAirData = C.px4_msgs__msg__VehicleAirData
type CVehicleAirData__Sequence = C.px4_msgs__msg__VehicleAirData__Sequence

func VehicleAirData__Sequence_to_Go(goSlice *[]VehicleAirData, cSlice CVehicleAirData__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleAirData, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleAirData__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAirData * uintptr(i)),
		))
		(*goSlice)[i] = VehicleAirData{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func VehicleAirData__Sequence_to_C(cSlice *CVehicleAirData__Sequence, goSlice []VehicleAirData) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleAirData)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleAirData * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleAirData)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAirData * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__VehicleAirData)(v.AsCStruct())
	}
}
func VehicleAirData__Array_to_Go(goSlice []VehicleAirData, cSlice []CVehicleAirData) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleAirData__Array_to_C(cSlice []CVehicleAirData, goSlice []VehicleAirData) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__VehicleAirData)(goSlice[i].AsCStruct())
	}
}


