// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: grpc_OpenWeather/protobuf/open_weather.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type Temperature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location  string  `protobuf:"bytes,1,opt,name=Location,proto3" json:"Location,omitempty"`
	Temp      float64 `protobuf:"fixed64,2,opt,name=Temp,proto3" json:"Temp,omitempty"`                      // current Temperature
	TempMin   float64 `protobuf:"fixed64,3,opt,name=Temp_min,json=TempMin,proto3" json:"Temp_min,omitempty"` // lowest Temperatures
	TempMax   float64 `protobuf:"fixed64,4,opt,name=Temp_max,json=TempMax,proto3" json:"Temp_max,omitempty"` // higest Temperatures
	TimeStamp int64   `protobuf:"varint,5,opt,name=TimeStamp,proto3" json:"TimeStamp,omitempty"`
}

func (x *Temperature) Reset() {
	*x = Temperature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_OpenWeather_protobuf_open_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Temperature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Temperature) ProtoMessage() {}

func (x *Temperature) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_OpenWeather_protobuf_open_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Temperature.ProtoReflect.Descriptor instead.
func (*Temperature) Descriptor() ([]byte, []int) {
	return file_grpc_OpenWeather_protobuf_open_weather_proto_rawDescGZIP(), []int{0}
}

func (x *Temperature) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Temperature) GetTemp() float64 {
	if x != nil {
		return x.Temp
	}
	return 0
}

func (x *Temperature) GetTempMin() float64 {
	if x != nil {
		return x.TempMin
	}
	return 0
}

func (x *Temperature) GetTempMax() float64 {
	if x != nil {
		return x.TempMax
	}
	return 0
}

func (x *Temperature) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

// A feature names something at a given point.
//
// If a feature could not be named, the name is empty.
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"` // The Location
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_OpenWeather_protobuf_open_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_OpenWeather_protobuf_open_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_grpc_OpenWeather_protobuf_open_weather_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

var File_grpc_OpenWeather_protobuf_open_weather_proto protoreflect.FileDescriptor

var file_grpc_OpenWeather_protobuf_open_weather_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x4f, 0x70, 0x65, 0x6e, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x5f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x6f, 0x70, 0x65, 0x6e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x22, 0x91, 0x01, 0x0a, 0x0b,
	0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x54,
	0x65, 0x6d, 0x70, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x54,
	0x65, 0x6d, 0x70, 0x4d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x5f, 0x6d,
	0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x54, 0x65, 0x6d, 0x70, 0x4d, 0x61,
	0x78, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x26, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x5c, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x6e, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x6d, 0x61, 0x7a, 0x61, 0x6e, 0x44, 0x65, 0x6d, 0x69, 0x72,
	0x63, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_OpenWeather_protobuf_open_weather_proto_rawDescOnce sync.Once
	file_grpc_OpenWeather_protobuf_open_weather_proto_rawDescData = file_grpc_OpenWeather_protobuf_open_weather_proto_rawDesc
)

func file_grpc_OpenWeather_protobuf_open_weather_proto_rawDescGZIP() []byte {
	file_grpc_OpenWeather_protobuf_open_weather_proto_rawDescOnce.Do(func() {
		file_grpc_OpenWeather_protobuf_open_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_OpenWeather_protobuf_open_weather_proto_rawDescData)
	})
	return file_grpc_OpenWeather_protobuf_open_weather_proto_rawDescData
}

var file_grpc_OpenWeather_protobuf_open_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_OpenWeather_protobuf_open_weather_proto_goTypes = []interface{}{
	(*Temperature)(nil), // 0: openweather.Temperature
	(*Location)(nil),    // 1: openweather.Location
}
var file_grpc_OpenWeather_protobuf_open_weather_proto_depIdxs = []int32{
	1, // 0: openweather.OpenWeather.GetTemperatureByLocation:input_type -> openweather.Location
	0, // 1: openweather.OpenWeather.GetTemperatureByLocation:output_type -> openweather.Temperature
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_OpenWeather_protobuf_open_weather_proto_init() }
func file_grpc_OpenWeather_protobuf_open_weather_proto_init() {
	if File_grpc_OpenWeather_protobuf_open_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_OpenWeather_protobuf_open_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Temperature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_OpenWeather_protobuf_open_weather_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_OpenWeather_protobuf_open_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_OpenWeather_protobuf_open_weather_proto_goTypes,
		DependencyIndexes: file_grpc_OpenWeather_protobuf_open_weather_proto_depIdxs,
		MessageInfos:      file_grpc_OpenWeather_protobuf_open_weather_proto_msgTypes,
	}.Build()
	File_grpc_OpenWeather_protobuf_open_weather_proto = out.File
	file_grpc_OpenWeather_protobuf_open_weather_proto_rawDesc = nil
	file_grpc_OpenWeather_protobuf_open_weather_proto_goTypes = nil
	file_grpc_OpenWeather_protobuf_open_weather_proto_depIdxs = nil
}