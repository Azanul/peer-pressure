// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pkg/pressure/pb/pressure.proto

package pb

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

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Len      *int32  `protobuf:"varint,1,opt,name=len,proto3,oneof" json:"len,omitempty"` // No. of bytes of data in this chunk
	Index    int32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`   // Index of the chunk being sent
	Data     []byte  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Filename *string `protobuf:"bytes,4,opt,name=filename,proto3,oneof" json:"filename,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pressure_pb_pressure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pressure_pb_pressure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_pkg_pressure_pb_pressure_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetLen() int32 {
	if x != nil && x.Len != nil {
		return *x.Len
	}
	return 0
}

func (x *Chunk) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Chunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Chunk) GetFilename() string {
	if x != nil && x.Filename != nil {
		return *x.Filename
	}
	return ""
}

type ChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"` // Index of the chunk that we want
}

func (x *ChunkRequest) Reset() {
	*x = ChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pressure_pb_pressure_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkRequest) ProtoMessage() {}

func (x *ChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pressure_pb_pressure_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkRequest.ProtoReflect.Descriptor instead.
func (*ChunkRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pressure_pb_pressure_proto_rawDescGZIP(), []int{1}
}

func (x *ChunkRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NChunks  int32  `protobuf:"varint,1,opt,name=n_chunks,json=nChunks,proto3" json:"n_chunks,omitempty"` // No. of chunks in the file
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Progress int32  `protobuf:"varint,3,opt,name=progress,proto3" json:"progress,omitempty"` // Percentage of the file received
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pressure_pb_pressure_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pressure_pb_pressure_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_pkg_pressure_pb_pressure_proto_rawDescGZIP(), []int{2}
}

func (x *Index) GetNChunks() int32 {
	if x != nil {
		return x.NChunks
	}
	return 0
}

func (x *Index) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Index) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

var File_pkg_pressure_pb_pressure_proto protoreflect.FileDescriptor

var file_pkg_pressure_pb_pressure_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x62, 0x22, 0x7e, 0x0a,
	0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x15, 0x0a, 0x03, 0x6c, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x6c, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6c, 0x65, 0x6e,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a,
	0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x22, 0x5a, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08,
	0x6e, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72,
	0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pressure_pb_pressure_proto_rawDescOnce sync.Once
	file_pkg_pressure_pb_pressure_proto_rawDescData = file_pkg_pressure_pb_pressure_proto_rawDesc
)

func file_pkg_pressure_pb_pressure_proto_rawDescGZIP() []byte {
	file_pkg_pressure_pb_pressure_proto_rawDescOnce.Do(func() {
		file_pkg_pressure_pb_pressure_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pressure_pb_pressure_proto_rawDescData)
	})
	return file_pkg_pressure_pb_pressure_proto_rawDescData
}

var file_pkg_pressure_pb_pressure_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_pressure_pb_pressure_proto_goTypes = []interface{}{
	(*Chunk)(nil),        // 0: pressure.pb.Chunk
	(*ChunkRequest)(nil), // 1: pressure.pb.ChunkRequest
	(*Index)(nil),        // 2: pressure.pb.Index
}
var file_pkg_pressure_pb_pressure_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pressure_pb_pressure_proto_init() }
func file_pkg_pressure_pb_pressure_proto_init() {
	if File_pkg_pressure_pb_pressure_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pressure_pb_pressure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_pkg_pressure_pb_pressure_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkRequest); i {
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
		file_pkg_pressure_pb_pressure_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
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
	file_pkg_pressure_pb_pressure_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_pressure_pb_pressure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_pressure_pb_pressure_proto_goTypes,
		DependencyIndexes: file_pkg_pressure_pb_pressure_proto_depIdxs,
		MessageInfos:      file_pkg_pressure_pb_pressure_proto_msgTypes,
	}.Build()
	File_pkg_pressure_pb_pressure_proto = out.File
	file_pkg_pressure_pb_pressure_proto_rawDesc = nil
	file_pkg_pressure_pb_pressure_proto_goTypes = nil
	file_pkg_pressure_pb_pressure_proto_depIdxs = nil
}
