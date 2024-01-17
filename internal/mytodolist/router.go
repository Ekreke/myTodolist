package mytodolist

import (
	"github.com/ekreke/myTodolist/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func installRouters(g *gin.Engine) error {
	// g.NoRoute(func(c *gin.Context) {
	// 	core.WriteResponse(c, errno.ErrPageNotFound, nil)
	// })
	// 注册 pprof 路由
	// pprof.Register(g)
	// uc := user.New(store.S)
	// ug := g.Group("/user")
	// {
	// 	ug.POST("/login", uc.Login)
	// 	ug.POST("/register", uc.Register)
	// 	ug.Use(middleware.Authn())
	// 	// ug.GET("/auth")
	// 	ug.GET("/info", middleware.Authn(), uc.Info)
	// 	ug.GET("/token", func(ctx *gin.Context) {
	// 		// resp := ctx.Request.Header.Get("Authorization")
	// 		// token, err := token.ParseRequest(ctx)
	// 		header := middleware.GetHeader(ctx)
	// 		// if err != nil {
	// 		// 	ctx.JSON(400, gin.H{
	// 		// 		"resp": err.Error()})
	// 		// }
	// 		ctx.JSON(200, gin.H{
	// 			"resp:": header})
	// 	})
	// }
	g.GET("/test", middleware.Authn())

	return nil
}
