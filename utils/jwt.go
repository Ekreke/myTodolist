package utils

import (
	"time"

	"github.com/ekreke/myTodolist/conf"
	"github.com/ekreke/myTodolist/pkg/logging"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecrect = []byte(conf.LoadJwtSecrect())

type UserClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// generate token with username and password
func GenerateUserToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24)
	claims := UserClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "myTodolist",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecrect)
	if err != nil {
		logging.Fatal(err)
	}
	return token, nil
}

// parse token to get username and password
func ParseUserToken(token string) (*UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecrect, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
