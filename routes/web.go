package routes

import (
	"apier/internal/global/consts"
	"apier/internal/global/variable"
	validatorFactory "apier/internal/http/validator/factory"
	"apier/internal/utils/gin_release"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func InitWebRouter() *gin.Engine {
	var router *gin.Engine

	if variable.ConfigYaml.GetBool("AppDebug") == false {
		// 1. gin自行记录接口访问日志，不需要nginx，如果开启以下3行，那么请屏蔽第 34 行代码
		gin.DisableConsoleColor()
		f, _ := os.Create(variable.BasePath + variable.ConfigYaml.GetString("Logs.GinLogName"))
		gin.DefaultWriter = io.MultiWriter(f)
		router = gin_release.ReleaseRouter()

	} else {
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		pprof.Register(router)
	}

	variable.ZapLog.Info("后台管理服务路由注册中......")

	admin := router.Group("/apier/api/admin/")
	{

		noAuthToSuperAdmin := admin.Group("super_admin/")
		{
			noAuthToSuperAdmin.GET("login", validatorFactory.Create(consts.ValidatorPrefix+"SuperAdminLogin"))

			// 一般不需要的时候要把超级管理员的注册入口关闭
			noAuthToSuperAdmin.POST("register", validatorFactory.Create(consts.ValidatorPrefix+"SuperAdminRegister"))
		}

		authToSuperAdmin := admin.Group("super_admin/")
		{
			authToSuperAdmin.POST("/")
			authToSuperAdmin.GET("/")
			authToSuperAdmin.DELETE("/")
			authToSuperAdmin.PUT("/")
		}

		authToAdmin := admin.Group("")
		{
			authToAdmin.POST("/")
		}

	}

	// 可以在这里继续添加其他的路由组和路由注册

	variable.ZapLog.Info("后台管理服务路由注册完成")

	return router
}
