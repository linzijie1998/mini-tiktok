package main

import (
	"context"
	"github.com/linzijie1998/mini-tiktok/cmd/user/pack"
	"github.com/linzijie1998/mini-tiktok/cmd/user/service"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/user"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 1. 检查请求格式是否正确
	if len(req.Username) == 0 || len(req.Password) < 6 {
		return pack.BuildRegisterResp(0, "", errno.ParamErr), nil
	}
	// 2. 处理请求
	id, token, err := service.NewUserRegisterService(ctx).UserRegister(req)
	if err != nil {
		return pack.BuildRegisterResp(0, "", err), nil
	}
	return pack.BuildRegisterResp(id, token, nil), nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	// 1. 检查请求格式是否正确
	if len(req.Username) == 0 || len(req.Password) < 6 {
		return pack.BuildLoginResp(0, "", errno.ParamErr), nil
	}
	// 2. 处理请求
	id, token, err := service.NewUserLoginService(ctx).UserLogin(req)
	if err != nil {
		return pack.BuildLoginResp(0, "", err), nil
	}
	return pack.BuildLoginResp(id, token, nil), nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.InfoRequest) (*user.InfoResponse, error) {
	// 1. 检查请求格式是否正确
	if req.UserId <= 0 {
		return pack.BuildInfoResp(nil, errno.ParamErr), nil
	}
	// 2. 处理请求
	userInfo, err := service.NewUserInfoService(ctx).UserInfo(req)
	if err != nil {
		return pack.BuildInfoResp(nil, err), nil
	}
	return pack.BuildInfoResp(userInfo, nil), nil
}
