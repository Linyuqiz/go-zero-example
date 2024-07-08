package ping

import (
	"context"

	"go-zero-example/server/internal/svc"
	"go-zero-example/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPingLogic {
	return &GetPingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPingLogic) GetPing(req *types.PingRequest) (resp *types.PingResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
