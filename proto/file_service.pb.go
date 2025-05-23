// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/file_service.proto

package proto

import (
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

type FileChunk struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Content       []byte                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	mi := &file_proto_file_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{0}
}

func (x *FileChunk) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileChunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type FileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	mi := &file_proto_file_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{1}
}

func (x *FileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type UploadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	mi := &file_proto_file_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{2}
}

func (x *UploadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_proto_file_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{3}
}

type FileInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	CreatedAt     int64                  `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     int64                  `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	mi := &file_proto_file_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{4}
}

func (x *FileInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *FileInfo) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Files         []*FileInfo            `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_proto_file_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListResponse) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_proto_file_service_proto protoreflect.FileDescriptor

const file_proto_file_service_proto_rawDesc = "" +
	"\n" +
	"\x18proto/file_service.proto\x12\vfileservice\"A\n" +
	"\tFileChunk\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\x12\x18\n" +
	"\acontent\x18\x02 \x01(\fR\acontent\")\n" +
	"\vFileRequest\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\"*\n" +
	"\x0eUploadResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"\r\n" +
	"\vListRequest\"d\n" +
	"\bFileInfo\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\x12\x1d\n" +
	"\n" +
	"created_at\x18\x02 \x01(\x03R\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x03 \x01(\x03R\tupdatedAt\";\n" +
	"\fListResponse\x12+\n" +
	"\x05files\x18\x01 \x03(\v2\x15.fileservice.FileInfoR\x05files2\xd8\x01\n" +
	"\vFileService\x12C\n" +
	"\n" +
	"UploadFile\x12\x16.fileservice.FileChunk\x1a\x1b.fileservice.UploadResponse(\x01\x12B\n" +
	"\fDownloadFile\x12\x18.fileservice.FileRequest\x1a\x16.fileservice.FileChunk0\x01\x12@\n" +
	"\tListFiles\x12\x18.fileservice.ListRequest\x1a\x19.fileservice.ListResponseB\tZ\a./protob\x06proto3"

var (
	file_proto_file_service_proto_rawDescOnce sync.Once
	file_proto_file_service_proto_rawDescData []byte
)

func file_proto_file_service_proto_rawDescGZIP() []byte {
	file_proto_file_service_proto_rawDescOnce.Do(func() {
		file_proto_file_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_file_service_proto_rawDesc), len(file_proto_file_service_proto_rawDesc)))
	})
	return file_proto_file_service_proto_rawDescData
}

var file_proto_file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_file_service_proto_goTypes = []any{
	(*FileChunk)(nil),      // 0: fileservice.FileChunk
	(*FileRequest)(nil),    // 1: fileservice.FileRequest
	(*UploadResponse)(nil), // 2: fileservice.UploadResponse
	(*ListRequest)(nil),    // 3: fileservice.ListRequest
	(*FileInfo)(nil),       // 4: fileservice.FileInfo
	(*ListResponse)(nil),   // 5: fileservice.ListResponse
}
var file_proto_file_service_proto_depIdxs = []int32{
	4, // 0: fileservice.ListResponse.files:type_name -> fileservice.FileInfo
	0, // 1: fileservice.FileService.UploadFile:input_type -> fileservice.FileChunk
	1, // 2: fileservice.FileService.DownloadFile:input_type -> fileservice.FileRequest
	3, // 3: fileservice.FileService.ListFiles:input_type -> fileservice.ListRequest
	2, // 4: fileservice.FileService.UploadFile:output_type -> fileservice.UploadResponse
	0, // 5: fileservice.FileService.DownloadFile:output_type -> fileservice.FileChunk
	5, // 6: fileservice.FileService.ListFiles:output_type -> fileservice.ListResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_file_service_proto_init() }
func file_proto_file_service_proto_init() {
	if File_proto_file_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_file_service_proto_rawDesc), len(file_proto_file_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_file_service_proto_goTypes,
		DependencyIndexes: file_proto_file_service_proto_depIdxs,
		MessageInfos:      file_proto_file_service_proto_msgTypes,
	}.Build()
	File_proto_file_service_proto = out.File
	file_proto_file_service_proto_goTypes = nil
	file_proto_file_service_proto_depIdxs = nil
}
