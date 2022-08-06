package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"seckill/service/coupon/api/internal/logic"
	"seckill/service/coupon/api/internal/svc"
	"seckill/service/coupon/api/internal/types"
)

func addCouponHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCouponRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddCouponLogic(r.Context(), svcCtx)
		resp, err := l.AddCoupon(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
