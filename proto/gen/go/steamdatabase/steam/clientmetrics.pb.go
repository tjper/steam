// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: clientmetrics.proto

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

type CClientMetrics_ClientBootstrap_RequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalHostname    *string `protobuf:"bytes,1,opt,name=original_hostname,json=originalHostname" json:"original_hostname,omitempty"`
	ActualHostname      *string `protobuf:"bytes,2,opt,name=actual_hostname,json=actualHostname" json:"actual_hostname,omitempty"`
	Path                *string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	BaseName            *string `protobuf:"bytes,4,opt,name=base_name,json=baseName" json:"base_name,omitempty"`
	Success             *bool   `protobuf:"varint,5,opt,name=success" json:"success,omitempty"`
	StatusCode          *uint32 `protobuf:"varint,6,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	AddressOfRequestUrl *string `protobuf:"bytes,7,opt,name=address_of_request_url,json=addressOfRequestUrl" json:"address_of_request_url,omitempty"`
	ResponseTimeMs      *uint32 `protobuf:"varint,8,opt,name=response_time_ms,json=responseTimeMs" json:"response_time_ms,omitempty"`
	BytesReceived       *uint64 `protobuf:"varint,9,opt,name=bytes_received,json=bytesReceived" json:"bytes_received,omitempty"`
	NumRetries          *uint32 `protobuf:"varint,10,opt,name=num_retries,json=numRetries" json:"num_retries,omitempty"`
}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) Reset() {
	*x = CClientMetrics_ClientBootstrap_RequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientmetrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClientMetrics_ClientBootstrap_RequestInfo) ProtoMessage() {}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_clientmetrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClientMetrics_ClientBootstrap_RequestInfo.ProtoReflect.Descriptor instead.
func (*CClientMetrics_ClientBootstrap_RequestInfo) Descriptor() ([]byte, []int) {
	return file_clientmetrics_proto_rawDescGZIP(), []int{0}
}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) GetOriginalHostname() string {
	if x != nil && x.OriginalHostname != nil {
		return *x.OriginalHostname
	}
	return ""
}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) GetActualHostname() string {
	if x != nil && x.ActualHostname != nil {
		return *x.ActualHostname
	}
	return ""
}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) GetBaseName() string {
	if x != nil && x.BaseName != nil {
		return *x.BaseName
	}
	return ""
}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) GetStatusCode() uint32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) GetAddressOfRequestUrl() string {
	if x != nil && x.AddressOfRequestUrl != nil {
		return *x.AddressOfRequestUrl
	}
	return ""
}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) GetResponseTimeMs() uint32 {
	if x != nil && x.ResponseTimeMs != nil {
		return *x.ResponseTimeMs
	}
	return 0
}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) GetBytesReceived() uint64 {
	if x != nil && x.BytesReceived != nil {
		return *x.BytesReceived
	}
	return 0
}

func (x *CClientMetrics_ClientBootstrap_RequestInfo) GetNumRetries() uint32 {
	if x != nil && x.NumRetries != nil {
		return *x.NumRetries
	}
	return 0
}

type CClientMetrics_ClientBootstrap_Summary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LauncherType      *uint32                                       `protobuf:"varint,1,opt,name=launcher_type,json=launcherType" json:"launcher_type,omitempty"`
	SteamRealm        *uint32                                       `protobuf:"varint,2,opt,name=steam_realm,json=steamRealm" json:"steam_realm,omitempty"`
	BetaName          *string                                       `protobuf:"bytes,3,opt,name=beta_name,json=betaName" json:"beta_name,omitempty"`
	DownloadCompleted *bool                                         `protobuf:"varint,4,opt,name=download_completed,json=downloadCompleted" json:"download_completed,omitempty"`
	TotalTimeMs       *uint32                                       `protobuf:"varint,6,opt,name=total_time_ms,json=totalTimeMs" json:"total_time_ms,omitempty"`
	ManifestRequests  []*CClientMetrics_ClientBootstrap_RequestInfo `protobuf:"bytes,7,rep,name=manifest_requests,json=manifestRequests" json:"manifest_requests,omitempty"`
	PackageRequests   []*CClientMetrics_ClientBootstrap_RequestInfo `protobuf:"bytes,8,rep,name=package_requests,json=packageRequests" json:"package_requests,omitempty"`
}

func (x *CClientMetrics_ClientBootstrap_Summary) Reset() {
	*x = CClientMetrics_ClientBootstrap_Summary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientmetrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClientMetrics_ClientBootstrap_Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClientMetrics_ClientBootstrap_Summary) ProtoMessage() {}

func (x *CClientMetrics_ClientBootstrap_Summary) ProtoReflect() protoreflect.Message {
	mi := &file_clientmetrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClientMetrics_ClientBootstrap_Summary.ProtoReflect.Descriptor instead.
func (*CClientMetrics_ClientBootstrap_Summary) Descriptor() ([]byte, []int) {
	return file_clientmetrics_proto_rawDescGZIP(), []int{1}
}

func (x *CClientMetrics_ClientBootstrap_Summary) GetLauncherType() uint32 {
	if x != nil && x.LauncherType != nil {
		return *x.LauncherType
	}
	return 0
}

func (x *CClientMetrics_ClientBootstrap_Summary) GetSteamRealm() uint32 {
	if x != nil && x.SteamRealm != nil {
		return *x.SteamRealm
	}
	return 0
}

func (x *CClientMetrics_ClientBootstrap_Summary) GetBetaName() string {
	if x != nil && x.BetaName != nil {
		return *x.BetaName
	}
	return ""
}

func (x *CClientMetrics_ClientBootstrap_Summary) GetDownloadCompleted() bool {
	if x != nil && x.DownloadCompleted != nil {
		return *x.DownloadCompleted
	}
	return false
}

func (x *CClientMetrics_ClientBootstrap_Summary) GetTotalTimeMs() uint32 {
	if x != nil && x.TotalTimeMs != nil {
		return *x.TotalTimeMs
	}
	return 0
}

func (x *CClientMetrics_ClientBootstrap_Summary) GetManifestRequests() []*CClientMetrics_ClientBootstrap_RequestInfo {
	if x != nil {
		return x.ManifestRequests
	}
	return nil
}

func (x *CClientMetrics_ClientBootstrap_Summary) GetPackageRequests() []*CClientMetrics_ClientBootstrap_RequestInfo {
	if x != nil {
		return x.PackageRequests
	}
	return nil
}

type CClientMetrics_ContentDownloadResponse_Counts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class_100    *uint32 `protobuf:"varint,1,opt,name=class_100,json=class100" json:"class_100,omitempty"`
	Class_200    *uint32 `protobuf:"varint,2,opt,name=class_200,json=class200" json:"class_200,omitempty"`
	Class_300    *uint32 `protobuf:"varint,3,opt,name=class_300,json=class300" json:"class_300,omitempty"`
	Class_400    *uint32 `protobuf:"varint,4,opt,name=class_400,json=class400" json:"class_400,omitempty"`
	Class_500    *uint32 `protobuf:"varint,5,opt,name=class_500,json=class500" json:"class_500,omitempty"`
	NoResponse   *uint32 `protobuf:"varint,6,opt,name=no_response,json=noResponse" json:"no_response,omitempty"`
	ClassUnknown *uint32 `protobuf:"varint,7,opt,name=class_unknown,json=classUnknown" json:"class_unknown,omitempty"`
}

func (x *CClientMetrics_ContentDownloadResponse_Counts) Reset() {
	*x = CClientMetrics_ContentDownloadResponse_Counts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientmetrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClientMetrics_ContentDownloadResponse_Counts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClientMetrics_ContentDownloadResponse_Counts) ProtoMessage() {}

func (x *CClientMetrics_ContentDownloadResponse_Counts) ProtoReflect() protoreflect.Message {
	mi := &file_clientmetrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClientMetrics_ContentDownloadResponse_Counts.ProtoReflect.Descriptor instead.
func (*CClientMetrics_ContentDownloadResponse_Counts) Descriptor() ([]byte, []int) {
	return file_clientmetrics_proto_rawDescGZIP(), []int{2}
}

func (x *CClientMetrics_ContentDownloadResponse_Counts) GetClass_100() uint32 {
	if x != nil && x.Class_100 != nil {
		return *x.Class_100
	}
	return 0
}

func (x *CClientMetrics_ContentDownloadResponse_Counts) GetClass_200() uint32 {
	if x != nil && x.Class_200 != nil {
		return *x.Class_200
	}
	return 0
}

func (x *CClientMetrics_ContentDownloadResponse_Counts) GetClass_300() uint32 {
	if x != nil && x.Class_300 != nil {
		return *x.Class_300
	}
	return 0
}

func (x *CClientMetrics_ContentDownloadResponse_Counts) GetClass_400() uint32 {
	if x != nil && x.Class_400 != nil {
		return *x.Class_400
	}
	return 0
}

func (x *CClientMetrics_ContentDownloadResponse_Counts) GetClass_500() uint32 {
	if x != nil && x.Class_500 != nil {
		return *x.Class_500
	}
	return 0
}

func (x *CClientMetrics_ContentDownloadResponse_Counts) GetNoResponse() uint32 {
	if x != nil && x.NoResponse != nil {
		return *x.NoResponse
	}
	return 0
}

func (x *CClientMetrics_ContentDownloadResponse_Counts) GetClassUnknown() uint32 {
	if x != nil && x.ClassUnknown != nil {
		return *x.ClassUnknown
	}
	return 0
}

type CClientMetrics_ContentDownloadResponse_HostCounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname   *string                                        `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	SourceType *uint32                                        `protobuf:"varint,2,opt,name=source_type,json=sourceType" json:"source_type,omitempty"`
	Counts     *CClientMetrics_ContentDownloadResponse_Counts `protobuf:"bytes,3,opt,name=counts" json:"counts,omitempty"`
}

