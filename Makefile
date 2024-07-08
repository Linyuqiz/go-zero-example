PROJECT_ROOT = $(shell pwd)

.PHONY: format-api

dependency-up:
	@docker-compose up -d

dependency-down:
	@docker-compose down

fmt-api:
	@goctl api format --dir=api

gen-api:
	@goctl api go --api=api/server.api --dir=server --style=go_zero

gen-grpc:
	@goctl rpc protoc proto/example/example.proto --multiple --go_out=micro/example --go-grpc_out=micro/example --zrpc_out=micro/example --style=go_zero
	@goctl rpc protoc proto/hello/hello.proto --multiple --go_out=micro/hello --go-grpc_out=micro/hello --zrpc_out=micro/hello --style=go_zero
