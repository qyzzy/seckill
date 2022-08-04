package logic

import (
	"context"

	"seckill/service/coupon/rpc/coupon"
	"seckill/service/coupon/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCouponByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindCouponByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCouponByIdLogic {
	return &FindCouponByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindCouponByIdLogic) FindCouponById(in *coupon.FindCouponRequest) (*coupon.FindCouponResponse, error) {
	// todo: add your logic here and delete this line

	return &coupon.FindCouponResponse{}, nil
}
