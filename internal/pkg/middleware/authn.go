package middleware

import (
	"fmt"

	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/known"
	"github.com/ekreke/myTodolist/pkg/token"
	"github.com/gin-gonic/gin"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk := c.Request.Header.Get("Authorization")
		if len(tk) == 0 {
			core.WriteResponse(c, token.ErrMissingHeader, nil)
		}
		var t string
		fmt.Sscanf(tk, "Bearer %s", &t)
		// TODO: delete
		_, err := token.Parse(t, token.K)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		id, err := store.GetUserIdByUserName(tk)
		if err != nil {
			core.WriteResponse(c, errno.ErrUserNotFound, nil)
			c.Abort()
			return
		}
		c.Set(known.XUsernameKey, tk)
		c.Set(known.XUserID, id)
		c.Next()
	}
}

// func ParseRequest(c *gin.Context) (string, error) {
// 	header := c.Request.Header.Get("Authorization")
// 	var t string
// 	// 从请求头中取出 token
// 	fmt.Sscanf(header, "Bearer %s", &t)
// 	return Parse(t, config.key)
// }
