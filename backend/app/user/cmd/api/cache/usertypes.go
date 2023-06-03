package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	//MongoDB中的UUID
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
	//唯一身份标志符
	StuffCode string `bson:"stuffCode" json:"stuffCode"`
	Name      string `bson:"name" json:"name"`
	Password  string `bson:"password" json:"password"`
	//教职工身份: 不细分就是admin || user
	Role string `bson:"role" json:"role"`
}
