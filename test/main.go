package main

import (
	"crypto/rand"
	"fmt"
	"strings"
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
	pwd, err := GenerateRandomProjectPwd()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pwd)
	}
}
