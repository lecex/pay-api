// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/notify/notify.proto

package notify

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

// Client API for Notify service

type NotifyService interface {
	// Notify 异步通知
	Notify(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type notifyService struct {
	c    client.Client
	name string
}

func NewNotifyService(name string, c client.Client) NotifyService {
	return &notifyService{
		c:    c,
		name: name,
	}
}

func (c *notifyService) Notify(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Notify.Notify", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notify service

type NotifyHandler interface {
	// Notify 异步通知
	Notify(context.Context, *Request, *Response) error
}

func RegisterNotifyHandler(s server.Server, hdlr NotifyHandler, opts ...server.HandlerOption) error {
	type notify interface {
		Notify(ctx context.Context, in *Request, out *Response) error
	}
	type Notify struct {
		notify
	}
	h := &notifyHandler{hdlr}
	return s.Handle(s.NewHandler(&Notify{h}, opts...))
}

type notifyHandler struct {
	NotifyHandler
}

func (h *notifyHandler) Notify(ctx context.Context, in *Request, out *Response) error {
	return h.NotifyHandler.Notify(ctx, in, out)
}