package handler

import (
	"context"
	"errors"
	"fmt"

	client "github.com/lecex/core/client"
	"github.com/micro/go-micro/v2/metadata"

	pb "github.com/lecex/pay-api/proto/order"
)

// Order 订单结构
type Order struct {
	ServiceName string
}

// SelfAmount 查询自己的订单收款总金额
func (srv *Order) SelfAmount(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// meta["Userid"] 通过 meta 获取用户 id --- So this function needs token to use
	meta, _ := metadata.FromContext(ctx)
	if userID, ok := meta["Userid"]; ok {
		where := "true"
		if req.ListQuery.Where != "" {
			where = req.ListQuery.Where
		}
		req.ListQuery.Where = fmt.Sprintf("%s AND store_id = '%s' AND stauts = true", where, userID)
		return client.Call(ctx, srv.ServiceName, "Orders.Amount", req, res)
	} else {
		return errors.New("获取用户失败,未找到用户ID")
	}
}

// SelfList 查询自己的订单列表
func (srv *Order) SelfList(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// meta["Userid"] 通过 meta 获取用户 id --- So this function needs token to use
	meta, _ := metadata.FromContext(ctx)
	if userID, ok := meta["Userid"]; ok {
		where := "true"
		if req.ListQuery.Where != "" {
			where = req.ListQuery.Where
		}
		req.ListQuery.Where = fmt.Sprintf("%s AND store_id = '%s'", where, userID)
		return client.Call(ctx, srv.ServiceName, "Orders.List", req, res)
	} else {
		return errors.New("获取用户失败,未找到用户ID")
	}
}

// List 订单列表
func (srv *Order) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Orders.List", req, res)
}

// Get 获取订单
func (srv *Order) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Orders.Get", req, res)
}

// // Create 创建订单
// func (srv *Order) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
// 	return client.Call(ctx, srv.ServiceName, "Orders.Create", req, res)
// }

// Update 更新订单
func (srv *Order) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Orders.Update", req, res)
}

// Delete 删除订单
// func (srv *Order) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
// 	return client.Call(ctx, srv.ServiceName, "Orders.Delete", req, res)
// }
