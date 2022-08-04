package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"seckill/service/user/rpc/user"
	"seckill/service/user_login/api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
