package routes

import (
	"apier/internal/global/variable"
	"apier/internal/utils/gin_release"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func InitWebRouter() *gin.Engine {
	var router *gin.Engine

	if variable.ConfigYaml.GetBool("AppDebug") == false {

		//1.gin自行记录接口访问日志，不需要nginx，如果开启以下3行，那么请屏蔽第 34 行代码
		gin.DisableConsoleColor()
		f, _ := os.Create(variable.BasePath + variable.ConfigYaml.GetString("Logs.GinLogName"))
		gin.DefaultWriter = io.MultiWriter(f)
		router = gin_release.ReleaseRouter()

	} else {

		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		pprof.Register(router)
	}

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
