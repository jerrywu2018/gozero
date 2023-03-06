package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-micro/mall/order/api/internal/logic"
	"go-zero-micro/mall/order/api/internal/svc"
	"go-zero-micro/mall/order/api/internal/types"
)

func getOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetOrderLogic(r.Context(), svcCtx)
		resp, err := l.GetOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}