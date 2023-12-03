// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: offline_ticket.proto

package steam

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

type Offline_Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncryptedTicket []byte `protobuf:"bytes,1,opt,name=encrypted_ticket,json=encryptedTicket" json:"encrypted_ticket,omitempty"`
	Signature       []byte `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	Kdf1            *int32 `protobuf:"varint,3,opt,name=kdf1" json:"kdf1,omitempty"`
	Salt1           []byte `protobuf:"bytes,4,opt,name=salt1" json:"salt1,omitempty"`
	Kdf2            *int32 `protobuf:"varint,5,opt,name=kdf2" json:"kdf2,omitempty"`
	Salt2           []byte `protobuf:"bytes,6,opt,name=salt2" json:"salt2,omitempty"`
}

func (x *Offline_Ticket) Reset() {
	*x = Offline_Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offline_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offline_Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offline_Ticket) ProtoMessage() {}

func (x *Offline_Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_offline_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offline_Ticket.ProtoReflect.Descriptor instead.
func (*Offline_Ticket) Descriptor() ([]byte, []int) {
	return file_offline_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *Offline_Ticket) GetEncryptedTicket() []byte {
	if x != nil {
		return x.EncryptedTicket
	}
	return nil
}

func (x *Offline_Ticket) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Offline_Ticket) GetKdf1() int32 {
	if x != nil && x.Kdf1 != nil {
		return *x.Kdf1
	}
	return 0
}

func (x *Offline_Ticket) GetSalt1() []byte {
	if x != nil {
		return x.Salt1
	}
	return nil
}

func (x *Offline_Ticket) GetKdf2() int32 {
	if x != nil && x.Kdf2 != nil {
		return *x.Kdf2
	}
	return 0
}

func (x *Offline_Ticket) GetSalt2() []byte {
	if x != nil {
		return x.Salt2
	}
	return nil
}

var File_offline_ticket_proto protoreflect.FileDescriptor

var file_offline_ticket_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x0e, 0x4f, 0x66, 0x66, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x64, 0x66, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x6b, 0x64, 0x66, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x61, 0x6c, 0x74, 0x31, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x61, 0x6c, 0x74, 0x31, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x64, 0x66, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6b, 0x64, 0x66, 0x32,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x61, 0x6c, 0x74, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x73, 0x61, 0x6c, 0x74, 0x32, 0x42, 0x2b, 0x42, 0x12, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e,
	0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x13,
	0x73, 0x74, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74,
	0x65, 0x61, 0x6d,
}

var (
	file_offline_ticket_proto_rawDescOnce sync.Once
	file_offline_ticket_proto_rawDescData = file_offline_ticket_proto_rawDesc
)

func file_offline_ticket_proto_rawDescGZIP() []byte {
	file_offline_ticket_proto_rawDescOnce.Do(func() {
		file_offline_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_offline_ticket_proto_rawDescData)
	})
	return file_offline_ticket_proto_rawDescData
}

var file_offline_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_offline_ticket_proto_goTypes = []interface{}{
	(*Offline_Ticket)(nil), // 0: Offline_Ticket
}
var file_offline_ticket_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_offline_ticket_proto_init() }
func file_offline_ticket_proto_init() {
	if File_offline_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_offline_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offline_Ticket); i {
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
			RawDescriptor: file_offline_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_offline_ticket_proto_goTypes,
		DependencyIndexes: file_offline_ticket_proto_depIdxs,
		MessageInfos:      file_offline_ticket_proto_msgTypes,
	}.Build()
	File_offline_ticket_proto = out.File
	file_offline_ticket_proto_rawDesc = nil
	file_offline_ticket_proto_goTypes = nil
	file_offline_ticket_proto_depIdxs = nil
}
