package router

import (
	"im/middlewares"
	"im/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	//用户登录
	r.POST("/login", service.Login)

	//创建一个以/user为前缀的路由组,该组下所有路由的路径，都会加上该前缀
	//middlewares.AuthCheck()表示该组的专属中间件，会作用于该组下的所有路由。
	auth := r.Group("/user", middlewares.AuthCheck())

	//用户详情
	//由于 auth路由组，接口路径会变成 "/user/user/detail"
	auth.GET("/user/detail", service.UserDetail)
	//在路由注册时写 service.UserDetail（不带 ()），这是把函数本身作为参数传给 auth.GET
	//（因为 Gin 需要的是一个 gin.HandlerFunc 类型的函数作为回调，而不是执行函数的结果）。

	return r
}
