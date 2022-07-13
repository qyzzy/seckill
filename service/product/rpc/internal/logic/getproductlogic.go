package logic

import (
	"context"

	"product_service/rpc/internal/svc"
	"product_service/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *product.IdReq) (*product.ProductInfoReply, error) {
	one, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &product.ProductInfoReply{
		Id:     one.Id,
		Name:   one.Name,
		Price:  one.Price,
		Stock:  one.Stock,
		TypeId: one.TypeId,
		SalesVolume: one.SalesVolume,
		SpuId:  one.SpuId,
	}, nil
}
