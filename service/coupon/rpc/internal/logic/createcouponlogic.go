package logic

import (
	"context"
	"seckill/common/timeparser"
	"seckill/service/coupon/model"

	"seckill/service/coupon/rpc/coupon"
	"seckill/service/coupon/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCouponLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCouponLogic {
	return &CreateCouponLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCouponLogic) CreateCoupon(in *coupon.CreateCouponRequest) (*coupon.CreateCouponResponse, error) {
	couponId, err := l.svcCtx.CouponModel.CreateCoupon(l.ctx, &model.Coupon{
		Name:           in.Name,
		TypeId:         in.TypeId,
		IsMutex:        in.IsMutex,
		ProductId:      in.ProductId,
		ShopId:         in.ShopId,
		CategoryId:     in.CategoryId,
		WithAmount:     in.WithAmount,
		UsedAmount:     in.UsedAmount,
		Stock:          in.Stock,
		ValidType:      in.ValidType,
		ValidDay:       in.ValidDay,
		ValidStartTime: timeparser.UnixToTime(in.ValidStartTime),
		ValidEndTime:   timeparser.UnixToTime(in.ValidStartTime),
		UpdateUserId:   in.UpdateUserId,
		Status:         in.Status,
	})

	if err != nil {
		return nil, err
	}

	return &coupon.CreateCouponResponse{
		Id: couponId,
	}, nil
}
