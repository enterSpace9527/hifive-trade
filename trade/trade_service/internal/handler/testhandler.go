package handler

import (
	"net/http"

	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/logic"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func testHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
