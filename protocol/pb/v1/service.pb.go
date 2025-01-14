// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: pb/v1/service.proto

package pb

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SampleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Guid          string                 `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Id            int32                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SampleRequest) Reset() {
	*x = SampleRequest{}
	mi := &file_pb_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleRequest) ProtoMessage() {}

func (x *SampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleRequest.ProtoReflect.Descriptor instead.
func (*SampleRequest) Descriptor() ([]byte, []int) {
	return file_pb_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *SampleRequest) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *SampleRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SampleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          string                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SampleResponse) Reset() {
	*x = SampleResponse{}
	mi := &file_pb_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SampleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleResponse) ProtoMessage() {}

func (x *SampleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleResponse.ProtoReflect.Descriptor instead.
func (*SampleResponse) Descriptor() ([]byte, []int) {
	return file_pb_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *SampleResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type StreamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Seed          int64                  `protobuf:"varint,1,opt,name=seed,proto3" json:"seed,omitempty"`
	Iteration     int32                  `protobuf:"varint,2,opt,name=iteration,proto3" json:"iteration,omitempty"`
	SectionLength int32                  `protobuf:"varint,3,opt,name=section_length,json=sectionLength,proto3" json:"section_length,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	mi := &file_pb_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_pb_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *StreamRequest) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *StreamRequest) GetIteration() int32 {
	if x != nil {
		return x.Iteration
	}
	return 0
}

func (x *StreamRequest) GetSectionLength() int32 {
	if x != nil {
		return x.SectionLength
	}
	return 0
}

type StreamResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Payload       []byte                 `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	mi := &file_pb_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_pb_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *StreamResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_pb_v1_service_proto protoreflect.FileDescriptor

var file_pb_v1_service_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb3, 0x01, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0b, 0xba, 0x48, 0x08, 0xd8, 0x01, 0x01, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x04, 0x67, 0x75,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07,
	0xba, 0x48, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x68, 0xba, 0x48, 0x65,
	0x1a, 0x63, 0x0a, 0x1f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x6f, 0x6e, 0x65,
	0x5f, 0x6f, 0x66, 0x12, 0x1f, 0x63, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x20, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x79, 0x20, 0x62, 0x6f, 0x74, 0x68, 0x20, 0x67, 0x75, 0x69, 0x64, 0x20, 0x61, 0x6e,
	0x64, 0x20, 0x69, 0x64, 0x1a, 0x1f, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x67, 0x75, 0x69, 0x64, 0x20,
	0x3d, 0x3d, 0x20, 0x27, 0x27, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x64,
	0x20, 0x3d, 0x3d, 0x20, 0x30, 0x22, 0x24, 0x0a, 0x0e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x0d, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x32, 0x9d, 0x03, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5e, 0x0a, 0x0e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a,
	0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2f, 0x7b, 0x73, 0x65, 0x65, 0x64, 0x7d, 0x30, 0x01, 0x12, 0x69, 0x0a, 0x12,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x74,
	0x74, 0x70, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x7b,
	0x73, 0x65, 0x65, 0x64, 0x7d, 0x30, 0x01, 0x12, 0x68, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x74, 0x74, 0x70, 0x12, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64,
	0x79, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a,
	0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x28,
	0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_pb_v1_service_proto_rawDescOnce sync.Once
	file_pb_v1_service_proto_rawDescData []byte
)

func file_pb_v1_service_proto_rawDescGZIP() []byte {
	file_pb_v1_service_proto_rawDescOnce.Do(func() {
		file_pb_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pb_v1_service_proto_rawDesc), len(file_pb_v1_service_proto_rawDesc)))
	})
	return file_pb_v1_service_proto_rawDescData
}

var file_pb_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_v1_service_proto_goTypes = []any{
	(*SampleRequest)(nil),     // 0: pb.SampleRequest
	(*SampleResponse)(nil),    // 1: pb.SampleResponse
	(*StreamRequest)(nil),     // 2: pb.StreamRequest
	(*StreamResponse)(nil),    // 3: pb.StreamResponse
	(*httpbody.HttpBody)(nil), // 4: google.api.HttpBody
}
var file_pb_v1_service_proto_depIdxs = []int32{
	0, // 0: pb.SampleService.RetrieveSampleData:input_type -> pb.SampleRequest
	2, // 1: pb.SampleService.StreamDownload:input_type -> pb.StreamRequest
	2, // 2: pb.SampleService.StreamDownloadHttp:input_type -> pb.StreamRequest
	4, // 3: pb.SampleService.StreamUploadHttp:input_type -> google.api.HttpBody
	1, // 4: pb.SampleService.RetrieveSampleData:output_type -> pb.SampleResponse
	3, // 5: pb.SampleService.StreamDownload:output_type -> pb.StreamResponse
	4, // 6: pb.SampleService.StreamDownloadHttp:output_type -> google.api.HttpBody
	1, // 7: pb.SampleService.StreamUploadHttp:output_type -> pb.SampleResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_v1_service_proto_init() }
func file_pb_v1_service_proto_init() {
	if File_pb_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pb_v1_service_proto_rawDesc), len(file_pb_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_v1_service_proto_goTypes,
		DependencyIndexes: file_pb_v1_service_proto_depIdxs,
		MessageInfos:      file_pb_v1_service_proto_msgTypes,
	}.Build()
	File_pb_v1_service_proto = out.File
	file_pb_v1_service_proto_goTypes = nil
	file_pb_v1_service_proto_depIdxs = nil
}
