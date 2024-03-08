// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/annotations/v1/annotations.proto

package pb

import (
	code "google.golang.org/genproto/googleapis/rpc/code"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MethodRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes             []code.Code `protobuf:"varint,1,rep,packed,name=codes,proto3,enum=google.rpc.Code" json:"codes,omitempty"`
	NeedAuthorization bool        `protobuf:"varint,2,opt,name=need_authorization,json=needAuthorization,proto3" json:"need_authorization,omitempty"`
	ResponseMetadata  []string    `protobuf:"bytes,3,rep,name=response_metadata,json=responseMetadata,proto3" json:"response_metadata,omitempty"`
}

func (x *MethodRule) Reset() {
	*x = MethodRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_annotations_v1_annotations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodRule) ProtoMessage() {}

func (x *MethodRule) ProtoReflect() protoreflect.Message {
	mi := &file_api_annotations_v1_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodRule.ProtoReflect.Descriptor instead.
func (*MethodRule) Descriptor() ([]byte, []int) {
	return file_api_annotations_v1_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *MethodRule) GetCodes() []code.Code {
	if x != nil {
		return x.Codes
	}
	return nil
}

func (x *MethodRule) GetNeedAuthorization() bool {
	if x != nil {
		return x.NeedAuthorization
	}
	return false
}

func (x *MethodRule) GetResponseMetadata() []string {
	if x != nil {
		return x.ResponseMetadata
	}
	return nil
}

type ServiceRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes []code.Code `protobuf:"varint,1,rep,packed,name=codes,proto3,enum=google.rpc.Code" json:"codes,omitempty"`
}

func (x *ServiceRule) Reset() {
	*x = ServiceRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_annotations_v1_annotations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRule) ProtoMessage() {}

func (x *ServiceRule) ProtoReflect() protoreflect.Message {
	mi := &file_api_annotations_v1_annotations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceRule.ProtoReflect.Descriptor instead.
func (*ServiceRule) Descriptor() ([]byte, []int) {
	return file_api_annotations_v1_annotations_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceRule) GetCodes() []code.Code {
	if x != nil {
		return x.Codes
	}
	return nil
}

var file_api_annotations_v1_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodRule)(nil),
		Field:         71599529,
		Name:          "api.annotations.v1.method_rule",
		Tag:           "bytes,71599529,opt,name=method_rule",
		Filename:      "api/annotations/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*ServiceRule)(nil),
		Field:         51562530,
		Name:          "api.annotations.v1.service_rule",
		Tag:           "bytes,51562530,opt,name=service_rule",
		Filename:      "api/annotations/v1/annotations.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional api.annotations.v1.MethodRule method_rule = 71599529;
	E_MethodRule = &file_api_annotations_v1_annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional api.annotations.v1.ServiceRule service_rule = 51562530;
	E_ServiceRule = &file_api_annotations_v1_annotations_proto_extTypes[1]
)

var File_api_annotations_v1_annotations_proto protoreflect.FileDescriptor

var file_api_annotations_v1_annotations_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x6e, 0x65,
	0x65, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6e, 0x65, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x3a, 0x62, 0x0a,
	0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa9, 0x8b, 0x92,
	0x22, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x75, 0x6c,
	0x65, 0x3a, 0x66, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xa2, 0x90, 0xcb, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x65, 0x72, 0x67, 0x73, 0x4c, 0x61, 0x77,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_annotations_v1_annotations_proto_rawDescOnce sync.Once
	file_api_annotations_v1_annotations_proto_rawDescData = file_api_annotations_v1_annotations_proto_rawDesc
)

func file_api_annotations_v1_annotations_proto_rawDescGZIP() []byte {
	file_api_annotations_v1_annotations_proto_rawDescOnce.Do(func() {
		file_api_annotations_v1_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_annotations_v1_annotations_proto_rawDescData)
	})
	return file_api_annotations_v1_annotations_proto_rawDescData
}

var file_api_annotations_v1_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_annotations_v1_annotations_proto_goTypes = []interface{}{
	(*MethodRule)(nil),                  // 0: api.annotations.v1.MethodRule
	(*ServiceRule)(nil),                 // 1: api.annotations.v1.ServiceRule
	(code.Code)(0),                      // 2: google.rpc.Code
	(*descriptorpb.MethodOptions)(nil),  // 3: google.protobuf.MethodOptions
	(*descriptorpb.ServiceOptions)(nil), // 4: google.protobuf.ServiceOptions
}
var file_api_annotations_v1_annotations_proto_depIdxs = []int32{
	2, // 0: api.annotations.v1.MethodRule.codes:type_name -> google.rpc.Code
	2, // 1: api.annotations.v1.ServiceRule.codes:type_name -> google.rpc.Code
	3, // 2: api.annotations.v1.method_rule:extendee -> google.protobuf.MethodOptions
	4, // 3: api.annotations.v1.service_rule:extendee -> google.protobuf.ServiceOptions
	0, // 4: api.annotations.v1.method_rule:type_name -> api.annotations.v1.MethodRule
	1, // 5: api.annotations.v1.service_rule:type_name -> api.annotations.v1.ServiceRule
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	4, // [4:6] is the sub-list for extension type_name
	2, // [2:4] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_annotations_v1_annotations_proto_init() }
func file_api_annotations_v1_annotations_proto_init() {
	if File_api_annotations_v1_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_annotations_v1_annotations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodRule); i {
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
		file_api_annotations_v1_annotations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceRule); i {
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
			RawDescriptor: file_api_annotations_v1_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_api_annotations_v1_annotations_proto_goTypes,
		DependencyIndexes: file_api_annotations_v1_annotations_proto_depIdxs,
		MessageInfos:      file_api_annotations_v1_annotations_proto_msgTypes,
		ExtensionInfos:    file_api_annotations_v1_annotations_proto_extTypes,
	}.Build()
	File_api_annotations_v1_annotations_proto = out.File
	file_api_annotations_v1_annotations_proto_rawDesc = nil
	file_api_annotations_v1_annotations_proto_goTypes = nil
	file_api_annotations_v1_annotations_proto_depIdxs = nil
}
