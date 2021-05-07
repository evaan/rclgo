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
#include <px4_msgs/msg/cellular_status.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/CellularStatus", &CellularStatus{})
}
const (
	CellularStatus_CELLULAR_STATUS_FLAG_UNKNOWN uint8 = 0// State unknown or not reportable
	CellularStatus_CELLULAR_STATUS_FLAG_FAILED uint8 = 1// velocity setpoint
	CellularStatus_CELLULAR_STATUS_FLAG_INITIALIZING uint8 = 2// Modem is being initialized
	CellularStatus_CELLULAR_STATUS_FLAG_LOCKED uint8 = 3// Modem is locked
	CellularStatus_CELLULAR_STATUS_FLAG_DISABLED uint8 = 4// Modem is not enabled and is powered down
	CellularStatus_CELLULAR_STATUS_FLAG_DISABLING uint8 = 5// Modem is currently transitioning to the CELLULAR_STATUS_FLAG_DISABLED state
	CellularStatus_CELLULAR_STATUS_FLAG_ENABLING uint8 = 6// Modem is currently transitioning to the CELLULAR_STATUS_FLAG_ENABLED state
	CellularStatus_CELLULAR_STATUS_FLAG_ENABLED uint8 = 7// Modem is enabled and powered on but not registered with a network provider and not available for data connections
	CellularStatus_CELLULAR_STATUS_FLAG_SEARCHING uint8 = 8// Modem is searching for a network provider to register
	CellularStatus_CELLULAR_STATUS_FLAG_REGISTERED uint8 = 9// Modem is registered with a network provider, and data connections and messaging may be available for use
	CellularStatus_CELLULAR_STATUS_FLAG_DISCONNECTING uint8 = 10// Modem is disconnecting and deactivating the last active packet data bearer. This state will not be entered if more than one packet data bearer is active and one of the active bearers is deactivated
	CellularStatus_CELLULAR_STATUS_FLAG_CONNECTING uint8 = 11// Modem is activating and connecting the first packet data bearer. Subsequent bearer activations when another bearer is already active do not cause this state to be entered
	CellularStatus_CELLULAR_STATUS_FLAG_CONNECTED uint8 = 12// One or more packet data bearers is active and connected
	CellularStatus_CELLULAR_NETWORK_FAILED_REASON_NONE uint8 = 0// No error
	CellularStatus_CELLULAR_NETWORK_FAILED_REASON_UNKNOWN uint8 = 1// Error state is unknown
	CellularStatus_CELLULAR_NETWORK_FAILED_REASON_SIM_MISSING uint8 = 2// SIM is required for the modem but missing
	CellularStatus_CELLULAR_NETWORK_FAILED_REASON_SIM_ERROR uint8 = 3// SIM is available, but not usuable for connection
)

// Do not create instances of this type directly. Always use NewCellularStatus
// function instead.
type CellularStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Status uint16 `yaml:"status"`// Status bitmap 1: Roaming is active
	FailureReason uint8 `yaml:"failure_reason"`// Failure reason when status in in CELLUAR_STATUS_FAILED
	Type uint8 `yaml:"type"`// Cellular network radio type 0: none 1: gsm 2: cdma 3: wcdma 4: lte
	Quality uint8 `yaml:"quality"`// Cellular network RSSI/RSRP in dBm, absolute value
	Mcc uint16 `yaml:"mcc"`// Mobile country code. If unknown, set to: UINT16_MAX
	Mnc uint16 `yaml:"mnc"`// Mobile network code. If unknown, set to: UINT16_MAX
	Lac uint16 `yaml:"lac"`// Location area code. If unknown, set to: 0
}

// NewCellularStatus creates a new CellularStatus with default values.
func NewCellularStatus() *CellularStatus {
	self := CellularStatus{}
	self.SetDefaults(nil)
	return &self
}

func (t *CellularStatus) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *CellularStatus) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__CellularStatus())
}
func (t *CellularStatus) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__CellularStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__CellularStatus__create())
}
func (t *CellularStatus) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__CellularStatus__destroy((*C.px4_msgs__msg__CellularStatus)(pointer_to_free))
}
func (t *CellularStatus) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__CellularStatus)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.status = C.uint16_t(t.Status)
	mem.failure_reason = C.uint8_t(t.FailureReason)
	mem._type = C.uint8_t(t.Type)
	mem.quality = C.uint8_t(t.Quality)
	mem.mcc = C.uint16_t(t.Mcc)
	mem.mnc = C.uint16_t(t.Mnc)
	mem.lac = C.uint16_t(t.Lac)
	return unsafe.Pointer(mem)
}
func (t *CellularStatus) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__CellularStatus)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.Status = uint16(mem.status)
	t.FailureReason = uint8(mem.failure_reason)
	t.Type = uint8(mem._type)
	t.Quality = uint8(mem.quality)
	t.Mcc = uint16(mem.mcc)
	t.Mnc = uint16(mem.mnc)
	t.Lac = uint16(mem.lac)
}
func (t *CellularStatus) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CCellularStatus = C.px4_msgs__msg__CellularStatus
type CCellularStatus__Sequence = C.px4_msgs__msg__CellularStatus__Sequence

func CellularStatus__Sequence_to_Go(goSlice *[]CellularStatus, cSlice CCellularStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CellularStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__CellularStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CellularStatus * uintptr(i)),
		))
		(*goSlice)[i] = CellularStatus{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func CellularStatus__Sequence_to_C(cSlice *CCellularStatus__Sequence, goSlice []CellularStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__CellularStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__CellularStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__CellularStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CellularStatus * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__CellularStatus)(v.AsCStruct())
	}
}
func CellularStatus__Array_to_Go(goSlice []CellularStatus, cSlice []CCellularStatus) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func CellularStatus__Array_to_C(cSlice []CCellularStatus, goSlice []CellularStatus) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__CellularStatus)(goSlice[i].AsCStruct())
	}
}


