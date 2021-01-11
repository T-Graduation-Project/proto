// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: demo_server/demo.proto

package proto

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

type HelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req string `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_server_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_demo_server_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_demo_server_demo_proto_rawDescGZIP(), []int{0}
}

func (x *HelloReq) GetReq() string {
	if x != nil {
		return x.Req
	}
	return ""
}

type HelloRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res string `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *HelloRsp) Reset() {
	*x = HelloRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_server_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRsp) ProtoMessage() {}

func (x *HelloRsp) ProtoReflect() protoreflect.Message {
	mi := &file_demo_server_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRsp.ProtoReflect.Descriptor instead.
func (*HelloRsp) Descriptor() ([]byte, []int) {
	return file_demo_server_demo_proto_rawDescGZIP(), []int{1}
}

func (x *HelloRsp) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

var File_demo_server_demo_proto protoreflect.FileDescriptor

var file_demo_server_demo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x65, 0x6d, 0x6f, 0x22, 0x1c,
	0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65,
	0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x1c, 0x0a, 0x08,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x73, 0x32, 0x32, 0x0a, 0x05, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x29, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0e, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_demo_server_demo_proto_rawDescOnce sync.Once
	file_demo_server_demo_proto_rawDescData = file_demo_server_demo_proto_rawDesc
)

func file_demo_server_demo_proto_rawDescGZIP() []byte {
	file_demo_server_demo_proto_rawDescOnce.Do(func() {
		file_demo_server_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_server_demo_proto_rawDescData)
	})
	return file_demo_server_demo_proto_rawDescData
}

var file_demo_server_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_demo_server_demo_proto_goTypes = []interface{}{
	(*HelloReq)(nil), // 0: demo.HelloReq
	(*HelloRsp)(nil), // 1: demo.HelloRsp
}
var file_demo_server_demo_proto_depIdxs = []int32{
	0, // 0: demo.Hello.Hello:input_type -> demo.HelloReq
	1, // 1: demo.Hello.Hello:output_type -> demo.HelloRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_demo_server_demo_proto_init() }
func file_demo_server_demo_proto_init() {
	if File_demo_server_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demo_server_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq); i {
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
		file_demo_server_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRsp); i {
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
			RawDescriptor: file_demo_server_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_server_demo_proto_goTypes,
		DependencyIndexes: file_demo_server_demo_proto_depIdxs,
		MessageInfos:      file_demo_server_demo_proto_msgTypes,
	}.Build()
	File_demo_server_demo_proto = out.File
	file_demo_server_demo_proto_rawDesc = nil
	file_demo_server_demo_proto_goTypes = nil
	file_demo_server_demo_proto_depIdxs = nil
}
