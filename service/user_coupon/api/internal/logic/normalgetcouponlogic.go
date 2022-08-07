package logic

import (
	"context"
	"fmt"
	"seckill/service/coupon/rpc/coupon"
	"seckill/service/user_coupon/model"

	"seckill/service/user_coupon/api/internal/svc"
	"seckill/service/user_coupon/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NormalGetCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNormalGetCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NormalGetCouponLogic {
	return &NormalGetCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NormalGetCouponLogic) NormalGetCoupon(req *types.NormalGetCouponRequest) (*types.NormalGetCouponResponse, error) {
	response, err := l.svcCtx.CouponRpc.AddCouponTakeCount(l.ctx, &coupon.AddCouponTakeCountRequest{
		Id: req.CouponId,
	})
	fmt.Println(response)
	if err != nil {
		return nil, err
	}

	uuid, status, err := l.svcCtx.UserCouponModel.CreateUserCoupon(l.ctx, &model.UserCoupon{
		UserId:   req.UserId,
		CouponId: req.CouponId,
	})

	if err != nil {
		return nil, err
	}

	return &types.NormalGetCouponResponse{
		Uuid:   uuid,
		Status: status,
	}, nil
}
