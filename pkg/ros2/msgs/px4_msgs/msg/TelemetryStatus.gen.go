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
#include <px4_msgs/msg/telemetry_status.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/TelemetryStatus", &TelemetryStatus{})
}
const (
	TelemetryStatus_LINK_TYPE_GENERIC uint8 = 0
	TelemetryStatus_LINK_TYPE_UBIQUITY_BULLET uint8 = 1
	TelemetryStatus_LINK_TYPE_WIRE uint8 = 2
	TelemetryStatus_LINK_TYPE_USB uint8 = 3
	TelemetryStatus_LINK_TYPE_IRIDIUM uint8 = 4
	TelemetryStatus_HEARTBEAT_TIMEOUT_US uint64 = 1500000// Heartbeat timeout 1.5 seconds
)

// Do not create instances of this type directly. Always use NewTelemetryStatus
// function instead.
type TelemetryStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Type uint8 `yaml:"type"`// type of the radio hardware (LINK_TYPE_*)
	Mode uint8 `yaml:"mode"`
	FlowControl bool `yaml:"flow_control"`
	Forwarding bool `yaml:"forwarding"`
	MavlinkV2 bool `yaml:"mavlink_v2"`
	Ftp bool `yaml:"ftp"`
	Streams uint8 `yaml:"streams"`
	DataRate float32 `yaml:"data_rate"`// configured maximum data rate (Bytes/s)
	RateMultiplier float32 `yaml:"rate_multiplier"`
	TxRateAvg float32 `yaml:"tx_rate_avg"`// transmit rate average (Bytes/s)
	TxErrorRateAvg float32 `yaml:"tx_error_rate_avg"`// transmit error rate average (Bytes/s)
	TxMessageCount uint32 `yaml:"tx_message_count"`// total message sent count
	TxBufferOverruns uint32 `yaml:"tx_buffer_overruns"`// number of TX buffer overruns
	RxRateAvg float32 `yaml:"rx_rate_avg"`// transmit rate average (Bytes/s)
	RxMessageCount uint32 `yaml:"rx_message_count"`// count of total messages received
	RxMessageCountSupported uint32 `yaml:"rx_message_count_supported"`// count of total messages received from supported systems and components (for loss statistics)
	RxMessageLostCount uint32 `yaml:"rx_message_lost_count"`
	RxBufferOverruns uint32 `yaml:"rx_buffer_overruns"`// number of RX buffer overruns
	RxParseErrors uint32 `yaml:"rx_parse_errors"`// number of parse errors
	RxPacketDropCount uint32 `yaml:"rx_packet_drop_count"`// number of packet drops
	RxMessageLostRate float32 `yaml:"rx_message_lost_rate"`
	HeartbeatTypeAntennaTracker bool `yaml:"heartbeat_type_antenna_tracker"`// MAV_TYPE_ANTENNA_TRACKER. Heartbeats per type
	HeartbeatTypeGcs bool `yaml:"heartbeat_type_gcs"`// MAV_TYPE_GCS. Heartbeats per type
	HeartbeatTypeOnboardController bool `yaml:"heartbeat_type_onboard_controller"`// MAV_TYPE_ONBOARD_CONTROLLER. Heartbeats per type
	HeartbeatTypeGimbal bool `yaml:"heartbeat_type_gimbal"`// MAV_TYPE_GIMBAL. Heartbeats per type
	HeartbeatTypeAdsb bool `yaml:"heartbeat_type_adsb"`// MAV_TYPE_ADSB. Heartbeats per type
	HeartbeatTypeCamera bool `yaml:"heartbeat_type_camera"`// MAV_TYPE_CAMERA. Heartbeats per type
	HeartbeatComponentTelemetryRadio bool `yaml:"heartbeat_component_telemetry_radio"`// MAV_COMP_ID_TELEMETRY_RADIO. Heartbeats per component
	HeartbeatComponentLog bool `yaml:"heartbeat_component_log"`// MAV_COMP_ID_LOG. Heartbeats per component
	HeartbeatComponentOsd bool `yaml:"heartbeat_component_osd"`// MAV_COMP_ID_OSD. Heartbeats per component
	HeartbeatComponentObstacleAvoidance bool `yaml:"heartbeat_component_obstacle_avoidance"`// MAV_COMP_ID_OBSTACLE_AVOIDANCE. Heartbeats per component
	HeartbeatComponentVio bool `yaml:"heartbeat_component_vio"`// MAV_COMP_ID_VISUAL_INERTIAL_ODOMETRY. Heartbeats per component
	HeartbeatComponentPairingManager bool `yaml:"heartbeat_component_pairing_manager"`// MAV_COMP_ID_PAIRING_MANAGER. Heartbeats per component
	HeartbeatComponentUdpBridge bool `yaml:"heartbeat_component_udp_bridge"`// MAV_COMP_ID_UDP_BRIDGE. Heartbeats per component
	HeartbeatComponentUartBridge bool `yaml:"heartbeat_component_uart_bridge"`// MAV_COMP_ID_UART_BRIDGE. Heartbeats per component
	AvoidanceSystemHealthy bool `yaml:"avoidance_system_healthy"`// Misc component health
}

