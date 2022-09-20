package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindAll(ctx context.Context) ([]*User, error)
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
	// var userBuilderQueryRows = strings.Join(builder.RawFieldNames(&User{}), ",")
	// var userBuilderQueryRows = "'id', 'userna'"
	// fmt.Println(userBuilderQueryRows, "JJJ")

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
