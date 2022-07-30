package logic

import (
	"context"

	"seckill/service/user/rpc/internal/svc"
	"seckill/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateRequest) (*user.UpdateResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateResponse{}, nil
}
