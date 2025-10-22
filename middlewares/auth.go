package middlewares

import (
	"im/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
身份认证中间件，验证客户端请求中的Token的合法性，确保携带有效Token的用户才能访问受保护的接口
*/

//HandlerFunc是Context的子类型吗???

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		userClaims, err := helper.AnalyseToken(token)
		if err != nil {
			c.Abort() //这是?????????
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户认证不通过",
			})
			return
		}
		c.Set("user_claims", userClaims) //userClaims的定义在hepler.go中，含有JWT相关声明
		c.Next()
	}
}
