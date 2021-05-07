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
#include <px4_msgs/msg/system_power.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/SystemPower", &SystemPower{})
}
const (
	SystemPower_BRICK1_VALID_SHIFTS uint8 = 0
	SystemPower_BRICK1_VALID_MASK uint8 = 1
	SystemPower_BRICK2_VALID_SHIFTS uint8 = 1
	SystemPower_BRICK2_VALID_MASK uint8 = 2
	SystemPower_BRICK3_VALID_SHIFTS uint8 = 2
	SystemPower_BRICK3_VALID_MASK uint8 = 4
	SystemPower_BRICK4_VALID_SHIFTS uint8 = 3
	SystemPower_BRICK4_VALID_MASK uint8 = 8
)

// Do not create instances of this type directly. Always use NewSystemPower
// function instead.
type SystemPower struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Voltage5vV float32 `yaml:"voltage5v_v"`// peripheral 5V rail voltage
	Sensors3v3 [4]float32 `yaml:"sensors3v3"`// Sensors 3V3 rail voltage
	Sensors3v3Valid uint8 `yaml:"sensors3v3_valid"`// Sensors 3V3 rail voltage was read (bitfield).
	UsbConnected uint8 `yaml:"usb_connected"`// USB is connected when 1
	BrickValid uint8 `yaml:"brick_valid"`// brick bits power is good when bit 1
	UsbValid uint8 `yaml:"usb_valid"`// USB is valid when 1
	ServoValid uint8 `yaml:"servo_valid"`// servo power is good when 1
	Periph5vOc uint8 `yaml:"periph_5v_oc"`// peripheral overcurrent when 1
	Hipower5vOc uint8 `yaml:"hipower_5v_oc"`// high power peripheral overcurrent when 1
	Comp5vValid uint8 `yaml:"comp_5v_valid"`// 5V to companion valid
	Can1Gps15vValid uint8 `yaml:"can1_gps1_5v_valid"`// 5V for CAN1/GPS1 valid
}

// NewSystemPower creates a new SystemPower with default values.
func NewSystemPower() *SystemPower {
	self := SystemPower{}
	self.SetDefaults(nil)
	return &self
}

func (t *SystemPower) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *SystemPower) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__SystemPower())
}
func (t *SystemPower) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__SystemPower
	return (unsafe.Pointer)(C.px4_msgs__msg__SystemPower__create())
}
func (t *SystemPower) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__SystemPower__destroy((*C.px4_msgs__msg__SystemPower)(pointer_to_free))
}
func (t *SystemPower) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__SystemPower)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.voltage5v_v = C.float(t.Voltage5vV)
	cSlice_sensors3v3 := mem.sensors3v3[:]
	rosidl_runtime_c.Float32__Array_to_C(*(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_sensors3v3)), t.Sensors3v3[:])
	mem.sensors3v3_valid = C.uint8_t(t.Sensors3v3Valid)
	mem.usb_connected = C.uint8_t(t.UsbConnected)
	mem.brick_valid = C.uint8_t(t.BrickValid)
	mem.usb_valid = C.uint8_t(t.UsbValid)
	mem.servo_valid = C.uint8_t(t.ServoValid)
	mem.periph_5v_oc = C.uint8_t(t.Periph5vOc)
	mem.hipower_5v_oc = C.uint8_t(t.Hipower5vOc)
	mem.comp_5v_valid = C.uint8_t(t.Comp5vValid)
	mem.can1_gps1_5v_valid = C.uint8_t(t.Can1Gps15vValid)
	return unsafe.Pointer(mem)
}
func (t *SystemPower) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__SystemPower)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.Voltage5vV = float32(mem.voltage5v_v)
	cSlice_sensors3v3 := mem.sensors3v3[:]
	rosidl_runtime_c.Float32__Array_to_Go(t.Sensors3v3[:], *(*[]rosidl_runtime_c.CFloat32)(unsafe.Pointer(&cSlice_sensors3v3)))
	t.Sensors3v3Valid = uint8(mem.sensors3v3_valid)
	t.UsbConnected = uint8(mem.usb_connected)
	t.BrickValid = uint8(mem.brick_valid)
	t.UsbValid = uint8(mem.usb_valid)
	t.ServoValid = uint8(mem.servo_valid)
	t.Periph5vOc = uint8(mem.periph_5v_oc)
	t.Hipower5vOc = uint8(mem.hipower_5v_oc)
	t.Comp5vValid = uint8(mem.comp_5v_valid)
	t.Can1Gps15vValid = uint8(mem.can1_gps1_5v_valid)
}
func (t *SystemPower) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CSystemPower = C.px4_msgs__msg__SystemPower
type CSystemPower__Sequence = C.px4_msgs__msg__SystemPower__Sequence

func SystemPower__Sequence_to_Go(goSlice *[]SystemPower, cSlice CSystemPower__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SystemPower, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__SystemPower__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SystemPower * uintptr(i)),
		))
		(*goSlice)[i] = SystemPower{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func SystemPower__Sequence_to_C(cSlice *CSystemPower__Sequence, goSlice []SystemPower) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__SystemPower)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__SystemPower * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__SystemPower)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SystemPower * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__SystemPower)(v.AsCStruct())
	}
}
func SystemPower__Array_to_Go(goSlice []SystemPower, cSlice []CSystemPower) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func SystemPower__Array_to_C(cSlice []CSystemPower, goSlice []SystemPower) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__SystemPower)(goSlice[i].AsCStruct())
	}
}