// NewTelemetryStatus creates a new TelemetryStatus with default values.
func NewTelemetryStatus() *TelemetryStatus {
	self := TelemetryStatus{}
	self.SetDefaults(nil)
	return &self
}

func (t *TelemetryStatus) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *TelemetryStatus) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__TelemetryStatus())
}
func (t *TelemetryStatus) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__TelemetryStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__TelemetryStatus__create())
}
func (t *TelemetryStatus) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__TelemetryStatus__destroy((*C.px4_msgs__msg__TelemetryStatus)(pointer_to_free))
}
func (t *TelemetryStatus) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__TelemetryStatus)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem._type = C.uint8_t(t.Type)
	mem.mode = C.uint8_t(t.Mode)
	mem.flow_control = C.bool(t.FlowControl)
	mem.forwarding = C.bool(t.Forwarding)
	mem.mavlink_v2 = C.bool(t.MavlinkV2)
	mem.ftp = C.bool(t.Ftp)
	mem.streams = C.uint8_t(t.Streams)
	mem.data_rate = C.float(t.DataRate)
	mem.rate_multiplier = C.float(t.RateMultiplier)
	mem.tx_rate_avg = C.float(t.TxRateAvg)
	mem.tx_error_rate_avg = C.float(t.TxErrorRateAvg)
	mem.tx_message_count = C.uint32_t(t.TxMessageCount)
	mem.tx_buffer_overruns = C.uint32_t(t.TxBufferOverruns)
	mem.rx_rate_avg = C.float(t.RxRateAvg)
	mem.rx_message_count = C.uint32_t(t.RxMessageCount)
	mem.rx_message_count_supported = C.uint32_t(t.RxMessageCountSupported)
	mem.rx_message_lost_count = C.uint32_t(t.RxMessageLostCount)
	mem.rx_buffer_overruns = C.uint32_t(t.RxBufferOverruns)
	mem.rx_parse_errors = C.uint32_t(t.RxParseErrors)
	mem.rx_packet_drop_count = C.uint32_t(t.RxPacketDropCount)
	mem.rx_message_lost_rate = C.float(t.RxMessageLostRate)
	mem.heartbeat_type_antenna_tracker = C.bool(t.HeartbeatTypeAntennaTracker)
	mem.heartbeat_type_gcs = C.bool(t.HeartbeatTypeGcs)
	mem.heartbeat_type_onboard_controller = C.bool(t.HeartbeatTypeOnboardController)
	mem.heartbeat_type_gimbal = C.bool(t.HeartbeatTypeGimbal)
	mem.heartbeat_type_adsb = C.bool(t.HeartbeatTypeAdsb)
	mem.heartbeat_type_camera = C.bool(t.HeartbeatTypeCamera)
	mem.heartbeat_component_telemetry_radio = C.bool(t.HeartbeatComponentTelemetryRadio)
	mem.heartbeat_component_log = C.bool(t.HeartbeatComponentLog)
	mem.heartbeat_component_osd = C.bool(t.HeartbeatComponentOsd)
	mem.heartbeat_component_obstacle_avoidance = C.bool(t.HeartbeatComponentObstacleAvoidance)
	mem.heartbeat_component_vio = C.bool(t.HeartbeatComponentVio)
	mem.heartbeat_component_pairing_manager = C.bool(t.HeartbeatComponentPairingManager)
	mem.heartbeat_component_udp_bridge = C.bool(t.HeartbeatComponentUdpBridge)
	mem.heartbeat_component_uart_bridge = C.bool(t.HeartbeatComponentUartBridge)
	mem.avoidance_system_healthy = C.bool(t.AvoidanceSystemHealthy)
	return unsafe.Pointer(mem)
}
func (t *TelemetryStatus) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__TelemetryStatus)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.Type = uint8(mem._type)
	t.Mode = uint8(mem.mode)
	t.FlowControl = bool(mem.flow_control)
	t.Forwarding = bool(mem.forwarding)
	t.MavlinkV2 = bool(mem.mavlink_v2)
	t.Ftp = bool(mem.ftp)
	t.Streams = uint8(mem.streams)
	t.DataRate = float32(mem.data_rate)
	t.RateMultiplier = float32(mem.rate_multiplier)
	t.TxRateAvg = float32(mem.tx_rate_avg)
	t.TxErrorRateAvg = float32(mem.tx_error_rate_avg)
	t.TxMessageCount = uint32(mem.tx_message_count)
	t.TxBufferOverruns = uint32(mem.tx_buffer_overruns)
	t.RxRateAvg = float32(mem.rx_rate_avg)
	t.RxMessageCount = uint32(mem.rx_message_count)
	t.RxMessageCountSupported = uint32(mem.rx_message_count_supported)
	t.RxMessageLostCount = uint32(mem.rx_message_lost_count)
	t.RxBufferOverruns = uint32(mem.rx_buffer_overruns)
	t.RxParseErrors = uint32(mem.rx_parse_errors)
	t.RxPacketDropCount = uint32(mem.rx_packet_drop_count)
	t.RxMessageLostRate = float32(mem.rx_message_lost_rate)
	t.HeartbeatTypeAntennaTracker = bool(mem.heartbeat_type_antenna_tracker)
	t.HeartbeatTypeGcs = bool(mem.heartbeat_type_gcs)
	t.HeartbeatTypeOnboardController = bool(mem.heartbeat_type_onboard_controller)
	t.HeartbeatTypeGimbal = bool(mem.heartbeat_type_gimbal)
	t.HeartbeatTypeAdsb = bool(mem.heartbeat_type_adsb)
	t.HeartbeatTypeCamera = bool(mem.heartbeat_type_camera)
	t.HeartbeatComponentTelemetryRadio = bool(mem.heartbeat_component_telemetry_radio)
	t.HeartbeatComponentLog = bool(mem.heartbeat_component_log)
	t.HeartbeatComponentOsd = bool(mem.heartbeat_component_osd)
	t.HeartbeatComponentObstacleAvoidance = bool(mem.heartbeat_component_obstacle_avoidance)
	t.HeartbeatComponentVio = bool(mem.heartbeat_component_vio)
	t.HeartbeatComponentPairingManager = bool(mem.heartbeat_component_pairing_manager)
	t.HeartbeatComponentUdpBridge = bool(mem.heartbeat_component_udp_bridge)
	t.HeartbeatComponentUartBridge = bool(mem.heartbeat_component_uart_bridge)
	t.AvoidanceSystemHealthy = bool(mem.avoidance_system_healthy)
}
func (t *TelemetryStatus) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CTelemetryStatus = C.px4_msgs__msg__TelemetryStatus
type CTelemetryStatus__Sequence = C.px4_msgs__msg__TelemetryStatus__Sequence

func TelemetryStatus__Sequence_to_Go(goSlice *[]TelemetryStatus, cSlice CTelemetryStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TelemetryStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__TelemetryStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TelemetryStatus * uintptr(i)),
		))
		(*goSlice)[i] = TelemetryStatus{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func TelemetryStatus__Sequence_to_C(cSlice *CTelemetryStatus__Sequence, goSlice []TelemetryStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__TelemetryStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__TelemetryStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__TelemetryStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TelemetryStatus * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__TelemetryStatus)(v.AsCStruct())
	}
}
func TelemetryStatus__Array_to_Go(goSlice []TelemetryStatus, cSlice []CTelemetryStatus) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func TelemetryStatus__Array_to_C(cSlice []CTelemetryStatus, goSlice []TelemetryStatus) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__TelemetryStatus)(goSlice[i].AsCStruct())
	}
}


