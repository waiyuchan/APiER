package routes

import (
	"apier/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	adminRoutes := router.Group("/api/admin")
	{
		adminRoutes.POST("/super_admin_login", handlers.SuperAdminLogin)
		adminRoutes.POST("/", handlers.CreateAdmin)
		adminRoutes.GET("/", handlers.ListAdmins) // 这将是我们接下来要实现的接口
	}
	// 可以在这里继续添加其他的路由组和路由注册
}
