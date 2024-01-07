package util

import (
	"time"

	"github.com/ekreke/myTodolist/conf"
	"github.com/golang-jwt/jwt/v4"
)

// TODO add jwt

var jwtSecret = []byte(conf.LoadJwtSecrect())

type UserClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateUserToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := UserClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "cmall",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}
