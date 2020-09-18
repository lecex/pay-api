package handler

import (
	"context"
	"errors"

	client "github.com/lecex/core/client"
	"github.com/micro/go-micro/v2/metadata"

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

// Cancel 撤销接口
func (srv *Pay) Cancel(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// meta["Userid"] 通过 meta 获取用户 id --- So this function needs token to use
	meta, _ := metadata.FromContext(ctx)
	if userID, ok := meta["Userid"]; ok {
		req.Order.StoreId = userID
		return client.Call(ctx, srv.ServiceName, "Pays.Cancel", req, res)
	} else {
		return errors.New("获取用户失败,未找到用户ID")
	}
}

// Refund 退款接口
func (srv *Pay) Refund(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// meta["Userid"] 通过 meta 获取用户 id --- So this function needs token to use
	meta, _ := metadata.FromContext(ctx)
	if userID, ok := meta["Userid"]; ok {
		req.Order.StoreId = userID
		return client.Call(ctx, srv.ServiceName, "Pays.Refund", req, res)
	} else {
		return errors.New("获取用户失败,未找到用户ID")
	}
}

// OpenRefund 退款接口
func (srv *Pay) OpenRefund(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	req.Order.Verify = true // 开放退款接口需要 AffirmRefund 接口确认完成退款
	return client.Call(ctx, srv.ServiceName, "Pays.Refund", req, res)
}

// AffirmRefund 确认退款接口
func (srv *Pay) AffirmRefund(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// meta["Userid"] 通过 meta 获取用户 id --- So this function needs token to use
	meta, _ := metadata.FromContext(ctx)
	if userID, ok := meta["Userid"]; ok {
		req.Order.StoreId = userID
		return client.Call(ctx, srv.ServiceName, "Pays.AffirmRefund", req, res)
	} else {
		return errors.New("获取用户失败,未找到用户ID")
	}
}
