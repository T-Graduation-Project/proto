// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: admin_server/admin_server.proto

package admin_server

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SignUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignUpReq) Reset() {
	*x = SignUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_server_admin_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpReq) ProtoMessage() {}

func (x *SignUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_server_admin_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpReq.ProtoReflect.Descriptor instead.
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return file_admin_server_admin_server_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignUpReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignUpRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignUpRsp) Reset() {
	*x = SignUpRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_server_admin_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRsp) ProtoMessage() {}

func (x *SignUpRsp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_server_admin_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRsp.ProtoReflect.Descriptor instead.
func (*SignUpRsp) Descriptor() ([]byte, []int) {
	return file_admin_server_admin_server_proto_rawDescGZIP(), []int{1}
}

var File_admin_server_admin_server_proto protoreflect.FileDescriptor

var file_admin_server_admin_server_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22,
	0x43, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x0b, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x73,
	0x70, 0x32, 0x4c, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3c, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x17, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55,
	0x70, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42,
	0x10, 0x5a, 0x0e, 0x2e, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_server_admin_server_proto_rawDescOnce sync.Once
	file_admin_server_admin_server_proto_rawDescData = file_admin_server_admin_server_proto_rawDesc
)

func file_admin_server_admin_server_proto_rawDescGZIP() []byte {
	file_admin_server_admin_server_proto_rawDescOnce.Do(func() {
		file_admin_server_admin_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_server_admin_server_proto_rawDescData)
	})
	return file_admin_server_admin_server_proto_rawDescData
}

var file_admin_server_admin_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_admin_server_admin_server_proto_goTypes = []interface{}{
	(*SignUpReq)(nil), // 0: admin_server.SignUpReq
	(*SignUpRsp)(nil), // 1: admin_server.SignUpRsp
}
var file_admin_server_admin_server_proto_depIdxs = []int32{
	0, // 0: admin_server.AdminService.SignUp:input_type -> admin_server.SignUpReq
	1, // 1: admin_server.AdminService.SignUp:output_type -> admin_server.SignUpRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_server_admin_server_proto_init() }
func file_admin_server_admin_server_proto_init() {
	if File_admin_server_admin_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_server_admin_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpReq); i {
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
		file_admin_server_admin_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRsp); i {
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
			RawDescriptor: file_admin_server_admin_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_server_admin_server_proto_goTypes,
		DependencyIndexes: file_admin_server_admin_server_proto_depIdxs,
		MessageInfos:      file_admin_server_admin_server_proto_msgTypes,
	}.Build()
	File_admin_server_admin_server_proto = out.File
	file_admin_server_admin_server_proto_rawDesc = nil
	file_admin_server_admin_server_proto_goTypes = nil
	file_admin_server_admin_server_proto_depIdxs = nil
}
