package logic

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"zero-strange/service/authentication/api/internal/svc"
	"zero-strange/service/authentication/api/internal/types"
	"zero-strange/service/authentication/model"
	"zero-strange/service/broker/rpc/types/log"

	"github.com/golang-jwt/jwt/v4"
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
	// todo: add your logic here and delete this line

	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("invalid parameter")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("user does not exist")
	default:
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	if userInfo.Password != req.Password {
		return nil, errors.New("username or password is incorrect")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	if err != nil {
		return nil, err
	}

	// use login rpc
	logReply, err := l.svcCtx.LoginRpc.LogLogin(l.ctx, &log.LoginReq{
		Username: userInfo.Username,
	})
	fmt.Println("LOG REPLY", logReply)
	if err != nil {
		fmt.Println("ERRROR GETTING", err.Error())
		return nil, err
	}

	return &types.LoginReply{
		Id:           userInfo.Id,
		Name:         userInfo.Name,
		Gender:       userInfo.Gender,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
