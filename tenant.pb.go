// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.1
// source: tenant.proto

package proto

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

type Tenant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description              string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Domain                   string  `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	DomainVerificationStatus string  `protobuf:"bytes,5,opt,name=domain_verification_status,json=domainVerificationStatus,proto3" json:"domain_verification_status,omitempty"`
	AgentId                  string  `protobuf:"bytes,6,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	CreatedAt                *string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt                *string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *Tenant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tenant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tenant) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Tenant) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Tenant) GetDomainVerificationStatus() string {
	if x != nil {
		return x.DomainVerificationStatus
	}
	return ""
}

func (x *Tenant) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *Tenant) GetCreatedAt() string {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return ""
}

func (x *Tenant) GetUpdatedAt() string {
	if x != nil && x.UpdatedAt != nil {
		return *x.UpdatedAt
	}
	return ""
}

type FindOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *FindOneRequest) Reset() {
	*x = FindOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneRequest) ProtoMessage() {}

func (x *FindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneRequest.ProtoReflect.Descriptor instead.
func (*FindOneRequest) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *FindOneRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type FindOneResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Tenant `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *FindOneResult) Reset() {
	*x = FindOneResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneResult) ProtoMessage() {}

func (x *FindOneResult) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneResult.ProtoReflect.Descriptor instead.
func (*FindOneResult) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{2}
}

func (x *FindOneResult) GetResult() *Tenant {
	if x != nil {
		return x.Result
	}
	return nil
}

type CreateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Tenant `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateTenantRequest) Reset() {
	*x = CreateTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantRequest) ProtoMessage() {}

func (x *CreateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantRequest.ProtoReflect.Descriptor instead.
func (*CreateTenantRequest) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTenantRequest) GetData() *Tenant {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateTenantResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Tenant `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Status string  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateTenantResult) Reset() {
	*x = CreateTenantResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenantResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantResult) ProtoMessage() {}

func (x *CreateTenantResult) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantResult.ProtoReflect.Descriptor instead.
func (*CreateTenantResult) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTenantResult) GetResult() *Tenant {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *CreateTenantResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_tenant_proto protoreflect.FileDescriptor

var file_tenant_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0xa5, 0x02, 0x0a, 0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x3c, 0x0a, 0x1a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x18, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x2d,
	0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x47, 0x0a,
	0x0d, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x49, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x64, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xd2, 0x01, 0x0a, 0x0d, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x07, 0x46, 0x69, 0x6e,
	0x64, 0x4f, 0x6e, 0x65, 0x12, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67,
	0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x67, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tenant_proto_rawDescOnce sync.Once
	file_tenant_proto_rawDescData = file_tenant_proto_rawDesc
)

func file_tenant_proto_rawDescGZIP() []byte {
	file_tenant_proto_rawDescOnce.Do(func() {
		file_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_tenant_proto_rawDescData)
	})
	return file_tenant_proto_rawDescData
}

var file_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tenant_proto_goTypes = []interface{}{
	(*Tenant)(nil),              // 0: goosechase.data.tenant.Tenant
	(*FindOneRequest)(nil),      // 1: goosechase.data.tenant.FindOneRequest
	(*FindOneResult)(nil),       // 2: goosechase.data.tenant.FindOneResult
	(*CreateTenantRequest)(nil), // 3: goosechase.data.tenant.CreateTenantRequest
	(*CreateTenantResult)(nil),  // 4: goosechase.data.tenant.CreateTenantResult
}
var file_tenant_proto_depIdxs = []int32{
	0, // 0: goosechase.data.tenant.FindOneResult.result:type_name -> goosechase.data.tenant.Tenant
	0, // 1: goosechase.data.tenant.CreateTenantRequest.data:type_name -> goosechase.data.tenant.Tenant
	0, // 2: goosechase.data.tenant.CreateTenantResult.result:type_name -> goosechase.data.tenant.Tenant
	1, // 3: goosechase.data.tenant.TenantService.FindOne:input_type -> goosechase.data.tenant.FindOneRequest
	3, // 4: goosechase.data.tenant.TenantService.CreateTenant:input_type -> goosechase.data.tenant.CreateTenantRequest
	2, // 5: goosechase.data.tenant.TenantService.FindOne:output_type -> goosechase.data.tenant.FindOneResult
	4, // 6: goosechase.data.tenant.TenantService.CreateTenant:output_type -> goosechase.data.tenant.CreateTenantResult
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tenant_proto_init() }
func file_tenant_proto_init() {
	if File_tenant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tenant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tenant); i {
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
		file_tenant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneRequest); i {
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
		file_tenant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneResult); i {
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
		file_tenant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTenantRequest); i {
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
		file_tenant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTenantResult); i {
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
	file_tenant_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tenant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tenant_proto_goTypes,
		DependencyIndexes: file_tenant_proto_depIdxs,
		MessageInfos:      file_tenant_proto_msgTypes,
	}.Build()
	File_tenant_proto = out.File
	file_tenant_proto_rawDesc = nil
	file_tenant_proto_goTypes = nil
	file_tenant_proto_depIdxs = nil
}
