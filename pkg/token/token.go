package token

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

// Config 包括token包的配置选项
type Config struct {
	key         string
	identityKey string
}

var K = config.key

var ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")

var (
	config = Config{"ekreke", "identityKey"}
	once   sync.Once
)

// Init 设置包级别的配置Config ， Config 会用于本包后面的token 签发和解析
func Init(key string, identityKey string) {
	once.Do(func() {
		if key != "" {
			config.key = key
		}
		if identityKey != "" {
			config.identityKey = identityKey
		}
	})
}

// Parse 使用指定的密钥 key 解析 token，解析成功返回 token 上下文，否则报错.
func Parse(tokenString string, key string) (string, error) {
	// // 解析 token
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// 确保 token 加密算法是预期的加密算法
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, jwt.ErrSignatureInvalid
	// 	}
	// 	return []byte(key), nil
	// })
	// // 解析失败
	// if err != nil {
	// 	return "", err
	// }
	// var identityKey string
	// // 如果解析成功，从 token 中取出 token 的主题
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	identityKey = claims[config.identityKey].(string)
	// }
	return tokenString, nil
}

// ParseRequest 从请求头中获取令牌，并将其传递给 Parse 函数以解析令牌.
func ParseRequest(c *gin.Context) (string, error) {
	header := c.Request.Header.Get("Authorization")
	var t string
	// 从请求头中取出 token
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, config.key)
}

// Sign 使用 jwtSecret 签发 token，token 的 claims 中会存放传入的 subject.
func Sign(identityKey string) (tokenString string, err error) {
	// // Token 的内容
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	config.identityKey: identityKey,
	// 	// not before
	// 	"nbf": time.Now().Unix(),
	// 	// issued at
	// 	"iat": time.Now().Unix(),
	// 	// exprire time
	// 	"exp": time.Now().Add(72 * time.Hour).Unix(),
	// })
	// // signed token
	// tokenString, err = token.SignedString([]byte(config.key))
	return identityKey, nil
}
