package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
*
拿到mongoDB对应的database
*/
var Mongo = InitMongoDB()

func InitMongoDB() *mongo.Database {
	//canel是“手动取消函数”，用于在超时前主动终止上下文（比如操作提前完成时，避免等待到超时才释放资源）
	ctx, canel := context.WithTimeout(context.Background(), 10*time.Second)

	// 确保函数结束时无论是否超时，都主动释放资源
	defer canel()

	//options 包是 MongoDB Go 驱动的 “配置中心”，几乎所有需要自定义的行为都通过它提供的类型来设置。
	//options.Credential是options包中定义好的结构体，用于存储数据库登录认证信息
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "root",
		Password: "123456",
	}).ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Println("连接 MongoDB 出现错误", err)
		return nil
	}
	return client.Database("im")
}
