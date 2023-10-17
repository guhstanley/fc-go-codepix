// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pixkey.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PixKeyRegistration struct {
	Kind                 string   `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	AccountId            string   `protobuf:"bytes,3,opt,name=accountId,proto3" json:"accountId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PixKeyRegistration) Reset()         { *m = PixKeyRegistration{} }
func (m *PixKeyRegistration) String() string { return proto.CompactTextString(m) }
func (*PixKeyRegistration) ProtoMessage()    {}
func (*PixKeyRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c629aa9982df412, []int{0}
}

func (m *PixKeyRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PixKeyRegistration.Unmarshal(m, b)
}
func (m *PixKeyRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PixKeyRegistration.Marshal(b, m, deterministic)
}
func (m *PixKeyRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PixKeyRegistration.Merge(m, src)
}
func (m *PixKeyRegistration) XXX_Size() int {
	return xxx_messageInfo_PixKeyRegistration.Size(m)
}
func (m *PixKeyRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_PixKeyRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_PixKeyRegistration proto.InternalMessageInfo

func (m *PixKeyRegistration) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *PixKeyRegistration) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PixKeyRegistration) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type PixKey struct {
	Kind                 string   `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PixKey) Reset()         { *m = PixKey{} }
func (m *PixKey) String() string { return proto.CompactTextString(m) }
func (*PixKey) ProtoMessage()    {}
func (*PixKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c629aa9982df412, []int{1}
}

func (m *PixKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PixKey.Unmarshal(m, b)
}
func (m *PixKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PixKey.Marshal(b, m, deterministic)
}
func (m *PixKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PixKey.Merge(m, src)
}
func (m *PixKey) XXX_Size() int {
	return xxx_messageInfo_PixKey.Size(m)
}
func (m *PixKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PixKey.DiscardUnknown(m)
}

var xxx_messageInfo_PixKey proto.InternalMessageInfo

func (m *PixKey) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *PixKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Account struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	AccountNumber        string   `protobuf:"bytes,2,opt,name=accountNumber,proto3" json:"accountNumber,omitempty"`
	BankId               string   `protobuf:"bytes,3,opt,name=bankId,proto3" json:"bankId,omitempty"`
	BankName             string   `protobuf:"bytes,4,opt,name=bankName,proto3" json:"bankName,omitempty"`
	OwnerName            string   `protobuf:"bytes,5,opt,name=OwnerName,proto3" json:"OwnerName,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c629aa9982df412, []int{2}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Account) GetAccountNumber() string {
	if m != nil {
		return m.AccountNumber
	}
	return ""
}

func (m *Account) GetBankId() string {
	if m != nil {
		return m.BankId
	}
	return ""
}

func (m *Account) GetBankName() string {
	if m != nil {
		return m.BankName
	}
	return ""
}

func (m *Account) GetOwnerName() string {
	if m != nil {
		return m.OwnerName
	}
	return ""
}

