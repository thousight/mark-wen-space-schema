// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/components/navbar.proto

package components

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

type Navbar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tabs   []*NavbarTab `protobuf:"bytes,1,rep,name=tabs,proto3" json:"tabs,omitempty"`
	Logo   *Image       `protobuf:"bytes,2,opt,name=logo,proto3" json:"logo,omitempty"`
	Themes []string     `protobuf:"bytes,3,rep,name=themes,proto3" json:"themes,omitempty"`
}

func (x *Navbar) Reset() {
	*x = Navbar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_components_navbar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Navbar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Navbar) ProtoMessage() {}

func (x *Navbar) ProtoReflect() protoreflect.Message {
	mi := &file_proto_components_navbar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Navbar.ProtoReflect.Descriptor instead.
func (*Navbar) Descriptor() ([]byte, []int) {
	return file_proto_components_navbar_proto_rawDescGZIP(), []int{0}
}

func (x *Navbar) GetTabs() []*NavbarTab {
	if x != nil {
		return x.Tabs
	}
	return nil
}

func (x *Navbar) GetLogo() *Image {
	if x != nil {
		return x.Logo
	}
	return nil
}

func (x *Navbar) GetThemes() []string {
	if x != nil {
		return x.Themes
	}
	return nil
}

type NavbarTab struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DisplayName  string       `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	RelativePath string       `protobuf:"bytes,3,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
	Tabs         []*NavbarTab `protobuf:"bytes,4,rep,name=tabs,proto3" json:"tabs,omitempty"`
}

func (x *NavbarTab) Reset() {
	*x = NavbarTab{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_components_navbar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NavbarTab) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NavbarTab) ProtoMessage() {}

func (x *NavbarTab) ProtoReflect() protoreflect.Message {
	mi := &file_proto_components_navbar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NavbarTab.ProtoReflect.Descriptor instead.
func (*NavbarTab) Descriptor() ([]byte, []int) {
	return file_proto_components_navbar_proto_rawDescGZIP(), []int{1}
}

func (x *NavbarTab) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NavbarTab) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *NavbarTab) GetRelativePath() string {
	if x != nil {
		return x.RelativePath
	}
	return ""
}

func (x *NavbarTab) GetTabs() []*NavbarTab {
	if x != nil {
		return x.Tabs
	}
	return nil
}

var File_proto_components_navbar_proto protoreflect.FileDescriptor

var file_proto_components_navbar_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x6e, 0x61, 0x76, 0x62, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x06, 0x4e, 0x61, 0x76,
	0x62, 0x61, 0x72, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x61, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4e,
	0x61, 0x76, 0x62, 0x61, 0x72, 0x54, 0x61, 0x62, 0x52, 0x04, 0x74, 0x61, 0x62, 0x73, 0x12, 0x25,
	0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x22, 0x8e, 0x01,
	0x0a, 0x09, 0x4e, 0x61, 0x76, 0x62, 0x61, 0x72, 0x54, 0x61, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x61, 0x62, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4e,
	0x61, 0x76, 0x62, 0x61, 0x72, 0x54, 0x61, 0x62, 0x52, 0x04, 0x74, 0x61, 0x62, 0x73, 0x42, 0x40,
	0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x6f,
	0x75, 0x73, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x2d, 0x77, 0x65, 0x6e, 0x2d,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_components_navbar_proto_rawDescOnce sync.Once
	file_proto_components_navbar_proto_rawDescData = file_proto_components_navbar_proto_rawDesc
)

func file_proto_components_navbar_proto_rawDescGZIP() []byte {
	file_proto_components_navbar_proto_rawDescOnce.Do(func() {
		file_proto_components_navbar_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_components_navbar_proto_rawDescData)
	})
	return file_proto_components_navbar_proto_rawDescData
}

var file_proto_components_navbar_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_components_navbar_proto_goTypes = []interface{}{
	(*Navbar)(nil),    // 0: components.Navbar
	(*NavbarTab)(nil), // 1: components.NavbarTab
	(*Image)(nil),     // 2: components.Image
}
var file_proto_components_navbar_proto_depIdxs = []int32{
	1, // 0: components.Navbar.tabs:type_name -> components.NavbarTab
	2, // 1: components.Navbar.logo:type_name -> components.Image
	1, // 2: components.NavbarTab.tabs:type_name -> components.NavbarTab
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_components_navbar_proto_init() }
func file_proto_components_navbar_proto_init() {
	if File_proto_components_navbar_proto != nil {
		return
	}
	file_proto_components_image_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_components_navbar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Navbar); i {
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
		file_proto_components_navbar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NavbarTab); i {
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
			RawDescriptor: file_proto_components_navbar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_components_navbar_proto_goTypes,
		DependencyIndexes: file_proto_components_navbar_proto_depIdxs,
		MessageInfos:      file_proto_components_navbar_proto_msgTypes,
	}.Build()
	File_proto_components_navbar_proto = out.File
	file_proto_components_navbar_proto_rawDesc = nil
	file_proto_components_navbar_proto_goTypes = nil
	file_proto_components_navbar_proto_depIdxs = nil
}
