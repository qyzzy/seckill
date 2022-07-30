package logic

import (
	"context"
	"fmt"
	"seckill/common/cryptx"
	"seckill/service/user/rpc/user"

	"seckill/service/register/api/internal/svc"
	"seckill/service/register/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	response, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password),
		Age:         req.Age,
		Avatar:      req.Avatar,
		Gender:      req.Gender,
	})

	if err != nil {
		return nil, nil
	}

	fmt.Println(response)
	return
}
