// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: cable/grpc/proto/service.proto

package grpc

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

// Every response contains a status field with one of the following values
type Status int32

const (
	// RPC called failed unexpectedly
	Status_ERROR Status = 0
	// RPC called succeed, action was performed
	Status_SUCCESS Status = 1
	// RPC called succeed but actions was rejected by the application (e.g., rejected subscription/connection)
	Status_FAILURE Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "ERROR",
		1: "SUCCESS",
		2: "FAILURE",
	}
	Status_value = map[string]int32{
		"ERROR":   0,
		"SUCCESS": 1,
		"FAILURE": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_cable_grpc_proto_service_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_cable_grpc_proto_service_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_cable_grpc_proto_service_proto_rawDescGZIP(), []int{0}
}

// Env represents a client connection information passed to RPC server
type Env struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Underlying HTTP request URL
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Underlying HTTP request headers
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Connection-level metadata
	Cstate map[string]string `protobuf:"bytes,3,rep,name=cstate,proto3" json:"cstate,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Channel-level metadata (only set for Command calls, contains data for the affected subscription)
	Istate map[string]string `protobuf:"bytes,4,rep,name=istate,proto3" json:"istate,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Env) Reset() {
	*x = Env{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cable_grpc_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Env) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Env) ProtoMessage() {}

func (x *Env) ProtoReflect() protoreflect.Message {
	mi := &file_cable_grpc_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Env.ProtoReflect.Descriptor instead.
func (*Env) Descriptor() ([]byte, []int) {
	return file_cable_grpc_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *Env) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Env) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Env) GetCstate() map[string]string {
	if x != nil {
		return x.Cstate
	}
	return nil
}

func (x *Env) GetIstate() map[string]string {
	if x != nil {
		return x.Istate
	}
	return nil
}

// EnvResponse contains the changes made to the connection or channel state,
// that must be applied to the state
type EnvResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cstate map[string]string `protobuf:"bytes,1,rep,name=cstate,proto3" json:"cstate,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Istate map[string]string `protobuf:"bytes,2,rep,name=istate,proto3" json:"istate,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EnvResponse) Reset() {
	*x = EnvResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cable_grpc_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvResponse) ProtoMessage() {}

func (x *EnvResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cable_grpc_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvResponse.ProtoReflect.Descriptor instead.
func (*EnvResponse) Descriptor() ([]byte, []int) {
	return file_cable_grpc_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *EnvResponse) GetCstate() map[string]string {
	if x != nil {
		return x.Cstate
	}
	return nil
}

func (x *EnvResponse) GetIstate() map[string]string {
	if x != nil {
		return x.Istate
	}
	return nil
}

// ConnectionRequest describes a payload for the Connect call
type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Env *Env `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cable_grpc_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cable_grpc_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_cable_grpc_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectionRequest) GetEnv() *Env {
	if x != nil {
		return x.Env
	}
	return nil
}

// ConnectionResponse describes a response of the Connect call
type ConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=anycable.Status" json:"status,omitempty"`
	// Connection identifiers passed as string (in most cases, JSON)
	Identifiers string `protobuf:"bytes,2,opt,name=identifiers,proto3" json:"identifiers,omitempty"`
	// Messages to be sent to the client
	Transmissions []string `protobuf:"bytes,3,rep,name=transmissions,proto3" json:"transmissions,omitempty"`
	// Error message in case status is ERROR or FAILURE
	ErrorMsg string       `protobuf:"bytes,4,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	Env      *EnvResponse `protobuf:"bytes,5,opt,name=env,proto3" json:"env,omitempty"`
}

func (x *ConnectionResponse) Reset() {
	*x = ConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cable_grpc_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionResponse) ProtoMessage() {}

func (x *ConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cable_grpc_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionResponse.ProtoReflect.Descriptor instead.
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return file_cable_grpc_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectionResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ERROR
}

func (x *ConnectionResponse) GetIdentifiers() string {
	if x != nil {
		return x.Identifiers
	}
	return ""
}

func (x *ConnectionResponse) GetTransmissions() []string {
	if x != nil {
		return x.Transmissions
	}
	return nil
}

func (x *ConnectionResponse) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *ConnectionResponse) GetEnv() *EnvResponse {
	if x != nil {
		return x.Env
	}
	return nil
}

// ConnectionMesssage describes a payload for the Command call
type CommandMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the command ("subscribe", "unsubscribe", "message" for Action Cable)
	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	// Subscription identifier (channel id, channel params)
	Identifier string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Client's connection identifiers (received in ConnectionResponse)
	ConnectionIdentifiers string `protobuf:"bytes,3,opt,name=connection_identifiers,json=connectionIdentifiers,proto3" json:"connection_identifiers,omitempty"`
	// Command payload
	Data string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Env  *Env   `protobuf:"bytes,5,opt,name=env,proto3" json:"env,omitempty"`
}

func (x *CommandMessage) Reset() {
	*x = CommandMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cable_grpc_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandMessage) ProtoMessage() {}

func (x *CommandMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cable_grpc_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandMessage.ProtoReflect.Descriptor instead.
func (*CommandMessage) Descriptor() ([]byte, []int) {
	return file_cable_grpc_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *CommandMessage) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CommandMessage) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *CommandMessage) GetConnectionIdentifiers() string {
	if x != nil {
		return x.ConnectionIdentifiers
	}
	return ""
}

func (x *CommandMessage) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *CommandMessage) GetEnv() *Env {
	if x != nil {
		return x.Env
	}
	return nil
}

// CommandResponse describes a response of the Command call
type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=anycable.Status" json:"status,omitempty"`
	// If true, the client must be disconencted
	Disconnect bool `protobuf:"varint,2,opt,name=disconnect,proto3" json:"disconnect,omitempty"`
	// If true, the client must be unsubscribed from all streams from this subscription
	StopStreams bool `protobuf:"varint,3,opt,name=stop_streams,json=stopStreams,proto3" json:"stop_streams,omitempty"`
	// List of the streams to subscribe the client to
	Streams []string `protobuf:"bytes,4,rep,name=streams,proto3" json:"streams,omitempty"`
	// Messages to be sent to the client
	Transmissions []string     `protobuf:"bytes,5,rep,name=transmissions,proto3" json:"transmissions,omitempty"`
	ErrorMsg      string       `protobuf:"bytes,6,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	Env           *EnvResponse `protobuf:"bytes,7,opt,name=env,proto3" json:"env,omitempty"`
	// List of the stream to unsubscribe the client from
	StoppedStreams []string `protobuf:"bytes,8,rep,name=stopped_streams,json=stoppedStreams,proto3" json:"stopped_streams,omitempty"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cable_grpc_proto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cable_grpc_proto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_cable_grpc_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *CommandResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ERROR
}

