package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/util/log"

	server "github.com/micro/go-micro/v2/server"

	client "github.com/lecex/core/client"
	"github.com/lecex/core/env"
	configPB "github.com/lecex/pay-api/proto/config"
	healthPB "github.com/lecex/pay-api/proto/health"
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
	PayService := env.Getenv("PAY_NAME", "pay")
	orderPB.RegisterOrdersHandler(srv.Server, &Order{PayService})
	configPB.RegisterConfigsHandler(srv.Server, &Config{PayService})
	payPB.RegisterPaysHandler(srv.Server, &Pay{PayService})
	healthPB.RegisterHealthHandler(srv.Server, &Health{})

	go Sync() // 同步前端权限
}

// Sync 同步
func Sync() {
	time.Sleep(5 * time.Second)
	req := &PB.Request{
		Permissions: Conf.Permissions,
	}
	res := &PB.Response{}
	err := client.Call(context.TODO(), Conf.UserService, "Permissions.Sync", req, res)
	if err != nil {
		log.Log(err)
		Sync()
	}
}
