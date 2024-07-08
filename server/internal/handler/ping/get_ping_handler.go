package ping

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-example/server/internal/logic/ping"
	"go-zero-example/server/internal/svc"
	"go-zero-example/server/internal/types"
)

func GetPingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ping.NewGetPingLogic(r.Context(), svcCtx)
		resp, err := l.GetPing(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
