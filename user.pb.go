// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.1
// source: user.proto

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

type Roles int32

const (
	Roles_RoleAdmin Roles = 0
	Roles_RoleUser  Roles = 1
)

// Enum value maps for Roles.
var (
	Roles_name = map[int32]string{
		0: "RoleAdmin",
		1: "RoleUser",
	}
	Roles_value = map[string]int32{
		"RoleAdmin": 0,
		"RoleUser":  1,
	}
)

func (x Roles) Enum() *Roles {
	p := new(Roles)
	*p = x
	return p
}

func (x Roles) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Roles) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (Roles) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x Roles) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Roles.Descriptor instead.
func (Roles) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName              string  `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName               string  `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	TenantId               string  `protobuf:"bytes,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Email                  string  `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Sub                    string  `protobuf:"bytes,6,opt,name=sub,proto3" json:"sub,omitempty"`
	Bio                    *string `protobuf:"bytes,7,opt,name=bio,proto3,oneof" json:"bio,omitempty"`
	IsProduction           bool    `protobuf:"varint,8,opt,name=is_production,json=isProduction,proto3" json:"is_production,omitempty"`
	OverwriteProspectEmail string  `protobuf:"bytes,9,opt,name=overwrite_prospect_email,json=overwriteProspectEmail,proto3" json:"overwrite_prospect_email,omitempty"`
	Role                   Roles   `protobuf:"varint,10,opt,name=role,proto3,enum=goosechase.data.user.Roles" json:"role,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *User) GetBio() string {
	if x != nil && x.Bio != nil {
		return *x.Bio
	}
	return ""
}

func (x *User) GetIsProduction() bool {
	if x != nil {
		return x.IsProduction
	}
	return false
}

func (x *User) GetOverwriteProspectEmail() string {
	if x != nil {
		return x.OverwriteProspectEmail
	}
	return ""
}

func (x *User) GetRole() Roles {
	if x != nil {
		return x.Role
	}
	return Roles_RoleAdmin
}

type FindOneUserResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *User `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *FindOneUserResult) Reset() {
	*x = FindOneUserResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneUserResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneUserResult) ProtoMessage() {}

func (x *FindOneUserResult) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneUserResult.ProtoReflect.Descriptor instead.
func (*FindOneUserResult) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *FindOneUserResult) GetResult() *User {
	if x != nil {
		return x.Result
	}
	return nil
}

type FindManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    *string `protobuf:"bytes,1,opt,name=query,proto3,oneof" json:"query,omitempty"`
	TenantId string  `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *FindManyRequest) Reset() {
	*x = FindManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindManyRequest) ProtoMessage() {}

func (x *FindManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindManyRequest.ProtoReflect.Descriptor instead.
func (*FindManyRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *FindManyRequest) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

func (x *FindManyRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type FindManyResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*User `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *FindManyResult) Reset() {
	*x = FindManyResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindManyResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindManyResult) ProtoMessage() {}

func (x *FindManyResult) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindManyResult.ProtoReflect.Descriptor instead.
func (*FindManyResult) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *FindManyResult) GetResults() []*User {
	if x != nil {
		return x.Results
	}
	return nil
}

type SetDemoModeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsProduction           bool   `protobuf:"varint,1,opt,name=is_production,json=isProduction,proto3" json:"is_production,omitempty"`
	UserId                 string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OverwriteProspectEmail string `protobuf:"bytes,3,opt,name=overwrite_prospect_email,json=overwriteProspectEmail,proto3" json:"overwrite_prospect_email,omitempty"`
}

func (x *SetDemoModeRequest) Reset() {
	*x = SetDemoModeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDemoModeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDemoModeRequest) ProtoMessage() {}

func (x *SetDemoModeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDemoModeRequest.ProtoReflect.Descriptor instead.
func (*SetDemoModeRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *SetDemoModeRequest) GetIsProduction() bool {
	if x != nil {
		return x.IsProduction
	}
	return false
}

func (x *SetDemoModeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SetDemoModeRequest) GetOverwriteProspectEmail() string {
	if x != nil {
		return x.OverwriteProspectEmail
	}
	return ""
}

type SetDemoModeResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *User  `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetDemoModeResult) Reset() {
	*x = SetDemoModeResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDemoModeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDemoModeResult) ProtoMessage() {}

func (x *SetDemoModeResult) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDemoModeResult.ProtoReflect.Descriptor instead.
func (*SetDemoModeResult) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *SetDemoModeResult) GetResult() *User {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *SetDemoModeResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type InviteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName              string  `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName               string  `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	TenantId               string  `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Email                  string  `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Sub                    string  `protobuf:"bytes,5,opt,name=sub,proto3" json:"sub,omitempty"`
	Bio                    *string `protobuf:"bytes,6,opt,name=bio,proto3,oneof" json:"bio,omitempty"`
	IsProduction           bool    `protobuf:"varint,7,opt,name=is_production,json=isProduction,proto3" json:"is_production,omitempty"`
	OverwriteProspectEmail string  `protobuf:"bytes,8,opt,name=overwrite_prospect_email,json=overwriteProspectEmail,proto3" json:"overwrite_prospect_email,omitempty"`
	Role                   Roles   `protobuf:"varint,9,opt,name=role,proto3,enum=goosechase.data.user.Roles" json:"role,omitempty"`
}

func (x *InviteUserRequest) Reset() {
	*x = InviteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteUserRequest) ProtoMessage() {}

func (x *InviteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteUserRequest.ProtoReflect.Descriptor instead.
func (*InviteUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *InviteUserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *InviteUserRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *InviteUserRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *InviteUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *InviteUserRequest) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *InviteUserRequest) GetBio() string {
	if x != nil && x.Bio != nil {
		return *x.Bio
	}
	return ""
}

func (x *InviteUserRequest) GetIsProduction() bool {
	if x != nil {
		return x.IsProduction
	}
	return false
}

func (x *InviteUserRequest) GetOverwriteProspectEmail() string {
	if x != nil {
		return x.OverwriteProspectEmail
	}
	return ""
}

func (x *InviteUserRequest) GetRole() Roles {
	if x != nil {
		return x.Role
	}
	return Roles_RoleAdmin
}

type InviteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *User  `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *InviteUserResponse) Reset() {
	*x = InviteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteUserResponse) ProtoMessage() {}

