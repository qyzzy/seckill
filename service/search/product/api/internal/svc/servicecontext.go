package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"product-search-service/api/internal/config"
	"product-search-service/rpc/product"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
