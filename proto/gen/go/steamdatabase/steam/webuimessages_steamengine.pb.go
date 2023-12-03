// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: webuimessages_steamengine.proto

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

type CSteamEngine_UpdateTextFilterDictionary_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language *string `protobuf:"bytes,1,req,name=language" json:"language,omitempty"`
	Type     *string `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
}

func (x *CSteamEngine_UpdateTextFilterDictionary_Notification) Reset() {
	*x = CSteamEngine_UpdateTextFilterDictionary_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_steamengine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamEngine_UpdateTextFilterDictionary_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamEngine_UpdateTextFilterDictionary_Notification) ProtoMessage() {}

func (x *CSteamEngine_UpdateTextFilterDictionary_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_steamengine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamEngine_UpdateTextFilterDictionary_Notification.ProtoReflect.Descriptor instead.
func (*CSteamEngine_UpdateTextFilterDictionary_Notification) Descriptor() ([]byte, []int) {
	return file_webuimessages_steamengine_proto_rawDescGZIP(), []int{0}
}

func (x *CSteamEngine_UpdateTextFilterDictionary_Notification) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

func (x *CSteamEngine_UpdateTextFilterDictionary_Notification) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

type CSteamEngine_GetTextFilterDictionary_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language *string `protobuf:"bytes,1,req,name=language" json:"language,omitempty"`
	Type     *string `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
}

func (x *CSteamEngine_GetTextFilterDictionary_Request) Reset() {
	*x = CSteamEngine_GetTextFilterDictionary_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_steamengine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamEngine_GetTextFilterDictionary_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamEngine_GetTextFilterDictionary_Request) ProtoMessage() {}

func (x *CSteamEngine_GetTextFilterDictionary_Request) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_steamengine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamEngine_GetTextFilterDictionary_Request.ProtoReflect.Descriptor instead.
func (*CSteamEngine_GetTextFilterDictionary_Request) Descriptor() ([]byte, []int) {
	return file_webuimessages_steamengine_proto_rawDescGZIP(), []int{1}
}

func (x *CSteamEngine_GetTextFilterDictionary_Request) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

func (x *CSteamEngine_GetTextFilterDictionary_Request) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

type CSteamEngine_GetTextFilterDictionary_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dictionary *string `protobuf:"bytes,1,opt,name=dictionary" json:"dictionary,omitempty"`
}

func (x *CSteamEngine_GetTextFilterDictionary_Response) Reset() {
	*x = CSteamEngine_GetTextFilterDictionary_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_steamengine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamEngine_GetTextFilterDictionary_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamEngine_GetTextFilterDictionary_Response) ProtoMessage() {}

func (x *CSteamEngine_GetTextFilterDictionary_Response) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_steamengine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamEngine_GetTextFilterDictionary_Response.ProtoReflect.Descriptor instead.
func (*CSteamEngine_GetTextFilterDictionary_Response) Descriptor() ([]byte, []int) {
	return file_webuimessages_steamengine_proto_rawDescGZIP(), []int{2}
}

func (x *CSteamEngine_GetTextFilterDictionary_Response) GetDictionary() string {
	if x != nil && x.Dictionary != nil {
		return *x.Dictionary
	}
	return ""
}

type CSteamEngine_TextFilterDictionaryChanged_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language *string `protobuf:"bytes,1,req,name=language" json:"language,omitempty"`
	Type     *string `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
}

func (x *CSteamEngine_TextFilterDictionaryChanged_Notification) Reset() {
	*x = CSteamEngine_TextFilterDictionaryChanged_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_steamengine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamEngine_TextFilterDictionaryChanged_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamEngine_TextFilterDictionaryChanged_Notification) ProtoMessage() {}

func (x *CSteamEngine_TextFilterDictionaryChanged_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_steamengine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamEngine_TextFilterDictionaryChanged_Notification.ProtoReflect.Descriptor instead.
func (*CSteamEngine_TextFilterDictionaryChanged_Notification) Descriptor() ([]byte, []int) {
	return file_webuimessages_steamengine_proto_rawDescGZIP(), []int{3}
}

func (x *CSteamEngine_TextFilterDictionaryChanged_Notification) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

func (x *CSteamEngine_TextFilterDictionaryChanged_Notification) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

type CSteamEngine_GetGameIDForPID_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid *uint32 `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
}

