// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: envoy/config/filter/http/modsec/v2/modsec.proto

package ModSecurity_envoy

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Remote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Remote) Reset() {
	*x = Remote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_modsec_v2_modsec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Remote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Remote) ProtoMessage() {}

func (x *Remote) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_modsec_v2_modsec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Remote.ProtoReflect.Descriptor instead.
func (*Remote) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDescGZIP(), []int{0}
}

func (x *Remote) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Remote) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Decoder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If set, rules are loaded from this path
	RulesPath string `protobuf:"bytes,1,rep,name=rules_path,json=rulesPath,proto3" json:"rules_path,omitempty"`
	// If set, rules are loaded from this inline configuration.
	// Note, if both rules_path and rules_inline are set, rules_path is first loaded and afterwards rules_inline is loaded
	RulesInline string `protobuf:"bytes,2,rep,name=rules_inline,json=rulesInline,proto3" json:"rules_inline,omitempty"`
	// If set, it will takes rules from url set with a http header ModSec-key set to key
	Remotes []*Remote `protobuf:"bytes,3,rep,name=remotes,proto3" json:"remotes,omitempty"`
	// If set to true, if no errors occured during remote download, those rules will overwrite all rules.
	RemotesOverwriteOnSuccess bool `protobuf:"varint,4,opt,name=remotes_overwrite_on_success,json=remotesOverwriteOnSuccess,proto3" json:"remotes_overwrite_on_success,omitempty"`
}

func (x *Decoder) Reset() {
	*x = Decoder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_modsec_v2_modsec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Decoder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Decoder) ProtoMessage() {}

func (x *Decoder) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_modsec_v2_modsec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Decoder.ProtoReflect.Descriptor instead.
func (*Decoder) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDescGZIP(), []int{1}
}

func (x *Decoder) GetRulesPath() string {
	if x != nil {
		return x.RulesPath
	}
	return ""
}

func (x *Decoder) GetRulesInline() string {
	if x != nil {
		return x.RulesInline
	}
	return ""
}

func (x *Decoder) GetRemotes() []*Remote {
	if x != nil {
		return x.Remotes
	}
	return nil
}

func (x *Decoder) GetRemotesOverwriteOnSuccess() bool {
	if x != nil {
		return x.RemotesOverwriteOnSuccess
	}
	return false
}

var File_envoy_config_filter_http_modsec_v2_modsec_proto protoreflect.FileDescriptor

var file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x73, 0x65,
	0x63, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xbb, 0x01, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x73, 0x12, 0x3f, 0x0a, 0x1c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x76,
	0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x73, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x42, 0x77, 0x0a, 0x3a, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x42, 0x10, 0x4d, 0x6f, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x6f, 0x68, 0x68, 0x6e, 0x79, 0x64, 0x69, 0x6e, 0x68, 0x2f, 0x4d, 0x6f, 0x64, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDescOnce sync.Once
	file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDescData = file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDesc
)

func file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDescData)
	})
	return file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDescData
}

var file_envoy_config_filter_http_modsec_v2_modsec_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_filter_http_modsec_v2_modsec_proto_goTypes = []interface{}{
	(*Remote)(nil),  // 0: modsecurity.Remote
	(*Decoder)(nil), // 1: modsecurity.Decoder
}
var file_envoy_config_filter_http_modsec_v2_modsec_proto_depIdxs = []int32{
	0, // 0: modsecurity.Decoder.remotes:type_name -> modsecurity.Remote
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_http_modsec_v2_modsec_proto_init() }
func file_envoy_config_filter_http_modsec_v2_modsec_proto_init() {
	if File_envoy_config_filter_http_modsec_v2_modsec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_http_modsec_v2_modsec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Remote); i {
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
		file_envoy_config_filter_http_modsec_v2_modsec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Decoder); i {
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
			RawDescriptor: file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_http_modsec_v2_modsec_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_http_modsec_v2_modsec_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_http_modsec_v2_modsec_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_http_modsec_v2_modsec_proto = out.File
	file_envoy_config_filter_http_modsec_v2_modsec_proto_rawDesc = nil
	file_envoy_config_filter_http_modsec_v2_modsec_proto_goTypes = nil
	file_envoy_config_filter_http_modsec_v2_modsec_proto_depIdxs = nil
}
