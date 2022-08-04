package logic

import (
	"context"
	"seckill/common/timeparser"
	"seckill/service/coupon/api/internal/svc"
	"seckill/service/coupon/api/internal/types"
	"seckill/service/coupon/model"

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
	id, err := l.svcCtx.FullDiscountCouponModel.Insert(l.ctx, &model.FullDiscountCoupon{
		Name:           req.Name,
		WithAmount:     req.WithAmount,
		UsedAmount:     req.UsedAmount,
		Stock:          req.Stock,
		StartTime:      timeparser.UnixToTime(req.StartTime),
		EndTime:        timeparser.UnixToTime(req.EndTime),
		Status:         1,
		ValidEndTime:   timeparser.UnixToTime(req.ValidEndTime),
		ValidStartTime: timeparser.UnixToTime(req.ValidStartTime),
	})

	return &types.AddCouponResponse{
		Id: id,
	}, err
}