func (x *CClientMetrics_ContentDownloadResponse_HostCounts) Reset() {
	*x = CClientMetrics_ContentDownloadResponse_HostCounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientmetrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClientMetrics_ContentDownloadResponse_HostCounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClientMetrics_ContentDownloadResponse_HostCounts) ProtoMessage() {}

func (x *CClientMetrics_ContentDownloadResponse_HostCounts) ProtoReflect() protoreflect.Message {
	mi := &file_clientmetrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClientMetrics_ContentDownloadResponse_HostCounts.ProtoReflect.Descriptor instead.
func (*CClientMetrics_ContentDownloadResponse_HostCounts) Descriptor() ([]byte, []int) {
	return file_clientmetrics_proto_rawDescGZIP(), []int{3}
}

func (x *CClientMetrics_ContentDownloadResponse_HostCounts) GetHostname() string {
	if x != nil && x.Hostname != nil {
		return *x.Hostname
	}
	return ""
}

func (x *CClientMetrics_ContentDownloadResponse_HostCounts) GetSourceType() uint32 {
	if x != nil && x.SourceType != nil {
		return *x.SourceType
	}
	return 0
}

func (x *CClientMetrics_ContentDownloadResponse_HostCounts) GetCounts() *CClientMetrics_ContentDownloadResponse_Counts {
	if x != nil {
		return x.Counts
	}
	return nil
}

type CClientMetrics_ContentDownloadResponse_Hosts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts []*CClientMetrics_ContentDownloadResponse_HostCounts `protobuf:"bytes,1,rep,name=hosts" json:"hosts,omitempty"`
}

