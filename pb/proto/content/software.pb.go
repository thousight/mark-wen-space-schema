// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/content/software.proto

package content

import (
	components "github.com/thousight/mark-wen-space-schema/pb/proto/components"
	primitives "github.com/thousight/mark-wen-space-schema/pb/proto/primitives"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Software struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string                    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description  string                    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Image        *components.Image         `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	TimeRange    *primitives.DateTimeRange `protobuf:"bytes,5,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	Icons        []*components.Image       `protobuf:"bytes,6,rep,name=icons,proto3" json:"icons,omitempty"`
	Platforms    []string                  `protobuf:"bytes,7,rep,name=platforms,proto3" json:"platforms,omitempty"`
	Frameworks   []string                  `protobuf:"bytes,8,rep,name=frameworks,proto3" json:"frameworks,omitempty"`
	Links        []*components.Link        `protobuf:"bytes,9,rep,name=links,proto3" json:"links,omitempty"`
	DateCreated  *timestamppb.Timestamp    `protobuf:"bytes,10,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	LastUpdated  *timestamppb.Timestamp    `protobuf:"bytes,11,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	CreatedBy    string                    `protobuf:"bytes,12,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	LastEditedBy string                    `protobuf:"bytes,13,opt,name=last_edited_by,json=lastEditedBy,proto3" json:"last_edited_by,omitempty"`
}

func (x *Software) Reset() {
	*x = Software{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_content_software_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Software) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Software) ProtoMessage() {}

func (x *Software) ProtoReflect() protoreflect.Message {
	mi := &file_proto_content_software_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Software.ProtoReflect.Descriptor instead.
func (*Software) Descriptor() ([]byte, []int) {
	return file_proto_content_software_proto_rawDescGZIP(), []int{0}
}

func (x *Software) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Software) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Software) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Software) GetImage() *components.Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Software) GetTimeRange() *primitives.DateTimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *Software) GetIcons() []*components.Image {
	if x != nil {
		return x.Icons
	}
	return nil
}

func (x *Software) GetPlatforms() []string {
	if x != nil {
		return x.Platforms
	}
	return nil
}

func (x *Software) GetFrameworks() []string {
	if x != nil {
		return x.Frameworks
	}
	return nil
}

func (x *Software) GetLinks() []*components.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *Software) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

func (x *Software) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *Software) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Software) GetLastEditedBy() string {
	if x != nil {
		return x.LastEditedBy
	}
	return ""
}

var File_proto_content_software_proto protoreflect.FileDescriptor

var file_proto_content_software_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f,
	0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x73, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x87, 0x04, 0x0a, 0x08, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x38, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73,
	0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x63, 0x6f,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x63, 0x6f,
	0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x12, 0x26, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x64,
	0x69, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x65, 0x64, 0x42, 0x79, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x6f, 0x75, 0x73, 0x69,
	0x67, 0x68, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x2d, 0x77, 0x65, 0x6e, 0x2d, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_content_software_proto_rawDescOnce sync.Once
	file_proto_content_software_proto_rawDescData = file_proto_content_software_proto_rawDesc
)

func file_proto_content_software_proto_rawDescGZIP() []byte {
	file_proto_content_software_proto_rawDescOnce.Do(func() {
		file_proto_content_software_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_content_software_proto_rawDescData)
	})
	return file_proto_content_software_proto_rawDescData
}

var file_proto_content_software_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_content_software_proto_goTypes = []interface{}{
	(*Software)(nil),                 // 0: content.Software
	(*components.Image)(nil),         // 1: components.Image
	(*primitives.DateTimeRange)(nil), // 2: primitives.DateTimeRange
	(*components.Link)(nil),          // 3: components.Link
	(*timestamppb.Timestamp)(nil),    // 4: google.protobuf.Timestamp
}
var file_proto_content_software_proto_depIdxs = []int32{
	1, // 0: content.Software.image:type_name -> components.Image
	2, // 1: content.Software.time_range:type_name -> primitives.DateTimeRange
	1, // 2: content.Software.icons:type_name -> components.Image
	3, // 3: content.Software.links:type_name -> components.Link
	4, // 4: content.Software.date_created:type_name -> google.protobuf.Timestamp
	4, // 5: content.Software.last_updated:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_content_software_proto_init() }
func file_proto_content_software_proto_init() {
	if File_proto_content_software_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_content_software_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Software); i {
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
			RawDescriptor: file_proto_content_software_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_content_software_proto_goTypes,
		DependencyIndexes: file_proto_content_software_proto_depIdxs,
		MessageInfos:      file_proto_content_software_proto_msgTypes,
	}.Build()
	File_proto_content_software_proto = out.File
	file_proto_content_software_proto_rawDesc = nil
	file_proto_content_software_proto_goTypes = nil
	file_proto_content_software_proto_depIdxs = nil
}
