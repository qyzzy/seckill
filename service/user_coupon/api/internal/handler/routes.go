// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"seckill/service/user_coupon/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/user/coupon",
					Handler: normalGetCouponHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/user/coupon/:uuid",
					Handler: getUserCouponHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