func (m *Account) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type PixKeyInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Account              *Account `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PixKeyInfo) Reset()         { *m = PixKeyInfo{} }
func (m *PixKeyInfo) String() string { return proto.CompactTextString(m) }
func (*PixKeyInfo) ProtoMessage()    {}
func (*PixKeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c629aa9982df412, []int{3}
}

func (m *PixKeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PixKeyInfo.Unmarshal(m, b)
}
func (m *PixKeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PixKeyInfo.Marshal(b, m, deterministic)
}
func (m *PixKeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PixKeyInfo.Merge(m, src)
}
func (m *PixKeyInfo) XXX_Size() int {
	return xxx_messageInfo_PixKeyInfo.Size(m)
}
func (m *PixKeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PixKeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PixKeyInfo proto.InternalMessageInfo

func (m *PixKeyInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PixKeyInfo) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *PixKeyInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PixKeyInfo) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *PixKeyInfo) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type PixKeyCreatedResult struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PixKeyCreatedResult) Reset()         { *m = PixKeyCreatedResult{} }
func (m *PixKeyCreatedResult) String() string { return proto.CompactTextString(m) }
func (*PixKeyCreatedResult) ProtoMessage()    {}
func (*PixKeyCreatedResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c629aa9982df412, []int{4}
}

func (m *PixKeyCreatedResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PixKeyCreatedResult.Unmarshal(m, b)
}
func (m *PixKeyCreatedResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PixKeyCreatedResult.Marshal(b, m, deterministic)
}
func (m *PixKeyCreatedResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PixKeyCreatedResult.Merge(m, src)
}
func (m *PixKeyCreatedResult) XXX_Size() int {
	return xxx_messageInfo_PixKeyCreatedResult.Size(m)
}
func (m *PixKeyCreatedResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PixKeyCreatedResult.DiscardUnknown(m)
}

var xxx_messageInfo_PixKeyCreatedResult proto.InternalMessageInfo

func (m *PixKeyCreatedResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PixKeyCreatedResult) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *PixKeyCreatedResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*PixKeyRegistration)(nil), "github.com.guhstanley.codepix.PixKeyRegistration")
	proto.RegisterType((*PixKey)(nil), "github.com.guhstanley.codepix.PixKey")
	proto.RegisterType((*Account)(nil), "github.com.guhstanley.codepix.Account")
	proto.RegisterType((*PixKeyInfo)(nil), "github.com.guhstanley.codepix.PixKeyInfo")
	proto.RegisterType((*PixKeyCreatedResult)(nil), "github.com.guhstanley.codepix.PixKeyCreatedResult")
}

func init() {
	proto.RegisterFile("pixkey.proto", fileDescriptor_9c629aa9982df412)
}

var fileDescriptor_9c629aa9982df412 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x6d, 0xd2, 0x36, 0xb5, 0xd7, 0x5a, 0x64, 0x94, 0x12, 0x8a, 0x82, 0x04, 0x15, 0x7d, 0x19,
	0xb1, 0x3e, 0xfa, 0x62, 0x15, 0x84, 0x22, 0x54, 0x49, 0x5f, 0xa4, 0x6f, 0xf9, 0xb8, 0xdb, 0x0e,
	0x69, 0x33, 0x61, 0x32, 0xd9, 0x6d, 0xfe, 0xce, 0xfe, 0x8f, 0xfd, 0x35, 0xfb, 0x47, 0x96, 0xcc,
	0x4c, 0x53, 0xd2, 0xc2, 0x6e, 0xde, 0xee, 0x3d, 0x37, 0xf7, 0xcc, 0x39, 0x67, 0x26, 0x30, 0xca,
	0xd8, 0x21, 0xc1, 0x92, 0x66, 0x82, 0x4b, 0x4e, 0xde, 0x6e, 0x98, 0xdc, 0x16, 0x21, 0x8d, 0xf8,
	0x9e, 0x6e, 0x8a, 0x6d, 0x2e, 0x83, 0x74, 0x87, 0x25, 0x8d, 0x78, 0x8c, 0x19, 0x3b, 0x78, 0xff,
	0x81, 0xfc, 0x63, 0x87, 0x3f, 0x58, 0xfa, 0xb8, 0x61, 0xb9, 0x14, 0x81, 0x64, 0x3c, 0x25, 0x04,
	0x7a, 0x09, 0x4b, 0x63, 0xd7, 0x7a, 0x67, 0x7d, 0x1a, 0xfa, 0xaa, 0x26, 0x2f, 0xa1, 0x9b, 0x60,
	0xe9, 0xda, 0x0a, 0xaa, 0x4a, 0xf2, 0x06, 0x86, 0x41, 0x14, 0xf1, 0x22, 0x95, 0x8b, 0xd8, 0xed,
	0x2a, 0xfc, 0x04, 0x78, 0x14, 0x1c, 0xcd, 0xdc, 0x8e, 0xcd, 0xbb, 0xb3, 0x60, 0x30, 0xd7, 0xdb,
	0x4d, 0x66, 0xeb, 0x8c, 0x99, 0xbc, 0x87, 0x17, 0xa6, 0x59, 0x16, 0xfb, 0x10, 0x85, 0x61, 0x69,
	0x82, 0x64, 0x02, 0x4e, 0x18, 0xa4, 0x49, 0x2d, 0xcd, 0x74, 0x64, 0x0a, 0xcf, 0xaa, 0x6a, 0x19,
	0xec, 0xd1, 0xed, 0xa9, 0x49, 0xdd, 0x57, 0xe7, 0xfe, 0xbd, 0x49, 0x51, 0xa8, 0x61, 0x5f, 0x9f,
	0x5b, 0x03, 0xd5, 0x34, 0x12, 0x18, 0x48, 0x8c, 0xe7, 0xd2, 0x75, 0xf4, 0xb4, 0x06, 0xbc, 0x5b,
	0x0b, 0x40, 0x1b, 0x5e, 0xa4, 0x57, 0x9c, 0x8c, 0xc1, 0x66, 0x47, 0xed, 0x36, 0x8b, 0xeb, 0x10,
	0xec, 0xcb, 0x10, 0xba, 0xa7, 0x48, 0x7f, 0xc0, 0xc0, 0xb8, 0x50, 0xda, 0x9e, 0xcf, 0x3e, 0xd2,
	0x47, 0xef, 0x8f, 0x9a, 0xc4, 0xfc, 0xe3, 0x5a, 0x53, 0x64, 0xff, 0x5c, 0xe4, 0x0a, 0x5e, 0x69,
	0x8d, 0xbf, 0x34, 0xe4, 0x63, 0x5e, 0xec, 0xe4, 0x85, 0xd8, 0x09, 0x38, 0xb9, 0x0c, 0x64, 0x91,
	0x1b, 0xb9, 0xa6, 0x23, 0xaf, 0xa1, 0x8f, 0x42, 0x70, 0x61, 0x24, 0xeb, 0x66, 0x76, 0xaf, 0x9d,
	0xaf, 0x50, 0x5c, 0xb3, 0x08, 0x49, 0x09, 0x63, 0xfd, 0x98, 0x50, 0x98, 0x07, 0xf0, 0xf5, 0x09,
	0x13, 0x97, 0x2f, 0x70, 0x3a, 0x6b, 0xb5, 0xd2, 0x70, 0xe1, 0x75, 0xc8, 0x1a, 0x7a, 0xbf, 0xab,
	0x60, 0x3f, 0xb4, 0xda, 0x9e, 0x7e, 0x6e, 0xf5, 0x59, 0x75, 0x9d, 0x5e, 0xe7, 0xe7, 0x68, 0x0d,
	0xf4, 0x8b, 0xfa, 0xa7, 0xbe, 0x67, 0x61, 0xe8, 0xa8, 0xea, 0xdb, 0x43, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x35, 0xc7, 0xc5, 0xca, 0x6d, 0x03, 0x00, 0x00,
}
