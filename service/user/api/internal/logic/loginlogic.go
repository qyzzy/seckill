package logic

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"user_service/api/internal/svc"
	"user_service/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	if len(strings.TrimSpace(req.Email)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, 1)
	fmt.Println(userInfo, err)

	userInfo, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	fmt.Println("userInfo :", userInfo, err)
	if userInfo == nil {
		return
	}
	if userInfo.Password != req.Password {
		return nil, errors.New("用户密码不正确")
	}

	//// ---start---
	//now := time.Now().Unix()
	//accessExpire := l.svcCtx.Config.Auth.AccessExpire
	//jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	//if err != nil {
	//	return nil, err
	//}
	//// ---end---

	return &types.LoginReply{
		Id: userInfo.Id,
		//AccessToken:  jwtToken,
		//AccessExpire: now + accessExpire,
		//RefreshAfter: now + accessExpire/2,
	}, nil
}
