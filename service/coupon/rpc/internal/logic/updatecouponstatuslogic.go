package logic

import (
	"context"

	"seckill/service/coupon/rpc/coupon"
	"seckill/service/coupon/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCouponStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCouponStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponStatusLogic {
	return &UpdateCouponStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCouponStatusLogic) UpdateCouponStatus(in *coupon.UpdateCouponStatusRequest) (*coupon.UpdateCouponStatusResponse, error) {
	// todo: add your logic here and delete this line

	return &coupon.UpdateCouponStatusResponse{}, nil
}
