package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"seckill/service/user_coupon/api/internal/logic"
	"seckill/service/user_coupon/api/internal/svc"
	"seckill/service/user_coupon/api/internal/types"
)

func normalGetCouponHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NormalGetCouponRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewNormalGetCouponLogic(r.Context(), svcCtx)
		resp, err := l.NormalGetCoupon(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
