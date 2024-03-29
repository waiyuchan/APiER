package gin_release

import (
	"apier/internal/global/variable"
	"apier/internal/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
)

// ReleaseRouter 根据 gin 路由包官方的建议，gin 路由引擎如果在生产模式使用，官方建议设置为 release 模式
func ReleaseRouter() *gin.Engine {
	// 切换到生产模式禁用 gin 输出接口访问日志，经过并发测试验证，可以提升5%的性能
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	engine := gin.New()
	// 载入gin的中间件，关键是第二个中间件，我们对它进行了自定义重写，将可能的 panic 异常等，统一使用 zaplog 接管，保证全局日志打印统一
	engine.Use(gin.Logger(), CustomRecovery())
	return engine
}

// CustomRecovery 自定义错误(panic等)拦截中间件、对可能发生的错误进行拦截、统一记录
func CustomRecovery() gin.HandlerFunc {
	return gin.CustomRecoveryWithWriter(ioutil.Discard, func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(error); ok {
			// Log the error using zap
			variable.ZapLog.Error("Panic recovered", zap.Error(err))
			// Respond with a generic error message
			response.ErrorSystem(c, "", "Internal server error")
		}
	})
}
