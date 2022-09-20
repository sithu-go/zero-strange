package logic

import (
	"context"
	"fmt"

	"zero-strange/service/broker/model"
	"zero-strange/service/broker/rpc/internal/svc"
	"zero-strange/service/broker/rpc/types/log"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogLoginLogic {
	return &LogLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogLoginLogic) LogLogin(in *log.LoginReq) (*log.LoginInfoReply, error) {
	// todo: add your logic here and delete this line
	var userLogin model.UserLogin
	userLogin.Username = in.Username
	userLogin.Message = fmt.Sprintf("%v logged in", in.Username)
	err := l.svcCtx.UserLoginModel.Insert(l.ctx, &userLogin)
	if err != nil {
		fmt.Printf("error inserting log %v\n", err)
		return nil, err
	}

	return &log.LoginInfoReply{
		Username: in.Username,
	}, nil
}
