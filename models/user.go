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

//查询某邮箱是否已经被注册

func GetUserByEmail(email string) (int64, error) {
	return Mongo.Collection(User{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"email", email}})
	//bson.D 是 MongoDB Go 驱动中定义的一种有序键值对数据结构，用于表示 BSON 格式的查询条件（BSON 是 MongoDB 使用的二进制 JSON 格式，类似 JSON 但支持更多数据类型）。
	//bson.D{{"email", email}} 表示一个查询条件：筛选出 email 字段的值等于变量 email 的文档（即 “查询邮箱为 xxx 的用户”）。
	//兼容性：MongoDB Go 驱动的所有数据库操作方法（如 CountDocuments、Find、InsertOne 等）都强制要求传入 context，这是驱动设计的规范，确保操作可被外部控制。
}