func (x *CSteamEngine_GetGameIDForPID_Request) Reset() {
	*x = CSteamEngine_GetGameIDForPID_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_steamengine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamEngine_GetGameIDForPID_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamEngine_GetGameIDForPID_Request) ProtoMessage() {}

func (x *CSteamEngine_GetGameIDForPID_Request) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_steamengine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamEngine_GetGameIDForPID_Request.ProtoReflect.Descriptor instead.
func (*CSteamEngine_GetGameIDForPID_Request) Descriptor() ([]byte, []int) {
	return file_webuimessages_steamengine_proto_rawDescGZIP(), []int{4}
}

func (x *CSteamEngine_GetGameIDForPID_Request) GetPid() uint32 {
	if x != nil && x.Pid != nil {
		return *x.Pid
	}
	return 0
}

type CSteamEngine_GetGameIDForPID_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gameid *uint64 `protobuf:"varint,1,opt,name=gameid" json:"gameid,omitempty"`
}

func (x *CSteamEngine_GetGameIDForPID_Response) Reset() {
	*x = CSteamEngine_GetGameIDForPID_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_steamengine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamEngine_GetGameIDForPID_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamEngine_GetGameIDForPID_Response) ProtoMessage() {}

func (x *CSteamEngine_GetGameIDForPID_Response) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_steamengine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamEngine_GetGameIDForPID_Response.ProtoReflect.Descriptor instead.
func (*CSteamEngine_GetGameIDForPID_Response) Descriptor() ([]byte, []int) {
	return file_webuimessages_steamengine_proto_rawDescGZIP(), []int{5}
}

func (x *CSteamEngine_GetGameIDForPID_Response) GetGameid() uint64 {
	if x != nil && x.Gameid != nil {
		return *x.Gameid
	}
	return 0
}

var File_webuimessages_steamengine_proto protoreflect.FileDescriptor

var file_webuimessages_steamengine_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x77, 0x65, 0x62, 0x75, 0x69, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x73, 0x74, 0x65, 0x61, 0x6d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x77, 0x65, 0x62,
	0x75, 0x69, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x34, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79,
	0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5e, 0x0a,
	0x2c, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x44, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4f, 0x0a,
	0x2d, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x44, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x22, 0x67,
	0x0a, 0x35, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x54,
	0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x38, 0x0a, 0x24, 0x43, 0x53, 0x74, 0x65, 0x61,
	0x6d, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x49,
	0x44, 0x46, 0x6f, 0x72, 0x50, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69,
	0x64, 0x22, 0x3f, 0x0a, 0x25, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x5f, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x46, 0x6f, 0x72, 0x50, 0x49,
	0x44, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65,
	0x69, 0x64, 0x32, 0xc5, 0x03, 0x0a, 0x0b, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x12, 0x65, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79,
	0x12, 0x35, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x57, 0x65, 0x62, 0x55, 0x49, 0x4e,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x78, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x72, 0x79, 0x12, 0x2d, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x21, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x65, 0x78,
	0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x36, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61,
	0x6d, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x10, 0x2e, 0x57, 0x65, 0x62, 0x55, 0x49, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x60, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x46,
	0x6f, 0x72, 0x50, 0x49, 0x44, 0x12, 0x25, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x46, 0x6f,
	0x72, 0x50, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x43,
	0x53, 0x74, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x47,
	0x61, 0x6d, 0x65, 0x49, 0x44, 0x46, 0x6f, 0x72, 0x50, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x04, 0x80, 0x97, 0x22, 0x02, 0x42, 0x3b, 0x42, 0x1d, 0x57, 0x65,
	0x62, 0x75, 0x69, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x53, 0x74, 0x65, 0x61, 0x6d,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x01, 0x50, 0x01, 0x5a,
	0x13, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x80, 0x01, 0x01,
}

