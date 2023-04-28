// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: protobuf/github_webhook.proto

package __

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

type GithubHook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UseSecret bool `protobuf:"varint,1,opt,name=use_secret,json=useSecret,proto3" json:"use_secret,omitempty"`
}

func (x *GithubHook) Reset() {
	*x = GithubHook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_github_webhook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubHook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubHook) ProtoMessage() {}

func (x *GithubHook) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_github_webhook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubHook.ProtoReflect.Descriptor instead.
func (*GithubHook) Descriptor() ([]byte, []int) {
	return file_protobuf_github_webhook_proto_rawDescGZIP(), []int{0}
}

func (x *GithubHook) GetUseSecret() bool {
	if x != nil {
		return x.UseSecret
	}
	return false
}

var File_protobuf_github_webhook_proto protoreflect.FileDescriptor

var file_protobuf_github_webhook_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2b, 0x0a, 0x0a, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x75, 0x73, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x03, 0x5a, 0x01,
	0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_github_webhook_proto_rawDescOnce sync.Once
	file_protobuf_github_webhook_proto_rawDescData = file_protobuf_github_webhook_proto_rawDesc
)

func file_protobuf_github_webhook_proto_rawDescGZIP() []byte {
	file_protobuf_github_webhook_proto_rawDescOnce.Do(func() {
		file_protobuf_github_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_github_webhook_proto_rawDescData)
	})
	return file_protobuf_github_webhook_proto_rawDescData
}

var file_protobuf_github_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_github_webhook_proto_goTypes = []interface{}{
	(*GithubHook)(nil), // 0: GithubHook
}
var file_protobuf_github_webhook_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_github_webhook_proto_init() }
func file_protobuf_github_webhook_proto_init() {
	if File_protobuf_github_webhook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_github_webhook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubHook); i {
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
			RawDescriptor: file_protobuf_github_webhook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_github_webhook_proto_goTypes,
		DependencyIndexes: file_protobuf_github_webhook_proto_depIdxs,
		MessageInfos:      file_protobuf_github_webhook_proto_msgTypes,
	}.Build()
	File_protobuf_github_webhook_proto = out.File
	file_protobuf_github_webhook_proto_rawDesc = nil
	file_protobuf_github_webhook_proto_goTypes = nil
	file_protobuf_github_webhook_proto_depIdxs = nil
}
