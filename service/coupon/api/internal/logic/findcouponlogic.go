package logic

import (
	"context"

	"seckill/service/coupon/api/internal/svc"
	"seckill/service/coupon/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCouponLogic {
	return &FindCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCouponLogic) FindCoupon(req *types.FindCouponRequest) (*types.FindCouponResponse, error) {
	response, err := l.svcCtx.CouponModel.FindOne(l.ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &types.FindCouponResponse{
		Name:   response.Name,
		TypeId: response.TypeId,
	}, nil
}