var (
	file_webuimessages_steamengine_proto_rawDescOnce sync.Once
	file_webuimessages_steamengine_proto_rawDescData = file_webuimessages_steamengine_proto_rawDesc
)

func file_webuimessages_steamengine_proto_rawDescGZIP() []byte {
	file_webuimessages_steamengine_proto_rawDescOnce.Do(func() {
		file_webuimessages_steamengine_proto_rawDescData = protoimpl.X.CompressGZIP(file_webuimessages_steamengine_proto_rawDescData)
	})
	return file_webuimessages_steamengine_proto_rawDescData
}

var file_webuimessages_steamengine_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_webuimessages_steamengine_proto_goTypes = []interface{}{
	(*CSteamEngine_UpdateTextFilterDictionary_Notification)(nil),  // 0: CSteamEngine_UpdateTextFilterDictionary_Notification
	(*CSteamEngine_GetTextFilterDictionary_Request)(nil),          // 1: CSteamEngine_GetTextFilterDictionary_Request
	(*CSteamEngine_GetTextFilterDictionary_Response)(nil),         // 2: CSteamEngine_GetTextFilterDictionary_Response
	(*CSteamEngine_TextFilterDictionaryChanged_Notification)(nil), // 3: CSteamEngine_TextFilterDictionaryChanged_Notification
	(*CSteamEngine_GetGameIDForPID_Request)(nil),                  // 4: CSteamEngine_GetGameIDForPID_Request
	(*CSteamEngine_GetGameIDForPID_Response)(nil),                 // 5: CSteamEngine_GetGameIDForPID_Response
	(*WebUINoResponse)(nil),                                       // 6: WebUINoResponse
}
var file_webuimessages_steamengine_proto_depIdxs = []int32{
	0, // 0: SteamEngine.UpdateTextFilterDictionary:input_type -> CSteamEngine_UpdateTextFilterDictionary_Notification
	1, // 1: SteamEngine.GetTextFilterDictionary:input_type -> CSteamEngine_GetTextFilterDictionary_Request
	3, // 2: SteamEngine.NotifyTextFilterDictionaryChanged:input_type -> CSteamEngine_TextFilterDictionaryChanged_Notification
	4, // 3: SteamEngine.GetGameIDForPID:input_type -> CSteamEngine_GetGameIDForPID_Request
	6, // 4: SteamEngine.UpdateTextFilterDictionary:output_type -> WebUINoResponse
	2, // 5: SteamEngine.GetTextFilterDictionary:output_type -> CSteamEngine_GetTextFilterDictionary_Response
	6, // 6: SteamEngine.NotifyTextFilterDictionaryChanged:output_type -> WebUINoResponse
	5, // 7: SteamEngine.GetGameIDForPID:output_type -> CSteamEngine_GetGameIDForPID_Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_webuimessages_steamengine_proto_init() }
func file_webuimessages_steamengine_proto_init() {
	if File_webuimessages_steamengine_proto != nil {
		return
	}
	file_steammessages_base_proto_init()
	file_webuimessages_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_webuimessages_steamengine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSteamEngine_UpdateTextFilterDictionary_Notification); i {
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
		file_webuimessages_steamengine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSteamEngine_GetTextFilterDictionary_Request); i {
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
		file_webuimessages_steamengine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSteamEngine_GetTextFilterDictionary_Response); i {
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
		file_webuimessages_steamengine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSteamEngine_TextFilterDictionaryChanged_Notification); i {
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
		file_webuimessages_steamengine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSteamEngine_GetGameIDForPID_Request); i {
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
		file_webuimessages_steamengine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSteamEngine_GetGameIDForPID_Response); i {
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
			RawDescriptor: file_webuimessages_steamengine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_webuimessages_steamengine_proto_goTypes,
		DependencyIndexes: file_webuimessages_steamengine_proto_depIdxs,
		MessageInfos:      file_webuimessages_steamengine_proto_msgTypes,
	}.Build()
	File_webuimessages_steamengine_proto = out.File
	file_webuimessages_steamengine_proto_rawDesc = nil
	file_webuimessages_steamengine_proto_goTypes = nil
	file_webuimessages_steamengine_proto_depIdxs = nil
}
