package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ UserLoginModel = (*customUserLoginModel)(nil)

type (
	// UserLoginModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLoginModel.
	UserLoginModel interface {
		userLoginModel
	}

	customUserLoginModel struct {
		*defaultUserLoginModel
	}
)

// NewUserLoginModel returns a model for the mongo.
func NewUserLoginModel(url, db, collection string) UserLoginModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customUserLoginModel{
		defaultUserLoginModel: newDefaultUserLoginModel(conn),
	}
}
