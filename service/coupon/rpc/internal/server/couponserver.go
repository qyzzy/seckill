// Code generated by goctl. DO NOT EDIT!
// Source: coupon.proto

package server

import (
	"context"

	"seckill/service/coupon/rpc/coupon"
	"seckill/service/coupon/rpc/internal/logic"
	"seckill/service/coupon/rpc/internal/svc"
)

type CouponServer struct {
	svcCtx *svc.ServiceContext
	coupon.UnimplementedCouponServer
}

func NewCouponServer(svcCtx *svc.ServiceContext) *CouponServer {
	return &CouponServer{
		svcCtx: svcCtx,
	}
}

func (s *CouponServer) CreateCoupon(ctx context.Context, in *coupon.CreateCouponRequest) (*coupon.CreateCouponResponse, error) {
	l := logic.NewCreateCouponLogic(ctx, s.svcCtx)
	return l.CreateCoupon(in)
}

func (s *CouponServer) UpdateCouponStatus(ctx context.Context, in *coupon.UpdateCouponStatusRequest) (*coupon.UpdateCouponStatusResponse, error) {
	l := logic.NewUpdateCouponStatusLogic(ctx, s.svcCtx)
	return l.UpdateCouponStatus(in)
}

func (s *CouponServer) AddCouponTakeCount(ctx context.Context, in *coupon.AddCouponTakeCountRequest) (*coupon.AddCouponTakeCountResponse, error) {
	l := logic.NewAddCouponTakeCountLogic(ctx, s.svcCtx)
	return l.AddCouponTakeCount(in)
}

func (s *CouponServer) AddCouponUsedCount(ctx context.Context, in *coupon.AddCouponUsedCountRequest) (*coupon.AddCouponUsedCountResponse, error) {
	l := logic.NewAddCouponUsedCountLogic(ctx, s.svcCtx)
	return l.AddCouponUsedCount(in)
}