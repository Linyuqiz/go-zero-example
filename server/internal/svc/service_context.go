package svc

import (
	"go-zero-example/micro/example/client/example"
	"go-zero-example/micro/hello/client/hello"
	"go-zero-example/server/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	ExampleZRpc example.Example
	HelloZRpc   hello.Hello
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		ExampleZRpc: example.NewExample(zrpc.MustNewClient(c.ExampleZRpc)),
		HelloZRpc:   hello.NewHello(zrpc.MustNewClient(c.HelloZRpc)),
	}
}