func (x *CommandResponse) GetDisconnect() bool {
	if x != nil {
		return x.Disconnect
	}
	return false
}

func (x *CommandResponse) GetStopStreams() bool {
	if x != nil {
		return x.StopStreams
	}
	return false
}

func (x *CommandResponse) GetStreams() []string {
	if x != nil {
		return x.Streams
	}
	return nil
}

func (x *CommandResponse) GetTransmissions() []string {
	if x != nil {
		return x.Transmissions
	}
	return nil
}

func (x *CommandResponse) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *CommandResponse) GetEnv() *EnvResponse {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *CommandResponse) GetStoppedStreams() []string {
	if x != nil {
		return x.StoppedStreams
	}
	return nil
}

// DisconnectRequest describes a payload for the Disconnect call
type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifiers string `protobuf:"bytes,1,opt,name=identifiers,proto3" json:"identifiers,omitempty"`
	// List of a client's subscriptions (identifiers).
	// Required to call `unsubscribe` callbacks.
	Subscriptions []string `protobuf:"bytes,2,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	Env           *Env     `protobuf:"bytes,5,opt,name=env,proto3" json:"env,omitempty"`
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cable_grpc_proto_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cable_grpc_proto_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_cable_grpc_proto_service_proto_rawDescGZIP(), []int{6}
}

func (x *DisconnectRequest) GetIdentifiers() string {
	if x != nil {
		return x.Identifiers
	}
	return ""
}

func (x *DisconnectRequest) GetSubscriptions() []string {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

func (x *DisconnectRequest) GetEnv() *Env {
	if x != nil {
		return x.Env
	}
	return nil
}

// DisconnectResponse describes a response of the Disconnect call
type DisconnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   Status `protobuf:"varint,1,opt,name=status,proto3,enum=anycable.Status" json:"status,omitempty"`
	ErrorMsg string `protobuf:"bytes,2,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
}

func (x *DisconnectResponse) Reset() {
	*x = DisconnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cable_grpc_proto_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectResponse) ProtoMessage() {}

func (x *DisconnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cable_grpc_proto_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectResponse.ProtoReflect.Descriptor instead.
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return file_cable_grpc_proto_service_proto_rawDescGZIP(), []int{7}
}

func (x *DisconnectResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ERROR
}

func (x *DisconnectResponse) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

var File_cable_grpc_proto_service_proto protoreflect.FileDescriptor

var file_cable_grpc_proto_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xe5, 0x02, 0x0a, 0x03, 0x45,
	0x6e, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x34, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x45, 0x6e, 0x76, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x63, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x6e, 0x79,
	0x63, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x2e, 0x43, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x31, 0x0a,
	0x06, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x2e, 0x49, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b,
	0x43, 0x73, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x49, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xf9, 0x01, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x63, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45, 0x6e,
	0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a,
	0x06, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x73, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x49, 0x73, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x34,
	0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x52,
	0x03, 0x65, 0x6e, 0x76, 0x22, 0xcc, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x6e,
	0x79, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x27, 0x0a, 0x03, 0x65, 0x6e,
	0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03,
	0x65, 0x6e, 0x76, 0x22, 0xb6, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x35, 0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x15, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x03, 0x65,
	0x6e, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x22, 0xad, 0x02, 0x0a,
	0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74,
	0x6f, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x73, 0x74, 0x6f, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x27, 0x0a, 0x03, 0x65, 0x6e,
	0x76, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03,
	0x65, 0x6e, 0x76, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x74,
	0x6f, 0x70, 0x70, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x7c, 0x0a, 0x11,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x03, 0x65, 0x6e, 0x76,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x45, 0x6e, 0x76, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x22, 0x5b, 0x0a, 0x12, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x2a, 0x2d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x32, 0xda, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x46,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x6e, 0x79, 0x63,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x18, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x19, 0x2e, 0x61, 0x6e,
	0x79, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x6e, 0x79, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cable_grpc_proto_service_proto_rawDescOnce sync.Once
	file_cable_grpc_proto_service_proto_rawDescData = file_cable_grpc_proto_service_proto_rawDesc
)

