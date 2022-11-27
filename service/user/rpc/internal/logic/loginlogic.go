package logic

import (
	"context"

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
	uid, authorityId, err := l.svcCtx.UserModel.Login(l.ctx, in.PhoneNumber, in.Password)
	if err != nil {
		return nil, err
	}

	return &user.LoginResponse{
		Id:          uid,
		AuthorityId: authorityId,
	}, nil
}
