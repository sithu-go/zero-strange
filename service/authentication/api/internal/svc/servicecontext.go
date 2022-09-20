package svc

import (
	"zero-strange/service/authentication/api/internal/config"
	"zero-strange/service/authentication/model"
	"zero-strange/service/broker/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	LoginRpc  user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		LoginRpc:  user.NewUser(zrpc.MustNewClient(c.LoginRpc)),
	}
}
