// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/pay/pay.proto

package pay

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Order struct {
	StoreId     string `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	StoreName   string `protobuf:"bytes,2,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	Method      string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	AuthCode    string `protobuf:"bytes,4,opt,name=auth_code,json=authCode,proto3" json:"auth_code,omitempty"`
	Title       string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	TotalAmount int64  `protobuf:"varint,6,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	OrderNo     string `protobuf:"bytes,7,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	OperatorId  string `protobuf:"bytes,8,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	TerminalId  string `protobuf:"bytes,9,opt,name=terminal_id,json=terminalId,proto3" json:"terminal_id,omitempty"`
	CreatedAt   string `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a0282dd5dc82e7b, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetStoreId() string {
	if m != nil {
		return m.StoreId
	}
	return ""
}

func (m *Order) GetStoreName() string {
	if m != nil {
		return m.StoreName
	}
	return ""
}

func (m *Order) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Order) GetAuthCode() string {
	if m != nil {
		return m.AuthCode
	}
	return ""
}

func (m *Order) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Order) GetTotalAmount() int64 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

func (m *Order) GetOrderNo() string {
	if m != nil {
		return m.OrderNo
	}
	return ""
}

func (m *Order) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

func (m *Order) GetTerminalId() string {
	if m != nil {
		return m.TerminalId
	}
	return ""
}

func (m *Order) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Order) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Request struct {
	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a0282dd5dc82e7b, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type Response struct {
	Valid   bool   `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Error   *Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a0282dd5dc82e7b, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Response) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code   string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a0282dd5dc82e7b, []int{3}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Error.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return m.Size()
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func init() {
	proto.RegisterType((*Order)(nil), "pay.Order")
	proto.RegisterType((*Request)(nil), "pay.Request")
	proto.RegisterType((*Response)(nil), "pay.Response")
	proto.RegisterType((*Error)(nil), "pay.Error")
}

func init() { proto.RegisterFile("proto/pay/pay.proto", fileDescriptor_3a0282dd5dc82e7b) }

