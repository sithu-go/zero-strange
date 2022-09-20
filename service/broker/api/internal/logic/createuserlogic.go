package logic

import (
	"context"
	"time"

	"zero-strange/service/authentication/model"
	"zero-strange/service/broker/api/internal/svc"
	"zero-strange/service/broker/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type CreateuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateuserLogic {
	return &CreateuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateuserLogic) Createuser(req *types.UserCreateReq) (resp *types.UserReply, err error) {
	// todo: add your logic here and delete this line
	user := model.User{
		Name:      req.Name,
		Username:  req.Username,
		Gender:    req.Gender,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := l.svcCtx.UserModel.Insert(l.ctx, &user)
	switch err {
	case nil:
		insertId, _ := result.LastInsertId()
		var responseUser *model.User
		responseUser, err = l.svcCtx.UserModel.FindOne(l.ctx, insertId)
		if err != nil {
			return nil, err
		}
		resp.Id = responseUser.Id
		resp.Name = responseUser.Name
		resp.Username = responseUser.Username
		resp.Gender = responseUser.Gender
		resp.CreatedAt = responseUser.CreatedAt.String()
		resp.UpdatedAt = responseUser.UpdatedAt.String()
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlc.ErrNotFound
	default:
		return nil, err
	}
}
