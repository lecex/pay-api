# pay-api 支付模块API

## MACOS开发环境
```
    # enable go modules
    export GO111MODULE=on

    go get github.com/micro/micro/v2

    # protobuf 开发环境
    go get -u github.com/golang/·/protoc-gen-go
    go get -u github.com/micro/protoc-gen-micro/v2

    go get -u github.com/gogo/protobuf/proto
    go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
    go get -u github.com/gogo/protobuf/gogoproto
```
## 依赖模块
- [pay](https://github.com/lecex/pay)
## Pays 支付服务
    - AopF2F            商家扫用户
## Orders 订单服务
    - List              订单列表
    - Get               查询订单
    - Update            更新订单
## Configs 配置服务
    - SelfUpdate        登陆更新配置
    - List              配置列表
    - Get               查询配置
    - Create            增加配置
    - Update            更新配置
    - Delete            删除配置
## Health 健康检查
    - Health            健康检查