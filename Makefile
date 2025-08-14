run:
	go run cmd/main.go --addr=5001 --robotkey=xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx


.PHONY: build
build:
	go build -o alertmanager-wechat-webhook cmd/main.go