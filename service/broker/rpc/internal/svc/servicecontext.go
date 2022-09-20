package svc

import (
	"zero-strange/service/broker/model"
	"zero-strange/service/broker/rpc/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	UserLoginModel model.UserLoginModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := c.Mongo.DataSource
	return &ServiceContext{
		Config:         c,
		UserLoginModel: model.NewUserLoginModel(conn, "logs", "log"),
	}
}
