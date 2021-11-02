// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: runtime/internal/examplepb/proto2.proto

package examplepb

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

type Proto2Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FloatValue    *float32 `protobuf:"fixed32,42,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	DoubleValue   *float64 `protobuf:"fixed64,43,opt,name=double_value,json=doubleValue" json:"double_value,omitempty"`
	Int64Value    *int64   `protobuf:"varint,3,opt,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	Int32Value    *int32   `protobuf:"varint,4,opt,name=int32_value,json=int32Value" json:"int32_value,omitempty"`
	Uint64Value   *uint64  `protobuf:"varint,5,opt,name=uint64_value,json=uint64Value" json:"uint64_value,omitempty"`
	Uint32Value   *uint32  `protobuf:"varint,6,opt,name=uint32_value,json=uint32Value" json:"uint32_value,omitempty"`
	BoolValue     *bool    `protobuf:"varint,7,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	StringValue   *string  `protobuf:"bytes,8,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	BytesValue    []byte   `protobuf:"bytes,9,opt,name=bytes_value,json=bytesValue" json:"bytes_value,omitempty"`
	RepeatedValue []string `protobuf:"bytes,10,rep,name=repeated_value,json=repeatedValue" json:"repeated_value,omitempty"`
}

func (x *Proto2Message) Reset() {
	*x = Proto2Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runtime_internal_examplepb_proto2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proto2Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proto2Message) ProtoMessage() {}

func (x *Proto2Message) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_internal_examplepb_proto2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proto2Message.ProtoReflect.Descriptor instead.
func (*Proto2Message) Descriptor() ([]byte, []int) {
	return file_runtime_internal_examplepb_proto2_proto_rawDescGZIP(), []int{0}
}

func (x *Proto2Message) GetFloatValue() float32 {
	if x != nil && x.FloatValue != nil {
		return *x.FloatValue
	}
	return 0
}

func (x *Proto2Message) GetDoubleValue() float64 {
	if x != nil && x.DoubleValue != nil {
		return *x.DoubleValue
	}
	return 0
}

func (x *Proto2Message) GetInt64Value() int64 {
	if x != nil && x.Int64Value != nil {
		return *x.Int64Value
	}
	return 0
}

func (x *Proto2Message) GetInt32Value() int32 {
	if x != nil && x.Int32Value != nil {
		return *x.Int32Value
	}
	return 0
}

func (x *Proto2Message) GetUint64Value() uint64 {
	if x != nil && x.Uint64Value != nil {
		return *x.Uint64Value
	}
	return 0
}

func (x *Proto2Message) GetUint32Value() uint32 {
	if x != nil && x.Uint32Value != nil {
		return *x.Uint32Value
	}
	return 0
}

func (x *Proto2Message) GetBoolValue() bool {
	if x != nil && x.BoolValue != nil {
		return *x.BoolValue
	}
	return false
}

func (x *Proto2Message) GetStringValue() string {
	if x != nil && x.StringValue != nil {
		return *x.StringValue
	}
	return ""
}

func (x *Proto2Message) GetBytesValue() []byte {
	if x != nil {
		return x.BytesValue
	}
	return nil
}

func (x *Proto2Message) GetRepeatedValue() []string {
	if x != nil {
		return x.RepeatedValue
	}
	return nil
}

var File_runtime_internal_examplepb_proto2_proto protoreflect.FileDescriptor

var file_runtime_internal_examplepb_proto2_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x22, 0xe5, 0x02, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63,
	0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62,
}

var (
	file_runtime_internal_examplepb_proto2_proto_rawDescOnce sync.Once
	file_runtime_internal_examplepb_proto2_proto_rawDescData = file_runtime_internal_examplepb_proto2_proto_rawDesc
)

func file_runtime_internal_examplepb_proto2_proto_rawDescGZIP() []byte {
	file_runtime_internal_examplepb_proto2_proto_rawDescOnce.Do(func() {
		file_runtime_internal_examplepb_proto2_proto_rawDescData = protoimpl.X.CompressGZIP(file_runtime_internal_examplepb_proto2_proto_rawDescData)
	})
	return file_runtime_internal_examplepb_proto2_proto_rawDescData
}

var file_runtime_internal_examplepb_proto2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_runtime_internal_examplepb_proto2_proto_goTypes = []interface{}{
	(*Proto2Message)(nil), // 0: grpc.gateway.runtime.internal.examplepb.Proto2Message
}
var file_runtime_internal_examplepb_proto2_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_runtime_internal_examplepb_proto2_proto_init() }
func file_runtime_internal_examplepb_proto2_proto_init() {
	if File_runtime_internal_examplepb_proto2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runtime_internal_examplepb_proto2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proto2Message); i {
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
			RawDescriptor: file_runtime_internal_examplepb_proto2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_runtime_internal_examplepb_proto2_proto_goTypes,
		DependencyIndexes: file_runtime_internal_examplepb_proto2_proto_depIdxs,
		MessageInfos:      file_runtime_internal_examplepb_proto2_proto_msgTypes,
	}.Build()
	File_runtime_internal_examplepb_proto2_proto = out.File
	file_runtime_internal_examplepb_proto2_proto_rawDesc = nil
	file_runtime_internal_examplepb_proto2_proto_goTypes = nil
	file_runtime_internal_examplepb_proto2_proto_depIdxs = nil
}
