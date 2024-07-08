package hello

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-example/server/internal/logic/hello"
	"go-zero-example/server/internal/svc"
	"go-zero-example/server/internal/types"
)

func GetHelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := hello.NewGetHelloLogic(r.Context(), svcCtx)
		resp, err := l.GetHello(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
