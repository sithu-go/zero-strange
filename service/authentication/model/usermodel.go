package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

var r redis.Redis = redis.Redis{
	Addr: "127.0.0.1:6379",
}

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindAll(ctx context.Context) ([]*User, error)
		FindOneByUsername(ctx context.Context, username string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *customUserModel) FindAll(ctx context.Context) ([]*User, error) {
	var resp = []*User{}
	err := m.QueryRowsNoCacheCtx(ctx, &resp, "SELECT * FROM `user`;")

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	userUsernameKey := fmt.Sprintf("%s%v", cacheUserUsernamePrefix, username)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, userUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		cuname := fmt.Sprintf("cache:user:username:%v", "strange")
		err := r.ExpireCtx(ctx, cuname, 10)
		if err != nil {
			return &resp, err
		}
		err = r.ExpireCtx(ctx, fmt.Sprintf("cache:user:id:%v", 1), 10)
		if err != nil {
			return &resp, err
		}
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
