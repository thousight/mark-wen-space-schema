// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/services/pages/page.proto

package services

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_services_pages_page_proto protoreflect.FileDescriptor

var file_proto_services_pages_page_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x24, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xe1, 0x01, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x08,
	0x48, 0x6f, 0x6d, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x48,
	0x6f, 0x6d, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x47, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x53, 0x6f, 0x66, 0x74,
	0x77, 0x61, 0x72, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x6f, 0x75, 0x73, 0x69, 0x67, 0x68, 0x74, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x2d, 0x77, 0x65, 0x6e, 0x2d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2d, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_services_pages_page_proto_goTypes = []interface{}{
	(*HomePageRequest)(nil),      // 0: services.HomePageRequest
	(*ResumePageRequest)(nil),    // 1: services.ResumePageRequest
	(*SoftwarePageRequest)(nil),  // 2: services.SoftwarePageRequest
	(*HomePageResponse)(nil),     // 3: services.HomePageResponse
	(*ResumePageResponse)(nil),   // 4: services.ResumePageResponse
	(*SoftwarePageResponse)(nil), // 5: services.SoftwarePageResponse
}
var file_proto_services_pages_page_proto_depIdxs = []int32{
	0, // 0: services.Page.HomePage:input_type -> services.HomePageRequest
	1, // 1: services.Page.ResumePage:input_type -> services.ResumePageRequest
	2, // 2: services.Page.SoftwarePage:input_type -> services.SoftwarePageRequest
	3, // 3: services.Page.HomePage:output_type -> services.HomePageResponse
	4, // 4: services.Page.ResumePage:output_type -> services.ResumePageResponse
	5, // 5: services.Page.SoftwarePage:output_type -> services.SoftwarePageResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_services_pages_page_proto_init() }
func file_proto_services_pages_page_proto_init() {
	if File_proto_services_pages_page_proto != nil {
		return
	}
	file_proto_services_pages_home_page_proto_init()
	file_proto_services_pages_resume_page_proto_init()
	file_proto_services_pages_software_page_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_services_pages_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_services_pages_page_proto_goTypes,
		DependencyIndexes: file_proto_services_pages_page_proto_depIdxs,
	}.Build()
	File_proto_services_pages_page_proto = out.File
	file_proto_services_pages_page_proto_rawDesc = nil
	file_proto_services_pages_page_proto_goTypes = nil
	file_proto_services_pages_page_proto_depIdxs = nil
}
