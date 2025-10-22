package helper

import (
	"crypto/md5"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserClaims struct {
	Identity           primitive.ObjectID `json:"identity"`
	Email              string             `json:"email"`
	jwt.StandardClaims                    //嵌入JWT标准声明（包含过期时间、签发时间等默认字段）
}

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// JWT签名密钥（生产环境需使用复杂密钥，且避免硬编码）
var myKey = []byte("im")

// GenerateToken
// 生成 token
// 生成 JWT 令牌
func GenerateToken(identity, email string) (string, error) {
	//把后端传来的字符串类型的identity转换成ObjectID类型
	objectID, err := primitive.ObjectIDFromHex(identity)
	if err != nil {
		return "", err
	}
	// 构建自定义声明（包含用户信息和标准声明）
	UserClaim := &UserClaims{
		Identity:       objectID,
		Email:          email,
		StandardClaims: jwt.StandardClaims{},
	}
	// 创建令牌：指定签名算法为HS256（HMAC-SHA256），并传入声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	// 用密钥签名令牌，生成最终的令牌字符串
	//为啥要传入一个密钥 myKey ?????????????????????????????????????
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token,提取其中的用户信息
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil // 回调函数返回签名密钥，用于验证令牌
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}
