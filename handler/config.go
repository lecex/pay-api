package handler

import (
	"context"
	"errors"

	"github.com/micro/go-micro/v2/metadata"

	client "github.com/lecex/core/client"
	pb "github.com/lecex/pay-api/proto/config"
)

// Config 配置结构
type Config struct {
	ServiceName string
}

// SelfUpdate 配置列表
func (srv *Config) SelfUpdate(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// meta["Userid"] 通过 meta 获取用户 id --- So this function needs token to use
	meta, _ := metadata.FromContext(ctx)
	if userID, ok := meta["Userid"]; ok {
		req.Config.Id = userID
		if Username, ok := meta["Username"]; ok {
			req.Config.StoreName = Username
		}
		return client.Call(ctx, srv.ServiceName, "Configs.SelfUpdate", req, res)
	}
	return errors.New("更新用户失败,未找到用户ID")
}

// Info 获取用户支付信息
func (srv *Config) Info(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.New("no auth meta-data found in request")
	}
	if userID, ok := meta["Userid"]; ok {
		// 获取用户信息
		if req.Config == nil {
			req.Config = &pb.Config{}
		}
		req.Config.Id = userID
		err = client.Call(ctx, srv.ServiceName, "Configs.Get", req, res)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Empty userID")
	}
	return err
}

// List 配置列表
func (srv *Config) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Configs.List", req, res)
}

// Get 获取配置
func (srv *Config) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Configs.Get", req, res)
}

// Create 创建配置
func (srv *Config) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Configs.Create", req, res)
}

// Update 更新配置
func (srv *Config) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Configs.Update", req, res)
}

// Delete 删除配置
func (srv *Config) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Configs.Delete", req, res)
}
