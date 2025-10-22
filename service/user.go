package service

import (
	"im/helper"
	"im/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
Gin 框架中的 *gin.Context 是一个自定义类型，它嵌入了标准库的 context.Context 接口，因此具备了标准 Context 的所有方法（Deadline()、Done() 等）。但 *gin.Context 还扩展了大量 Gin 框架特有的方法，用于处理 HTTP 请求，例如：
PostForm(key string) string：获取 POST 表单中的字段值（你代码中使用的方法）
Query(key string) string：获取 URL 查询参数
JSON(code int, obj interface{})：返回 JSON 响应
*/

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码不能为空",
		})
		return
	}
	ub, err := models.GetUserByAccountAndPassword(account, helper.GetMd5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	token, err := helper.GenerateToken(ub.Identity, ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误" + err.Error(),
		})
		return
	}
	/**
	gin.H 的主要用途是快速构建 JSON 响应体。,它的定义非常简单：
	type H map[string]interface{}
	也就是说，gin.H 本质上是一个 map[string]interface{}（键为字符串、值为任意类型的映射），只是被 Gin 框架起了一个更简洁的别名，方便在编写响应时使用。
	*/
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func UserDetail(c *gin.Context) {
	//从Context中获取User_claims对象
	u, _ := c.Get("user_claims")
	//从User_claims对象中取出user的identity
	uc := u.(*helper.UserClaims)
	//根据identity从数据库中查询用户信息
	user, err := models.GetUserByIdentity(uc.Identity)
	if err != nil {
		log.Printf("[DB数据库 ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据查询成功",
		"data": user,
	})
}
