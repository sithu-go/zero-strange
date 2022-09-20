package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLogin struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	Username string    `bson:"username" json:"username"`
	Message  string    `bson:"message" json:"message"`
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
