package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"seckill/service/coupon/api/internal/logic"
	"seckill/service/coupon/api/internal/svc"
	"seckill/service/coupon/api/internal/types"
)

func findCouponHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindCouponRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFindCouponLogic(r.Context(), svcCtx)
		resp, err := l.FindCoupon(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
