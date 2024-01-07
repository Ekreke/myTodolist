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

		// TODO: get user's projects
		v1.GET("user/GetProjectsIds", api.GetProjectsIds)

		// TODO: get user's apartments
		v1.GET("user/GetApartmentIds", api.GetApartmentIds)

		// TODO: set userinfo
		authed := v1.Group("/")
		authed.Use(middleware.USER_JWT())
		{
			//TODO: check token
			authed.GET("/ping", api.CheckToken)
			// change user info
			authed.POST("user/SetUserInfo", api.SetUserInfo)
		}
	}

	return r
}
