package logic

import (
	"context"
	"seckill/service/coupon/rpc/coupon"
	"seckill/service/coupon_activity/model"

	"seckill/service/coupon_activity/api/internal/svc"
	"seckill/service/coupon_activity/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponActivitylogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewcouponActivitylogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponActivitylogic {
	return &CouponActivitylogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponActivitylogic) CouponActivity(req *types.CreateCouponActivityRequest) (*types.CreateCouponActivityResponse, error) {
	activityId, err := l.svcCtx.CouponActivityModel.CreateCouponActivity(l.ctx, &model.CouponActivity{
		Name:         req.ActivityName,
		CouponTypeId: req.TypeId,
	})

	if err != nil {
		return nil, err
	}

	response, err := l.svcCtx.CouponRpc.CreateCoupon(l.ctx, &coupon.CreateCouponRequest{
		Name:           req.CouponName,
		TypeId:         req.TypeId,
		IsMutex:        req.IsMutex,
		ProductId:      req.ProductId,
		ShopId:         req.ProductId,
		CategoryId:     req.CategoryId,
		WithAmount:     req.WithAmount,
		UsedAmount:     req.UsedAmount,
		Stock:          req.Stock,
		ValidType:      req.ValidType,
		ValidDay:       req.ValidDay,
		ValidStartTime: req.ValidStartTime,
		ValidEndTime:   req.ValidEndTime,
		UpdateUserId:   req.UpdateUserId,
		Status:         req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateCouponActivityResponse{
		ActivityId: activityId,
		CouponId:   response.Id,
	}, nil
}
