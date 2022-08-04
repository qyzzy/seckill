package logic

import (
	"context"

	"seckill/service/coupon/rpc/coupon"
	"seckill/service/coupon/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductStockByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductStockByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductStockByIdLogic {
	return &DeductStockByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductStockByIdLogic) DeductStockById(in *coupon.DeductStockRequest) (*coupon.DeductStockResponse, error) {
	// todo: add your logic here and delete this line

	return &coupon.DeductStockResponse{}, nil
}
