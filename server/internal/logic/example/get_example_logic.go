package example

import (
	"context"

	"go-zero-example/server/internal/svc"
	"go-zero-example/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExampleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExampleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExampleLogic {
	return &GetExampleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExampleLogic) GetExample(req *types.ExampleRequest) (resp *types.ExampleResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
