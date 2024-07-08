package examplelogic

import (
	"context"

	"go-zero-example/micro/example/internal/svc"
	"go-zero-example/micro/example/proto/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HelloLogic) Hello(in *example.ExampleRequest) (*example.ExampleResponse, error) {
	// todo: add your logic here and delete this line

	return &example.ExampleResponse{}, nil
}