func (x *CClientMetrics_ContentDownloadResponse_Hosts) Reset() {
	*x = CClientMetrics_ContentDownloadResponse_Hosts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientmetrics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClientMetrics_ContentDownloadResponse_Hosts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClientMetrics_ContentDownloadResponse_Hosts) ProtoMessage() {}

func (x *CClientMetrics_ContentDownloadResponse_Hosts) ProtoReflect() protoreflect.Message {
	mi := &file_clientmetrics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClientMetrics_ContentDownloadResponse_Hosts.ProtoReflect.Descriptor instead.
func (*CClientMetrics_ContentDownloadResponse_Hosts) Descriptor() ([]byte, []int) {
	return file_clientmetrics_proto_rawDescGZIP(), []int{4}
}

func (x *CClientMetrics_ContentDownloadResponse_Hosts) GetHosts() []*CClientMetrics_ContentDownloadResponse_HostCounts {
	if x != nil {
		return x.Hosts
	}
	return nil
}

var File_clientmetrics_proto protoreflect.FileDescriptor

var file_clientmetrics_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x03, 0x0a, 0x2a, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x75,
	0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4f,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x90, 0x03,
	0x0a, 0x26, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x5f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x5f, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x75, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x65, 0x74, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x65, 0x74, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x58,
	0x0a, 0x11, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x43, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x56, 0x0a, 0x10, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x5f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x22, 0x86, 0x02, 0x0a, 0x2d, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x5f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x31, 0x30, 0x30, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x31, 0x30, 0x30, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x32, 0x30, 0x30, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x32, 0x30, 0x30, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x33, 0x30, 0x30, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x33, 0x30, 0x30, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x34, 0x30, 0x30, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x34, 0x30, 0x30, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x35, 0x30, 0x30, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x35, 0x30, 0x30, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x75, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x22, 0xb8, 0x01, 0x0a, 0x31, 0x43, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x06,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x43,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x22, 0x78, 0x0a, 0x2c, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x48,
	0x6f, 0x73, 0x74, 0x73, 0x12, 0x48, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x5f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x48, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x30,
	0x42, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x48, 0x01, 0x50, 0x01, 0x5a, 0x13, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x80, 0x01, 0x00,
}

