package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/linzijie1998/mini-tiktok/cmd/user/constant"
	"github.com/linzijie1998/mini-tiktok/cmd/user/dal/cache"
	"github.com/linzijie1998/mini-tiktok/cmd/user/dal/db"
	"github.com/linzijie1998/mini-tiktok/cmd/user/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/user"
	"github.com/linzijie1998/mini-tiktok/model"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRegisterService struct {
	ctx context.Context
}

func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{ctx: ctx}
}

func (s *UserRegisterService) UserRegister(req *user.RegisterRequest) (int64, string, error) {
	// 1. 判断用户是否已经注册, 未注册将返回`gorm.ErrRecordNotFound`错误
	res, err := db.QueryFirstUserInfoByUsername(s.ctx, req.Username, constant.UserRegisterInfoQueryString)
	if res != nil {
		// 查询到用户信息
		return 0, "", errno.UserAlreadyExistErr
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 错误信息不是ErrRecordNotFound
		return 0, "", err
	}
	// 2. 对用户密码进行加密
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, "", errno.ServiceErr.WithMessage(fmt.Sprintf("密码加密错误, %s", err.Error()))
	}
	// 3. 创建用户, 并且写入数据库
	newUserInfo := &model.User{
		//UUID:     uuid.New(),
		Username: req.Username,
		Password: string(bytes),
		Nickname: req.Username,
	}
	if err = db.CreateUserInfos(s.ctx, []*model.User{newUserInfo}); err != nil {
		return 0, "", errno.ServiceErr.WithMessage(fmt.Sprintf("数据库添加错误, %s", err.Error()))
	}
	// 4. 若该UID或Username存在空值缓存, 则将其删除
	if err = cache.DelUserInfoNullKey(s.ctx, newUserInfo.Id); err != nil {
		return 0, "", err
	}
	if err = cache.DelUserLoginNullKey(s.ctx, req.Username); err != nil {
		return 0, "", err
	}
	// 5. 获得用户id, 并且颁发token
	claims, err := jwt.BuildCustomClaims(
		newUserInfo.Id, global.Configs.JWT.ExpiresTime, global.Configs.JWT.Issuer, global.Configs.JWT.Subject)
	if err != nil {
		return 0, "", errno.ServiceErr.WithMessage(fmt.Sprintf("JWT创建错误, %s", err.Error()))
	}
	token, err := jwt.NewJWT(global.Configs.JWT.SigningKey).CreateToken(claims)
	if err != nil {
		return 0, "", errno.ServiceErr.WithMessage(fmt.Sprintf("JWT加密错误, %s", err.Error()))
	}
	// 5. 返回用户id和token
	return newUserInfo.Id, token, nil
}
