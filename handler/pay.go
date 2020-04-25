package handler

import (
	"context"

	client "github.com/lecex/core/client"

	pb "github.com/lecex/pay-api/proto/pay"
)

// Pay 支付结构
type Pay struct {
	ServiceName string
}

// AopF2F 商家扫用户
func (srv *Pay) AopF2F(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Pays.AopF2F", req, res)
}

// Query 订单查询
func (srv *Pay) Query(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Pays.Query", req, res)
}
