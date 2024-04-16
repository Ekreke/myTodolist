package main

import (
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/ekreke/myTodolist/pkg/token"
)

func GenerateRandomProjectPwd() (pwd string, err error) {
	const pwdLength = 15                                                                // 密码长度
	const validChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 可用字符集
	b := make([]byte, pwdLength)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	// 将生成的随机字节转换为字符串
	pwd = strings.Join(strings.Fields(fmt.Sprintf("%x", b)), "")
	// 确保密码长度不超过15个字符
	if len(pwd) > pwdLength {
		pwd = pwd[:pwdLength]
	}
	return pwd, nil
}

func main() {
	// tokenString, err := token.Sign("ekreke34555")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(tokenString)
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTM1Mzc5OTgsImlhdCI6MTcxMzI3ODc5OCwiaWRlbnRpdHlLZXkiOiJla3Jla2UzNDU1NSIsIm5iZiI6MTcxMzI3ODc5OH0.9IvQkFmwF8podTRW5oc2fZXLV_K3vwyqp6l3Q5MAQA4"

	fmt.Sscanf(tokenString, "Bearer %s", &tokenString)
	username, err := token.Parse(tokenString, token.K)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(username)
}
