package logic

import (
	"context"
	"seckill/common/cryptx"
	"seckill/common/jwt"
	"seckill/service/user/rpc/user"
	"time"

	"seckill/service/user_login/api/internal/svc"
	"seckill/service/user_login/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (*types.LoginResponse, error) {
	response, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		PhoneNumber: req.PhoneNumber,
		Password:    cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password),
	})

	if err != nil {
		return nil, err
	}

	jwtToken, err := jwt.GetToken(
		l.svcCtx.Config.Auth.AccessSecret,
		time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire,
		response.Id,
	)

	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Id:          response.Id,
		AuthorityId: response.AuthorityId,
		AccessToken: jwtToken,
	}, nil
}
