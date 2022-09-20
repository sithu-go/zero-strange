package logic

import (
	"context"

	"zero-strange/service/authentication/model"
	"zero-strange/service/broker/api/internal/svc"
	"zero-strange/service/broker/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetusersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetusersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetusersLogic {
	return &GetusersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetusersLogic) Getusers() (resp []types.UserReply, err error) {
	// todo: add your logic here and delete this line
	var users []*model.User
	users, err = l.svcCtx.UserModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		var userReply types.UserReply
		userReply.Id = user.Id
		userReply.Name = user.Name
		userReply.Username = user.Username
		userReply.Gender = user.Gender
		userReply.CreatedAt = user.CreatedAt.String()
		userReply.UpdatedAt = user.UpdatedAt.String()
		resp = append(resp, userReply)
	}
	return resp, nil
}
