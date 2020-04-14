package handler

import (
	"context"
	"fmt"

	pb "github.com/lecex/pay-api/proto/notify"
)

// Notify 用户结构
type Notify struct {
}

// Notify 用户是否存在
func (srv *Notify) Notify(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	fmt.Println(req)
	res.StatusCode = 200
	res.Body = string("true")
	return nil
}
