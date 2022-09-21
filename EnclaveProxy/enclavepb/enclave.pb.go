//python -m grpc_tools.protoc -I . --python_out=. --grpc_python_out=. enclave.proto
//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative enclave.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: enclave.proto

package enclavepb

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enclave_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_enclave_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_enclave_proto_rawDescGZIP(), []int{0}
}

type OnboardTxn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash      string `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	TxnHash        string `protobuf:"bytes,2,opt,name=txnHash,proto3" json:"txnHash,omitempty"`
	Identification string `protobuf:"bytes,3,opt,name=identification,proto3" json:"identification,omitempty"`
	Batch          int64  `protobuf:"varint,4,opt,name=batch,proto3" json:"batch,omitempty"`
}

func (x *OnboardTxn) Reset() {
	*x = OnboardTxn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enclave_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnboardTxn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnboardTxn) ProtoMessage() {}

func (x *OnboardTxn) ProtoReflect() protoreflect.Message {
	mi := &file_enclave_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnboardTxn.ProtoReflect.Descriptor instead.
func (*OnboardTxn) Descriptor() ([]byte, []int) {
	return file_enclave_proto_rawDescGZIP(), []int{1}
}

func (x *OnboardTxn) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *OnboardTxn) GetTxnHash() string {
	if x != nil {
		return x.TxnHash
	}
	return ""
}

func (x *OnboardTxn) GetIdentification() string {
	if x != nil {
		return x.Identification
	}
	return ""
}

func (x *OnboardTxn) GetBatch() int64 {
	if x != nil {
		return x.Batch
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enclave_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_enclave_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_enclave_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type OffboardTxn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash      string `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	TxnHash        string `protobuf:"bytes,2,opt,name=txnHash,proto3" json:"txnHash,omitempty"`
	Identification string `protobuf:"bytes,3,opt,name=identification,proto3" json:"identification,omitempty"`
	Batch          int64  `protobuf:"varint,4,opt,name=batch,proto3" json:"batch,omitempty"`
}

func (x *OffboardTxn) Reset() {
	*x = OffboardTxn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enclave_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OffboardTxn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OffboardTxn) ProtoMessage() {}

func (x *OffboardTxn) ProtoReflect() protoreflect.Message {
	mi := &file_enclave_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OffboardTxn.ProtoReflect.Descriptor instead.
func (*OffboardTxn) Descriptor() ([]byte, []int) {
	return file_enclave_proto_rawDescGZIP(), []int{3}
}

func (x *OffboardTxn) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *OffboardTxn) GetTxnHash() string {
	if x != nil {
		return x.TxnHash
	}
	return ""
}

func (x *OffboardTxn) GetIdentification() string {
	if x != nil {
		return x.Identification
	}
	return ""
}

func (x *OffboardTxn) GetBatch() int64 {
	if x != nil {
		return x.Batch
	}
	return 0
}

var File_enclave_proto protoreflect.FileDescriptor

var file_enclave_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x6e, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x65, 0x6e, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x70, 0x62, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x54,
	0x78, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x78, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x78, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x4f,
	0x66, 0x66, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x54, 0x78, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x78, 0x6e, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x6e, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x32, 0x8d, 0x01, 0x0a, 0x07, 0x45, 0x6e, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x12, 0x3f, 0x0a, 0x11,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x54, 0x78,
	0x6e, 0x12, 0x15, 0x2e, 0x65, 0x6e, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x70, 0x62, 0x2e, 0x4f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x54, 0x78, 0x6e, 0x1a, 0x11, 0x2e, 0x65, 0x6e, 0x63, 0x6c, 0x61,
	0x76, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x12, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4f, 0x66, 0x66, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x54, 0x78, 0x6e, 0x12, 0x16, 0x2e, 0x65, 0x6e, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x70, 0x62, 0x2e,
	0x4f, 0x66, 0x66, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x54, 0x78, 0x6e, 0x1a, 0x11, 0x2e, 0x65, 0x6e,
	0x63, 0x6c, 0x61, 0x76, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x65, 0x6e, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enclave_proto_rawDescOnce sync.Once
	file_enclave_proto_rawDescData = file_enclave_proto_rawDesc
)

func file_enclave_proto_rawDescGZIP() []byte {
	file_enclave_proto_rawDescOnce.Do(func() {
		file_enclave_proto_rawDescData = protoimpl.X.CompressGZIP(file_enclave_proto_rawDescData)
	})
	return file_enclave_proto_rawDescData
}

var file_enclave_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_enclave_proto_goTypes = []interface{}{
	(*Empty)(nil),       // 0: enclavepb.Empty
	(*OnboardTxn)(nil),  // 1: enclavepb.OnboardTxn
	(*Status)(nil),      // 2: enclavepb.Status
	(*OffboardTxn)(nil), // 3: enclavepb.OffboardTxn
}
var file_enclave_proto_depIdxs = []int32{
	1, // 0: enclavepb.Enclave.ReceiveOnboardTxn:input_type -> enclavepb.OnboardTxn
	3, // 1: enclavepb.Enclave.ReceiveOffboardTxn:input_type -> enclavepb.OffboardTxn
	2, // 2: enclavepb.Enclave.ReceiveOnboardTxn:output_type -> enclavepb.Status
	2, // 3: enclavepb.Enclave.ReceiveOffboardTxn:output_type -> enclavepb.Status
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enclave_proto_init() }
func file_enclave_proto_init() {
	if File_enclave_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enclave_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_enclave_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnboardTxn); i {
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
		file_enclave_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_enclave_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OffboardTxn); i {
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
			RawDescriptor: file_enclave_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_enclave_proto_goTypes,
		DependencyIndexes: file_enclave_proto_depIdxs,
		MessageInfos:      file_enclave_proto_msgTypes,
	}.Build()
	File_enclave_proto = out.File
	file_enclave_proto_rawDesc = nil
	file_enclave_proto_goTypes = nil
	file_enclave_proto_depIdxs = nil
}