func (x *InviteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteUserResponse.ProtoReflect.Descriptor instead.
func (*InviteUserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *InviteUserResponse) GetResult() *User {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *InviteUserResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x6f,
	0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc6, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x15, 0x0a, 0x03, 0x62, 0x69,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x18, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x2f, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x62, 0x69, 0x6f, 0x22, 0x47, 0x0a, 0x11, 0x46, 0x69, 0x6e,
	0x64, 0x4f, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x53, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x46, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4d,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22,
	0x8c, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x5f,
	0x0a, 0x11, 0x53, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xc3, 0x02, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x15, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a,
	0x0d, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x18, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2f, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x62, 0x69, 0x6f, 0x22, 0x60, 0x0a, 0x12, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x24, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x12, 0x0d, 0x0a, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x10, 0x01, 0x32, 0x89, 0x03,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a,
	0x0b, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x67,
	0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x4f, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x57, 0x0a,
	0x08, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x73,
	0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x79,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x60, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x44, 0x65, 0x6d,
	0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74,
	0x44, 0x65, 0x6d, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x5f, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_proto_goTypes = []interface{}{
	(Roles)(0),                 // 0: goosechase.data.user.Roles
	(*User)(nil),               // 1: goosechase.data.user.User
	(*FindOneUserResult)(nil),  // 2: goosechase.data.user.FindOneUserResult
	(*FindManyRequest)(nil),    // 3: goosechase.data.user.FindManyRequest
	(*FindManyResult)(nil),     // 4: goosechase.data.user.FindManyResult
	(*SetDemoModeRequest)(nil), // 5: goosechase.data.user.SetDemoModeRequest
	(*SetDemoModeResult)(nil),  // 6: goosechase.data.user.SetDemoModeResult
	(*InviteUserRequest)(nil),  // 7: goosechase.data.user.InviteUserRequest
	(*InviteUserResponse)(nil), // 8: goosechase.data.user.InviteUserResponse
	(*FindOneRequest)(nil),     // 9: goosechase.data.common.FindOneRequest
}
var file_user_proto_depIdxs = []int32{
	0,  // 0: goosechase.data.user.User.role:type_name -> goosechase.data.user.Roles
	1,  // 1: goosechase.data.user.FindOneUserResult.result:type_name -> goosechase.data.user.User
	1,  // 2: goosechase.data.user.FindManyResult.results:type_name -> goosechase.data.user.User
	1,  // 3: goosechase.data.user.SetDemoModeResult.result:type_name -> goosechase.data.user.User
	0,  // 4: goosechase.data.user.InviteUserRequest.role:type_name -> goosechase.data.user.Roles
	1,  // 5: goosechase.data.user.InviteUserResponse.result:type_name -> goosechase.data.user.User
	9,  // 6: goosechase.data.user.UserService.FindOneUser:input_type -> goosechase.data.common.FindOneRequest
	3,  // 7: goosechase.data.user.UserService.FindMany:input_type -> goosechase.data.user.FindManyRequest
	5,  // 8: goosechase.data.user.UserService.SetDemoMode:input_type -> goosechase.data.user.SetDemoModeRequest
	7,  // 9: goosechase.data.user.UserService.InviteUser:input_type -> goosechase.data.user.InviteUserRequest
	2,  // 10: goosechase.data.user.UserService.FindOneUser:output_type -> goosechase.data.user.FindOneUserResult
	4,  // 11: goosechase.data.user.UserService.FindMany:output_type -> goosechase.data.user.FindManyResult
	6,  // 12: goosechase.data.user.UserService.SetDemoMode:output_type -> goosechase.data.user.SetDemoModeResult
	8,  // 13: goosechase.data.user.UserService.InviteUser:output_type -> goosechase.data.user.InviteUserResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneUserResult); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindManyRequest); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindManyResult); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDemoModeRequest); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDemoModeResult); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteUserRequest); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteUserResponse); i {
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
	file_user_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_user_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_user_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
