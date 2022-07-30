package logic

import (
	"context"

	"seckill/service/user/rpc/internal/svc"
	"seckill/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIdLogic {
	return &GetUserInfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByIdLogic) GetUserInfoById(in *user.GetByIdRequest) (*user.GetByIdResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetByIdResponse{}, nil
}
