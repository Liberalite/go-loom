// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/coin/coin.proto

/*
Package coin is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/coin/coin.proto

It has these top-level messages:
	Economy
	Account
	Allowance
	InitRequest
	TotalSupplyRequest
	TotalSupplyResponse
	BalanceOfRequest
	BalanceOfResponse
	AllowanceRequest
	AllowanceResponse
	ApproveRequest
	ApproveResponse
	TransferRequest
	TransferResponse
	TransferFromRequest
	TransferFromResponse
*/
package coin

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Economy struct {
	TotalSupply *types.BigUInt `protobuf:"bytes,1,opt,name=total_supply,json=totalSupply" json:"total_supply,omitempty"`
}

func (m *Economy) Reset()                    { *m = Economy{} }
func (m *Economy) String() string            { return proto.CompactTextString(m) }
func (*Economy) ProtoMessage()               {}
func (*Economy) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{0} }

func (m *Economy) GetTotalSupply() *types.BigUInt {
	if m != nil {
		return m.TotalSupply
	}
	return nil
}

type Account struct {
	Owner   *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Balance *types.BigUInt `protobuf:"bytes,2,opt,name=balance" json:"balance,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{1} }

func (m *Account) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Account) GetBalance() *types.BigUInt {
	if m != nil {
		return m.Balance
	}
	return nil
}

type Allowance struct {
	Owner   *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Spender *types.Address `protobuf:"bytes,2,opt,name=spender" json:"spender,omitempty"`
	Amount  *types.BigUInt `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *Allowance) Reset()                    { *m = Allowance{} }
func (m *Allowance) String() string            { return proto.CompactTextString(m) }
func (*Allowance) ProtoMessage()               {}
func (*Allowance) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{2} }

func (m *Allowance) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Allowance) GetSpender() *types.Address {
	if m != nil {
		return m.Spender
	}
	return nil
}

func (m *Allowance) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type InitRequest struct {
	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
}

func (m *InitRequest) Reset()                    { *m = InitRequest{} }
func (m *InitRequest) String() string            { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()               {}
func (*InitRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{3} }

func (m *InitRequest) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// read only
type TotalSupplyRequest struct {
}

func (m *TotalSupplyRequest) Reset()                    { *m = TotalSupplyRequest{} }
func (m *TotalSupplyRequest) String() string            { return proto.CompactTextString(m) }
func (*TotalSupplyRequest) ProtoMessage()               {}
func (*TotalSupplyRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{4} }

type TotalSupplyResponse struct {
	TotalSupply *types.BigUInt `protobuf:"bytes,1,opt,name=total_supply,json=totalSupply" json:"total_supply,omitempty"`
}

func (m *TotalSupplyResponse) Reset()                    { *m = TotalSupplyResponse{} }
func (m *TotalSupplyResponse) String() string            { return proto.CompactTextString(m) }
func (*TotalSupplyResponse) ProtoMessage()               {}
func (*TotalSupplyResponse) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{5} }

func (m *TotalSupplyResponse) GetTotalSupply() *types.BigUInt {
	if m != nil {
		return m.TotalSupply
	}
	return nil
}

type BalanceOfRequest struct {
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
}

func (m *BalanceOfRequest) Reset()                    { *m = BalanceOfRequest{} }
func (m *BalanceOfRequest) String() string            { return proto.CompactTextString(m) }
func (*BalanceOfRequest) ProtoMessage()               {}
func (*BalanceOfRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{6} }

func (m *BalanceOfRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

type BalanceOfResponse struct {
	Balance *types.BigUInt `protobuf:"bytes,1,opt,name=balance" json:"balance,omitempty"`
}

func (m *BalanceOfResponse) Reset()                    { *m = BalanceOfResponse{} }
func (m *BalanceOfResponse) String() string            { return proto.CompactTextString(m) }
func (*BalanceOfResponse) ProtoMessage()               {}
func (*BalanceOfResponse) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{7} }

func (m *BalanceOfResponse) GetBalance() *types.BigUInt {
	if m != nil {
		return m.Balance
	}
	return nil
}

type AllowanceRequest struct {
	Owner   *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Spender *types.Address `protobuf:"bytes,2,opt,name=spender" json:"spender,omitempty"`
}

func (m *AllowanceRequest) Reset()                    { *m = AllowanceRequest{} }
func (m *AllowanceRequest) String() string            { return proto.CompactTextString(m) }
func (*AllowanceRequest) ProtoMessage()               {}
func (*AllowanceRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{8} }

func (m *AllowanceRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *AllowanceRequest) GetSpender() *types.Address {
	if m != nil {
		return m.Spender
	}
	return nil
}

type AllowanceResponse struct {
	Amount *types.BigUInt `protobuf:"bytes,1,opt,name=amount" json:"amount,omitempty"`
}

