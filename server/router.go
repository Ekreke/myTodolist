package server

import (
	"github.com/ekreke/myTodolist/api"
	"github.com/ekreke/myTodolist/docs"
	_ "github.com/ekreke/myTodolist/docs"
	"github.com/ekreke/myTodolist/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Gin router

func NewRouter() *gin.Engine {
	r := gin.Default()
	// TODO: add middleware
	// r.Use(middleware.Cors())

	// swagger
	r.Use(middleware.Cors())
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		// TODO: session
		// store := cookie.NewStore([]byte(sdk.VERSION))
		// r.Use(sessions.Sessions("myssion", store))
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		v1.POST("user/login", api.UserLogin)

		v1.POST("user/register", api.UserRegister)

		authed := v1.Group("/")
		authed.Use(middleware.USER_JWT())
		{
			authed.GET("/ping", api.CheckToken)
			authed.POST("/getPageToken", api.GetPageToken)
			// change user info
			authed.POST("user/SetUserInfo", api.SetUserInfo)
			// TODO: user check myday
			authed.GET("user/CheckMyday", api.UserCheckMyDay)
			// TODO: user get projects
			// authed.GET("user/GetProjects", api.UserGetProjectsIds)
		}
	}

	return r
}