var (
	file_clientmetrics_proto_rawDescOnce sync.Once
	file_clientmetrics_proto_rawDescData = file_clientmetrics_proto_rawDesc
)

func file_clientmetrics_proto_rawDescGZIP() []byte {
	file_clientmetrics_proto_rawDescOnce.Do(func() {
		file_clientmetrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_clientmetrics_proto_rawDescData)
	})
	return file_clientmetrics_proto_rawDescData
}

var file_clientmetrics_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_clientmetrics_proto_goTypes = []interface{}{
	(*CClientMetrics_ClientBootstrap_RequestInfo)(nil),        // 0: CClientMetrics_ClientBootstrap_RequestInfo
	(*CClientMetrics_ClientBootstrap_Summary)(nil),            // 1: CClientMetrics_ClientBootstrap_Summary
	(*CClientMetrics_ContentDownloadResponse_Counts)(nil),     // 2: CClientMetrics_ContentDownloadResponse_Counts
	(*CClientMetrics_ContentDownloadResponse_HostCounts)(nil), // 3: CClientMetrics_ContentDownloadResponse_HostCounts
	(*CClientMetrics_ContentDownloadResponse_Hosts)(nil),      // 4: CClientMetrics_ContentDownloadResponse_Hosts
}
var file_clientmetrics_proto_depIdxs = []int32{
	0, // 0: CClientMetrics_ClientBootstrap_Summary.manifest_requests:type_name -> CClientMetrics_ClientBootstrap_RequestInfo
	0, // 1: CClientMetrics_ClientBootstrap_Summary.package_requests:type_name -> CClientMetrics_ClientBootstrap_RequestInfo
	2, // 2: CClientMetrics_ContentDownloadResponse_HostCounts.counts:type_name -> CClientMetrics_ContentDownloadResponse_Counts
	3, // 3: CClientMetrics_ContentDownloadResponse_Hosts.hosts:type_name -> CClientMetrics_ContentDownloadResponse_HostCounts
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_clientmetrics_proto_init() }
func file_clientmetrics_proto_init() {
	if File_clientmetrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clientmetrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CClientMetrics_ClientBootstrap_RequestInfo); i {
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
		file_clientmetrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CClientMetrics_ClientBootstrap_Summary); i {
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
		file_clientmetrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CClientMetrics_ContentDownloadResponse_Counts); i {
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
		file_clientmetrics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CClientMetrics_ContentDownloadResponse_HostCounts); i {
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
		file_clientmetrics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CClientMetrics_ContentDownloadResponse_Hosts); i {
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
			RawDescriptor: file_clientmetrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_clientmetrics_proto_goTypes,
		DependencyIndexes: file_clientmetrics_proto_depIdxs,
		MessageInfos:      file_clientmetrics_proto_msgTypes,
	}.Build()
	File_clientmetrics_proto = out.File
	file_clientmetrics_proto_rawDesc = nil
	file_clientmetrics_proto_goTypes = nil
	file_clientmetrics_proto_depIdxs = nil
}
