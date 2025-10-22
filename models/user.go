package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
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

func (User) CollectionName() string {
	return "t_user"
}

func GetUserByAccountAndPassword(account, password string) (*User, error) {
	ub := new(User)
	err := Mongo.Collection(User{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).
		Decode(ub)
	return ub, err
}

// 根据_id从MongoDB中返回User用户信息

func GetUserByIdentity(identity primitive.ObjectID) (*User, error) {
	ub := new(User)
	err := Mongo.Collection(User{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", identity}}).
		Decode(ub) //把根据identity从数据库中查询到的User对象，返回到ub中，并return
	return ub, err
}
