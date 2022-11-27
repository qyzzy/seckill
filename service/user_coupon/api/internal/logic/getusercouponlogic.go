package logic

import (
	"context"

	"seckill/service/user_coupon/api/internal/svc"
	"seckill/service/user_coupon/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCouponLogic {
	return &GetUserCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserCouponLogic) GetUserCoupon(req *types.GetUserCouponRequest) (resp *types.GetUserCouponResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
