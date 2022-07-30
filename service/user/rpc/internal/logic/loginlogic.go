package logic

import (
	"context"
	"fmt"

	"seckill/service/user/rpc/internal/svc"
	"seckill/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	resp, err := l.svcCtx.UserModel.FindOneByPhoneNumber(l.ctx, in.PhoneNumber)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return &user.LoginResponse{
		Id:          resp.Id,
		AuthorityId: resp.AuthorityId,
	}, nil
}
