package logic

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"zero-strange/service/authentication/model"
	"zero-strange/service/broker/api/internal/svc"
	"zero-strange/service/broker/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetuserLogic {
	return &GetuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetuserLogic) Getuser(r *http.Request) (resp *types.UserReply, err error) {
	// todo: add your logic here and delete this line
	username := strings.TrimPrefix(r.URL.Path, "/v1/users/")
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, username)
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

	return &types.UserReply{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		Gender:    user.Gender,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}
