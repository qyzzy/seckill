package logic

import (
	"context"

	"seckill/service/coupon/rpc/coupon"
	"seckill/service/coupon/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCouponUsedCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCouponUsedCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCouponUsedCountLogic {
	return &AddCouponUsedCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCouponUsedCountLogic) AddCouponUsedCount(in *coupon.AddCouponUsedCountRequest) (*coupon.AddCouponUsedCountResponse, error) {
	usedCount, err := l.svcCtx.CouponModel.AddCouponUsedCount(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &coupon.AddCouponUsedCountResponse{
		UsedCount: usedCount,
	}, nil
}