func file_cable_grpc_proto_service_proto_rawDescGZIP() []byte {
	file_cable_grpc_proto_service_proto_rawDescOnce.Do(func() {
		file_cable_grpc_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cable_grpc_proto_service_proto_rawDescData)
	})
	return file_cable_grpc_proto_service_proto_rawDescData
}

var file_cable_grpc_proto_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cable_grpc_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_cable_grpc_proto_service_proto_goTypes = []interface{}{
	(Status)(0),                // 0: anycable.Status
	(*Env)(nil),                // 1: anycable.Env
	(*EnvResponse)(nil),        // 2: anycable.EnvResponse
	(*ConnectionRequest)(nil),  // 3: anycable.ConnectionRequest
	(*ConnectionResponse)(nil), // 4: anycable.ConnectionResponse
	(*CommandMessage)(nil),     // 5: anycable.CommandMessage
	(*CommandResponse)(nil),    // 6: anycable.CommandResponse
	(*DisconnectRequest)(nil),  // 7: anycable.DisconnectRequest
	(*DisconnectResponse)(nil), // 8: anycable.DisconnectResponse
	nil,                        // 9: anycable.Env.HeadersEntry
	nil,                        // 10: anycable.Env.CstateEntry
	nil,                        // 11: anycable.Env.IstateEntry
	nil,                        // 12: anycable.EnvResponse.CstateEntry
	nil,                        // 13: anycable.EnvResponse.IstateEntry
}
var file_cable_grpc_proto_service_proto_depIdxs = []int32{
	9,  // 0: anycable.Env.headers:type_name -> anycable.Env.HeadersEntry
	10, // 1: anycable.Env.cstate:type_name -> anycable.Env.CstateEntry
	11, // 2: anycable.Env.istate:type_name -> anycable.Env.IstateEntry
	12, // 3: anycable.EnvResponse.cstate:type_name -> anycable.EnvResponse.CstateEntry
	13, // 4: anycable.EnvResponse.istate:type_name -> anycable.EnvResponse.IstateEntry
	1,  // 5: anycable.ConnectionRequest.env:type_name -> anycable.Env
	0,  // 6: anycable.ConnectionResponse.status:type_name -> anycable.Status
	2,  // 7: anycable.ConnectionResponse.env:type_name -> anycable.EnvResponse
	1,  // 8: anycable.CommandMessage.env:type_name -> anycable.Env
	0,  // 9: anycable.CommandResponse.status:type_name -> anycable.Status
	2,  // 10: anycable.CommandResponse.env:type_name -> anycable.EnvResponse
	1,  // 11: anycable.DisconnectRequest.env:type_name -> anycable.Env
	0,  // 12: anycable.DisconnectResponse.status:type_name -> anycable.Status
	3,  // 13: anycable.RPC.Connect:input_type -> anycable.ConnectionRequest
	5,  // 14: anycable.RPC.Command:input_type -> anycable.CommandMessage
	7,  // 15: anycable.RPC.Disconnect:input_type -> anycable.DisconnectRequest
	4,  // 16: anycable.RPC.Connect:output_type -> anycable.ConnectionResponse
	6,  // 17: anycable.RPC.Command:output_type -> anycable.CommandResponse
	8,  // 18: anycable.RPC.Disconnect:output_type -> anycable.DisconnectResponse
	16, // [16:19] is the sub-list for method output_type
	13, // [13:16] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_cable_grpc_proto_service_proto_init() }
func file_cable_grpc_proto_service_proto_init() {
	if File_cable_grpc_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cable_grpc_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Env); i {
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
		file_cable_grpc_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvResponse); i {
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
		file_cable_grpc_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
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
		file_cable_grpc_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionResponse); i {
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
		file_cable_grpc_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandMessage); i {
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
		file_cable_grpc_proto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResponse); i {
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
		file_cable_grpc_proto_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectRequest); i {
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
		file_cable_grpc_proto_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectResponse); i {
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
			RawDescriptor: file_cable_grpc_proto_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cable_grpc_proto_service_proto_goTypes,
		DependencyIndexes: file_cable_grpc_proto_service_proto_depIdxs,
		EnumInfos:         file_cable_grpc_proto_service_proto_enumTypes,
		MessageInfos:      file_cable_grpc_proto_service_proto_msgTypes,
	}.Build()
	File_cable_grpc_proto_service_proto = out.File
	file_cable_grpc_proto_service_proto_rawDesc = nil
	file_cable_grpc_proto_service_proto_goTypes = nil
	file_cable_grpc_proto_service_proto_depIdxs = nil
}
