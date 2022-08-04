package logic

import (
	"context"
	"seckill/common/cryptx"
	"seckill/service/user/rpc/user"

	"seckill/service/user_register/api/internal/svc"
	"seckill/service/user_register/api/internal/types"

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

func (l *RegisterLogic) Register(req *types.RegisterRequest) (*types.RegisterResponse, error) {
	response, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Gender:      req.Gender,
		Age:         req.Age,
		Avatar:      req.Avatar,
		Password:    cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password),
	})

	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{
		Id:          response.Id,
		AuthorityId: response.AuthorityId,
		Status:      response.Status,
	}, nil
}
