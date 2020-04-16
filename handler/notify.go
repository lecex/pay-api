package handler

import (
	"context"

	client "github.com/lecex/core/client"
	pb "github.com/lecex/pay-api/proto/notify"
)

// Notify 用户结构
type Notify struct {
	ServiceName string
}

// Notify 用户是否存在
func (srv *Notify) Notify(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Notify.Notify", req, res)
}
