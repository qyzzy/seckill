package logic

import (
	"context"
	"fmt"
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
	newUser := &model.User{
		Name:        in.Name,
		PhoneNumber: in.PhoneNumber,
		Password:    in.Password,
		Age:         in.Age,
		Avatar:      in.Avatar,
		Gender:      in.Gender,
	}
	newUser.AuthorityId = 22
	newUser.Status = 1
	res, err := l.svcCtx.UserModel.Insert(l.ctx, newUser)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return &user.RegisterResponse{}, nil
}
