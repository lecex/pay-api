package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/util/log"

	server "github.com/micro/go-micro/v2/server"

	client "github.com/lecex/core/client"
	configPB "github.com/lecex/pay-api/proto/config"
	healthPB "github.com/lecex/pay-api/proto/health"
	notifyPB "github.com/lecex/pay-api/proto/notify"
	orderPB "github.com/lecex/pay-api/proto/order"
	payPB "github.com/lecex/pay-api/proto/pay"

	"github.com/lecex/pay-api/config"
	PB "github.com/lecex/user/proto/permission"
)

// Handler 注册方法
type Handler struct {
	Server server.Server
}

var Conf = config.Conf

// Register 注册
func (srv *Handler) Register() {
	orderPB.RegisterOrdersHandler(srv.Server, &Order{Conf.Service["pay"]})
	configPB.RegisterConfigsHandler(srv.Server, &Config{Conf.Service["pay"]})
	payPB.RegisterPaysHandler(srv.Server, &Pay{Conf.Service["pay"]})
	healthPB.RegisterHealthHandler(srv.Server, &Health{})
	notifyPB.RegisterNotifyHandler(srv.Server, &Notify{})

	go Sync() // 同步前端权限
}

// Sync 同步
func Sync() {
	time.Sleep(5 * time.Second)
	req := &PB.Request{
		Permissions: Conf.Permissions,
	}
	res := &PB.Response{}
	err := client.Call(context.TODO(), Conf.Service["user"], "Permissions.Sync", req, res)
	if err != nil {
		log.Log(err)
		Sync()
	}
}
