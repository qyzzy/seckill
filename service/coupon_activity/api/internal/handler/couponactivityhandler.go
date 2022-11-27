package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"seckill/service/coupon_activity/api/internal/logic"
	"seckill/service/coupon_activity/api/internal/svc"
	"seckill/service/coupon_activity/api/internal/types"
)

func coupon_activityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateCouponActivityRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewcouponActivitylogic(r.Context(), svcCtx)
		resp, err := l.CouponActivity(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
