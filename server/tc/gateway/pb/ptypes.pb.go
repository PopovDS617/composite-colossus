// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: gateway/pb/ptypes.proto

package pb

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

type GetInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OBUID int32 `protobuf:"varint,1,opt,name=OBUID,proto3" json:"OBUID,omitempty"`
}

func (x *GetInvoiceRequest) Reset() {
	*x = GetInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_pb_ptypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoiceRequest) ProtoMessage() {}

func (x *GetInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_pb_ptypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoiceRequest.ProtoReflect.Descriptor instead.
func (*GetInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_gateway_pb_ptypes_proto_rawDescGZIP(), []int{0}
}

func (x *GetInvoiceRequest) GetOBUID() int32 {
	if x != nil {
		return x.OBUID
	}
	return 0
}

var File_gateway_pb_ptypes_proto protoreflect.FileDescriptor

var file_gateway_pb_ptypes_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x4f, 0x42, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4f,
	0x42, 0x55, 0x49, 0x44, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_pb_ptypes_proto_rawDescOnce sync.Once
	file_gateway_pb_ptypes_proto_rawDescData = file_gateway_pb_ptypes_proto_rawDesc
)

func file_gateway_pb_ptypes_proto_rawDescGZIP() []byte {
	file_gateway_pb_ptypes_proto_rawDescOnce.Do(func() {
		file_gateway_pb_ptypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_pb_ptypes_proto_rawDescData)
	})
	return file_gateway_pb_ptypes_proto_rawDescData
}

var file_gateway_pb_ptypes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gateway_pb_ptypes_proto_goTypes = []interface{}{
	(*GetInvoiceRequest)(nil), // 0: GetInvoiceRequest
}
var file_gateway_pb_ptypes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gateway_pb_ptypes_proto_init() }
func file_gateway_pb_ptypes_proto_init() {
	if File_gateway_pb_ptypes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_pb_ptypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvoiceRequest); i {
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
			RawDescriptor: file_gateway_pb_ptypes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_pb_ptypes_proto_goTypes,
		DependencyIndexes: file_gateway_pb_ptypes_proto_depIdxs,
		MessageInfos:      file_gateway_pb_ptypes_proto_msgTypes,
	}.Build()
	File_gateway_pb_ptypes_proto = out.File
	file_gateway_pb_ptypes_proto_rawDesc = nil
	file_gateway_pb_ptypes_proto_goTypes = nil
	file_gateway_pb_ptypes_proto_depIdxs = nil
}
