package hellologic

import (
	"context"

	"go-zero-example/micro/hello/internal/svc"
	"go-zero-example/micro/hello/proto/hello"

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

func (l *HelloLogic) Hello(in *hello.HelloRequest) (*hello.HelloResponse, error) {
	// todo: add your logic here and delete this line

	return &hello.HelloResponse{}, nil
}
