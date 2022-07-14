package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"product-search-service/rpc/product"

	"product-search-service/api/internal/svc"
	"product-search-service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchReply, err error) {
	fmt.Println(req, l.ctx.Value("productId"))
	productIdNumber := json.Number(fmt.Sprintf("%d", req.Id))
	logx.Infof("userId: %s", productIdNumber)
	productId, err := productIdNumber.Int64()
	if err != nil {
		return nil, err
	}

	// 使用user rpc
	productInfoReply, err := l.svcCtx.ProductRpc.GetProduct(l.ctx, &product.IdReq{
		Id: productId,
	})
	if err != nil {
		return nil, err
	}

	return &types.SearchReply{
		Name:  productInfoReply.Name,
		Price: productInfoReply.Price,
		Stock: productInfoReply.Stock,
	}, nil
	return
}
