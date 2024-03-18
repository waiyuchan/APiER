package cmd

import (
	"apier_platform/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init() // 初始化数据库连接

	r := gin.Default()
	// 设置路由...
	r.Run() // 在localhost:8080上启动服务
}
