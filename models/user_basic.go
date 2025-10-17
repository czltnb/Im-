package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserBasic struct {
	Identity   string `bson:"_id"`
	Account    string `bson:"account"`
	Password   string `bson:"password"`
	Nickname   string `bson:"nickname"`
	Sex        int    `bson:"sex"`
	Email      string `bson:"email"`
	Avatar     string `bson:"avatar"`
	CreateTime int64  `bson:"create_time"`
	UpdateTime int64  `bson:"update_time"`
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}

func GetUserBasicByAccountAndPassword(account, password string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).
		Decode(ub)
	return ub, err
}
