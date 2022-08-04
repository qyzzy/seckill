package logic

import (
	"context"
	"seckill/service/user/model"

	"seckill/service/user/rpc/internal/svc"
	"seckill/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	uid, authorityId, status, err := l.svcCtx.UserModel.RegisterByPhoneNumber(l.ctx, &model.User{
		Name:        in.Name,
		PhoneNumber: in.PhoneNumber,
		Password:    in.Password,
		Avatar:      in.Avatar,
		Age:         in.Age,
		Gender:      in.Gender,
		Status:      1,
		AuthorityId: 22,
	})

	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{
		Id:          uid,
		AuthorityId: authorityId,
		Status:      status,
	}, nil
}
