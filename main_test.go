package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/micro/go-micro/v2/metadata"

	"github.com/lecex/pay-api/handler"
	payPB "github.com/lecex/pay-api/proto/pay"
)

func TestPermissions(t *testing.T) {
	// Sync 同步权限
	// permissions, _ := json.Marshal(Conf.Permissions)
	// fmt.Println(permissions)
}

func TestMobileBuild(t *testing.T) {
	h := handler.Pay{}
	req := &payPB.Request{
		Verify: "632541",
		Uuid:   "asfdasfasfasafsafs",
	}
	res := &payPB.Response{}
	err := h.MobileBuild(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestPayGet(t *testing.T) {
	h := handler.Pay{}
	req := &payPB.Request{
		Pay: &payPB.Pay{
			Payname:  `bvbv0111`,
			Password: `1234asdasdasdasdasdas561234asdasdasdasdasdas561234asdasdasdasdasdas561234asdasdasdasdasdas561234asdasdasdasdasdas56`,
			Mobile:   `13953186114`,
			Email:    `bvbv0a1@qq.com`,
			Origin:   `pay-api`,
		},
	}
	res := &payPB.Response{}
	err := h.Get(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestPayInfo(t *testing.T) {
	h := handler.Pay{}
	req := &payPB.Request{}
	res := &payPB.Response{}
	meta := map[string]string{
		"Payid": "fed9c9a5-c03f-4d8d-be7c-95b9f24c9299",
	}
	ctx := metadata.NewContext(context.TODO(), meta)
	err := h.Info(ctx, req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}
