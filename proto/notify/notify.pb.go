// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/notify/notify.proto

package notify

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

type Pair struct {
	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4f26e246bead917, []int{0}
}
func (m *Pair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return m.Size()
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// A HTTP request as RPC
// Forward by the api handler
type Request struct {
	Method string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path   string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Get    map[string]*Pair `protobuf:"bytes,4,rep,name=get,proto3" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Post   map[string]*Pair `protobuf:"bytes,5,rep,name=post,proto3" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body   string           `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Url    string           `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4f26e246bead917, []int{1}
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

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// A HTTP response as RPC
// Expected response for the api handler
type Response struct {
	StatusCode int32            `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Header     map[string]*Pair `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body       string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4f26e246bead917, []int{2}
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

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "notify.Pair")
	proto.RegisterType((*Request)(nil), "notify.Request")
	proto.RegisterMapType((map[string]*Pair)(nil), "notify.Request.GetEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "notify.Request.HeaderEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "notify.Request.PostEntry")
	proto.RegisterType((*Response)(nil), "notify.Response")
	proto.RegisterMapType((map[string]*Pair)(nil), "notify.Response.HeaderEntry")
}

func init() { proto.RegisterFile("proto/notify/notify.proto", fileDescriptor_b4f26e246bead917) }

var fileDescriptor_b4f26e246bead917 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcb, 0x4e, 0xf2, 0x40,
	0x14, 0xc7, 0x7b, 0xa3, 0xc0, 0xe1, 0x4b, 0x3e, 0x32, 0x0b, 0x33, 0xa0, 0x69, 0x48, 0x57, 0xc4,
	0x04, 0x34, 0xe0, 0xc2, 0xb8, 0xf3, 0x42, 0x70, 0x65, 0x48, 0x37, 0xae, 0x8b, 0x1d, 0x2c, 0xb1,
	0x32, 0xb5, 0x33, 0x35, 0xe9, 0x5b, 0xf8, 0x42, 0x2e, 0xdc, 0xb9, 0x64, 0xe9, 0xd2, 0xc0, 0x8b,
	0x98, 0x99, 0x0e, 0x97, 0x34, 0x2c, 0x8c, 0xba, 0xe2, 0x5c, 0xfe, 0xe7, 0xf0, 0x9f, 0xdf, 0x49,
	0xa1, 0x11, 0x27, 0x94, 0xd3, 0xa3, 0x19, 0xe5, 0xd3, 0x49, 0xa6, 0x7e, 0xba, 0xb2, 0x86, 0xec,
	0x3c, 0x73, 0x8f, 0xc1, 0x1a, 0xf9, 0xd3, 0x04, 0xd5, 0xc1, 0x7c, 0x20, 0x19, 0xd6, 0x5b, 0x7a,
	0xbb, 0xea, 0x89, 0x10, 0xed, 0x81, 0xfd, 0xec, 0x47, 0x29, 0x61, 0xd8, 0x68, 0x99, 0xed, 0xaa,
	0xa7, 0x32, 0xf7, 0xcd, 0x84, 0xb2, 0x47, 0x9e, 0x52, 0xc2, 0xb8, 0xd0, 0x3c, 0x12, 0x1e, 0xd2,
	0x40, 0x0d, 0xaa, 0x0c, 0x21, 0xb0, 0x62, 0x9f, 0x87, 0xd8, 0x90, 0x55, 0x19, 0xa3, 0x3e, 0xd8,
	0x21, 0xf1, 0x03, 0x92, 0x60, 0xb3, 0x65, 0xb6, 0x6b, 0xbd, 0xfd, 0xae, 0x32, 0xa4, 0x96, 0x75,
	0xaf, 0x65, 0x77, 0x30, 0xe3, 0x49, 0xe6, 0x29, 0x29, 0x3a, 0x04, 0xf3, 0x9e, 0x70, 0x6c, 0xc9,
	0x09, 0x5c, 0x9c, 0x18, 0x12, 0x9e, 0xcb, 0x85, 0x08, 0x75, 0xc0, 0x8a, 0x29, 0xe3, 0xb8, 0x24,
	0xc5, 0x8d, 0xa2, 0x78, 0x44, 0x99, 0x52, 0x4b, 0x99, 0xf0, 0x38, 0xa6, 0x41, 0x86, 0xed, 0xdc,
	0xa3, 0x88, 0x05, 0x85, 0x34, 0x89, 0x70, 0x39, 0xa7, 0x90, 0x26, 0x51, 0x73, 0x08, 0xb5, 0x2d,
	0x5f, 0x3b, 0x30, 0xb9, 0x50, 0x92, 0x60, 0xe4, 0x5b, 0x6b, 0xbd, 0x7f, 0xab, 0xbf, 0x15, 0x54,
	0xbd, 0xbc, 0x75, 0x66, 0x9c, 0xea, 0xcd, 0x2b, 0xa8, 0xac, 0xec, 0xfe, 0x62, 0xcb, 0x00, 0xaa,
	0xeb, 0x77, 0xfc, 0x7c, 0x8d, 0xfb, 0xaa, 0x43, 0xc5, 0x23, 0x2c, 0xa6, 0x33, 0x46, 0x90, 0x03,
	0xc0, 0xb8, 0xcf, 0x53, 0x76, 0x49, 0x03, 0x22, 0xb7, 0x95, 0xbc, 0xad, 0x0a, 0x3a, 0x59, 0x1f,
	0xce, 0x90, 0x64, 0x0f, 0x36, 0x64, 0xf3, 0x0d, 0x3b, 0x2f, 0xb7, 0xc2, 0x6b, 0x6e, 0xf0, 0xfe,
	0x19, 0xcc, 0xde, 0x04, 0xec, 0x1b, 0xd9, 0x41, 0x1d, 0xb0, 0xcf, 0xa3, 0x69, 0xec, 0x67, 0xe8,
	0x7f, 0xe1, 0xe0, 0xcd, 0x7a, 0xd1, 0xa7, 0xab, 0x09, 0xf9, 0x2d, 0xb9, 0x0b, 0x7d, 0xfe, 0x2d,
	0xf9, 0x05, 0x7e, 0x5f, 0x38, 0xfa, 0x7c, 0xe1, 0xe8, 0x9f, 0x0b, 0x47, 0x7f, 0x59, 0x3a, 0xda,
	0x7c, 0xe9, 0x68, 0x1f, 0x4b, 0x47, 0x1b, 0xdb, 0xf2, 0x33, 0xea, 0x7f, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x00, 0x6e, 0x23, 0x92, 0x63, 0x03, 0x00, 0x00,
}

