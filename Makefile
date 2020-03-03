.PHONY: git
git:
	git add .
	git commit -m"自动提交 git 代码"
	git push
.PHONY: tag
tag:
	git push --tags
.PHONY: micro
micro:
	micro api --enable_rpc=true

.PHONY: proto
proto:
	protoc -I . --micro_out=. --gogofaster_out=. proto/pay/pay.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/config/config.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/order/order.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/health/health.proto

.PHONY: docker
docker:
	docker build -f Dockerfile  -t pay-api .
.PHONY: run
run:
	go run main.go