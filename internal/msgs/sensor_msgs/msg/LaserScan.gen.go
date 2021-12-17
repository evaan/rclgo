/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/laser_scan.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/LaserScan", LaserScanTypeSupport)
}

// Do not create instances of this type directly. Always use NewLaserScan
// function instead.
type LaserScan struct {
	Header std_msgs_msg.Header `yaml:"header"`// timestamp in the header is the acquisition time of
	AngleMin float32 `yaml:"angle_min"`// start angle of the scan [rad]
	AngleMax float32 `yaml:"angle_max"`// end angle of the scan [rad]
	AngleIncrement float32 `yaml:"angle_increment"`// angular distance between measurements [rad]
	TimeIncrement float32 `yaml:"time_increment"`// time between measurements [seconds] - if your scanner
	ScanTime float32 `yaml:"scan_time"`// time between scans [seconds]. is moving, this will be used in interpolating positionof 3d points
	RangeMin float32 `yaml:"range_min"`// minimum range value [m]
	RangeMax float32 `yaml:"range_max"`// maximum range value [m]
	Ranges []float32 `yaml:"ranges"`// range data [m]
	Intensities []float32 `yaml:"intensities"`// intensity data [device-specific units].  If your. (Note: values < range_min or > range_max should be discarded)
}

// NewLaserScan creates a new LaserScan with default values.
func NewLaserScan() *LaserScan {
	self := LaserScan{}
	self.SetDefaults()
	return &self
}

func (t *LaserScan) Clone() *LaserScan {
	c := &LaserScan{}
	c.Header = *t.Header.Clone()
	c.AngleMin = t.AngleMin
	c.AngleMax = t.AngleMax
	c.AngleIncrement = t.AngleIncrement
	c.TimeIncrement = t.TimeIncrement
	c.ScanTime = t.ScanTime
	c.RangeMin = t.RangeMin
	c.RangeMax = t.RangeMax
	if t.Ranges != nil {
		c.Ranges = make([]float32, len(t.Ranges))
		copy(c.Ranges, t.Ranges)
	}
	if t.Intensities != nil {
		c.Intensities = make([]float32, len(t.Intensities))
		copy(c.Intensities, t.Intensities)
	}
	return c
}

func (t *LaserScan) CloneMsg() types.Message {
	return t.Clone()
}

func (t *LaserScan) SetDefaults() {
	t.Header.SetDefaults()
	t.AngleMin = 0
	t.AngleMax = 0
	t.AngleIncrement = 0
	t.TimeIncrement = 0
	t.ScanTime = 0
	t.RangeMin = 0
	t.RangeMax = 0
	t.Ranges = nil
	t.Intensities = nil
}

// LaserScanPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type LaserScanPublisher struct {
	*rclgo.Publisher
}

// NewLaserScanPublisher creates and returns a new publisher for the
// LaserScan
func NewLaserScanPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*LaserScanPublisher, error) {
	pub, err := node.NewPublisher(topic_name, LaserScanTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &LaserScanPublisher{pub}, nil
}

func (p *LaserScanPublisher) Publish(msg *LaserScan) error {
	return p.Publisher.Publish(msg)
}

// LaserScanSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type LaserScanSubscription struct {
	*rclgo.Subscription
}

// NewLaserScanSubscription creates and returns a new subscription for the
// LaserScan
func NewLaserScanSubscription(node *rclgo.Node, topic_name string, subscriptionCallback rclgo.SubscriptionCallback) (*LaserScanSubscription, error) {
	sub, err := node.NewSubscription(topic_name, LaserScanTypeSupport, subscriptionCallback)
	if err != nil {
		return nil, err
	}
	return &LaserScanSubscription{sub}, nil
}

func (s *LaserScanSubscription) TakeMessage(out *LaserScan) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneLaserScanSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneLaserScanSlice(dst, src []LaserScan) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var LaserScanTypeSupport types.MessageTypeSupport = _LaserScanTypeSupport{}

type _LaserScanTypeSupport struct{}

func (t _LaserScanTypeSupport) New() types.Message {
	return NewLaserScan()
}

func (t _LaserScanTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__LaserScan
	return (unsafe.Pointer)(C.sensor_msgs__msg__LaserScan__create())
}

func (t _LaserScanTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__LaserScan__destroy((*C.sensor_msgs__msg__LaserScan)(pointer_to_free))
}

func (t _LaserScanTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*LaserScan)
	mem := (*C.sensor_msgs__msg__LaserScan)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.angle_min = C.float(m.AngleMin)
	mem.angle_max = C.float(m.AngleMax)
	mem.angle_increment = C.float(m.AngleIncrement)
	mem.time_increment = C.float(m.TimeIncrement)
	mem.scan_time = C.float(m.ScanTime)
	mem.range_min = C.float(m.RangeMin)
	mem.range_max = C.float(m.RangeMax)
	primitives.Float32__Sequence_to_C((*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.ranges)), m.Ranges)
	primitives.Float32__Sequence_to_C((*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.intensities)), m.Intensities)
}

func (t _LaserScanTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*LaserScan)
	mem := (*C.sensor_msgs__msg__LaserScan)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.AngleMin = float32(mem.angle_min)
	m.AngleMax = float32(mem.angle_max)
	m.AngleIncrement = float32(mem.angle_increment)
	m.TimeIncrement = float32(mem.time_increment)
	m.ScanTime = float32(mem.scan_time)
	m.RangeMin = float32(mem.range_min)
	m.RangeMax = float32(mem.range_max)
	primitives.Float32__Sequence_to_Go(&m.Ranges, *(*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.ranges)))
	primitives.Float32__Sequence_to_Go(&m.Intensities, *(*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.intensities)))
}

func (t _LaserScanTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__LaserScan())
}

type CLaserScan = C.sensor_msgs__msg__LaserScan
type CLaserScan__Sequence = C.sensor_msgs__msg__LaserScan__Sequence

func LaserScan__Sequence_to_Go(goSlice *[]LaserScan, cSlice CLaserScan__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]LaserScan, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__LaserScan__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__LaserScan * uintptr(i)),
		))
		LaserScanTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func LaserScan__Sequence_to_C(cSlice *CLaserScan__Sequence, goSlice []LaserScan) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__LaserScan)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__LaserScan * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__LaserScan)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__LaserScan * uintptr(i)),
		))
		LaserScanTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func LaserScan__Array_to_Go(goSlice []LaserScan, cSlice []CLaserScan) {
	for i := 0; i < len(cSlice); i++ {
		LaserScanTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func LaserScan__Array_to_C(cSlice []CLaserScan, goSlice []LaserScan) {
	for i := 0; i < len(goSlice); i++ {
		LaserScanTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
