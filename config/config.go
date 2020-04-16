package config

import (
	"github.com/lecex/core/config"
	"github.com/lecex/core/env"
	PB "github.com/lecex/user/proto/permission"
)

// 	Conf 配置
// 	Service // 服务名称
//	Method // 方法
//	Auth // 是否认证授权
//	Policy // 是否认证权限
//	Name // 权限名称
//	Description // 权限解释
var Conf config.Config = config.Config{
	Name:    env.Getenv("MICRO_API_NAMESPACE", "go.micro.api.") + "pay-api",
	Version: "latest",
	Service: map[string]string{
		"user": env.Getenv("USER_SERVICE", "go.micro.srv.user"),
		"pay":  env.Getenv("PAY_SERVICE", "go.micro.srv.pay"),
	},
	Permissions: []*PB.Permission{
		// 支付管理
		{Method: "Pays.AopF2F", Auth: false, Policy: false, Name: "商家扫用户", Description: "商家扫用户付款码付款。"},
		// 订单管理
		{Method: "Orders.Get", Auth: true, Policy: true, Name: "查询订单", Description: "查询订单信息权限。"},
		{Method: "Orders.Update", Auth: true, Policy: true, Name: "更新订单", Description: "更新订单信息。"},
		{Method: "Orders.List", Auth: true, Policy: true, Name: "订单列表", Description: "查询订单列表"},
		{Method: "Orders.SelfList", Auth: true, Policy: false, Name: "本账户订单列表", Description: "本账户订单列表。"},
		{Method: "Orders.SelfAmount", Auth: true, Policy: false, Name: "本账户订单统计", Description: "本账户订单统计。"},
		// 配置管理
		{Method: "Configs.SelfUpdate", Auth: true, Policy: false, Name: "更新登陆用户支付配置", Description: "更新登陆用户支付配置信息"},
		{Method: "Configs.Info", Auth: true, Policy: false, Name: "获取登陆用户支付配置", Description: "获取登陆用户支付配置信息"},
		{Method: "Configs.Create", Auth: true, Policy: true, Name: "创建支付配置", Description: "创建新角色权限。"},
		{Method: "Configs.Delete", Auth: true, Policy: true, Name: "删除支付配置", Description: "删除支付配置。"},
		{Method: "Configs.Update", Auth: true, Policy: true, Name: "更新支付配置", Description: "更新支付配置信息。"},
		{Method: "Configs.Get", Auth: true, Policy: true, Name: "查询支付配置", Description: "查询支付配置信息权限。"},
		{Method: "Configs.List", Auth: true, Policy: true, Name: "支付配置列表", Description: "查询支付配置列表"},
	},
}
