package logic

import (
	"context"

	"product-search-service/rpc/internal/svc"
	"product-search-service/rpc/types/product"

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
	// todo: add your logic here and delete this line

	return &product.ProductInfoReply{}, nil
}
