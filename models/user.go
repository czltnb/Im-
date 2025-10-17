package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

func GetUserBasicByAccountAndPassword(account, password string) (*User, error) {
	ub := new(User)
	err := Mongo.Collection(User{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).
		Decode(ub)
	return ub, err
}
