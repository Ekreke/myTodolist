package middleware

import (
	"time"

	"github.com/ekreke/myTodolist/pkg/e"
	"github.com/ekreke/myTodolist/utils"
	"github.com/gin-gonic/gin"
)

func USER_JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			// make sure that the remaining handlers won't be called
			c.Abort()
			return
		}
		// if handler use correctly , than call the next middleware handler
		c.Next()
	}
}
