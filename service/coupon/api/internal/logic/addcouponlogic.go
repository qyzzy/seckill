package logic

import (
	"context"

	"seckill/service/coupon/api/internal/svc"
	"seckill/service/coupon/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCouponLogic {
	return &AddCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCouponLogic) AddCoupon(req *types.AddCouponRequest) (resp *types.AddCouponResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
