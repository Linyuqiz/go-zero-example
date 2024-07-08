package example

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-example/server/internal/logic/example"
	"go-zero-example/server/internal/svc"
	"go-zero-example/server/internal/types"
)

func GetExampleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExampleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := example.NewGetExampleLogic(r.Context(), svcCtx)
		resp, err := l.GetExample(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
