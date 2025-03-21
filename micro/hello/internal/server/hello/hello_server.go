// Code generated by goctl. DO NOT EDIT.
// Source: hello.proto

package server

import (
	"context"

	"go-zero-example/micro/hello/internal/logic/hello"
	"go-zero-example/micro/hello/internal/svc"
	"go-zero-example/micro/hello/proto/hello"
)

type HelloServer struct {
	svcCtx *svc.ServiceContext
	hello.UnimplementedHelloServer
}

func NewHelloServer(svcCtx *svc.ServiceContext) *HelloServer {
	return &HelloServer{
		svcCtx: svcCtx,
	}
}

func (s *HelloServer) Hello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	l := hellologic.NewHelloLogic(ctx, s.svcCtx)
	return l.Hello(in)
}
