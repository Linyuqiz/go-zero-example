// Code generated by goctl. DO NOT EDIT.
// Source: example.proto

package server

import (
	"context"

	"go-zero-example/micro/example/internal/logic/example"
	"go-zero-example/micro/example/internal/svc"
	"go-zero-example/micro/example/proto/example"
)

type ExampleServer struct {
	svcCtx *svc.ServiceContext
	example.UnimplementedExampleServer
}

func NewExampleServer(svcCtx *svc.ServiceContext) *ExampleServer {
	return &ExampleServer{
		svcCtx: svcCtx,
	}
}

func (s *ExampleServer) Hello(ctx context.Context, in *example.ExampleRequest) (*example.ExampleResponse, error) {
	l := examplelogic.NewHelloLogic(ctx, s.svcCtx)
	return l.Hello(in)
}