func (m *Pair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Values[iNdEx])
			copy(dAtA[i:], m.Values[iNdEx])
			i = encodeVarintNotify(dAtA, i, uint64(len(m.Values[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintNotify(dAtA, i, uint64(len(m.Key)))
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
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintNotify(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintNotify(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Post) > 0 {
		for k := range m.Post {
			v := m.Post[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintNotify(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintNotify(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintNotify(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Get) > 0 {
		for k := range m.Get {
			v := m.Get[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintNotify(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintNotify(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintNotify(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Header) > 0 {
		for k := range m.Header {
			v := m.Header[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintNotify(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintNotify(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintNotify(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintNotify(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Method) > 0 {
		i -= len(m.Method)
		copy(dAtA[i:], m.Method)
		i = encodeVarintNotify(dAtA, i, uint64(len(m.Method)))
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
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintNotify(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Header) > 0 {
		for k := range m.Header {
			v := m.Header[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintNotify(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintNotify(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintNotify(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.StatusCode != 0 {
		i = encodeVarintNotify(dAtA, i, uint64(m.StatusCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNotify(dAtA []byte, offset int, v uint64) int {
	offset -= sovNotify(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovNotify(uint64(l))
	}
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovNotify(uint64(l))
		}
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovNotify(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovNotify(uint64(l))
	}
	if len(m.Header) > 0 {
		for k, v := range m.Header {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovNotify(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovNotify(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovNotify(uint64(mapEntrySize))
		}
	}
	if len(m.Get) > 0 {
		for k, v := range m.Get {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovNotify(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovNotify(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovNotify(uint64(mapEntrySize))
		}
	}
	if len(m.Post) > 0 {
		for k, v := range m.Post {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovNotify(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovNotify(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovNotify(uint64(mapEntrySize))
		}
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovNotify(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovNotify(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StatusCode != 0 {
		n += 1 + sovNotify(uint64(m.StatusCode))
	}
	if len(m.Header) > 0 {
		for k, v := range m.Header {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovNotify(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovNotify(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovNotify(uint64(mapEntrySize))
		}
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovNotify(uint64(l))
	}
	return n
}

func sovNotify(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNotify(x uint64) (n int) {
	return sovNotify(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotify
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
			return fmt.Errorf("proto: Pair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotify
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
				return ErrInvalidLengthNotify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotify
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
				return ErrInvalidLengthNotify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotify(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotify
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotify
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
				return ErrIntOverflowNotify
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
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotify
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
				return ErrInvalidLengthNotify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotify
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
				return ErrInvalidLengthNotify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotify
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
				return ErrInvalidLengthNotify
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNotify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = make(map[string]*Pair)
			}
			var mapkey string
			var mapvalue *Pair
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNotify
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNotify
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthNotify
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthNotify
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNotify
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthNotify
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthNotify
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Pair{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipNotify(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthNotify
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Header[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Get", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotify
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
				return ErrInvalidLengthNotify
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNotify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Get == nil {
				m.Get = make(map[string]*Pair)
			}
			var mapkey string
			var mapvalue *Pair
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNotify
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNotify
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthNotify
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthNotify
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNotify
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthNotify
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthNotify
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Pair{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipNotify(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthNotify
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Get[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Post", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotify
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
				return ErrInvalidLengthNotify
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNotify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Post == nil {
				m.Post = make(map[string]*Pair)
			}
			var mapkey string
			var mapvalue *Pair
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNotify
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNotify
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthNotify
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthNotify
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNotify
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthNotify
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthNotify
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Pair{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipNotify(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthNotify
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Post[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotify
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
				return ErrInvalidLengthNotify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotify
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
				return ErrInvalidLengthNotify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotify(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotify
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotify
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
				return ErrIntOverflowNotify
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
				return fmt.Errorf("proto: wrong wireType = %d for field StatusCode", wireType)
			}
			m.StatusCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatusCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotify
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
				return ErrInvalidLengthNotify
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNotify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = make(map[string]*Pair)
			}
			var mapkey string
			var mapvalue *Pair
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNotify
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNotify
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthNotify
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthNotify
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNotify
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthNotify
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthNotify
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Pair{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipNotify(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthNotify
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Header[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotify
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
				return ErrInvalidLengthNotify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotify(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotify
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotify
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
func skipNotify(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNotify
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
					return 0, ErrIntOverflowNotify
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
					return 0, ErrIntOverflowNotify
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
				return 0, ErrInvalidLengthNotify
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNotify
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNotify
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNotify        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNotify          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNotify = fmt.Errorf("proto: unexpected end of group")
)
