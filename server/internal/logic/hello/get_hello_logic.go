package hello

import (
	"context"

	"go-zero-example/micro/hello/proto/hello"
	"go-zero-example/server/internal/svc"
	"go-zero-example/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHelloLogic {
	return &GetHelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHelloLogic) GetHello(req *types.HelloRequest) (resp *types.HelloResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.HelloZRpc.Hello(l.ctx, &hello.HelloRequest{})
	if err != nil {
		return nil, err
	}

	return &types.HelloResponse{}, nil
}
