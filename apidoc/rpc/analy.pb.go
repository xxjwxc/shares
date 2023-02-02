// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: shares/analy.proto

package shares

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	_ "rpc/common"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AnalyCodeReq 请求
type AnalyCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code"` // 股票代码
}

func (x *AnalyCodeReq) Reset() {
	*x = AnalyCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shares_analy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyCodeReq) ProtoMessage() {}

func (x *AnalyCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_shares_analy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyCodeReq.ProtoReflect.Descriptor instead.
func (*AnalyCodeReq) Descriptor() ([]byte, []int) {
	return file_shares_analy_proto_rawDescGZIP(), []int{0}
}

func (x *AnalyCodeReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

// AnalyCodeResp 返回
type AnalyCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msgs []string `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs"` // 内容
}

func (x *AnalyCodeResp) Reset() {
	*x = AnalyCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shares_analy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyCodeResp) ProtoMessage() {}

func (x *AnalyCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_shares_analy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyCodeResp.ProtoReflect.Descriptor instead.
func (*AnalyCodeResp) Descriptor() ([]byte, []int) {
	return file_shares_analy_proto_rawDescGZIP(), []int{1}
}

func (x *AnalyCodeResp) GetMsgs() []string {
	if x != nil {
		return x.Msgs
	}
	return nil
}

var File_shares_analy_proto protoreflect.FileDescriptor

var file_shares_analy_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x1a, 0x1a, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x23, 0x0a, 0x0d,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x73, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x73, 0x67,
	0x73, 0x32, 0x43, 0x0a, 0x05, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x09, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73,
	0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shares_analy_proto_rawDescOnce sync.Once
	file_shares_analy_proto_rawDescData = file_shares_analy_proto_rawDesc
)

func file_shares_analy_proto_rawDescGZIP() []byte {
	file_shares_analy_proto_rawDescOnce.Do(func() {
		file_shares_analy_proto_rawDescData = protoimpl.X.CompressGZIP(file_shares_analy_proto_rawDescData)
	})
	return file_shares_analy_proto_rawDescData
}

var file_shares_analy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shares_analy_proto_goTypes = []interface{}{
	(*AnalyCodeReq)(nil),  // 0: shares.AnalyCodeReq
	(*AnalyCodeResp)(nil), // 1: shares.AnalyCodeResp
}
var file_shares_analy_proto_depIdxs = []int32{
	0, // 0: shares.Analy.AnalyCode:input_type -> shares.AnalyCodeReq
	1, // 1: shares.Analy.AnalyCode:output_type -> shares.AnalyCodeResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shares_analy_proto_init() }
func file_shares_analy_proto_init() {
	if File_shares_analy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shares_analy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyCodeReq); i {
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
		file_shares_analy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyCodeResp); i {
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
			RawDescriptor: file_shares_analy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shares_analy_proto_goTypes,
		DependencyIndexes: file_shares_analy_proto_depIdxs,
		MessageInfos:      file_shares_analy_proto_msgTypes,
	}.Build()
	File_shares_analy_proto = out.File
	file_shares_analy_proto_rawDesc = nil
	file_shares_analy_proto_goTypes = nil
	file_shares_analy_proto_depIdxs = nil
}