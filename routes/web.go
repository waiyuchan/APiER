package routes

import (
	"github.com/gin-gonic/gin"
)

func InitWebRouter() *gin.Engine {
	var router *gin.Engine

	admin := router.Group("/api/admin/")
	{

		noAuthToSuperAdmin := admin.Group("super_admin/")
		{
			noAuthToSuperAdmin.POST("login")
		}

		authToSuperAdmin := admin.Group("super_admin/")
		{
			authToSuperAdmin.POST("/")
			authToSuperAdmin.GET("/")
			authToSuperAdmin.DELETE("/")
			authToSuperAdmin.PUT("/")

		}

	}
	// 可以在这里继续添加其他的路由组和路由注册
	return router
}
