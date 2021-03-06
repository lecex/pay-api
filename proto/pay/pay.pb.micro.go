// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/pay/pay.proto

package pay

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Pays service

type PaysService interface {
	// AopF2F 商家扫用户付款码
	AopF2F(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Query(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Cancel(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Refund(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	OpenRefund(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	AffirmRefund(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type paysService struct {
	c    client.Client
	name string
}

func NewPaysService(name string, c client.Client) PaysService {
	return &paysService{
		c:    c,
		name: name,
	}
}

func (c *paysService) AopF2F(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Pays.AopF2F", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paysService) Query(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Pays.Query", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paysService) Cancel(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Pays.Cancel", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paysService) Refund(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Pays.Refund", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paysService) OpenRefund(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Pays.OpenRefund", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paysService) AffirmRefund(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Pays.AffirmRefund", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pays service

type PaysHandler interface {
	// AopF2F 商家扫用户付款码
	AopF2F(context.Context, *Request, *Response) error
	Query(context.Context, *Request, *Response) error
	Cancel(context.Context, *Request, *Response) error
	Refund(context.Context, *Request, *Response) error
	OpenRefund(context.Context, *Request, *Response) error
	AffirmRefund(context.Context, *Request, *Response) error
}

func RegisterPaysHandler(s server.Server, hdlr PaysHandler, opts ...server.HandlerOption) error {
	type pays interface {
		AopF2F(ctx context.Context, in *Request, out *Response) error
		Query(ctx context.Context, in *Request, out *Response) error
		Cancel(ctx context.Context, in *Request, out *Response) error
		Refund(ctx context.Context, in *Request, out *Response) error
		OpenRefund(ctx context.Context, in *Request, out *Response) error
		AffirmRefund(ctx context.Context, in *Request, out *Response) error
	}
	type Pays struct {
		pays
	}
	h := &paysHandler{hdlr}
	return s.Handle(s.NewHandler(&Pays{h}, opts...))
}

type paysHandler struct {
	PaysHandler
}

func (h *paysHandler) AopF2F(ctx context.Context, in *Request, out *Response) error {
	return h.PaysHandler.AopF2F(ctx, in, out)
}

func (h *paysHandler) Query(ctx context.Context, in *Request, out *Response) error {
	return h.PaysHandler.Query(ctx, in, out)
}

func (h *paysHandler) Cancel(ctx context.Context, in *Request, out *Response) error {
	return h.PaysHandler.Cancel(ctx, in, out)
}

func (h *paysHandler) Refund(ctx context.Context, in *Request, out *Response) error {
	return h.PaysHandler.Refund(ctx, in, out)
}

func (h *paysHandler) OpenRefund(ctx context.Context, in *Request, out *Response) error {
	return h.PaysHandler.OpenRefund(ctx, in, out)
}

func (h *paysHandler) AffirmRefund(ctx context.Context, in *Request, out *Response) error {
	return h.PaysHandler.AffirmRefund(ctx, in, out)
}
