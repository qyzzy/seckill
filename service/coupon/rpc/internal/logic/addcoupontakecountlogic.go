package logic

import (
	"context"

	"seckill/service/coupon/rpc/coupon"
	"seckill/service/coupon/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCouponTakeCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCouponTakeCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCouponTakeCountLogic {
	return &AddCouponTakeCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCouponTakeCountLogic) AddCouponTakeCount(in *coupon.AddCouponTakeCountRequest) (*coupon.AddCouponTakeCountResponse, error) {
	takeCount, err := l.svcCtx.CouponModel.AddCouponTakeCount(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &coupon.AddCouponTakeCountResponse{
		TakeCount: takeCount,
	}, nil
}
