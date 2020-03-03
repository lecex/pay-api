package handler

import (
	"context"
	"errors"

	client "github.com/lecex/core/client"
	"github.com/micro/go-micro/v2/metadata"

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
		err = client.Call(ctx, srv.ServiceName, "Configs.Update", req, res)
		if err != nil {
			err = errors.New("更新配置信息失败")
			return err
		}
	} else {
		err = errors.New("更新配置失败,未找到用户ID")
		return err
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
	return client.Call(ctx, srv.ServiceName, "Configs.Delete", &pb.Config{
		Id: req.Config.Id,
	}, res)
}