func (m *AllowanceResponse) Reset()                    { *m = AllowanceResponse{} }
func (m *AllowanceResponse) String() string            { return proto.CompactTextString(m) }
func (*AllowanceResponse) ProtoMessage()               {}
func (*AllowanceResponse) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{9} }

func (m *AllowanceResponse) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

// volatile
type ApproveRequest struct {
	Spender *types.Address `protobuf:"bytes,1,opt,name=spender" json:"spender,omitempty"`
	Amount  *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *ApproveRequest) Reset()                    { *m = ApproveRequest{} }
func (m *ApproveRequest) String() string            { return proto.CompactTextString(m) }
func (*ApproveRequest) ProtoMessage()               {}
func (*ApproveRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{10} }

func (m *ApproveRequest) GetSpender() *types.Address {
	if m != nil {
		return m.Spender
	}
	return nil
}

func (m *ApproveRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type ApproveResponse struct {
}

func (m *ApproveResponse) Reset()                    { *m = ApproveResponse{} }
func (m *ApproveResponse) String() string            { return proto.CompactTextString(m) }
func (*ApproveResponse) ProtoMessage()               {}
func (*ApproveResponse) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{11} }

type TransferRequest struct {
	To     *types.Address `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	Amount *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *TransferRequest) Reset()                    { *m = TransferRequest{} }
func (m *TransferRequest) String() string            { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()               {}
func (*TransferRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{12} }

func (m *TransferRequest) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TransferRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type TransferResponse struct {
}

func (m *TransferResponse) Reset()                    { *m = TransferResponse{} }
func (m *TransferResponse) String() string            { return proto.CompactTextString(m) }
func (*TransferResponse) ProtoMessage()               {}
func (*TransferResponse) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{13} }

type TransferFromRequest struct {
	From   *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To     *types.Address `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Amount *types.BigUInt `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *TransferFromRequest) Reset()                    { *m = TransferFromRequest{} }
func (m *TransferFromRequest) String() string            { return proto.CompactTextString(m) }
func (*TransferFromRequest) ProtoMessage()               {}
func (*TransferFromRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{14} }

func (m *TransferFromRequest) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *TransferFromRequest) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TransferFromRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type TransferFromResponse struct {
}

func (m *TransferFromResponse) Reset()                    { *m = TransferFromResponse{} }
func (m *TransferFromResponse) String() string            { return proto.CompactTextString(m) }
func (*TransferFromResponse) ProtoMessage()               {}
func (*TransferFromResponse) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{15} }

func init() {
	proto.RegisterType((*Economy)(nil), "Economy")
	proto.RegisterType((*Account)(nil), "Account")
	proto.RegisterType((*Allowance)(nil), "Allowance")
	proto.RegisterType((*InitRequest)(nil), "InitRequest")
	proto.RegisterType((*TotalSupplyRequest)(nil), "TotalSupplyRequest")
	proto.RegisterType((*TotalSupplyResponse)(nil), "TotalSupplyResponse")
	proto.RegisterType((*BalanceOfRequest)(nil), "BalanceOfRequest")
	proto.RegisterType((*BalanceOfResponse)(nil), "BalanceOfResponse")
	proto.RegisterType((*AllowanceRequest)(nil), "AllowanceRequest")
	proto.RegisterType((*AllowanceResponse)(nil), "AllowanceResponse")
	proto.RegisterType((*ApproveRequest)(nil), "ApproveRequest")
	proto.RegisterType((*ApproveResponse)(nil), "ApproveResponse")
	proto.RegisterType((*TransferRequest)(nil), "TransferRequest")
	proto.RegisterType((*TransferResponse)(nil), "TransferResponse")
	proto.RegisterType((*TransferFromRequest)(nil), "TransferFromRequest")
	proto.RegisterType((*TransferFromResponse)(nil), "TransferFromResponse")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/coin/coin.proto", fileDescriptorCoin)
}

var fileDescriptorCoin = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xef, 0xeb, 0xd2, 0x40,
	0x18, 0x67, 0xfb, 0x96, 0xda, 0x63, 0xa4, 0x4e, 0x89, 0x11, 0x11, 0x72, 0xf4, 0x22, 0x88, 0x5c,
	0x28, 0xd5, 0x9b, 0xde, 0x4c, 0x28, 0xf0, 0x85, 0x04, 0xcb, 0x7c, 0x1b, 0xdb, 0x3c, 0x6d, 0xb4,
	0xdd, 0x73, 0xde, 0xdd, 0x12, 0xff, 0xfb, 0x70, 0xb7, 0xcd, 0xa9, 0x2c, 0xd7, 0x9b, 0x83, 0xfb,
	0x3c, 0xfb, 0xfc, 0xd8, 0xe7, 0xd9, 0xe0, 0xf3, 0x2e, 0x52, 0xbf, 0xd2, 0x60, 0x12, 0x62, 0xe2,
	0xc4, 0x88, 0x09, 0xa3, 0xea, 0x80, 0xe2, 0xb7, 0xb3, 0xc3, 0x77, 0xa7, 0xab, 0x13, 0xa4, 0x51,
	0xac, 0x22, 0xe6, 0xa8, 0x23, 0xa7, 0xd2, 0x09, 0x31, 0x62, 0xd9, 0x31, 0xe1, 0x02, 0x15, 0xbe,
	0x78, 0x7f, 0x87, 0xad, 0x59, 0xd9, 0xa9, 0x19, 0xe4, 0x23, 0xb4, 0xbf, 0x84, 0xc8, 0x30, 0x39,
	0x5a, 0x6f, 0xe1, 0xa9, 0x42, 0xe5, 0xc7, 0x3f, 0x65, 0xca, 0x79, 0x7c, 0xb4, 0x8d, 0xb1, 0xf1,
	0xa6, 0x3b, 0xed, 0x4c, 0xe6, 0xd1, 0xee, 0xc7, 0x82, 0x29, 0xaf, 0x9b, 0x4d, 0xbf, 0x67, 0x43,
	0xb2, 0x84, 0xb6, 0x1b, 0x86, 0x98, 0x32, 0x65, 0xbd, 0x82, 0xc7, 0x78, 0x60, 0x54, 0x94, 0x04,
	0x77, 0xb3, 0x11, 0x54, 0x4a, 0x4f, 0xc3, 0x16, 0x81, 0x76, 0xe0, 0xc7, 0x3e, 0x0b, 0xa9, 0x6d,
	0x5e, 0x49, 0x16, 0x03, 0xb2, 0x87, 0x27, 0x6e, 0x1c, 0xe3, 0xe1, 0x74, 0x69, 0x22, 0x28, 0x39,
	0x65, 0x1b, 0x2a, 0x4a, 0xc1, 0xe2, 0x89, 0x62, 0x60, 0x8d, 0xa1, 0xe5, 0x27, 0xa7, 0x78, 0xf6,
	0xc3, 0x95, 0x67, 0x8e, 0x93, 0x19, 0x74, 0x17, 0x2c, 0x52, 0x1e, 0xdd, 0xa7, 0x54, 0x2a, 0xeb,
	0x35, 0x74, 0x7c, 0xfd, 0x42, 0xd2, 0x36, 0xc6, 0x0f, 0x5a, 0x55, 0x03, 0x5e, 0x39, 0x21, 0x23,
	0xb0, 0x56, 0xe7, 0x16, 0x72, 0x2e, 0x99, 0xc3, 0xf0, 0x02, 0x95, 0x1c, 0x99, 0xa4, 0xff, 0x57,
	0xe8, 0x14, 0xfa, 0x73, 0x5d, 0xc6, 0xb7, 0x6d, 0x91, 0xe9, 0x4e, 0x11, 0xe4, 0x13, 0x0c, 0x2a,
	0x9c, 0xdc, 0xb5, 0x52, 0xb7, 0x51, 0x57, 0xf7, 0x1a, 0xfa, 0x65, 0xdd, 0x0d, 0xcd, 0x9a, 0xb4,
	0x4e, 0x3e, 0xc0, 0xa0, 0xa2, 0x9b, 0x07, 0x3a, 0xaf, 0xc2, 0xa8, 0x59, 0xc5, 0x1a, 0x9e, 0xb9,
	0x9c, 0x0b, 0xfc, 0x53, 0x86, 0xa9, 0x98, 0x19, 0xf7, 0x57, 0x6c, 0xd6, 0xe8, 0x0e, 0xa0, 0x57,
	0xea, 0xea, 0x30, 0x64, 0x09, 0xbd, 0x95, 0xf0, 0x99, 0xdc, 0x52, 0x51, 0x78, 0xd9, 0x60, 0x2a,
	0xbc, 0xb1, 0x31, 0x15, 0x36, 0x70, 0xb0, 0xa0, 0x7f, 0x96, 0xcb, 0x2d, 0x10, 0x86, 0x05, 0xf6,
	0x55, 0x60, 0x52, 0xd8, 0xbc, 0x84, 0x47, 0x5b, 0x81, 0xc9, 0x8d, 0x51, 0x86, 0xe6, 0x21, 0xcc,
	0x7f, 0x86, 0xa8, 0xfb, 0x92, 0x9f, 0xc3, 0xe8, 0xd2, 0x50, 0x07, 0x09, 0x5a, 0xd9, 0x2f, 0x3e,
	0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x5e, 0x09, 0xd0, 0x54, 0x04, 0x00, 0x00,
}