var fileDescriptor_3a0282dd5dc82e7b = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x8a, 0x13, 0x41,
	0x10, 0x86, 0x33, 0x9b, 0x4c, 0x32, 0xa9, 0xac, 0x97, 0x56, 0xa4, 0x55, 0x1c, 0xe3, 0x1c, 0x74,
	0x41, 0x58, 0x21, 0xfb, 0x04, 0x51, 0x5c, 0xd8, 0xcb, 0xaa, 0x73, 0xf1, 0x22, 0x0c, 0x6d, 0xba,
	0x60, 0x03, 0x33, 0x53, 0x63, 0x4f, 0x8d, 0x90, 0xb7, 0xf0, 0xb1, 0xc4, 0xd3, 0x1e, 0x3d, 0x4a,
	0xf2, 0x22, 0xd2, 0xd5, 0x3d, 0x67, 0x0f, 0x81, 0xfc, 0xff, 0x57, 0x14, 0x7f, 0xfd, 0x3d, 0xf0,
	0xb0, 0x73, 0xc4, 0xf4, 0xb6, 0x33, 0x07, 0xff, 0xbb, 0x14, 0xa5, 0xa6, 0x9d, 0x39, 0x14, 0xbf,
	0xcf, 0x20, 0xfd, 0xe8, 0x2c, 0x3a, 0xf5, 0x04, 0xb2, 0x9e, 0xc9, 0x61, 0xb5, 0xb7, 0x3a, 0x59,
	0x27, 0x17, 0xcb, 0x72, 0x21, 0xfa, 0xc6, 0xaa, 0xe7, 0x00, 0x01, 0xb5, 0xa6, 0x41, 0x7d, 0x26,
	0x70, 0x29, 0xce, 0xad, 0x69, 0x50, 0x3d, 0x86, 0x79, 0x83, 0x7c, 0x47, 0x56, 0x4f, 0x05, 0x45,
	0xa5, 0x9e, 0xc1, 0xd2, 0x0c, 0x7c, 0x57, 0xed, 0xc8, 0xa2, 0x9e, 0x09, 0xca, 0xbc, 0xf1, 0x9e,
	0x2c, 0xaa, 0x47, 0x90, 0xf2, 0x9e, 0x6b, 0xd4, 0xa9, 0x80, 0x20, 0xd4, 0x4b, 0x38, 0x67, 0x62,
	0x53, 0x57, 0xa6, 0xa1, 0xa1, 0x65, 0x3d, 0x5f, 0x27, 0x17, 0xd3, 0x72, 0x25, 0xde, 0x56, 0x2c,
	0x9f, 0x93, 0x7c, 0xe0, 0xaa, 0x25, 0xbd, 0x08, 0x39, 0x45, 0xdf, 0x92, 0x7a, 0x01, 0x2b, 0xea,
	0xd0, 0x19, 0x26, 0xe7, 0xaf, 0xc8, 0x84, 0xc2, 0x68, 0xdd, 0x58, 0x3f, 0xc0, 0xe8, 0x9a, 0x7d,
	0x6b, 0x6a, 0x3f, 0xb0, 0x0c, 0x03, 0xa3, 0x15, 0x2e, 0xdd, 0x39, 0x34, 0x8c, 0xb6, 0x32, 0xac,
	0x21, 0x5c, 0x1a, 0x9d, 0x2d, 0x7b, 0x3c, 0x74, 0x76, 0xc4, 0xab, 0x80, 0xa3, 0xb3, 0xe5, 0xe2,
	0x0d, 0x2c, 0x4a, 0xfc, 0x3e, 0x60, 0xcf, 0x6a, 0x0d, 0xa9, 0xa4, 0x92, 0x2a, 0x57, 0x1b, 0xb8,
	0xf4, 0xbd, 0x4b, 0xd1, 0x65, 0x00, 0xc5, 0x57, 0xc8, 0x4a, 0xec, 0x3b, 0x6a, 0x7b, 0x29, 0xe3,
	0x87, 0xa9, 0x63, 0xf1, 0x59, 0x19, 0x84, 0xd2, 0xb0, 0xd8, 0x51, 0xcb, 0xd8, 0x72, 0xec, 0x7c,
	0x94, 0x7e, 0x3b, 0x3a, 0x47, 0x4e, 0x0a, 0x1f, 0xb7, 0x7f, 0xf0, 0x4e, 0x19, 0x40, 0x71, 0x05,
	0xa9, 0x68, 0xa5, 0x60, 0x26, 0xfd, 0x87, 0x27, 0x95, 0xff, 0xfe, 0xc1, 0x2c, 0xb2, 0xd9, 0xd7,
	0x71, 0x6f, 0x54, 0x9b, 0x2f, 0x30, 0xfb, 0x64, 0x0e, 0xbd, 0x7a, 0x0d, 0xf3, 0x2d, 0x75, 0xd7,
	0x9b, 0x6b, 0x75, 0x2e, 0x9b, 0xe3, 0x51, 0x4f, 0x1f, 0x44, 0x15, 0x52, 0x17, 0x13, 0xf5, 0x0a,
	0xd2, 0xcf, 0x03, 0xba, 0xc3, 0x7f, 0xe6, 0xde, 0xe9, 0x5f, 0xc7, 0x3c, 0xb9, 0x3f, 0xe6, 0xc9,
	0xdf, 0x63, 0x9e, 0xfc, 0x3c, 0xe5, 0x93, 0xfb, 0x53, 0x3e, 0xf9, 0x73, 0xca, 0x27, 0xdf, 0xe6,
	0xf2, 0x2d, 0x5e, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xca, 0x95, 0x7d, 0x47, 0xa2, 0x02, 0x00,
	0x00,
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UpdatedAt) > 0 {
		i -= len(m.UpdatedAt)
		copy(dAtA[i:], m.UpdatedAt)
		i = encodeVarintPay(dAtA, i, uint64(len(m.UpdatedAt)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.CreatedAt) > 0 {
		i -= len(m.CreatedAt)
		copy(dAtA[i:], m.CreatedAt)
		i = encodeVarintPay(dAtA, i, uint64(len(m.CreatedAt)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.TerminalId) > 0 {
		i -= len(m.TerminalId)
		copy(dAtA[i:], m.TerminalId)
		i = encodeVarintPay(dAtA, i, uint64(len(m.TerminalId)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.OperatorId) > 0 {
		i -= len(m.OperatorId)
		copy(dAtA[i:], m.OperatorId)
		i = encodeVarintPay(dAtA, i, uint64(len(m.OperatorId)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.OrderNo) > 0 {
		i -= len(m.OrderNo)
		copy(dAtA[i:], m.OrderNo)
		i = encodeVarintPay(dAtA, i, uint64(len(m.OrderNo)))
		i--
		dAtA[i] = 0x3a
	}
	if m.TotalAmount != 0 {
		i = encodeVarintPay(dAtA, i, uint64(m.TotalAmount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintPay(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AuthCode) > 0 {
		i -= len(m.AuthCode)
		copy(dAtA[i:], m.AuthCode)
		i = encodeVarintPay(dAtA, i, uint64(len(m.AuthCode)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Method) > 0 {
		i -= len(m.Method)
		copy(dAtA[i:], m.Method)
		i = encodeVarintPay(dAtA, i, uint64(len(m.Method)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StoreName) > 0 {
		i -= len(m.StoreName)
		copy(dAtA[i:], m.StoreName)
		i = encodeVarintPay(dAtA, i, uint64(len(m.StoreName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.StoreId) > 0 {
		i -= len(m.StoreId)
		copy(dAtA[i:], m.StoreId)
		i = encodeVarintPay(dAtA, i, uint64(len(m.StoreId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Order != nil {
		{
			size, err := m.Order.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPay(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		{
			size, err := m.Error.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPay(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintPay(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x12
	}
	if m.Valid {
		i--
		if m.Valid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Detail) > 0 {
		i -= len(m.Detail)
		copy(dAtA[i:], m.Detail)
		i = encodeVarintPay(dAtA, i, uint64(len(m.Detail)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintPay(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPay(dAtA []byte, offset int, v uint64) int {
	offset -= sovPay(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StoreId)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	l = len(m.StoreName)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	l = len(m.AuthCode)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	if m.TotalAmount != 0 {
		n += 1 + sovPay(uint64(m.TotalAmount))
	}
	l = len(m.OrderNo)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	l = len(m.OperatorId)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	l = len(m.TerminalId)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Order != nil {
		l = m.Order.Size()
		n += 1 + l + sovPay(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Valid {
		n += 2
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovPay(uint64(l))
	}
	return n
}

func (m *Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	l = len(m.Detail)
	if l > 0 {
		n += 1 + l + sovPay(uint64(l))
	}
	return n
}

func sovPay(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPay(x uint64) (n int) {
	return sovPay(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPay
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoreId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoreName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAmount", wireType)
			}
			m.TotalAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalAmount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderNo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderNo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TerminalId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TerminalId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPay
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPay
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPay
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Order == nil {
				m.Order = &Order{}
			}
			if err := m.Order.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPay
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPay
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPay
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Valid = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPay
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPay
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPay
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Detail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Detail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPay
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPay
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPay(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPay
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPay
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPay
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPay
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPay
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPay
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPay        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPay          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPay = fmt.Errorf("proto: unexpected end of group")
)
