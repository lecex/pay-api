package handler

import (
	"context"

	client "github.com/lecex/core/client"

	pb "github.com/lecex/pay-api/proto/order"
)

// Order 订单结构
type Order struct {
	ServiceName string
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
// 	return client.Call(ctx, srv.ServiceName, "Orders.Delete", &pb.Order{
// 		Id: req.Order.Id,
// 	}, res)
// }
